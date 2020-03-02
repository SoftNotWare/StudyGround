package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/SoftNotWare/study-ground/model/table"
)

func TestCreateUser(t *testing.T) {
	InitDB()

	now := time.Now()
	user := &table.User{
		Email:    "test@test.com",
		Password: "123456",
		LoginAt:  now,
	}
	ok := CreateUser(user)
	t.Log(user)
	assert.True(t, ok)
}
