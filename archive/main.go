package main

import (
	"github.com/yby/todo_list/models"
	"github.com/yby/todo_list/router"
)

// main 程序入口
func main() {
	models.Setup()
	router.Setup()
}
