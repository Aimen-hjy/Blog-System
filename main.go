package main

import (
	"blogSystem/post"
	"blogSystem/user"
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func indexHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/login")
	//TODO:Remember account
}
func loginGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}
func loginPostHandler(c *gin.Context) {
	option := c.PostForm("option")
	username := c.PostForm("name")
	password := c.PostForm("password")
	if option == "login" {
		if user.UserMgr.Login(username, password) {
			c.Header("Refresh", "3;url=/dashboard")
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{"Name": username, "Password": password, "Error": "Invalid username and password."})
		}
	} else {
		c.Redirect(http.StatusFound, "/register")
	}
}
func registerPostHandler(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	if user.UserMgr.Register(name, password) {
		c.Header("Refresh", "3;url=/login")
		c.HTML(http.StatusOK, "register.html", gin.H{"Success": "Register successfully! Redirecting to login page in 3 seconds..."})
	} else {
		c.HTML(http.StatusOK, "register.html", gin.H{"Name": name, "Password": password, "Error": "Username already exists."})
	}
}
func registerGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}
func changePasswordGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "changepassword.html", gin.H{})
}
func changePasswordPostHandler(c *gin.Context) {
	oldPassword := c.PostForm("oldpassword")
	newPassword := c.PostForm("newpassword")
	confirmPassword := c.PostForm("confirmpassword")
	if newPassword != confirmPassword {
		c.HTML(http.StatusOK, "changepassword.html", gin.H{"Error": "The two new passwords do not match."})
		return
	}
	if user.UserMgr.ChangePassword(oldPassword, newPassword) {
		c.HTML(http.StatusOK, "changepassword.html", gin.H{"Success": "Password changed successfully!"})
	} else {
		c.HTML(http.StatusOK, "changepassword.html", gin.H{"Error": "Old password is incorrect."})
	}
}
func logoutGetHandler(c *gin.Context) {
	user.UserMgr.Logout()
	c.Redirect(http.StatusFound, "/login")
}
func dashboradGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", gin.H{})
	//TODO:Dashboard
}
func main() {
	post.PostMgr.Init()
	defer post.PostMgr.CloseDatabase()
	user.UserMgr.Init()
	defer user.UserMgr.CloseDatabase()
	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	r.Static("/static", "./static")
	r.GET("/", indexHandler)
	r.GET("/login", loginGetHandler)
	r.POST("/login", loginPostHandler)
	r.POST("/register", registerPostHandler)
	r.GET("/register", registerGetHandler)
	r.GET("dashboard", dashboradGetHandler)
	r.GET("/logout", logoutGetHandler)
	r.GET("/changepassword", changePasswordGetHandler)
	r.POST("/changepassword", changePasswordPostHandler)

	r.Run(":8080")
	return
}

func cmd() {
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
					log.Println("Error: invalid post ID")
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
					log.Println("Error: unknown argument")
					continue
				}
			} else if len(fields) == 5 {
				id, err := strconv.Atoi(fields[0])
				if err != nil {
					log.Println("Error: invalid post ID")
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
				log.Println(p)
			}
		} else if op == "searchPost" {
			var keyword string
			fmt.Scan(&keyword)
			if keyword == "-t" {
				var title string
				fmt.Scan(&title)
				res := post.PostMgr.SearchPostsByTitle(title)
				if len(res) == 0 {
					log.Println("No posts found with the given title.")
					continue
				}
				for _, p := range res {
					log.Println(p)
				}
			} else if keyword == "-ct" {
				var year, month, day int
				fmt.Scan(&year, &month, &day)
				res := post.PostMgr.SearchPostsByCreateTime(year, month, day)
				if len(res) == 0 {
					log.Println("No posts found with the given creation date.")
					continue
				}
				for _, p := range res {
					fmt.Println(p)
				}
			} else if keyword == "-ut" {
				var year, month, day int
				fmt.Scan(&year, &month, &day)
				res := post.PostMgr.SearchPostsByUpdateTime(year, month, day)
				if len(res) == 0 {
					log.Println("No posts found with the given update date.")
					continue
				}
				for _, p := range res {
					log.Println(p)
				}
			} else {
				log.Println("Error: unknown search option")
			}
		} else if op == "exit" {
			log.Println("Exiting...")
			break
		} else {
			log.Println("Error: unknown operation")
		}
	}
}
