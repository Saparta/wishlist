import React, { useState, useEffect } from 'react';
import { FaTimes } from 'react-icons/fa';


interface User {
  id: string;
  email: string;
  name: string;
  createdAt: string;
}

interface UserModalProps {
  onClose: () => void;
}

const UserModal: React.FC<UserModalProps> = ({ onClose }) => {
  const [userData, setUserData] = useState<User | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    // Fetch user data from the backend
    fetch('http://localhost:8080/users', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error('Failed to fetch user data');
        }
        return response.json();
      })
      .then((data) => {
        // Parse data for first user's data in the list
        if (Array.isArray(data) && data.length > 0) {
          const user = data[0];


          setUserData(
            {
              id: user.id,
              email: user.email,
              name: user.name,
              createdAt: user.createdAt
            }
          );
        }
        else {
          throw new Error('Failed to parse user data or response is empty');
        }
      })
      .catch((err) => {
        setError(err.message);
      });
  }, []);

  if (error) {
    return (
      <div className="modal-overlay">
        <div className="modal">
          <div className="modal-header">
            <h2>Error</h2>
            <FaTimes className="close-icon" onClick={onClose} />
          </div>
          <div className="modal-content">
            <p>{error}</p>
          </div>
        </div>
      </div>
    );
  }

  if (!userData) {
    return (
      <div className="modal-overlay">
        <div className="modal">
          <div className="modal-header">
            <h2>Loading...</h2>
            <FaTimes className="close-icon" onClick={onClose} />
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="modal-overlay">
      <div className="modal">
        <div className="modal-header">
          <h2>User Details</h2>
          <FaTimes className="close-icon" onClick={onClose} />
        </div>
        <div className="modal-content">
          <p><strong>ID:</strong> {userData.id}</p>
          <p><strong>Email:</strong> {userData.email}</p>
          <p><strong>Name:</strong> {userData.name}</p>
        </div>
      </div>
    </div>
  );
};

export default UserModal;
