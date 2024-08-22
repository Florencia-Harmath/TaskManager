# TaskManager API

TaskManager es una API RESTful para gestionar tareas, desarrollada en Go. Permite a los usuarios registrarse, autenticarse y manejar sus tareas de manera eficiente.

## Características

- **Autenticación**: Registro de usuarios, inicio de sesión y manejo de sesiones mediante tokens JWT.
- **Gestión de Tareas**: CRUD completo (Crear, Leer, Actualizar, Eliminar) para tareas.
- **Protección de Rutas**: Solo los usuarios autenticados pueden acceder a las rutas de gestión de tareas.

## Requisitos

- Go 1.18 o superior
- PostgreSQL
- Librerías:
  - `github.com/gorilla/mux`
  - `github.com/joho/godotenv`
  - `golang.org/x/crypto/bcrypt`
  - `gorm.io/gorm`
  - `gorm.io/driver/postgres`

## Instalación

1. Clona este repositorio:

    ```sh
    git clone https://github.com/tu-usuario/taskmanager.git
    cd taskmanager
    ```

2. Crea un archivo `.env` en la raíz del proyecto con las siguientes variables de entorno:

    ```plaintext
    DB_HOST=localhost
    DB_USER=tu-usuario
    DB_PASSWORD=tu-contraseña
    DB_NAME=taskmanager
    DB_PORT=5432
    JWT_SECRET_KEY=tu-secreto
    ```

3. Instala las dependencias:

    ```sh
    go mod tidy
    ```

4. Inicia el servidor:

    ```sh
    go run cmd/main.go
    ```

## Endpoints

### Autenticación

- **Registro de usuario**
  - `POST /register`
  - Cuerpo de la solicitud:
    ```json
    {
      "name": "John Doe",
      "email": "john.doe@example.com",
      "password": "your-password"
    }
    ```

- **Inicio de sesión**
  - `POST /login`
  - Cuerpo de la solicitud:
    ```json
    {
      "email": "john.doe@example.com",
      "password": "your-password"
    }
    ```

### Gestión de Tareas

- **Obtener todas las tareas**
  - `GET /tasks`

- **Obtener una tarea por ID**
  - `GET /tasks/{id}`

- **Crear una nueva tarea**
  - `POST /createTasks`
  - Cuerpo de la solicitud:
    ```json
    {
      "title": "New Task",
      "description": "Task description"
    }
    ```

- **Actualizar una tarea**
  - `PUT /tasks/{id}`
  - Cuerpo de la solicitud:
    ```json
    {
      "title": "Updated Task",
      "description": "Updated description"
    }
    ```

- **Eliminar una tarea**
  - `DELETE /tasks/{id}`

- **Completar una tarea**
  - `PUT /tasks/{id}/complete`

## Middleware de Autenticación

Para proteger las rutas, se utiliza un middleware que verifica la presencia y validez del token JWT en las solicitudes. Si el token no es válido o no está presente, la ruta devuelve un error 401 Unauthorized.

## Contribuciones

¡Las contribuciones son bienvenidas! Por favor, abre un issue o un pull request para discutir cualquier cambio.

