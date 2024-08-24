
import styles from './Task.module.css';
import PropTypes from 'prop-types';

const Task = ({ task, onComplete, onDelete }) => {
  return (
    <div className={styles.taskContainer}>
      <h3 className={styles.taskTitle}>{task.title}</h3>
      <p className={styles.taskDescription}>{task.description}</p>
      <div className={styles.taskActions}>
        <button
          className={styles.completeButton}
          onClick={() => onComplete(task.id)}
        >
          {task.completed ? 'Undo' : 'Complete'}
        </button>
        <button
          className={styles.deleteButton}
          onClick={() => onDelete(task.id)}
        >
          Delete
        </button>
      </div>
    </div>
  );
};

Task.propTypes = {
    task: PropTypes.shape({
      id: PropTypes.number.isRequired,
      title: PropTypes.string.isRequired,
      description: PropTypes.string.isRequired,
      completed: PropTypes.bool.isRequired,
    }).isRequired,
    onComplete: PropTypes.func.isRequired,
    onDelete: PropTypes.func.isRequired,
  };

export default Task;
