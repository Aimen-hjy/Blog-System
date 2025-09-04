package post

import (
	"blogSystem/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
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
		log.Println("Post Not Found")
		return "", false
	}
	if post.UserID != user.UserMgr.GetCurrentUser().ID {
		log.Println("It's not your post")
		return "", false
	}
	log.Println("Post Viewed:", post)
	return post.Content, true
}
func (PostMgr *PostManager) CreatePost(title, content string) (int64, bool) {
	var maxID int64
	PostMgr.dataBase.Model(&Post{}).Select("MAX(id)").Scan(&maxID)
	id := maxID + 1
	newPost := Post{ID: id, Title: title, Content: content, CreateTime: time.Now(), UpdateTime: time.Now(), UserID: user.UserMgr.GetCurrentUser().ID}
	result := PostMgr.dataBase.Create(&newPost)
	if result.Error != nil {
		log.Println("Error during creating post:", result.Error)
		return 0, false
	}
	log.Println("Post created:", newPost)
	return newPost.ID, true
} //return postID, success
func (PostMgr *PostManager) UpdatePost(postID int64, newtitle, newContent string) bool {
	var post Post
	post.ID = postID
	result := PostMgr.dataBase.First(&post)
	if result.Error == gorm.ErrRecordNotFound {
		log.Println("Post Not Found")
		return false
	}
	if post.UserID != user.UserMgr.GetCurrentUser().ID {
		log.Println("It's not your post")
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
	log.Println("Post updated:", post)
	return true
}
func (PostMgr *PostManager) DeletePost(postID int64) bool {
	var post Post
	post.ID = postID
	result := PostMgr.dataBase.First(&post)
	if result.Error == gorm.ErrRecordNotFound {
		log.Println("Post Not Found")
		return false
	}
	if post.UserID != user.UserMgr.GetCurrentUser().ID {
		log.Println("It's not your post")
		return false
	}
	PostMgr.dataBase.Delete(&post)
	log.Println("Post Deleted")
	return true
}
func (PostMgr *PostManager) ListPost() []Post {
	var res []Post
	PostMgr.dataBase.Where("user_id = ?", user.UserMgr.GetCurrentUser().ID).Find(&res)
	log.Println("Post List")
	return res
}
func (PostMgr *PostManager) GetPostCount() int64 {
	var count int64
	PostMgr.dataBase.Model(&Post{}).Count(&count)
	return count
}
func (PostMgr *PostManager) SearchPostsByTitle(title string) []Post {
	var posts []Post
	PostMgr.dataBase.Where("title = ? AND user_id = ?", title, user.UserMgr.GetCurrentUser().ID).Find(&posts)
	log.Println("Search by title:", title)
	return posts
}
func (PostMgr *PostManager) SearchPostsByCreateTime(year, month, day int) []Post {
	var posts []Post
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	next_date := date.AddDate(0, 0, 1)
	PostMgr.dataBase.Where("create_time >= ? AND create_time < ?", date, next_date).Find(&posts)
	log.Println("Search by create time:", date)
	return posts
}
func (PostMgr *PostManager) SearchPostsByUpdateTime(year, month, day int) []Post {
	var posts []Post
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	next_date := date.AddDate(0, 0, 1)
	PostMgr.dataBase.Where("update_time >= ? AND update_time < ?", date, next_date).Find(&posts)
	log.Println("Search by update time:", date)
	return posts
}
