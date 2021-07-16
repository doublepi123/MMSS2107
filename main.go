package main

import (
	"MMSS2107/LRUCache"
	"MMSS2107/entity"
	"MMSS2107/service"
	"MMSS2107/util"
	"time"
)

func main() {
	db := util.Database{}
	//连接数据库并初始化
	db.Connect()
	db.DB.AutoMigrate(entity.Paper{})
	db.DB.AutoMigrate(entity.Journal{})
	db.DB.AutoMigrate(entity.Teacher{})
	db.DB.AutoMigrate(entity.User{})
	db.DB.AutoMigrate(entity.Userid{})
	//初始化LRU缓存
	lru := LRUCache.LRU{Maxsize: 1000}
	lru.Init()
	//初始化服务
	sv := service.Service{
		DB:  db,
		LRU: lru,
	}
	//启动服务
	sv.Run()
	//保持主程序不退出
	for {
		time.Sleep(time.Hour)
	}
}
