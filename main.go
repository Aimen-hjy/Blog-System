package main

import (
	"blogSystem/post"
	"blogSystem/user"
	"fmt"
	"log"
)

func main() {
	post.PostMgr.Init()
	defer post.PostMgr.CloseDatabase()
	user.UserMgr.Init()
	defer user.UserMgr.CloseDatabase()
	for {
		var op string
		fmt.Scan(&op)
		if op == "register" {

		} else if op == "login" {

		} else if op == "logout" {

		} else if op == "createPost" {

		} else if op == "updatePost" {

		} else if op == "deletePost" {

		} else if op == "viewPost" {

		} else if op == "listPost" {

		} else if op == "exit" {

		} else {
			log.Println("Error: unknown operation")
		}
	}
}
