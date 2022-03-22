package mysql

import (
	"bluebell/models"
	"database/sql"
	"errors"
	"strings"

	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"
)

var (
	ErrorInsertFailed = errors.New("插入数据失败！")
)

// CreatePost 发布帖子
func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
		post_id, title, content, author_id, community_id)
		values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.PostID, p.Title, p.Content, p.AuthorId, p.CommunityID)
	if err != nil {
		err = ErrorInsertFailed
	}
	return
}

// PostDetail 帖子详情
func PostDetail(id int) (detail *models.Post, err error) {
	detail = new(models.Post)
	sqlStr := `select 
    	post_id, title, content, author_id, community_id, create_time, status
		from post where post_id = ?`

	err = db.Get(detail, sqlStr, id)
	if err == sql.ErrNoRows {
		err = ErrorInvalidID
	}
	return
}

// PostList 帖子列表
func PostList(offset, limit int64) (list []models.PostList, err error) {
	sqlStr := `select post_id, title, content, username, community_name, post.create_time, status
		from post 
		left join user on post.author_id = user.user_id
		left join community on post.community_id = community.id
		limit ?, ?`
	err = db.Select(&list, sqlStr, offset, limit)
	if err == sql.ErrNoRows {
		zap.L().Warn("no Post in db")
		err = nil
	}
	return
}

// PostListByIDs 根据id获取数据
func PostListByIDs(ids []string) (postList []models.PostList, err error) {
	sqlStr := `select post_id, title, content, username, community_name, post.create_time, status
		from post 
		left join user on post.author_id = user.user_id
		left join community on post.community_id = community.id
		where post_id in (?)
		order by FIND_IN_SET(post_id, ?)`

	// 动态填充id
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}
