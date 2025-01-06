import React, { useState } from "react";
import "./home.css";
import { FaUserCircle, FaCog } from "react-icons/fa";
import { RiAccountCircle2Line, RiCloseLine, RiUserLine, RiUserSettingsLine } from "react-icons/ri";
import UserModal from './userModal';
import "./userModal.css";
import "./wishlistModal.css"
import { mockWishlists, mocksharedWishlist } from "../mocks/mockData";
import WishlistModal from './wishlistModal';


interface Wishlist {
  id: number;
  title: string;
  items: string[];
  progress: number;
  description: string;
}

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

  const [activeTab, setactiveTab] = useState("my-wishlists");
  // Child tabs disabled
  // const [activeChildTab, setActiveChildTab] = useState("friends");

  // const renderChildTabContent = () => {
  //   if (activeTab === "my-wishlists") {
  //     switch (activeChildTab) {
  //       case "friends":
  //         return <div className="tab-content">Friends (Members) Content</div>;
  //       case "recent-gifts":
  //         return <div className="tab-content">Recent Gifts Content</div>;
  //       case "shop-now":
  //         return <div className="tab-content">Shop Now Content</div>;
  //       case "unreceived-gifts":
  //         return <div className="tab-content">Unreceived Gifts Content</div>;
  //       case "add-to-wishlist":
  //         return <div className="tab-content">Add More to Recent Wishlist Content</div>;
  //       default:
  //         return <div className="tab-content">Select a tab to view content</div>;
  //     }
  //   }

  //   if (activeTab === "shared-wishlists") {
  //     return <div className="tab-content">Shared Wishlists Content</div>;
  //   }

  //   return <div className="tab-content">No content available</div>;
  // };

  const [isModalOpen, setIsModalOpen] = useState(false);

  const renderWishlistCards = (wishlists: Wishlist[]) => {
    return (
      <div className="wishlist-grid">
        {wishlists.map((wishlist: Wishlist) => (
          <div
            key={wishlist.id}
            className="wishlist-card"
            onClick={() => handleCardClick(wishlist)} // Add the click handler here
          >
            <h3>{wishlist.title}</h3>
            <ul>
              {wishlist.items.map((item: string, index: number) => (
                <li key={index}>
                  <label>
                    <input type="checkbox" />
                    {item}
                  </label>
                </li>
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
        ))}
      </div>
    );
  };


  const [selectedWishlist, setSelectedWishlist] = useState<Wishlist | null>(null);



  const handleCardClick = (wishlist: Wishlist) => {
    console.log(wishlist);
    setSelectedWishlist(wishlist);
    setIsModalOpen(true);
  };


  const closeModal = () => {
    setIsModalOpen(false);
    setSelectedWishlist(null);
  };

  return (
    <div className="home-page">
      <header className="header">
        <div className="logo-container">
          <img src="/logo.svg" alt="Wishlist Logo" className="wishlist-logo" />
          <div className="title-header">Wishlist</div>
        </div>
        <div className="actions">
          <div>
            <button className="icon-button" onClick={() => setIsModalOpen(true)}>
              <RiUserLine />
            </button>
            {isModalOpen && <UserModal onClose={() => setIsModalOpen(false)} />}
          </div>
          <button className="icon-button" onClick={handleSettingsClick}>
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
              className={`tab-button ${activeTab === "my-wishlists" ? "active" : ""}`}
              onClick={() => setactiveTab("my-wishlists")}
            >
              My Wishlists
            </button>
            <button
              className={`tab-button ${activeTab === "shared-wishlists" ? "active" : ""}`}
              onClick={() => setactiveTab("shared-wishlists")}
            >
              Shared Wishlists
            </button>
          </div>

          {/* Child tabs  disabled for now */}
          {/* {activeParentTab === "my-wishlists" && (
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
          )} */}
          {/* <div className="content-area">{renderChildTabContent()}</div> */}

          {/* My Wishlists */}
          {activeTab === "my-wishlists" && renderWishlistCards(mockWishlists)}

          {/* Shared Wishlists */}
          {activeTab === "shared-wishlists" && renderWishlistCards(mocksharedWishlist)}


        </div>
      </main>

      {/* Wishlist Modal */}
      {isModalOpen && selectedWishlist && (
        <WishlistModal wishlist={selectedWishlist} onClose={closeModal} />
      )}

      <footer className="footer">
        © 2025 Wishlist, All Rights Reserved.
      </footer>
    </div>

  );


};

export default Home;
