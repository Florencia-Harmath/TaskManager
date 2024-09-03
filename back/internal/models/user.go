package models

import (
	"log"
	"taskmanager/internal/database"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterUser struct {
    ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    Name      string    `json:"name"`
    Email     string    `json:"email" gorm:"unique"`
    Password  string    `json:"-"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type LoginUser struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type UpdateUserRequest struct {
    Name string `json:"name"`
}

// HashPassword encripta la contraseña del usuario
func (u *RegisterUser) HashPassword() error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    log.Println("Hashed Password:", string(hashedPassword))
    u.Password = string(hashedPassword)
    return nil
}

// CheckPassword verifica la contraseña ingresada contra la almacenada
func (u *RegisterUser) CheckPassword(password string) bool {
    log.Println("Stored Password Hash:", u.Password)
    log.Println("Password for Comparison:", password)
    err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
    if err != nil {
        log.Println("Error comparing password:", err)
    }
    return err == nil
}

// BeforeCreate es un hook de GORM que se ejecuta antes de crear un registro en la BD
func (u *RegisterUser) BeforeCreate(tx *gorm.DB) (err error) {
    u.ID = uuid.New()
    return
}

// MigrateUsers realiza la migración automática para el modelo User
func MigrateUsers() {
    database.DB.AutoMigrate(&RegisterUser{})
}
