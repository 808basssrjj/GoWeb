package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

// GetCommunity 社区列表
func GetCommunity() (list []models.Community, err error) {
	return mysql.GetCommunity()
}

// CommunityDetail 社区详情
func CommunityDetail(id int64) (detail *models.CommunityDetail, err error) {
	return mysql.CommunityDetail(id)
}
