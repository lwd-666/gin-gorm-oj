package service

import (
	"github.com/gin-gonic/gin"
	"go_code/gin_gorm_oj/define"
	"go_code/gin_gorm_oj/models"
	"log"
	"net/http"
	"strconv"
)

//	GetSubmitList
//	@Tags 公共方法
//	@Summary 提交列表
//	@Param page query int false "page"
//	@Param size query int false "size"
//	@Param problem_identity query string false "problem_identity"
//	@Param user_identity query string false "user_identity"
//	@Param status query string false "status"
//  @Success 200 {string} json "{"code":"200","data":""}"
//	@Router /submit-list [get]
func GetSubmitList(c *gin.Context)  {
	size,err :=strconv.Atoi(c.DefaultQuery("size",define.DefaultSize))
	page ,err := strconv.Atoi(c.DefaultQuery("page",define.DefaultPage))
	if err != nil {
		log.Println("GetProblemList Page strconv err=",err)
		return
	}
	page = (page-1)*size
	var count int64
	list :=make([]models.SubmitBasic,0)
	problemidentity :=c.Query("problem_identity")
	useridentity :=c.Query("user_identity")
	status,_ :=strconv.Atoi(c.Query("status"))
	tx :=models.GetSubmitList(problemidentity,useridentity,status)

	err = tx.Count(&count).Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("GetProblemList  err=",err)
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"GetProblemList  err"+err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"count":count,
		"data": map[string]interface{}{
			"list":list,
		},
	})
}
