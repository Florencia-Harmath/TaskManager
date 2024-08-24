/* eslint-disable react/no-unescaped-entities */
import { useNavigate } from 'react-router-dom';
import styles from './Home.module.css';

const Home = () => {
  const navigate = useNavigate();

  return (
    <div className={styles.homeContainer}>
      <h1 className={styles.homeTitle}>Bienvenido a TaskManager</h1>
      <p className={styles.homeText}>
        En esta aplicación podrás organizar tus tareas! Podes crear, leer, actualizar y eliminar las tareas, y vas a poder cambiarle el estado de las mismas, "completadas" o "pendientes". <br />
        La Aplicacion esta creada con React y Vite del lado del front, y con Golang en el lado del back.
      </p>
      <div className={styles.buttonContainer}>
        <button onClick={() => navigate('/login')} className={styles.homeButton}>
          Iniciar Sesión
        </button>
        <button onClick={() => navigate('/register')} className={styles.homeButton}>
          Registrarse
        </button>
      </div>
      <img src="/src/assets/trabajando.jpeg" alt="" className={styles.homeImage} />
    </div>
  );
};

export default Home;
