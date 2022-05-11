package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/gin_gorm_oj/define"
	"go_code/gin_gorm_oj/helper"
	"go_code/gin_gorm_oj/models"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)


//	GetUserDetail
//	@Tags 公共方法
//	@Summary 用户详情
//	@Param identity query string false "problem identity"
//  @Success 200 {string} json "{"code":"200","data":""}"
//	@Router /user-detail [get]
func GetUserDetail(c *gin.Context)  {
	identity:=c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"用户唯一标识不能为空",
		})
		return
	}
	data := new(models.UserBasic)
	err := models.DB.Omit("password").Where("identity=?",identity).Find(&data).Error
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"GetUserDetail err="+err.Error()+identity,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":data,
	})


}

//	login
//	@Tags 公共方法
//	@Summary 用户登录
//	@Param username formData string false "username"
//	@Param password formData string false "password"
//  @Success 200 {string} json "{"code":"200","data":""}"
//	@Router /login [post]
func Login(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	//md5
	if username == ""|| password == "" {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"必填信息为空",
		})
	}
	password=helper.Getmd5(password)
	print(username,password)

	data := new(models.UserBasic)
	err :=models.DB.Where("name = ? AND password=?",username,password).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK,gin.H{
				"code":-1,
				"msg":"用户名密码错误"+err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"login err"+err.Error(),
		})
		return
	}

	token, err := helper.GenerateToken(data.Identity, data.Name)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":	"GenerateToken err"+err.Error(),
		})
	}
	c.JSON(http.StatusOK,gin.H{
		"code":-1,
		"data":map[string]interface{}{
			"token":token,
		},
	})













}

//	SendCode
//	@Tags 公共方法
//	@Summary 发送验证码
//	@Param email formData string true "email"
//  @Success 200 {string} json "{"code":"200","data":""}"
//	@Router /send-code [post]
func SendCode(c *gin.Context)  {
	fmt.Println("user.go-SendCode")
	email :=c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"参数不正确",
		})
	}
	code :="123"
	err :=helper.SendCode(email,code)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"SendCode err"+err.Error(),
		})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"验证码发送成功",
	})
}

//	GetRankList
//	@Tags 公共方法
//	@Summary 用户排行榜
//	@Param page formData string false "page"
//	@Param size formData string false "size"
//  @Success 200 {string} json "{"code":"200","data":""}"
//	@Router /rank-list [get]
func GetRankList(c *gin.Context)  {
	size,err :=strconv.Atoi(c.DefaultQuery("size",define.DefaultSize))
	page ,err := strconv.Atoi(c.DefaultQuery("page",define.DefaultPage))
	if err != nil {
		log.Println("GetProblemList Page strconv err=",err)
		return
	}
	page = (page-1)*size
	var count int64
	list :=make([]models.UserBasic,0)
	err=models.DB.Model(new(models.UserBasic)).Count(&count).Order("finsh_problem_num DESC,submit_num ASC").
		Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"GetRankList err="+err.Error(),
		})
	}

	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"data": map[string]interface{}{
			"list":list,
			"count":count,
		},
	})
}