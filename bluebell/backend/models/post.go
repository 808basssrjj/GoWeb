package models

import "time"

type Post struct {
	PostID      int64     `json:"post_id,string" db:"post_id"`
	AuthorId    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

type ApiPostDetail struct {
	AutoName string `json:"auto_name"`
	*Post
	*CommunityDetail `json:"community"`
}

type PostList struct {
	VoteNum       int64     `json:"vote_num"  db:"VoteNum"`
	PostID        int64     `json:"id,string" db:"post_id"`
	Status        int32     `json:"status" db:"status"`
	Title         string    `json:"title"  db:"title"`
	Content       string    `json:"content"  db:"content"`
	Username      string    `json:"username"  db:"username"`
	CommunityName string    `json:"community_name"  db:"community_name"`
	CreateTime    time.Time `json:"create_time"  db:"create_time"`
}

const (
	OrderTime  = "time"
	OrderScore = "score"
)
