package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"`
	Name string `gorm:"column:name;type:varchar(100);" json:"name"`
	Password string `gorm:"column:password;type:varchar(32);" json:"password"`
	Phone string `gorm:"column:phone;type:varchar(20);" json:"phone"`
	Mail string `gorm:"column:mail;type:varchar(100);" json:"Mail"`
	FinishProblemNum int64 `gorm:"column:finsh_problem_num;type:int(11);" json:"finsh_problem_num"`
	Submit int64	`gorm:"column:submit_num;type:int(11);" json:"submit_num"`
}
func (table *UserBasic) TableName()string {
	return "user_basic"
}