// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: wishlist-service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetUserWishlistsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserWishlistsRequest) Reset() {
	*x = GetUserWishlistsRequest{}
	mi := &file_wishlist_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserWishlistsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserWishlistsRequest) ProtoMessage() {}

func (x *GetUserWishlistsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wishlist_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserWishlistsRequest.ProtoReflect.Descriptor instead.
func (*GetUserWishlistsRequest) Descriptor() ([]byte, []int) {
	return file_wishlist_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserWishlistsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUserWishlistsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Wishlists     []*Wishlist            `protobuf:"bytes,1,rep,name=wishlists,proto3" json:"wishlists,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserWishlistsResponse) Reset() {
	*x = GetUserWishlistsResponse{}
	mi := &file_wishlist_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserWishlistsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserWishlistsResponse) ProtoMessage() {}

func (x *GetUserWishlistsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wishlist_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserWishlistsResponse.ProtoReflect.Descriptor instead.
func (*GetUserWishlistsResponse) Descriptor() ([]byte, []int) {
	return file_wishlist_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserWishlistsResponse) GetWishlists() []*Wishlist {
	if x != nil {
		return x.Wishlists
	}
	return nil
}

type Wishlist struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Title         string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	IsPublic      bool                   `protobuf:"varint,5,opt,name=isPublic,proto3" json:"isPublic,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	LastModified  *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=lastModified,proto3" json:"lastModified,omitempty"`
	LastOpened    *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=lastOpened,proto3" json:"lastOpened,omitempty"`
	Items         []*WishlistItem        `protobuf:"bytes,9,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Wishlist) Reset() {
	*x = Wishlist{}
	mi := &file_wishlist_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Wishlist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wishlist) ProtoMessage() {}

func (x *Wishlist) ProtoReflect() protoreflect.Message {
	mi := &file_wishlist_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wishlist.ProtoReflect.Descriptor instead.
func (*Wishlist) Descriptor() ([]byte, []int) {
	return file_wishlist_service_proto_rawDescGZIP(), []int{2}
}

func (x *Wishlist) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Wishlist) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Wishlist) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Wishlist) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Wishlist) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

func (x *Wishlist) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Wishlist) GetLastModified() *timestamppb.Timestamp {
	if x != nil {
		return x.LastModified
	}
	return nil
}

func (x *Wishlist) GetLastOpened() *timestamppb.Timestamp {
	if x != nil {
		return x.LastOpened
	}
	return nil
}

func (x *Wishlist) GetItems() []*WishlistItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type WishlistItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Url           string                 `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Price         float32                `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	IsGifted      bool                   `protobuf:"varint,5,opt,name=is_gifted,json=isGifted,proto3" json:"is_gifted,omitempty"`
	GiftedBy      string                 `protobuf:"bytes,6,opt,name=gifted_by,json=giftedBy,proto3" json:"gifted_by,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WishlistItem) Reset() {
	*x = WishlistItem{}
	mi := &file_wishlist_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WishlistItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WishlistItem) ProtoMessage() {}

func (x *WishlistItem) ProtoReflect() protoreflect.Message {
	mi := &file_wishlist_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WishlistItem.ProtoReflect.Descriptor instead.
func (*WishlistItem) Descriptor() ([]byte, []int) {
	return file_wishlist_service_proto_rawDescGZIP(), []int{3}
}

func (x *WishlistItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WishlistItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WishlistItem) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *WishlistItem) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *WishlistItem) GetIsGifted() bool {
	if x != nil {
		return x.IsGifted
	}
	return false
}

func (x *WishlistItem) GetGiftedBy() string {
	if x != nil {
		return x.GiftedBy
	}
	return ""
}

func (x *WishlistItem) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type CreateWishlistRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	IsPublic      bool                   `protobuf:"varint,4,opt,name=isPublic,proto3" json:"isPublic,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateWishlistRequest) Reset() {
	*x = CreateWishlistRequest{}
	mi := &file_wishlist_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateWishlistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWishlistRequest) ProtoMessage() {}

func (x *CreateWishlistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wishlist_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWishlistRequest.ProtoReflect.Descriptor instead.
func (*CreateWishlistRequest) Descriptor() ([]byte, []int) {
	return file_wishlist_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateWishlistRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateWishlistRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateWishlistRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateWishlistRequest) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

type CreateWishlistResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Title         string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	IsPublic      bool                   `protobuf:"varint,5,opt,name=isPublic,proto3" json:"isPublic,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	LastModified  *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=lastModified,proto3" json:"lastModified,omitempty"`
	LastOpened    *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=lastOpened,proto3" json:"lastOpened,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateWishlistResponse) Reset() {
	*x = CreateWishlistResponse{}
	mi := &file_wishlist_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateWishlistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWishlistResponse) ProtoMessage() {}

func (x *CreateWishlistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wishlist_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWishlistResponse.ProtoReflect.Descriptor instead.
func (*CreateWishlistResponse) Descriptor() ([]byte, []int) {
	return file_wishlist_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateWishlistResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateWishlistResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateWishlistResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateWishlistResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateWishlistResponse) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

func (x *CreateWishlistResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CreateWishlistResponse) GetLastModified() *timestamppb.Timestamp {
	if x != nil {
		return x.LastModified
	}
	return nil
}

func (x *CreateWishlistResponse) GetLastOpened() *timestamppb.Timestamp {
	if x != nil {
		return x.LastOpened
	}
	return nil
}

type ClearWishlistItemsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	WishlistId    string                 `protobuf:"bytes,2,opt,name=wishlistId,proto3" json:"wishlistId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClearWishlistItemsRequest) Reset() {
	*x = ClearWishlistItemsRequest{}
	mi := &file_wishlist_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClearWishlistItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearWishlistItemsRequest) ProtoMessage() {}

func (x *ClearWishlistItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wishlist_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearWishlistItemsRequest.ProtoReflect.Descriptor instead.
func (*ClearWishlistItemsRequest) Descriptor() ([]byte, []int) {
	return file_wishlist_service_proto_rawDescGZIP(), []int{6}
}

func (x *ClearWishlistItemsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ClearWishlistItemsRequest) GetWishlistId() string {
	if x != nil {
		return x.WishlistId
	}
	return ""
}

type ClearWishlistItemsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClearWishlistItemsResponse) Reset() {
	*x = ClearWishlistItemsResponse{}
	mi := &file_wishlist_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClearWishlistItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearWishlistItemsResponse) ProtoMessage() {}

func (x *ClearWishlistItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wishlist_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearWishlistItemsResponse.ProtoReflect.Descriptor instead.
func (*ClearWishlistItemsResponse) Descriptor() ([]byte, []int) {
	return file_wishlist_service_proto_rawDescGZIP(), []int{7}
}

type AddWishlistItemRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	WishlistId    string                 `protobuf:"bytes,2,opt,name=wishlistId,proto3" json:"wishlistId,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Url           string                 `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Price         float32                `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
	IsGifted      bool                   `protobuf:"varint,6,opt,name=isGifted,proto3" json:"isGifted,omitempty"`
	GiftedBy      string                 `protobuf:"bytes,7,opt,name=giftedBy,proto3" json:"giftedBy,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddWishlistItemRequest) Reset() {
	*x = AddWishlistItemRequest{}
	mi := &file_wishlist_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddWishlistItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddWishlistItemRequest) ProtoMessage() {}

func (x *AddWishlistItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wishlist_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddWishlistItemRequest.ProtoReflect.Descriptor instead.
func (*AddWishlistItemRequest) Descriptor() ([]byte, []int) {
	return file_wishlist_service_proto_rawDescGZIP(), []int{8}
}

func (x *AddWishlistItemRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddWishlistItemRequest) GetWishlistId() string {
	if x != nil {
		return x.WishlistId
	}
	return ""
}

func (x *AddWishlistItemRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddWishlistItemRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *AddWishlistItemRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *AddWishlistItemRequest) GetIsGifted() bool {
	if x != nil {
		return x.IsGifted
	}
	return false
}

func (x *AddWishlistItemRequest) GetGiftedBy() string {
	if x != nil {
		return x.GiftedBy
	}
	return ""
}

type AddWishlistItemResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	WishlistId    string                 `protobuf:"bytes,3,opt,name=wishlistId,proto3" json:"wishlistId,omitempty"`
	Name          string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Url           string                 `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	Price         float32                `protobuf:"fixed32,6,opt,name=price,proto3" json:"price,omitempty"`
	IsGifted      bool                   `protobuf:"varint,7,opt,name=isGifted,proto3" json:"isGifted,omitempty"`
	GiftedBy      string                 `protobuf:"bytes,8,opt,name=giftedBy,proto3" json:"giftedBy,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddWishlistItemResponse) Reset() {
	*x = AddWishlistItemResponse{}
	mi := &file_wishlist_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddWishlistItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddWishlistItemResponse) ProtoMessage() {}

func (x *AddWishlistItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wishlist_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddWishlistItemResponse.ProtoReflect.Descriptor instead.
func (*AddWishlistItemResponse) Descriptor() ([]byte, []int) {
	return file_wishlist_service_proto_rawDescGZIP(), []int{9}
}

func (x *AddWishlistItemResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AddWishlistItemResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddWishlistItemResponse) GetWishlistId() string {
	if x != nil {
		return x.WishlistId
	}
	return ""
}

func (x *AddWishlistItemResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddWishlistItemResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *AddWishlistItemResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *AddWishlistItemResponse) GetIsGifted() bool {
	if x != nil {
		return x.IsGifted
	}
	return false
}

func (x *AddWishlistItemResponse) GetGiftedBy() string {
	if x != nil {
		return x.GiftedBy
	}
	return ""
}

func (x *AddWishlistItemResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_wishlist_service_proto protoreflect.FileDescriptor

var file_wishlist_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x31, 0x0a, 0x17, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x73, 0x68, 0x6c,
	0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x18, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69,
	0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x09, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x69, 0x73, 0x68, 0x6c,
	0x69, 0x73, 0x74, 0x52, 0x09, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x22, 0xe7,
	0x02, 0x0a, 0x08, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x3e, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x12, 0x29, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xcf, 0x01, 0x0a, 0x0c, 0x57, 0x69, 0x73,
	0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x67, 0x69, 0x66, 0x74,
	0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x47, 0x69, 0x66, 0x74,
	0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x69, 0x66, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x69, 0x66, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x83, 0x01, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x22, 0xca, 0x02, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x69, 0x73, 0x68, 0x6c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x3e, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x22, 0x53, 0x0a,
	0x19, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74,
	0x49, 0x64, 0x22, 0x1c, 0x0a, 0x1a, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x57, 0x69, 0x73, 0x68, 0x6c,
	0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0xc4, 0x01, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x73, 0x47, 0x69, 0x66, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x69, 0x73, 0x47, 0x69, 0x66, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x67,
	0x69, 0x66, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67,
	0x69, 0x66, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x8f, 0x02, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x57,
	0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x77,
	0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x47, 0x69, 0x66,
	0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x47, 0x69, 0x66,
	0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x69, 0x66, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x69, 0x66, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12,
	0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xea, 0x02, 0x0a, 0x0f, 0x57, 0x69,
	0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x69,
	0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x69, 0x73, 0x68,
	0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55,
	0x0a, 0x10, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73,
	0x74, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x57, 0x69, 0x73, 0x68,
	0x6c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x64, 0x64, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x64, 0x64, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x12, 0x43, 0x6c, 0x65,
	0x61, 0x72, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x57, 0x69, 0x73,
	0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x57,
	0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wishlist_service_proto_rawDescOnce sync.Once
	file_wishlist_service_proto_rawDescData = file_wishlist_service_proto_rawDesc
)

func file_wishlist_service_proto_rawDescGZIP() []byte {
	file_wishlist_service_proto_rawDescOnce.Do(func() {
		file_wishlist_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_wishlist_service_proto_rawDescData)
	})
	return file_wishlist_service_proto_rawDescData
}

var file_wishlist_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_wishlist_service_proto_goTypes = []any{
	(*GetUserWishlistsRequest)(nil),    // 0: proto.getUserWishlistsRequest
	(*GetUserWishlistsResponse)(nil),   // 1: proto.getUserWishlistsResponse
	(*Wishlist)(nil),                   // 2: proto.Wishlist
	(*WishlistItem)(nil),               // 3: proto.WishlistItem
	(*CreateWishlistRequest)(nil),      // 4: proto.CreateWishlistRequest
	(*CreateWishlistResponse)(nil),     // 5: proto.CreateWishlistResponse
	(*ClearWishlistItemsRequest)(nil),  // 6: proto.ClearWishlistItemsRequest
	(*ClearWishlistItemsResponse)(nil), // 7: proto.ClearWishlistItemsResponse
	(*AddWishlistItemRequest)(nil),     // 8: proto.AddWishlistItemRequest
	(*AddWishlistItemResponse)(nil),    // 9: proto.AddWishlistItemResponse
	(*timestamppb.Timestamp)(nil),      // 10: google.protobuf.Timestamp
}
var file_wishlist_service_proto_depIdxs = []int32{
	2,  // 0: proto.getUserWishlistsResponse.wishlists:type_name -> proto.Wishlist
	10, // 1: proto.Wishlist.createdAt:type_name -> google.protobuf.Timestamp
	10, // 2: proto.Wishlist.lastModified:type_name -> google.protobuf.Timestamp
	10, // 3: proto.Wishlist.lastOpened:type_name -> google.protobuf.Timestamp
	3,  // 4: proto.Wishlist.items:type_name -> proto.WishlistItem
	10, // 5: proto.WishlistItem.created_at:type_name -> google.protobuf.Timestamp
	10, // 6: proto.CreateWishlistResponse.createdAt:type_name -> google.protobuf.Timestamp
	10, // 7: proto.CreateWishlistResponse.lastModified:type_name -> google.protobuf.Timestamp
	10, // 8: proto.CreateWishlistResponse.lastOpened:type_name -> google.protobuf.Timestamp
	10, // 9: proto.AddWishlistItemResponse.createdAt:type_name -> google.protobuf.Timestamp
	4,  // 10: proto.WishlistService.CreateWishlist:input_type -> proto.CreateWishlistRequest
	0,  // 11: proto.WishlistService.getUserWishlists:input_type -> proto.getUserWishlistsRequest
	8,  // 12: proto.WishlistService.AddWishlistItem:input_type -> proto.AddWishlistItemRequest
	6,  // 13: proto.WishlistService.ClearWishlistItems:input_type -> proto.ClearWishlistItemsRequest
	5,  // 14: proto.WishlistService.CreateWishlist:output_type -> proto.CreateWishlistResponse
	1,  // 15: proto.WishlistService.getUserWishlists:output_type -> proto.getUserWishlistsResponse
	9,  // 16: proto.WishlistService.AddWishlistItem:output_type -> proto.AddWishlistItemResponse
	7,  // 17: proto.WishlistService.ClearWishlistItems:output_type -> proto.ClearWishlistItemsResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_wishlist_service_proto_init() }
func file_wishlist_service_proto_init() {
	if File_wishlist_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wishlist_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wishlist_service_proto_goTypes,
		DependencyIndexes: file_wishlist_service_proto_depIdxs,
		MessageInfos:      file_wishlist_service_proto_msgTypes,
	}.Build()
	File_wishlist_service_proto = out.File
	file_wishlist_service_proto_rawDesc = nil
	file_wishlist_service_proto_goTypes = nil
	file_wishlist_service_proto_depIdxs = nil
}
