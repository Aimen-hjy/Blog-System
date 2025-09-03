package main

import (
	"blogSystem/post"
	"blogSystem/user"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
			var title, content string
			fmt.Scan(&title, &content)
			post.PostMgr.CreatePost(title, content)
		} else if op == "updatePost" {
			reader := bufio.NewReader(os.Stdin)
			line, _ := reader.ReadString('\n')
			fields := strings.Fields(line)
			if len(fields) == 3 {
				id, err := strconv.Atoi(fields[0])
				if err != nil {
					fmt.Println("Error: invalid post ID")
					continue
				}
				arg := fields[1]
				if arg == "-t" {
					newTitle := fields[2]
					post.PostMgr.UpdatePost(int64(id), newTitle, "")
				} else if arg == "-c" {
					newContent := fields[2]
					post.PostMgr.UpdatePost(int64(id), "", newContent)
				} else {
					fmt.Println("Error: unknown argument")
					continue
				}
			} else if len(fields) == 5 {
				id, err := strconv.Atoi(fields[0])
				if err != nil {
					fmt.Println("Error: invalid post ID")
					continue
				}
				arg1 := fields[1]
				arg2 := fields[3]
				if arg1 == "-t" && arg2 == "-c" {
					newTitle := fields[2]
					newContent := fields[4]
					post.PostMgr.UpdatePost(int64(id), newTitle, newContent)
				} else if arg1 == "-c" && arg2 == "-t" {
					newContent := fields[2]
					newTitle := fields[4]
					post.PostMgr.UpdatePost(int64(id), newTitle, newContent)
				} else {
					log.Println("Error: unknown arguments")
					continue
				}
			} else {
				log.Println("Error: invalid number of arguments")
				continue
			}
		} else if op == "deletePost" {
			var id int64
			fmt.Scan(&id)
			post.PostMgr.DeletePost(id)
		} else if op == "viewPost" {
			var id int64
			post.PostMgr.ViewPost(id)
		} else if op == "listPost" {
			res := post.PostMgr.ListPost()
			for _, p := range res {
				fmt.Println(p)
			}
		} else if op == "exit" {
			break
		} else {
			log.Println("Error: unknown operation")
		}
	}
	return
}
