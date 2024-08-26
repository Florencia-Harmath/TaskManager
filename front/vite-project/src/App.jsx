// src/App.js
import { useState, useEffect } from 'react';
import { Route, Routes } from 'react-router-dom';
import NavBar from './components/navbar/NavBar';
import Home from './components/home/Home';
import Login from './components/login/Login';
import Register from './components/register/Register';
import ContainerTasks from './components/containerTasks/containerTasks';
import CreateTask from './components/createTasks/CreateTasks';

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
        <Route path="/login" element={ <Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/tasks" element={<ContainerTasks />} />
        <Route path="/create-task" element={<CreateTask /> } />
      </Routes>
    </>
  );
};

export default App;
