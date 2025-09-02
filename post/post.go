package post

import (
	"database/sql"
	"fmt"
	"time"
)

var PostMgr = PostManager{}

type Post struct {
	ID         int64
	Title      string
	Content    string
	CreateTime time.Time
	UpdateTime time.Time
}
type PostManager struct {
	dataBase *sql.DB
}

func (PostMgr *PostManager) Init() {
	var err error
	PostMgr.dataBase, err = sql.Open("sqlite3", "./data/post.db")
	if err != nil {
		fmt.Println(err)
	}
	_, err = PostMgr.dataBase.Exec("CREATE TABLE IF NOT EXISTS Posts (ID INTEGER PRIMARY KEY AUTOINCREMENT, Title TEXT, Content TEXT, CreateTime DATETIME, UpdateTime DATETIME)")
	if err != nil {
		fmt.Println(err)
	}
}
func (PostMgr *PostManager) CloseDatabase() {
	PostMgr.dataBase.Close()
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
