package models

import "gorm.io/gorm"

type SubmitBasic struct {
	gorm.Model
	Identity string 			`gorm:"column:identity;type:varchar(36);" json:"identity"`
	ProblemIdentity string 		`gorm:"column:problem_identity;type:varchar(36);" json:"problem_identity"`
	ProblemBasic *ProblemBasic 	`gorm:"foreignKey:identity;references:problem_identity"`
	UserIdentity string 		`gorm:"column:user_identity;type:varchar(36);" json:"user_identity"`
	UserBasic *UserBasic		`gorm:"foreignKey:identity;references:user_identity"`
	Path string 				`gorm:"column:path;type:varchar(255);" json:"path"`
	Status int 					`gorm:"column:status;type:tinyint(1);" json:"status"`	//[-1-待判断，1-正确，2-错误，3-运行超时，4-超内存]
}

func (table *SubmitBasic) TableName()string {
	return "submit_basic"
}
func GetSubmitList(problemidentity,useridentity string,status int) *gorm.DB  {
	tx:=DB.Model(new(SubmitBasic)).Preload("ProblemBasic", func(db *gorm.DB) *gorm.DB {
		return db.Omit("content")
	}).Preload("UserBasic")
	if problemidentity!=""{
		tx.Where("problem_identity = ?",problemidentity)
	}
	if useridentity!="" {
		tx.Where("user_identity = ?",useridentity)
	}
	if status !=0 {
		tx.Where("status = ?",status)
	}
	return tx
}