package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/gin_gorm_oj/define"
	"go_code/gin_gorm_oj/models"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)
//	GetProblemList
//	@Tags 公共方法
//	@Summary 问题列表
//	@Param page query int false "page"
//	@Param size query int false "size"
//	@Param keyword query string false "keyword"
//	@Param category_identity query string false "category_identity"
//  @Success 200 {string} json "{"code":"200","data":""}"
//	@Router /problem-list [get]
func GetProblemList(c *gin.Context)  {
	size,err :=strconv.Atoi(c.DefaultQuery("size",define.DefaultSize))
	page ,err := strconv.Atoi(c.DefaultQuery("page",define.DefaultPage))
	if err != nil {
		log.Println("GetProblemList Page strconv err=",err)
		return
	}
	page = (page-1)*size
	var count int64
	keyword:=c.Query("keyword")
	categoryIdentity:=c.Query("category_identity")
	list:=make([] *models.ProblemBasic,0)
	tx := models.GetProblemList(keyword,categoryIdentity)
	err =tx.Count(&count).Omit("content").Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("Get Problem List err=",err)
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":map[string]interface{}{
			"list": list,
			"count":count,
		},
	})
}

//	GetProblemDetail
//	@Tags 公共方法
//	@Summary 问题详情
//	@Param identity query string false "problem identity"
//  @Success 200 {string} json "{"code":"200","data":""}"
//	@Router /problem-detail [get]
func GetProblemDetail(c *gin.Context) {
	fmt.Println("2222222222222222222222222")
	identity := c.Query("identity")
	if identity=="" {
		c.JSON(http.StatusOK,gin.H{
			"code": -1,
			"msg":"问题唯一标识不能为空",
		})
		return
	}
	data := new(models.ProblemBasic)
	err := models.DB.Where("identity=?",identity).
		Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").
		First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK,gin.H{
				"code": -1,
				"data":"当前问题不存在",
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"code": -1,
			"data":"GetProblemDetail err="+err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code": 200,
		"data":data,
	})
}