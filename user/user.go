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
	newUser := User{Name: username, Password: password}
	var count int64
	UsrMgr.dataBase.Model(&User{}).Count(&count)
	result := UsrMgr.dataBase.Where("name = ?", username).First(&newUser)
	newUser.ID = count + 1
	if result.Error == gorm.ErrRecordNotFound {
		UsrMgr.dataBase.Create(&newUser)
		fmt.Println("Registering new user:", username)
		return true
	} else if result.Error != nil {
		fmt.Println("Error during registration:", result.Error)
		return false
	} else {
		fmt.Println("Username already exists:", username)
		return false
	}
}
func (UsrMgr *UserManager) Login(username string, password string) bool {
	if UsrMgr.currentUser != nil && UsrMgr.currentUser.ID != 0 {
		fmt.Println("A user is already logged in:", UsrMgr.currentUser.Name)
		return false
	}
	var user User
	result := UsrMgr.dataBase.Where("name = ?", username).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		fmt.Println("User not found:", username)
		return false
	}
	if result.Error != nil {
		fmt.Println("Error during login:", result.Error)
		return false
	}
	if user.Password != password {
		fmt.Println("Wrong password:", user.Password)
		return false
	}
	UsrMgr.currentUser = &user
	fmt.Println("Logged in:", username)
	return true
}
func (UsrMgr *UserManager) Logout() bool {
	if UsrMgr.currentUser == nil || UsrMgr.currentUser.ID == 0 {
		fmt.Println("No user is currently logged in.")
		return false
	}
	UsrMgr.currentUser.ID = 0
	UsrMgr.currentUser.Name = ""
	UsrMgr.currentUser.Password = ""
	fmt.Println("Logged out.")
	return true
}
func (UsrMgr *UserManager) ChangePassword(old_password, new_password string) bool {
	if UsrMgr.currentUser == nil || UsrMgr.currentUser.ID == 0 {
		fmt.Println("No user is currently logged in.")
		return false
	}
	if UsrMgr.currentUser.Password != old_password {
		fmt.Println("Old password is incorrect.")
		return false
	}
	UsrMgr.dataBase.Model(&User{}).Where("name = ?", UsrMgr.currentUser.Name).Update("password", new_password)
	UsrMgr.currentUser.Password = new_password
	fmt.Println("Password changed successfully.")
	return true
}
func (UsrMgr *UserManager) GetCurrentUser() *User {
	return UsrMgr.currentUser
}
func (UsrMgr *UserManager) GetUserCount() int64 {
	var count int64
	UsrMgr.dataBase.Model(&User{}).Count(&count)
	return count
}
