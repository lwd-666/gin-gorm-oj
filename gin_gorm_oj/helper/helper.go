package helper

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"net/smtp"
)
type userClaims struct {
	Identity	string	`json:"identity"`
	Name 		string	`json:"name"`
	jwt.StandardClaims
}

//生成md5
func Getmd5(s string) string {
	return fmt.Sprintf("%x",md5.Sum([]byte(s)))
}

var mykey =[]byte("gin_gorm_oj")
//GenerateToken
//生成token
func GenerateToken(identity,name string) (string,error) {
	userClaim := &userClaims{
		Identity:       identity,
		Name:           name,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,userClaim)
	tokenstring,err:=token.SignedString(mykey)
	if err != nil {
		return "", err
	}

	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6InVzZXIxIiwibmFtZSI6Imx3ZCJ9.Dsvc1M50HdQ2VFnScy1wcjxcp4suzmjlyIe4j6oTF6E
	fmt.Println(tokenstring)
	return tokenstring,nil
}
//AnalyseToken
//解析token
func AnalyseToken(tokenstring string) (*userClaims ,error) {
	//tokenstring :="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6InVzZXIxIiwibmFtZSI6Imx3ZCJ9.Dsvc1M50HdQ2VFnScy1wcjxcp4suzmjlyIe4j6oTF6E"
	userClaim :=new(userClaims)
	claims ,err :=jwt.ParseWithClaims(tokenstring,userClaim, func(token *jwt.Token) (interface{}, error) {
		return mykey,nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, err
	}
	return userClaim,nil
}

//发送验证码
func SendCode(toUserEmail,code string) error {
	fmt.Println("helper.go-SendCode")
	e := email.NewEmail()
	e.From = "Get <getcharzhaopan@163.com>"
	//e.From = "Get <1124099628@qq.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("您的验证码是<b>123456</b>")
	//err :=e.Send("smtp.163.com:465", smtp.PlainAuth("", "getcharzhaopan@163.com", "XYQHBVISUAHXXRHQ", "smtp.163.com:465"))
	fmt.Println("SendCode err = ",toUserEmail,code)
	//返回EOF时，关闭SSL重试
	return e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "getcharzhaopan@163.com", "XYQHBVISUAHXXRHQ", "smtp.163.com"),&tls.Config{ServerName: "smtp.163.com",InsecureSkipVerify: true})
}

func GetUUID()string  {
	return uuid.NewV4().String()
}