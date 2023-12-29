// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameRedeemCode = "pikpak_redeem_code"

// RedeemCode mapped from table <pikpak_redeem_code>
type RedeemCode struct {
	AutoID     int64     `gorm:"column:auto_id;primaryKey;autoIncrement:true" json:"auto_id"`
	Code       string    `gorm:"column:code;not null" json:"code"`
	Status     string    `gorm:"column:status;not null;default:NOT_USED;comment:NOT_USED, USED, INVALID" json:"status"`
	UsedUserID string    `gorm:"column:used_user_id" json:"used_user_id"`
	CreatedAt  time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Error      string    `gorm:"column:error" json:"error"`
}

// TableName RedeemCode's table name
func (*RedeemCode) TableName() string {
	return TableNameRedeemCode
}
