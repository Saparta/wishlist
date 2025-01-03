import React from 'react';
import "./landingpage.css";
import { useNavigate } from 'react-router-dom';

const LandingPage: React.FC = () => {
  const navigate = useNavigate();

  const handleLogin = () => {
    // Redirect to home screen
    navigate('/home');
  };

  return (
    <div className="landing-page">
      <div className="login-widget">
        <h1 className="title">Wishlist</h1>
        <p className="subtitle">Your favorite app to manage and track wishlists!</p>
        <div className="form-container">
          <button onClick={handleLogin} className="login-button">Log In</button>
        </div>
      </div>
    </div>
  );
};

export default LandingPage;
