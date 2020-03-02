package model

import (
	"context"
	"database/sql"
	"errors"

	"github.com/SoftNotWare/study-ground/model/table"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
)

// QueryUserByEmail 通过email查询用户
func QueryUserByEmail(email string) *table.User {
	user, err := table.Users(
		table.UserWhere.Email.EQ(email),
	).One(context.Background(), db)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.WithFields(log.Fields{"email": email, "error": err}).Debug("cannot load user")
		return nil
	}

	log.WithFields(log.Fields{"email": email, "user": user}).Debug("query user by email")
	return user
}

// QueryUserWithPassword 查询用户
func QueryUserWithPassword(email string, password string) *table.User {
	user, err := table.Users(
		table.UserWhere.Email.EQ(email),
		table.UserWhere.Password.EQ(password),
	).One(context.Background(), db)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.WithFields(log.Fields{"email": email, "error": err}).Debug("cannot load user")
		return nil
	}

	log.WithFields(log.Fields{"email": email, "password": password, "user": user}).Debug("query user by with password")
	return user
}

// CreateUser 创建用户
func CreateUser(user *table.User) bool {
	err := user.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		log.WithFields(log.Fields{"id": user.ID, "user": user, "error": err}).Debug("cannot insert user")
		return false
	}
	return true
}

func Update(user *table.User) bool {
	return true
}
