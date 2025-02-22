import React from "react";
import { Wishlist } from "../mocks/mockData";
import { RiCloseFill } from "react-icons/ri";
import { CgTrash } from "react-icons/cg";

interface WishlistModalProps {
  wishlist: Wishlist | null;
  onClose: () => void;
}

const WishlistModal: React.FC<WishlistModalProps> = ({ wishlist, onClose }) => {
  if (!wishlist) return null;

  return (
    <div className="modal-overlay">
      <div className="wishlist-modal">
        <div className="wishlist-modal-header">
          <div className="close-button"></div>
          <div className="left-side-modal">
            <h2>{wishlist.title}</h2>
          </div>

          <div className="right-side-modal">
            <button onClick={onClose} className="trash-button">
              <CgTrash />
            </button>

            <button onClick={onClose} className="close-button">
              <RiCloseFill />
            </button>
          </div>

          {/* Edit trash Icon to delete wish list with new function */}
        </div>
        <p>{wishlist.description}</p>
        <ul>
          {wishlist.items.map((item) => (
            <li key={item.id}>
              <p>
                <strong>{item.name}</strong> - ${item.price.toFixed(2)}
              </p>
              <p>
                {item.is_gifted
                  ? `Gifted by: ${item.gifted_by}`
                  : "Not gifted yet"}
              </p>
              <a href={item.url} target="_blank" rel="noopener noreferrer">
                View Item
              </a>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default WishlistModal;
