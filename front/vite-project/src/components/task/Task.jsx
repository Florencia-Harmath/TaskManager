import styles from './Task.module.css';
import { useDispatch } from 'react-redux';
import { completeTask, deleteTask, incompleteTask } from '../../redux/tasksSlice';

const Task = ({task}) => {
  const dispatcher = useDispatch();

  const onComplete = () => {
    if (task.Completed) {
      dispatcher(incompleteTask(task.ID));
    } else {
      dispatcher(completeTask(task.ID));
    };
  };

  const onDelete = () => {
    dispatcher(deleteTask(task.ID))
  };

  return (
    <div className={styles.taskContainer}>
      <h3 className={styles.taskTitle}>{task.Title}</h3>
      <p className={styles.taskDescription}>{task.Description}</p>
      <div className={styles.taskActions}>
        <button
          className={task.Completed ? styles.undoButton : styles.completeButton}
          onClick={() => onComplete()}
        >
          {task.Completed ? 'Deshacer' : 'Completa'}
        </button>
        <button
          className={styles.deleteButton}
          onClick={() => onDelete()}
        >
          Eliminar
        </button>
      </div>
    </div>
  );
};

export default Task;
