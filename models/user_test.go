package tests

import (
    "testing"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
    db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    if err != nil {
        t.Fatalf("Failed to connect database: %v", err)
    }
    db.AutoMigrate(&User{})

    user := User{Username: "testuser", Password: "testpassword"}
    result := db.Create(&user)

    assert.Nil(t, result.Error)
    assert.Equal(t, int64(1), result.RowsAffected)
}
