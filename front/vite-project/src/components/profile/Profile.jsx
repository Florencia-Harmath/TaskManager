import { useState, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import styles from './profile.module.css';
import { fetchUserProfile, updateUserProfile } from '../../redux/authSlice';

const Profile = () => {
  const dispatch = useDispatch();
  const { user, status, error } = useSelector((state) => state.auth);
  const [newName, setNewName] = useState('');

  useEffect(() => {
    if (status === 'idle') {
      dispatch(fetchUserProfile());
    }
  }, [dispatch, status]);

  const handleUpdate = () => {
    dispatch(updateUserProfile(newName));
  };

  return (
    <div className={styles.profileContainer}>
      {status === 'loading' && <p>Loading...</p>}
      {status === 'failed' && <p className={styles.errorMessage}>{error}</p>}
      {status === 'succeeded' && user && (
        <div className={styles.profileForm}>
          <h1 className={styles.profileTitle}>Profile</h1>
          <div className={styles.formGroup}>
            <p><strong>Name:</strong> {user.name}</p>
            <p><strong>Email:</strong> {user.email}</p>
          </div>
          <div className={styles.formGroup}>
            <label htmlFor="newName">New Name:</label>
            <input
              id="newName"
              type="text"
              value={newName}
              onChange={(e) => setNewName(e.target.value)}
              placeholder="Enter new name"
            />
          </div>
          <button className={styles.saveButton} onClick={handleUpdate}>
            Update Name
          </button>
        </div>
      )}
    </div>
  );
};

export default Profile;
