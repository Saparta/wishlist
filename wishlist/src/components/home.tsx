// Home.tsx
import React from "react";
import "./home.css"

const Home: React.FC = () => {
  const handleLogout = () => {
    // Redirect to the landing page
    window.location.href = "/";
  };

  return (
    <div className="home-page">
      <header className="header">
        <div className="user-profile">
          <p>User Profile</p>
        </div>
        <div className="settings">
          <p>Settings</p>
        </div>
      </header>

      <main className="main-content">
        <section className="recent-lists">
          <h2>Recent Lists</h2>
          <div className="list-item">List 1</div>
          <div className="list-item">List 2</div>
        </section>

        <section className="aside-widget">
          <h2>Friends & Recent Gifts</h2>
          <div className="gift">Gift 1</div>
          <div className="gift">Gift 2</div>
        </section>

        <section className="carousel-widget">
          <h2>Unreceived Gifts</h2>
          {/* Your carousel logic goes here */}
        </section>

        <section className="retailers-widget">
          <h2>Shop Now</h2>
          <div className="retailer">Amazon</div>
          <div className="retailer">Best Buy</div>
        </section>
      </main>

      <footer className="footer">
        <button className="logout-btn" onClick={handleLogout}>Log Out</button>
      </footer>
    </div>
  );
};

export default Home;
