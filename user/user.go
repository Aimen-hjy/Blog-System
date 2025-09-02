package user

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var UserMgr = UserManager{}

type User struct {
	ID       int64
	Name     string
	Password string
}

type UserManager struct {
	dataBase    *gorm.DB
	currentUser *User
}

func (UsrMgr *UserManager) Init() {
	var err error
	UsrMgr.dataBase, err = gorm.Open(sqlite.Open("./data/user.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	UsrMgr.dataBase.AutoMigrate(&User{})
}

func (UsrMgr *UserManager) CloseDatabase() {
	sqlDB, _ := UsrMgr.dataBase.DB()
	sqlDB.Close()
}

func (UsrMgr *UserManager) Register(username string, password string) bool {

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
func (UsrMgr *UserManager) GetUserCount() int {
	//TODO
}
