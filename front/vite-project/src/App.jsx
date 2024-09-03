  import { useState, useEffect } from 'react';
import { Route, Routes } from 'react-router-dom';
import NavBar from './components/navbar/NavBar';
import Home from './components/home/Home';
import Login from './components/login/Login';
import Register from './components/register/Register';
import ContainerTasks from './components/containerTasks/containerTasks';
import CreateTask from './components/createTasks/CreateTasks';
import Profile from './components/profile/Profile';
import { useDispatch, useSelector } from 'react-redux';
import { login } from './redux/authSlice';

const App = () => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const { token } = useSelector(store=>store.auth);
  const dispatcher = useDispatch();

  useEffect(() => {
    const checkAuth = () => {
      const authToken = localStorage.getItem("isAuthenticated");
      if (authToken) {
        setIsAuthenticated(true);
        dispatcher(login(authToken));
      };
    };

    checkAuth();
  }, []);

  useEffect(()=> {
    if (token) {
      setIsAuthenticated(true);
    };
  }, [token]);

  return (
    <>
      {isAuthenticated && <NavBar />}
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/tasks" element={<ContainerTasks />} />
        <Route path="/create-tasks" element={<CreateTask />} />
        <Route path="/profile" element={<Profile />} />
      </Routes>
    </>
  );
};

export default App;
