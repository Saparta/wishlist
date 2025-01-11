package endpoints

import (
	"context"
	"time"

	pb "github.com/Saparta/wishlist/wishlist/services/wishlist-service/proto"
	"github.com/Saparta/wishlist/wishlist/services/wishlist-service/shared"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (w *WishlistService) GetWishlist(ctx context.Context, request *pb.GetWishlistRequest) (*pb.GetWishlistResponse, error) {
	dbPool, err := shared.ConnectToDatabase(ctx)
	if err != nil {
		return nil, err
	}

	var getSharedWith chan []string = make(chan []string, 1)
	var errChannel chan error = make(chan error, 1)
	go shared.GetWishlistSharedUsers(ctx, *request.WishlistId, getSharedWith, errChannel)

	var query string
	if *request.IsOwnWishlist {
		query = `
  WITH updated_wishlist AS (
    UPDATE wishlists
    SET last_opened = CURRENT_TIMESTAMP
    WHERE user_id = $1 AND id = $2
    RETURNING id, title, description, is_public, last_opened, last_modified
	)
	SELECT
    uw.id AS wishlist_id, 
    uw.title AS wishlist_title, 
    uw.description AS wishlist_description,
    uw.is_public AS is_public,
    uw.last_opened AS wishlist_last_opened,
    uw.last_modified AS wishlist_last_modified,
    i.id AS item_id,
    i.name AS item_name,
    i.url AS item_url,
    i.price AS item_price,
    i.is_gifted AS item_is_gifted,
    i.gifted_by AS item_gifted_by,
    i.created_at AS item_created_at
	FROM updated_wishlist uw
	LEFT JOIN items i ON uw.id = i.wishlist_id
	ORDER BY i.created_at;`
	} else {
		query = `
	WITH authorized_wishlists AS (
    SELECT w.id AS wishlist_id
    FROM wishlists w
    JOIN shared s ON w.id = s.wishlist_id
    WHERE w.id = $2
      AND (w.is_public = TRUE OR s.shared_with = $1)
	),
	updated_shared AS (
    UPDATE shared
    SET last_opened = CURRENT_TIMESTAMP
    WHERE shared_with = $1 
      AND wishlist_id = $2
      AND $2 IN (SELECT wishlist_id FROM authorized_wishlists)
    RETURNING wishlist_id
	)
	SELECT
    w.id AS wishlist_id, 
    w.title AS wishlist_title, 
    w.description AS wishlist_description,
    w.is_public AS is_public,
    w.last_opened AS wishlist_last_opened,
    w.last_modified AS wishlist_last_modified,
    i.id AS item_id,
    i.name AS item_name,
    i.url AS item_url,
    i.price AS item_price,
    i.is_gifted AS item_is_gifted,
    i.gifted_by AS item_gifted_by,
    i.created_at AS item_created_at
	FROM wishlists w
	LEFT JOIN items i ON w.id = i.wishlist_id
	WHERE w.id IN (SELECT wishlist_id FROM authorized_wishlists)
	ORDER BY i.created_at;`
	}

	rows, err := dbPool.Query(ctx, query, request.UserId, request.WishlistId)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to query wishlists: %v", err.Error())
	}

	var wishlist pb.Wishlist
	var wishLastOpened, wishLastModified, itemCreatedAt *time.Time
	items, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (*pb.WishlistItem, error) {
		var item pb.WishlistItem
		err := row.Scan(&wishlist.Id, &wishlist.Title, &wishlist.Description, &wishlist.IsPublic, &wishLastOpened, &wishLastModified, &item.Id, &item.Name, &item.Url, &item.Price, &item.IsGifted, &item.GiftedBy, &itemCreatedAt)
		if err != nil {
			return nil, err
		}
		wishlist.LastModified = timestamppb.New(*wishLastModified)
		wishlist.LastOpened = timestamppb.New(*wishLastOpened)
		var createdAt *timestamppb.Timestamp
		if itemCreatedAt == nil {
			createdAt = timestamppb.Now()
		} else {
			createdAt = timestamppb.New(*itemCreatedAt)
		}
		item.CreatedAt = createdAt
		return &item, nil
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to process query results: %v", err.Error())
	}

	wishlist.Items = []*pb.WishlistItem{}
	for _, item := range items {
		if item.Id != nil {
			if *request.IsOwnWishlist {
				item.GiftedBy = nil
				item.IsGifted = nil
			}
			wishlist.Items = append(wishlist.Items, item)
		}
	}

	select {
	case sharedWithUsers := <-getSharedWith:
		wishlist.SharedWith = sharedWithUsers
	case err = <-errChannel:
		return nil, err
	}

	return &pb.GetWishlistResponse{
		Wishlist: &wishlist,
	}, nil
}
