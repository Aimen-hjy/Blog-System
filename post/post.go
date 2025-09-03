package post

import (
	"blogSystem/user"
	"fmt"
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
	UserID     int64 `gorm:"index"`
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
func (PostMgr *PostManager) ViewPost(postID int64) (string, bool) {
	var post Post
	post.ID = postID
	result := PostMgr.dataBase.First(&post)
	if result.Error == gorm.ErrRecordNotFound {
		fmt.Println("Post Not Found")
		return "", false
	}
	if post.UserID != user.UserMgr.GetCurrentUser().ID {
		fmt.Println("It's not your post")
		return "", false
	}
	return post.Content, true
}
func (PostMgr *PostManager) CreatePost(title, content string) (int64, bool) {
	var maxID int64
	PostMgr.dataBase.Model(&Post{}).Select("MAX(id)").Scan(&maxID)
	id := maxID + 1
	newPost := Post{ID: id, Title: title, Content: content, CreateTime: time.Now(), UpdateTime: time.Now(), UserID: user.UserMgr.GetCurrentUser().ID}
	result := PostMgr.dataBase.Create(&newPost)
	if result.Error != nil {
		fmt.Println("Error during creating post:", result.Error)
		return 0, false
	}
	return newPost.ID, true
} //return postID, success
func (PostMgr *PostManager) UpdatePost(postID int64, newtitle, newContent string) bool {
	var post Post
	post.ID = postID
	result := PostMgr.dataBase.First(&post)
	if result.Error == gorm.ErrRecordNotFound {
		fmt.Println("Post Not Found")
		return false
	}
	if post.UserID != user.UserMgr.GetCurrentUser().ID {
		fmt.Println("It's not your post")
		return false
	}
	if newtitle != "" {
		post.Title = newtitle
	}
	if newContent != "" {
		post.Content = newContent
	}
	post.UpdateTime = time.Now()
	PostMgr.dataBase.Save(&post)
	return true
}
func (PostMgr *PostManager) DeletePost(postID int64) bool {
	var post Post
	post.ID = postID
	result := PostMgr.dataBase.First(&post)
	if result.Error == gorm.ErrRecordNotFound {
		fmt.Println("Post Not Found")
		return false
	}
	if post.UserID != user.UserMgr.GetCurrentUser().ID {
		fmt.Println("It's not your post")
		return false
	}
	PostMgr.dataBase.Delete(&post)
	fmt.Println("Post Deleted")
	return true
}
func (PostMgr *PostManager) ListPost() []Post {
	var res []Post
	PostMgr.dataBase.Where("user_id = ?", user.UserMgr.GetCurrentUser().ID).Find(&res)
	return res
}
func (PostMgr *PostManager) GetPostCount() int64 {
	var count int64
	PostMgr.dataBase.Model(&Post{}).Count(&count)
	return count
}
func searchPostsByTitle(title string) []Post {
	//TODO
}
func searchPostsByCreateTime() []Post {
	//TODO
}
func searchPostsByUpdateTime() []Post {
	//TODO
}
