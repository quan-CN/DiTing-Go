// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "user"

// User 用户表
type User struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:用户id" json:"id"`                                  // 用户id
	Password     string    `gorm:"column:password;comment:用户密码" json:"password"`                                                    // 用户密码
	Name         string    `gorm:"column:name;comment:用户昵称" json:"name"`                                                            // 用户昵称
	Avatar       string    `gorm:"column:avatar;comment:用户头像" json:"avatar"`                                                        // 用户头像
	Sex          int32     `gorm:"column:sex;comment:性别 1为男性，2为女性" json:"sex"`                                                      // 性别 1为男性，2为女性
	OpenID       string    `gorm:"column:open_id;not null;comment:微信openid用户标识" json:"open_id"`                                     // 微信openid用户标识
	ActiveStatus int32     `gorm:"column:active_status;default:2;comment:在线状态 1在线 2离线" json:"active_status"`                        // 在线状态 1在线 2离线
	LastOptTime  time.Time `gorm:"column:last_opt_time;not null;default:CURRENT_TIMESTAMP(3);comment:最后上下线时间" json:"last_opt_time"` // 最后上下线时间
	IPInfo       string    `gorm:"column:ip_info;comment:ip信息" json:"ip_info"`                                                      // ip信息
	ItemID       int64     `gorm:"column:item_id;comment:佩戴的徽章id" json:"item_id"`                                                   // 佩戴的徽章id
	Status       int32     `gorm:"column:status;comment:使用状态 0.正常 1拉黑" json:"status"`                                               // 使用状态 0.正常 1拉黑
	CreateTime   time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP(3);comment:创建时间" json:"create_time"`        // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP(3);comment:修改时间" json:"update_time"`        // 修改时间
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
