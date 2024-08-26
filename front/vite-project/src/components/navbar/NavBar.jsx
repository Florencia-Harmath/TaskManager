// src/components/navbar/NavBar.jsx
import { Link } from 'react-router-dom';
import styles from './NavBar.module.css';

const NavBar = () => {
  return (
    <>
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
    </>
  );
};


export default NavBar;
