package post

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

var PostMgr = PostManager{}

type Post struct {
	ID         int64 `gorm:"primary_key;autoIncrement"`
	Title      string
	Content    string
	CreateTime time.Time
	UpdateTime time.Time
}
type PostManager struct {
	dataBase *gorm.DB
}

func (PostMgr *PostManager) Init() {
	var err error
	PostMgr.dataBase, err = gorm.Open(sqlite.Open("./data/posts.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	PostMgr.dataBase.AutoMigrate(&Post{})
}
func (PostMgr *PostManager) CloseDatabase() {
	sqlDB, _ := PostMgr.dataBase.DB()
	sqlDB.Close()
}
func (PostMgr *PostManager) ViewPost() {
	//TODO
}
func (PostMgr *PostManager) CreatePost(post Post) {
	//TODO
}
func (PostMgr *PostManager) UpdatePost(post Post) {
	//TODO
}
func (PostMgr *PostManager) DeletePost(postID int64) {
	//TODO
}
func (PostMgr *PostManager) ListPost() {
	//TODO
}
func (PostMgr *PostManager) GetPostCount() int {
	//TODO
	return 0
}
