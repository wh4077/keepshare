// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameWorkerAccount = "pikpak_worker_account"

// WorkerAccount mapped from table <pikpak_worker_account>
type WorkerAccount struct {
	UserID            string    `gorm:"column:user_id;primaryKey" json:"user_id"`
	MasterUserID      string    `gorm:"column:master_user_id;not null" json:"master_user_id"`
	Email             string    `gorm:"column:email;not null" json:"email"`
	Password          string    `gorm:"column:password;not null" json:"password"`
	UsedSize          int64     `gorm:"column:used_size;not null" json:"used_size"`
	LimitSize         int64     `gorm:"column:limit_size;not null" json:"limit_size"`
	PremiumExpiration time.Time `gorm:"column:premium_expiration;not null;default:2000-01-01 00:00:00" json:"premium_expiration"`
	CreatedAt         time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName WorkerAccount's table name
func (*WorkerAccount) TableName() string {
	return TableNameWorkerAccount
}
