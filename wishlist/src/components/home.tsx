import React, { useState } from "react";
import "./home.css";
import { FaUserCircle, FaCog } from "react-icons/fa";

const Home: React.FC = () => {
  const handleLogout = () => {
    window.location.href = "/";
  };

  const handleUserClick = () => {
    console.log("User icon clicked!");
    // Add functionality here
  };

  const handleSettingsClick = () => {
    console.log("Settings icon clicked!");
    // Add functionality here
  };

  const [activeParentTab, setActiveParentTab] = useState("my-wishlists");
  const [activeChildTab, setActiveChildTab] = useState("friends");

  const renderChildTabContent = () => {
    if (activeParentTab === "my-wishlists") {
      switch (activeChildTab) {
        case "friends":
          return <div className="tab-content">Friends (Members) Content</div>;
        case "recent-gifts":
          return <div className="tab-content">Recent Gifts Content</div>;
        case "shop-now":
          return <div className="tab-content">Shop Now Content</div>;
        case "unreceived-gifts":
          return <div className="tab-content">Unreceived Gifts Content</div>;
        case "add-to-wishlist":
          return <div className="tab-content">Add More to Recent Wishlist Content</div>;
        default:
          return <div className="tab-content">Select a tab to view content</div>;
      }
    }

    if (activeParentTab === "draft-wishlist") {
      return <div className="tab-content">Draft Wishlist Content</div>;
    }

    if (activeParentTab === "shared-wishlists") {
      return <div className="tab-content">Shared Wishlists Content</div>;
    }

    return <div className="tab-content">No content available</div>;
  };

  return (
    <div className="home-page">
      <header className="header">
        <button className="icon-button" onClick={handleUserClick}>
          <FaUserCircle />
        </button>
        <button className="icon-button" onClick={handleSettingsClick}>
          <FaCog />
        </button>
        <button className="icon-button" onClick={handleLogout}>
          Logout
        </button>
      </header>

      <div className="title">
        <img src="/logo.svg" alt="Wishlist Logo" className="wishlist-logo" />
      </div>
      <div className="title-header">Wishlist</div>

      <main className="main-content">
        <div className="tabs-system">
          <div className="parent-tabs">
            <button
              className={`tab-button ${activeParentTab === "my-wishlists" ? "active" : ""}`}
              onClick={() => setActiveParentTab("my-wishlists")}
            >
              My Wishlists
            </button>
            <button
              className={`tab-button ${activeParentTab === "draft-wishlist" ? "active" : ""}`}
              onClick={() => setActiveParentTab("draft-wishlist")}
            >
              Draft Wishlist
            </button>
            <button
              className={`tab-button ${activeParentTab === "shared-wishlists" ? "active" : ""}`}
              onClick={() => setActiveParentTab("shared-wishlists")}
            >
              Shared Wishlists
            </button>
          </div>

          {activeParentTab === "my-wishlists" && (
            <div className="child-tabs">
              <button
                className={`tab-button ${activeChildTab === "friends" ? "active" : ""}`}
                onClick={() => setActiveChildTab("friends")}
              >
                Friends
              </button>
              <button
                className={`tab-button ${activeChildTab === "recent-gifts" ? "active" : ""}`}
                onClick={() => setActiveChildTab("recent-gifts")}
              >
                Recent Gifts
              </button>
              <button
                className={`tab-button ${activeChildTab === "shop-now" ? "active" : ""}`}
                onClick={() => setActiveChildTab("shop-now")}
              >
                Shop Now
              </button>
              <button
                className={`tab-button ${activeChildTab === "unreceived-gifts" ? "active" : ""}`}
                onClick={() => setActiveChildTab("unreceived-gifts")}
              >
                Unreceived Gifts
              </button>
              <button
                className={`tab-button ${activeChildTab === "add-to-wishlist" ? "active" : ""}`}
                onClick={() => setActiveChildTab("add-to-wishlist")}
              >
                Add More to Wishlist
              </button>
            </div>
          )}

          <div className="content-area">{renderChildTabContent()}</div>
        </div>
      </main>
    </div>
  );
};

export default Home;
