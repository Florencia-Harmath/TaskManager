// src/redux/tasksSlice.js
import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';

// Crea una instancia de Axios con la configuración base
const axiosInstance = axios.create({
  baseURL: 'http://localhost:3000/api', // Cambia al puerto correcto del backend
});

// Obtener todas las tareas
export const fetchTasks = createAsyncThunk(
  'tasks/fetchTasks',
  async (_, { getState, rejectWithValue }) => {
    try {
      const { auth } = getState();
      const response = await axiosInstance.get('/tasks', {
        headers: {
          Authorization: `Bearer ${auth.token}`,
        },
      });
      return response.data; // Debe ser un array de tareas
    } catch (error) {
      return rejectWithValue(error.response ? error.response.data : error.message);
    }
  }
);

// Obtener tarea por ID
export const fetchTaskById = createAsyncThunk(
  'tasks/fetchTaskById',
  async (taskId, { getState, rejectWithValue }) => {
    try {
      const { auth } = getState();
      const response = await axiosInstance.get(`/tasks/${taskId}`, {
        headers: {
          Authorization: `Bearer ${auth.token}`,
        },
      });
      return response.data; // Debe ser un objeto de tarea
    } catch (error) {
      return rejectWithValue(error.response ? error.response.data : error.message);
    }
  }
);

// Crear una tarea
export const addTask = createAsyncThunk(
  'tasks/addTask',
  async (taskData, { getState, rejectWithValue }) => {
    try {
      const { auth } = getState();
      const response = await axiosInstance.post('/tasks', taskData, {
        headers: {
          Authorization: `Bearer ${auth.token}`,
        },
      });
      return response.data; // Debe ser un objeto de tarea
    } catch (error) {
      return rejectWithValue(error.response ? error.response.data : error.message);
    }
  }
);

// Editar una tarea por ID
export const editTask = createAsyncThunk(
  'tasks/editTask',
  async ({ taskId, updatedTaskData }, { getState, rejectWithValue }) => {
    try {
      const { auth } = getState();
      const response = await axiosInstance.put(`/tasks/${taskId}`, updatedTaskData, {
        headers: {
          Authorization: `Bearer ${auth.token}`,
        },
      });
      return response.data; // Debe ser un objeto de tarea
    } catch (error) {
      return rejectWithValue(error.response ? error.response.data : error.message);
    }
  }
);

// Borrar una tarea por ID
export const deleteTask = createAsyncThunk(
  'tasks/deleteTask',
  async (taskId, { getState, rejectWithValue }) => {
    try {
      const { auth, flag } = getState();
      await axiosInstance.delete(`/tasks/${taskId}`, {
        headers: {
          Authorization: `Bearer ${auth.token}`,
        },
      });
      return taskId; // Debe ser un valor serializable
    } catch (error) {
      return rejectWithValue(error.response ? error.response.data : error.message);
    }
  }
);

// Cambiar estado de tarea a completo
export const completeTask = createAsyncThunk(
  'tasks/completeTask',
  async (taskId, { getState, rejectWithValue }) => {
    try {
      const { auth } = getState();
      const response = await axiosInstance.put(
        `/tasks/${taskId}/complete`,
        {},
        {
          headers: {
            Authorization: `Bearer ${auth.token}`,
          },
        }
      );
      return response.data; // Debe ser un objeto de tarea
    } catch (error) {
      return rejectWithValue(error.response ? error.response.data : error.message);
    }
  }
);

// Cambiar estado de tarea a incompleto
export const incompleteTask = createAsyncThunk(
  'tasks/uncompleteTask',
  async (taskId, { getState, rejectWithValue }) => {
    try {
      const { auth } = getState();
      const response = await axiosInstance.put(
        `/tasks/${taskId}/incomplete`,
        {},
        {
          headers: {
            Authorization: `Bearer ${auth.token}`,
          },
        }
      );
      return response.data; // Debe ser un objeto de tarea
    } catch (error) {
      return rejectWithValue(error.response ? error.response.data : error.message);
    }
  }
);

const tasksSlice = createSlice({
  name: 'tasks',
  initialState: {
    tasks: [], // Asegúrate de que `tasks` sea un array
    status: 'idle', // Estado de carga: 'idle', 'loading', 'succeeded', 'failed'
    error: null, // Para manejar errores
    flag: false, //Para actualizar la lista
  },
  reducers: {
    setTasks: (state, action) => {
      state.tasks = action.payload;
    }
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchTasks.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(fetchTasks.fulfilled, (state, action) => {
        state.status = 'succeeded';
        state.tasks = action.payload; // Asegúrate de que `action.payload` sea un array
      })
      .addCase(fetchTasks.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.payload;
      })
      .addCase(fetchTaskById.fulfilled, (state, action) => {
        const index = state.tasks.findIndex((task) => task.id === action.payload.id);
        if (index !== -1) {
          state.tasks[index] = action.payload;
        } else {
          state.tasks.push(action.payload);
        }
      })
      .addCase(addTask.fulfilled, (state, action) => {
        state.tasks.push(action.payload); // Asegúrate de que `action.payload` sea un objeto de tarea
      })
      .addCase(editTask.fulfilled, (state, action) => {
        const index = state.tasks.findIndex((task) => task.id === action.payload.id);
        if (index !== -1) {
          state.tasks[index] = action.payload; // Asegúrate de que `action.payload` sea un objeto de tarea
        }
      })
      .addCase(deleteTask.fulfilled, (state, action) => {
        state.flag = !state.flag;
        state.tasks = state.tasks.filter((task) => task.id !== action.payload);
      })
      .addCase(completeTask.fulfilled, (state, action) => {
        const index = state.tasks.findIndex((task) => task.ID === action.payload.ID);
        if (index !== -1) {
          state.tasks[index] = action.payload; // Asegúrate de que `action.payload` sea un objeto de tarea
        }
      })
      .addCase(incompleteTask.fulfilled, (state, action) => {
        const index = state.tasks.findIndex((task) => task.ID === action.payload.ID);
        if (index !== -1) {
          state.tasks[index] = action.payload; // Asegúrate de que `action.payload` sea un objeto de tarea
        }
      });
  },
});

export const {setTasks} = tasksSlice.actions;

export default tasksSlice.reducer;
