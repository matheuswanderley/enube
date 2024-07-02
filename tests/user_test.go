package tests

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/utils/tests"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect database: %v", err)
	}
	db.AutoMigrate(&tests.User{})

	user := tests.User{Username: "testuser", Password: "testpassword"}
	result := db.Create(&user)

	assert.Nil(t, result.Error)
	assert.Equal(t, int64(1), result.RowsAffected)
}
