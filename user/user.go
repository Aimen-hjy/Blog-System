package user

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
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
		log.Println(err)
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
		log.Println("Registering new user:", username)
		return true
	} else if result.Error != nil {
		log.Println("Error during registration:", result.Error)
		return false
	} else {
		log.Println("Username already exists:", username)
		return false
	}
}
func (UsrMgr *UserManager) Login(username string, password string) bool {
	if UsrMgr.currentUser != nil && UsrMgr.currentUser.ID != 0 {
		log.Println("A user is already logged in:", UsrMgr.currentUser.Name)
		return false
	}
	var user User
	result := UsrMgr.dataBase.Where("name = ?", username).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		log.Println("User not found:", username)
		return false
	}
	if result.Error != nil {
		log.Println("Error during login:", result.Error)
		return false
	}
	if user.Password != password {
		log.Println("Wrong password:", user.Password)
		return false
	}
	UsrMgr.currentUser = &user
	log.Println("Logged in:", username)
	return true
}
func (UsrMgr *UserManager) Logout() bool {
	if UsrMgr.currentUser == nil || UsrMgr.currentUser.ID == 0 {
		log.Println("No user is currently logged in.")
		return false
	}
	UsrMgr.currentUser.ID = 0
	UsrMgr.currentUser.Name = ""
	UsrMgr.currentUser.Password = ""
	log.Println("Logged out.")
	return true
}
func (UsrMgr *UserManager) ChangePassword(old_password, new_password string) bool {
	if UsrMgr.currentUser == nil || UsrMgr.currentUser.ID == 0 {
		log.Println("No user is currently logged in.")
		return false
	}
	if UsrMgr.currentUser.Password != old_password {
		log.Println("Old password is incorrect.")
		return false
	}
	UsrMgr.dataBase.Model(&User{}).Where("name = ?", UsrMgr.currentUser.Name).Update("password", new_password)
	UsrMgr.currentUser.Password = new_password
	log.Println("Password changed successfully.")
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
func (UsrMgr *UserManager) SetCurrentUserByName(name string) bool {
	var user User
	result := UsrMgr.dataBase.Where("name = ?", name).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		log.Println("User not found:", name)
		return false
	}
	if result.Error != nil {
		log.Println("Error during setting current user:", result.Error)
		return false
	}
	UsrMgr.currentUser = &user
	log.Println("Current user set to:", name)
	return true
}
