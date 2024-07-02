package db

import (
    "log"
    "time"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
    ID        uint      `gorm:"primaryKey"`
    Username  string    `gorm:"unique"`
    Password  string
    CreatedAt time.Time
    UpdatedAt time.Time
}

func InitDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database:", err)
    }

    // Migrate the schema
    err = DB.AutoMigrate(&User{})
    if err != nil {
        log.Fatal("failed to migrate database:", err)
    }
}

func CreateUser(username, password string) error {
    user := User{Username: username, Password: password}
    result := DB.Create(&user)
    return result.Error
}

func GetUser(username string) (User, error) {
    var user User
    result := DB.Where("username = ?", username).First(&user)
    return user, result.Error
}
