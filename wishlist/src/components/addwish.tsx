import React, { useState } from "react";
import "./addwish.css";
import "./home.css";
import { RiUserLine, RiUserSettingsLine } from "react-icons/ri";
import UserModal from "./userModal";
import "./userModal.css";
import "./wishlistModal.css";
import WishlistModal from "./wishlistModal";
import { MdOutlineAdd, MdDelete } from "react-icons/md";
import { MdArrowBackIosNew } from "react-icons/md";

const AddWish: React.FC = () => {
  const [isUserModalOpen, setIsUserModalOpen] = useState(false);
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [items, setItems] = useState([
    { itemName: "", itemLink: "", itemPrice: "" },
  ]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    console.log("New Wishlist:", { title, description, items });
  };

  const handleLogout = () => {
    window.location.href = "/";
  };

  const handlehome = () => {
    console.log("back to home page!");
    window.location.href = "/home";
  };

  const handleItemChange = (index: number, field: string, value: string) => {
    const newItems = [...items];
    newItems[index] = { ...newItems[index], [field]: value };
    setItems(newItems);
  };

  const addItem = () => {
    setItems([...items, { itemName: "", itemLink: "", itemPrice: "" }]);
  };

  const removeItem = (index: number) => {
    if (items.length > 1) {
      setItems(items.filter((_, i) => i !== index));
    }
  };

  return (
    <div className="add-page">
      <header className="header">
        <div className="logo-container">
          <img src="/logo.svg" alt="Wishlist Logo" className="wishlist-logo" />
          <div className="title-header">Wishlist</div>
        </div>
        <div className="actions">
          <button
            className="icon-button"
            onClick={() => setIsUserModalOpen(true)}
          >
            <RiUserLine />
          </button>
          {isUserModalOpen && (
            <UserModal onClose={() => setIsUserModalOpen(false)} />
          )}
          <button className="icon-button" onClick={handleLogout}>
            Logout
          </button>
        </div>
      </header>

      <div className="add-wish-head">
        <button className="back-button" onClick={handlehome}>
          <MdArrowBackIosNew />
        </button>
        <h1 className="text-2xl font-bold mb-4">Create a New Wishlist</h1>
      </div>

      <form onSubmit={handleSubmit} className="addwish-form">
        <div className="form-field">
          <label htmlFor="title" className="form-label">
            Wishlist Title
          </label>
          <input
            type="text"
            id="title"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            className="form-input"
            placeholder="Enter wishlist title"
          />
        </div>

        <div className="form-field">
          <label htmlFor="description" className="form-label">
            Description
          </label>
          <textarea
            id="description"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            className="form-input"
            placeholder="Enter wishlist description"
          />
        </div>

        <h2 className="text-xl font-bold mt-4">Wishlist Items</h2>

        {items.map((item, index) => (
          <div key={index} className="wishlist-item">
            <div className="form-field">
              <label className="form-label">Item Name</label>
              <input
                type="text"
                value={item.itemName}
                onChange={(e) =>
                  handleItemChange(index, "itemName", e.target.value)
                }
                className="form-input"
                placeholder="Enter item name"
              />
            </div>

            <div className="form-field">
              <label className="form-label">Item Link</label>
              <input
                type="url"
                value={item.itemLink}
                onChange={(e) =>
                  handleItemChange(index, "itemLink", e.target.value)
                }
                className="form-input"
                placeholder="Enter item link"
              />
            </div>

            <div className="form-field">
              <label className="form-label">Item Price ($)</label>
              <input
                type="number"
                value={item.itemPrice}
                onChange={(e) =>
                  handleItemChange(index, "itemPrice", e.target.value)
                }
                className="form-input"
                placeholder="Enter item price"
              />
            </div>

            {items.length > 1 && (
              <button
                type="button"
                className="remove-item-button"
                onClick={() => removeItem(index)}
              >
                <MdDelete /> Remove Item
              </button>
            )}
          </div>
        ))}

        <button type="button" className="add-item-button" onClick={addItem}>
          <MdOutlineAdd /> Add Another Item
        </button>

        <button type="submit" className="submit-button">
          Save Wishlist
        </button>
      </form>

      <footer className="footer">© 2025 Wishlist, All Rights Reserved.</footer>
    </div>
  );
};

export default AddWish;
