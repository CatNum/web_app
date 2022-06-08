package models

import "gorm.io/gorm"

type Community struct {
	CommunityID   uint64 `gorm:"not null"`
	CommunityName string `gorm:"not null"`
	gorm.Model
}

type CommunityDetail struct {
	Community
	Introduction  string `gorm:"not null"`
}
