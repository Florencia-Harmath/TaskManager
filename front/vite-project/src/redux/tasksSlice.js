// src/redux/tasksSlice.js
import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';

// Obtener todas las tareas
export const fetchTasks = createAsyncThunk(
  'tasks/fetchTasks',
  async (_, { getState }) => {
    const { auth } = getState();
    const response = await axios.get('/api/tasks', {
      headers: {
        Authorization: `Bearer ${auth.token}`,
      },
    });
    return response.data; // Debe ser un objeto serializable
  }
);

// Obtener tarea por ID
export const fetchTaskById = createAsyncThunk(
  'tasks/fetchTaskById',
  async (taskId, { getState }) => {
    const { auth } = getState();
    const response = await axios.get(`/api/tasks/${taskId}`, {
      headers: {
        Authorization: `Bearer ${auth.token}`,
      },
    });
    return response.data; // Debe ser un objeto serializable
  }
);

// Crear una tarea
export const addTask = createAsyncThunk(
  'tasks/addTask',
  async (taskData, { getState }) => {
    const { auth } = getState();
    const response = await axios.post('/api/createTasks', taskData, {
      headers: {
        Authorization: `Bearer ${auth.token}`,
      },
    });
    return response.data; // Debe ser un objeto serializable
  }
);

// Editar una tarea por ID
export const editTask = createAsyncThunk(
  'tasks/editTask',
  async ({ taskId, updatedTaskData }, { getState }) => {
    const { auth } = getState();
    const response = await axios.put(`/api/tasks/${taskId}`, updatedTaskData, {
      headers: {
        Authorization: `Bearer ${auth.token}`,
      },
    });
    return response.data; // Debe ser un objeto serializable
  }
);

// Borrar una tarea por ID
export const deleteTask = createAsyncThunk(
  'tasks/deleteTask',
  async (taskId, { getState }) => {
    const { auth } = getState();
    await axios.delete(`/api/tasks/${taskId}`, {
      headers: {
        Authorization: `Bearer ${auth.token}`,
      },
    });
    return taskId; // Debe ser un valor serializable
  }
);

// Cambiar estado de tarea a completo
export const completeTask = createAsyncThunk(
  'tasks/completeTask',
  async (taskId, { getState }) => {
    const { auth } = getState();
    const response = await axios.put(
      `/api/tasks/${taskId}/complete`,
      {},
      {
        headers: {
          Authorization: `Bearer ${auth.token}`,
        },
      }
    );
    return response.data; // Debe ser un objeto serializable
  }
);

// Cambiar estado de tarea a incompleto
export const uncompleteTask = createAsyncThunk(
  'tasks/uncompleteTask',
  async (taskId, { getState }) => {
    const { auth } = getState();
    const response = await axios.put(
      `/api/tasks/${taskId}/uncomplete`,
      {},
      {
        headers: {
          Authorization: `Bearer ${auth.token}`,
        },
      }
    );
    return response.data; // Debe ser un objeto serializable
  }
);

const tasksSlice = createSlice({
  name: 'tasks',
  initialState: [],
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchTasks.fulfilled, (state, action) => {
        return action.payload; // Asegúrate de que `action.payload` sea un array serializable
      })
      .addCase(fetchTaskById.fulfilled, (state, action) => {
        const index = state.findIndex((task) => task.id === action.payload.id);
        if (index !== -1) {
          state[index] = action.payload; // Asegúrate de que `action.payload` sea serializable
        } else {
          state.push(action.payload);
        }
      })
      .addCase(addTask.fulfilled, (state, action) => {
        state.push(action.payload); // Asegúrate de que `action.payload` sea serializable
      })
      .addCase(editTask.fulfilled, (state, action) => {
        const index = state.findIndex((task) => task.id === action.payload.id);
        if (index !== -1) {
          state[index] = action.payload; // Asegúrate de que `action.payload` sea serializable
        }
      })
      .addCase(deleteTask.fulfilled, (state, action) => {
        return state.filter((task) => task.id !== action.payload); // Asegúrate de que `action.payload` sea serializable
      })
      .addCase(completeTask.fulfilled, (state, action) => {
        const index = state.findIndex((task) => task.id === action.payload.id);
        if (index !== -1) {
          state[index] = action.payload; // Asegúrate de que `action.payload` sea serializable
        }
      })
      .addCase(uncompleteTask.fulfilled, (state, action) => {
        const index = state.findIndex((task) => task.id === action.payload.id);
        if (index !== -1) {
          state[index] = action.payload; // Asegúrate de que `action.payload` sea serializable
        }
      });
  },
});

export default tasksSlice.reducer;
