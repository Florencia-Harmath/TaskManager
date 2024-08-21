// Definimos el modelo TASK, que representa la tabla TASKS en la base de datos

package models

import (
	"taskmanager/internal/database"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title     string
	Description   string
	Completed bool
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	t.Completed = false
	return
}

//esta funcion realiza la migracion automatica de migracion despues de conectar a la base de atos
func MigrateTasks() {
	database.DB.AutoMigrate(&Task{})
}