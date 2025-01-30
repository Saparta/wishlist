import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import LandingPage from "./components/landingpage";
import Home from "./components/home";
import AddWish from "./components/addwish";
const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<LandingPage />} />
        <Route path="/home" element={<Home />} />
        <Route path="/addwish" element={<AddWish />} />
      </Routes>
    </Router>
  );
};

export default App;
