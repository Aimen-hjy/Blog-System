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
			var username, password string
			fmt.Scan(&username, &password)
			user.UserMgr.Register(username, password)
		} else if op == "login" {
			var username, password string
			fmt.Scan(&username, &password)
			user.UserMgr.Login(username, password)
		} else if op == "logout" {
			user.UserMgr.Logout()
		} else if op == "changePassword" {
			var oldPassword, newPassword string
			fmt.Scan(&oldPassword, &newPassword)
			user.UserMgr.ChangePassword(oldPassword, newPassword)
		} else if op == "createPost" {

		} else if op == "updatePost" {

		} else if op == "deletePost" {

		} else if op == "viewPost" {

		} else if op == "listPost" {

		} else if op == "exit" {
			break
		} else {
			log.Println("Error: unknown operation")
		}
	}
	return
}
