import { Link } from 'react-router-dom';
import styles from './NavBar.module.css';
import PropTypes from 'prop-types';

const NavBar = ({ isAuthenticated }) => {
  return (
    <>
      {isAuthenticated && (
        <nav className={styles.navBar}>
          <ul className={styles.navList}>
            <li className={styles.navItem}>
              <Link to="/tasks" className={styles.navLink}>Tareas</Link>
            </li>
            <li className={styles.navItem}>
              <Link to="/profile" className={styles.navLink}>Perfil</Link>
            </li>
          </ul>
        </nav>
      )}
    </>
  );
};

NavBar.propTypes = {
    isAuthenticated: PropTypes.bool.isRequired,
  };

export default NavBar;
