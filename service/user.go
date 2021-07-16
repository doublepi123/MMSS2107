package service

import (
	"MMSS2107/entity"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

func (s Service) WriteCookie(username string, key string, validtime time.Time, c *gin.Context) {
	s.DB.DB.Model(entity.Userid{}).Delete(&entity.Userid{}, username) // 删除旧的Cookie记录
	userid := entity.Userid{
		username,
		key,
		validtime,
	}
	s.DB.DB.Create(userid)
	b, _ := json.Marshal(userid)
	s.LRU.Set("userid_"+username, string(b))
	c.SetCookie("Username", username, int(time.Hour), "/", c.Request.Host, false, false)
	c.SetCookie("Userkey", key, int(time.Hour), "/", c.Request.Host, false, false)
}
