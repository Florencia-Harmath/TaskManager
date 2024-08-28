import { useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { addTask } from "../../redux/tasksSlice";
import { useNavigate } from "react-router-dom";
import styles from "./CreateTasks.module.css";

const CreateTask = () => {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const dispatch = useDispatch();
  const navigate = useNavigate();

  const handleSubmit = (e) => {
    e.preventDefault();
    dispatch(addTask({ title, description }))
      .unwrap()
      .then(() => {
        navigate("/tasks");
      })
      .catch((err) => {
        console.error("Failed to create task:", err);
      });
  };

  return (
    <div className={styles.createTaskContainer}>
      <h2>Crear una nueva tarea</h2>
      <form onSubmit={handleSubmit} className={styles.createTaskForm}>
        <div className={styles.formGroup}>
          <label>Título</label>
          <input
            type="text"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            required
          />
        </div>
        <div className={styles.formGroup}>
          <label>Descripción</label>
          <input
            type="text"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            required
          />
        </div>
        <button type="submit" className={styles.submitButton}>
          Crear Tarea
        </button>
      </form>
    </div>
  );
};

export default CreateTask;
