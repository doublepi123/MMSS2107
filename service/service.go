package service

import (
	"MMSS2107/LRUCache"
	"MMSS2107/entity"
	"MMSS2107/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Service struct {
	DB  util.Database
	LRU LRUCache.LRU
}

func (s Service) Run() {
	r := gin.Default()
	//设置404页面
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 Not Found")
	})
	//登录 /api/login POST
	r.POST("/api/login", func(c *gin.Context) {
		var user entity.User
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, fmt.Sprint(err))
			return
		}
		p := s.LRU.Get(user.Username)
		tu := entity.User{}
		if p == "" {
			err := s.DB.DB.Model(&entity.User{}).Where(&tu, user.Username)
			if err != nil {
				c.String(http.StatusForbidden, "username not found")
				return
			}
			if util.CmpPWD(tu.Password, user.Password) {
				s.WriteCookie(user.Username, util.GetPWD(user.Password), time.Now().Add(time.Hour), c)
				c.String(http.StatusOK, "success")
				return
			} else {
				c.String(http.StatusForbidden, "password wrong")
				return
			}
		}
	})
	//中间件，检查登入状态 /api
	api := r.Group("/api", func(c *gin.Context) {
		username, err := c.Cookie("Username")
		if err != nil {
			c.String(http.StatusForbidden, "not login")
			c.Abort()
			return
		}
		key, err := c.Cookie("Userkey")
		if err != nil {
			c.String(http.StatusForbidden, "not login")
			c.Abort()
			return
		}
		userid := entity.Userid{}
		t := s.LRU.Get("userid_" + username)
		if t != "" {
			err = json.Unmarshal([]byte(t), userid)
		} else {
			s.DB.DB.Find(userid, username)
		}
		if userid.Userkey == key {
			if userid.Validtime.After(time.Now()) {
				return
			}
		}
		c.String(http.StatusForbidden, "not login")
		c.Abort()
	})
	{
		//登出 /api/logout GET
		api.GET("/logout", func(c *gin.Context) {
			//获取当前用户名
			username, _ := c.Cookie("username")
			//从数据库中移除登入状态
			s.DB.DB.Delete(&entity.Userid{}, username)
			//从缓存中清除登入状态
			s.LRU.Remove(username)

			c.String(http.StatusOK, "success")
		})
	}
	r.Run(":8080")
}
