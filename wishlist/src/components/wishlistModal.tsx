import React from "react";
import { RiCloseLine } from "react-icons/ri";

interface Wishlist {
  id: number;
  title: string;
  items: string[];
  progress: number;
  description: string;
}

interface WishlistModalProps {
  wishlist: Wishlist;
  onClose: () => void;
}

const WishlistModal: React.FC<WishlistModalProps> = ({ wishlist, onClose }) => {
  return (
    <div className="modal-overlay">
      <div className="modal">
        <div className="modal-header">
          <h2>{wishlist.title}</h2>
          <RiCloseLine className="close-icon" onClick={onClose} />
        </div>

        <div className="modal-content">
          <p>{wishlist.description}</p>
          <h3>Items:</h3>
          <ul>
            {wishlist.items.map((item, index) => (
              <li key={index}>{item}</li>
            ))}
          </ul>
          <div className="progress-bar">
            <div
              className="progress"
              style={{ width: `${wishlist.progress}%` }}
            ></div>
          </div>
          <p>{wishlist.progress}% progress</p>
        </div>
      </div>
    </div>
  );
};

export default WishlistModal;
