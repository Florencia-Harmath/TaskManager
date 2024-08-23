import { useState, useEffect } from "react";
import { Route, Routes } from "react-router-dom";
import NavBar from "./components/navbar/NavBar";
import Home from "./components/home/Home";
import Login from "./components/login/Login";
import Register from "./components/register/Register";
// import Tasks from "./Tasks";
// import Profile from "./Profile";

const App = () => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  useEffect(() => {
    const checkAuth = () => {
      const authStatus = localStorage.getItem("isAuthenticated") === "true";
      setIsAuthenticated(authStatus);
    };

    checkAuth();
  }, []);

  return (
    <>
      <NavBar isAuthenticated={isAuthenticated} />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        {/* <Route
          path="/tasks"
          element={isAuthenticated ? <Tasks /> : <Login />}
        />
        <Route
          path="/profile"
          element={isAuthenticated ? <Profile /> : <Login />}
        /> */}
      </Routes>
    </>
  );
};

export default App;
