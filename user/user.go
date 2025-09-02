package user

import (
	"database/sql"
	"fmt"
)

var UserMgr = UserManager{}

type User struct {
	ID       int64
	Name     string
	Password string
}

type UserManager struct {
	dataBase    *sql.DB
	currentUser *User
}

func (UsrMgr *UserManager) Init() {
	var err error
	UsrMgr.dataBase, err = sql.Open("sqlite3", "./data/user.db")
	if err != nil {
		fmt.Println(err)
	}
	_, err = UsrMgr.dataBase.Exec("CREATE TABLE IF NOT EXISTS Users (ID INTEGER PRIMARY KEY AUTOINCREMENT, Name TEXT, Password TEXT)")
	if err != nil {
		fmt.Println(err)
	}
}

func (UsrMgr *UserManager) CloseDatabase() {
	UsrMgr.dataBase.Close()
}

func (UsrMgr *UserManager) Register(username string, password string) (err error) {
	//TODO
}
func (UsrMgr *UserManager) Login(username string, password string) (err error) {
	//TODO
}
func (UsrMgr *UserManager) Logout() {
	//TODO
}
func (UsrMgr *UserManager) ChangePassword(oldPassword string, newPassword string) (err error) {
	//TODO
}
func (UsrMgr *UserManager) GetCurrentUser() *User {
	//TODO
}
