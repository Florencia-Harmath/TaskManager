import { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { fetchTasks, setTasks } from "../../redux/tasksSlice";
import Task from "../task/Task";
import { Link } from "react-router-dom";
import styles from "./ContainerTasks.module.css";

const ContainerTasks = () => {
  const dispatch = useDispatch();
  const {tasks, status, flag} = useSelector((store) => store.tasks);

  useEffect(() => {
    dispatch(fetchTasks())
      .then(response=>{dispatch(setTasks(response.payload))})
      .catch(e => console.log(e));
    }, [flag]);

  if (status != "succeeded" && status != "failed") {
    return <p>Loading tasks...</p>;
  }

  if (status === "succeeded" && tasks.length > 0) {
    return (
      <div className={styles.containerTasks}>
        <h1 className={styles.TasksTitle}>Mis tareas:</h1>
        <div className={styles.tasksDiv}>
          {tasks.map((task) => (
            <Task key={task.ID} task={task} />
          ))}
        </div>
        <Link to="/create-tasks" className={styles.createTaskButton}>
          Crear una tarea
        </Link>
      </div>
    );
  }

  return (
    <div className={styles.noTasks}>
      <h1 className={styles.TasksTitle}>Mis tareas:</h1>
      <p>Aún no tienes tareas.</p>
      <Link to="/create-tasks" className={styles.createTaskButton}>
        Crear una tarea
      </Link>
    </div>
  );

};

export default ContainerTasks;
