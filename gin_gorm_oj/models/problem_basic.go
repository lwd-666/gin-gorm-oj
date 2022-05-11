package models

import (
	"fmt"
	"gorm.io/gorm"
)

type ProblemBasic struct {
	gorm.Model
	Identity string 						`gorm:"column:identity;type:varchar(36);" json:"identity"`
	ProblemCategories [] *ProblemCategory 	`gorm:"foreignKey:problem_id;references:id"`
	Title string 							`gorm:"column:title;type:varchar(255);" json:"title"`
	Content string 							`gorm:"column:content;type:text;" json:"content"`
	MaxRuntime int 							`gorm:"column:max_runtime;type:int(11);" json:"max_runtime"`
	MaxMem int 								`gorm:"column:max_mem;type:int(11);" json:"max_mem"`
}

func (table *ProblemBasic) TableName()string {
	return "problem_basic"
}
func GetProblemList(keyword ,categoryIdentity string) *gorm.DB{
	fmt.Println("keyword=",keyword)
	fmt.Println("11categoryIdentity=",categoryIdentity)
	tx := DB.Model(new(ProblemBasic)).
		Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").
		Where("title like ? OR content like ?","%"+keyword+"%","%"+keyword+"%")
	if categoryIdentity != ""{
		fmt.Println("22categoryIdentity=",categoryIdentity)
		tx.Joins("RIGHT JOIN problem_category pc on pc.problem_id=problem_basic.id").
			Where("pc.category_id = (SELECT cb.id FROM category_basic cb WHERE cb.identity = ? )",categoryIdentity)
	}
	return tx
}
