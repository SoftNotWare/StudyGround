package controller

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"regexp"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/SoftNotWare/study-ground/model"
	"github.com/SoftNotWare/study-ground/model/table"

	"github.com/gin-gonic/gin"
)

const trimSpace = "\t\r\n "
const emailReg = `^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`
const passwordReg = `^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,16}$`

func sha256String(v string) string {
	h := sha256.New()
	h.Write([]byte(v))
	return hex.EncodeToString(h.Sum(nil))
}

func makePassword(password string, sha bool) string {
	h := hmac.New(sha256.New, model.Secret)
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}

// Register 处理注册请求
func Register(ctx *gin.Context) {
	var p struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"pwd"`
	}
	err := ctx.ShouldBindJSON(&p)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		log.WithFields(log.Fields{"pwd": p.Password, "email": p.Email}).Debug("not found p")
		return
	}

	p.Email = strings.Trim(p.Email, trimSpace)
	p.Password = strings.Trim(p.Password, trimSpace)
	matched, err := regexp.MatchString(emailReg, p.Email)
	if err != nil || !matched {
		ctx.JSON(http.StatusBadRequest, nil)
		log.WithFields(log.Fields{"email": p.Email, "error": err}).Debug("email format wrong")
		return
	}

	if len(p.Password) != 64 {
		ctx.JSON(http.StatusBadRequest, nil)
		log.WithFields(log.Fields{"password": p.Password}).Debug("password format wrong")
		return
	}

	now := time.Now()
	user := model.QueryUserWithPassword(p.Email, makePassword(p.Password, false))
	if user == nil {
		user = &table.User{
			Email:     p.Email,
			Password:  makePassword(p.Password, false),
			CreatedAt: now,
			LoginAt:   now,
		}

		// 注册
		if !model.CreateUser(user) {
			ctx.JSON(http.StatusUnauthorized, nil)
			log.WithFields(log.Fields{"email": p.Email, "password": p.Password}).Info("create user failed")
			return
		}

		log.WithFields(log.Fields{"user": user}).Debug("register user succeed")
	} else {
		// 更新登录时间
		user.LoginAt = now
		if !model.Update(user) {
			ctx.JSON(http.StatusUnauthorized, nil)
			log.WithFields(log.Fields{"email": p.Email}).Info("update user failed")
			return
		}
	}

	type userInfo struct {
		Score     uint      `json:"score" `
		Avatar    string    `json:"avatar"`
		LoginAt   time.Time `json:"login_at"`
		CreatedAt time.Time ` json:"created_at"`
	}
	// 登录成功
	ctx.JSON(http.StatusOK, userInfo{
		Score:     user.Score,
		Avatar:    user.Avatar.String,
		LoginAt:   user.LoginAt,
		CreatedAt: user.CreatedAt,
	})
}
