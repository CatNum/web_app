package mysql

import (
	"errors"
	"go.uber.org/zap"
	"web_app/models"
)

// InitUser 建立社区表
func InitCommunity() (err error) {
	community := models.Community{}
	exist := DB.Migrator().HasTable(&community)
	if !exist {
		err = DB.AutoMigrate(&community)
	}
	return err
}

// InitUser 建立社区详情表
func InitCommunityDetail() (err error) {
	communityDetail := models.CommunityDetail{}
	exist := DB.Migrator().HasTable(&communityDetail)
	if !exist {
		err = DB.AutoMigrate(&communityDetail)
	}
	return err
}

func GetCommunityList() (communityList []*models.Community, err error) {
	// 获取全部记录
	result := DB.Find(&communityList)
	if result.Error != nil {
		err = result.Error
		return nil, err
	}
	if result.RowsAffected == 0 {
		zap.L().Warn("there is no community in db")
		err = nil
	}
	return communityList, err
}

// GetCommunityDetailByID 根据id查询社区详情
func GetCommunityDetailByID(id int64) (communityDetail *models.CommunityDetail, err error) {
	communityDetail = new(models.CommunityDetail)
	result := DB.Where("community_id = ?", id).First(&communityDetail)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("不存在")
	}
	return
}
