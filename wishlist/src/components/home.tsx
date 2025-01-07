import React, { useState } from "react";
import "./home.css";
import { RiUserLine, RiUserSettingsLine } from "react-icons/ri";
import UserModal from "./userModal";
import "./userModal.css";
import "./wishlistModal.css";
import {
  mockSharedWishlists,
  mockWishlists,
  Wishlist,
} from "../mocks/mockData";
import WishlistModal from "./wishlistModal";

const Home: React.FC = () => {
  const [isUserModalOpen, setIsUserModalOpen] = useState(false);
  const [isWishlistModalOpen, setIsWishlistModalOpen] = useState(false);
  const [selectedWishlist, setSelectedWishlist] = useState<Wishlist | null>(
    null
  );
  const [activeTab, setActiveTab] = useState("my-wishlists");

  const handleLogout = () => {
    window.location.href = "/";
  };

  const handleUserClick = () => {
    console.log("User icon clicked!");
    setIsUserModalOpen(true);
  };

  const handleCardClick = (wishlist: Wishlist) => {
    console.log("Opening wishlist modal for:", wishlist.title);
    setSelectedWishlist(wishlist);
    setIsWishlistModalOpen(true);
  };

  const closeModal = () => {
    console.log("Closing wishlist modal.");
    setIsWishlistModalOpen(false);
    setSelectedWishlist(null);
  };

  const calculateProgress = (items: Wishlist["items"]) => {
    if (!items || items.length === 0) return 0;
    const giftedCount = items.filter((item) => item.is_gifted).length;
    return Math.round((giftedCount / items.length) * 100);
  };

  const renderWishlistCards = (wishlists: Wishlist[]) => (
    <div className="wishlist-grid">
      {wishlists.map((wishlist) => {
        const progress = calculateProgress(wishlist.items);
        return (
          <div
            key={wishlist.id}
            className="wishlist-card"
            onClick={() => handleCardClick(wishlist)}
          >
            <h3>{wishlist.title}</h3>
            <ul>
              {wishlist.items.map((item) => (
                <li key={item.id}>
                  <label>
                    <input
                      type="checkbox"
                      checked={item.is_gifted}
                      readOnly
                      style={{
                        accentColor: item.is_gifted ? "green" : "gray",
                      }}
                    />
                    {item.name} - ${item.price.toFixed(2)}
                  </label>
                </li>
              ))}
            </ul>
            <div className="progress-bar">
              <div
                className="progress"
                style={{
                  width: `${progress}%`,
                  backgroundColor: progress === 100 ? "green" : "orange",
                }}
              ></div>
            </div>
            <p>{progress}% progress</p>
          </div>
        );
      })}
    </div>
  );

  return (
    <div className="home-page">
      <header className="header">
        <div className="logo-container">
          <img src="/logo.svg" alt="Wishlist Logo" className="wishlist-logo" />
          <div className="title-header">Wishlist</div>
        </div>
        <div className="actions">
          <button className="icon-button" onClick={handleUserClick}>
            <RiUserLine />
          </button>
          {isUserModalOpen && (
            <UserModal onClose={() => setIsUserModalOpen(false)} />
          )}
          <button className="icon-button" onClick={() => {}}>
            <RiUserSettingsLine />
          </button>
          <button className="icon-button" onClick={handleLogout}>
            Logout
          </button>
        </div>
      </header>

      <main className="main-content">
        <div className="tabs-system">
          <div className="parent-tabs">
            <button
              className={`tab-button ${
                activeTab === "my-wishlists" ? "active" : ""
              }`}
              onClick={() => setActiveTab("my-wishlists")}
            >
              My Wishlists
            </button>
            <button
              className={`tab-button ${
                activeTab === "shared-wishlists" ? "active" : ""
              }`}
              onClick={() => setActiveTab("shared-wishlists")}
            >
              Shared Wishlists
            </button>
          </div>

          {activeTab === "my-wishlists" && renderWishlistCards(mockWishlists)}
          {activeTab === "shared-wishlists" &&
            renderWishlistCards(mockSharedWishlists)}
        </div>
      </main>

      {/* Wishlist Modal */}
      {isWishlistModalOpen && selectedWishlist && (
        <WishlistModal wishlist={selectedWishlist} onClose={closeModal} />
      )}

      <footer className="footer">© 2025 Wishlist, All Rights Reserved.</footer>
    </div>
  );
};

export default Home;
