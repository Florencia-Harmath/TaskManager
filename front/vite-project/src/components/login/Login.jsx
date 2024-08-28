import { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { loginUser } from '../../redux/authSlice';
import { Link, useNavigate } from 'react-router-dom';
import styles from './Login.module.css';

const Login = () => {
  const [formData, setFormData] = useState({ email: '', password: '' });
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const { status, error } = useSelector((store)=>store.auth);

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    dispatch(loginUser(formData))
      .unwrap()
      .then((result) => {
        if (result.token) {
          navigate('/tasks');
        }
      })
      .catch((err) => {
        console.error('Failed to login:', err);
      });
  };

  return (
    <div className={styles.loginContainer}>
      <h2 className={styles.loginTitle}>Iniciá sesión:</h2>
      <form onSubmit={handleSubmit} className={styles.loginForm}>
        <div className={styles.formGroup}>
          <label>Email</label>
          <input
            type="email"
            name="email"
            value={formData.email}
            onChange={handleChange}
            required
          />
        </div>
        <div className={styles.formGroup}>
          <label>Password</label>
          <input
            type="password"
            name="password"
            value={formData.password}
            onChange={handleChange}
            required
          />
        </div>
        <button type="submit" disabled={status === 'loading'} className={styles.loginButton}>
          {status === 'loading' ? 'Logging in...' : 'Login'}
        </button>
      </form>
      {status === 'failed' && <p className={styles.errorMessage}>{error}</p>}
      <p className={styles.redirectText}>
        ¿No tienes cuenta? <Link to="/register" className={styles.redirectLink}>Regístrate</Link>
      </p>
      <img src="/src/assets/inicioregistro.jpeg" alt="" className={styles.loginImage} />
    </div>
  );
};

export default Login;
