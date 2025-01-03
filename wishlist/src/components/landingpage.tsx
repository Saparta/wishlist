import React from 'react';
import "./landingpage.css";
import { useNavigate } from 'react-router-dom';
import { FaGoogle } from 'react-icons/fa';

const LandingPage: React.FC = () => {
  const navigate = useNavigate();

  const handleLogin = () => {
    // Redirect to home screen for now
    navigate('/home');
  };

  return (
    <div className="landing-page">
      <div className="login-widget">
        <div className="logo-container">
          <img src="/logo.svg" alt="Wishlist Logo" className="logo" />
        </div>
        <h1 className="title">Wishlist</h1>
        <p className="subtitle">Your favorite app to manage and track wishlists!</p>
        <div className="form-container">
          <button onClick={handleLogin} className="login-button">
            <FaGoogle className="google-icon" />
            Log In with Google
          </button>
        </div>
      </div>
    </div>
  );
};

export default LandingPage;
