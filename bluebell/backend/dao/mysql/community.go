package mysql

import (
	"bluebell/models"
	"database/sql"
	"errors"
	"fmt"

	"go.uber.org/zap"
)

var (
	ErrorInvalidID = errors.New("无效的ID！")
)

// GetCommunity 社区列表
func GetCommunity() (list []models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	err = db.Select(&list, sqlStr)
	if err == sql.ErrNoRows {
		zap.L().Warn("no community in db")
		err = nil
	}
	return
}

// CommunityDetail 社区详情
func CommunityDetail(id int64) (detail *models.CommunityDetail, err error) {
	detail = new(models.CommunityDetail)
	sqlStr := `select 
       community_id, community_name, introduction, create_time
		from community 
		where community_id = ?`

	err = db.Get(detail, sqlStr, id)
	if err == sql.ErrNoRows {
		err = ErrorInvalidID
	}
	fmt.Println(detail)
	return
}
