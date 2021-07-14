package main

import (
	"MMSS2107/entity"
	"MMSS2107/util"
	"container/heap"
	"github.com/doublepi123/cachemap/priorityqueue"
)

func main() {
	//数据库初始化
	db := util.Database{}
	db.Connect()
	db.DB.AutoMigrate(entity.Paper{})
	db.DB.AutoMigrate(entity.Journal{})
	db.DB.AutoMigrate(entity.Teacher{})
	pq := make(priorityqueue.PQueue,10)
	heap.Init(&pq)
	it := priorityqueue.Item{
		10,
		"abc",
	}
	heap.Push(&pq,it)
	s := heap.Pop(&pq)
	print(s.(priorityqueue.Item).Value.(string))
}