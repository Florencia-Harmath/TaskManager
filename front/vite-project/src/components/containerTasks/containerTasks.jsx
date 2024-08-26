// src/components/containerTasks/ContainerTasks.jsx
import { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { fetchTasks } from "../../redux/tasksSlice";
import Task from "../task/Task";
import { Link } from "react-router-dom";
import styles from "./ContainerTasks.module.css";

const ContainerTasks = () => {
  const dispatch = useDispatch();
  const tasks = useSelector((state) => state.tasks);
  const status = useSelector((state) => state.status);

  useEffect(() => {
    dispatch(fetchTasks());
  }, [dispatch]);

  if (status === "loading") {
    return <p>Loading tasks...</p>;
  }

  // Verifica que tasks sea un array
  const tasksArray = Array.isArray(tasks) ? tasks : [];

  if (tasksArray.length === 0) {
    return (
      <div className={styles.noTasks}>
        <h1 className={styles.TasksTitle}>Mis tareas:</h1>
        <p>Aún no tienes tareas.</p>
        <Link to="/create-tasks" className={styles.createTaskButton}>
          Crear una tarea
        </Link>
      </div>
    );
  }

  return (
    <div className={styles.taskContainer}>
      <h1 className={styles.TasksTitle}>Mis tareas:</h1>

      {tasksArray.map((task) => (
        <Task key={task.id} task={task} />
      ))}

      <Link to="/create-tasks" className={styles.createTaskButton}>
        Crear una tarea
      </Link>
    </div>
  );
};

export default ContainerTasks;
