package test

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
)

type userClaims struct {
	Identity	string	`json:"identity"`
	Name 		string	`json:"name"`
	jwt.StandardClaims
}


var mykey =[]byte("gin_gorm_oj")
//生成token
func TestGenerateToken(t *testing.T)  {
	userClaim := &userClaims{
		Identity:       "user1",
		Name:           "lwd",
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,userClaim)
	tokenstring,err:=token.SignedString(mykey)
	if err != nil {
		t.Fatal(err)
	}
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6InVzZXIxIiwibmFtZSI6Imx3ZCJ9.Dsvc1M50HdQ2VFnScy1wcjxcp4suzmjlyIe4j6oTF6E
	fmt.Println(tokenstring)
}
//解析token
func TestAnalyseToken(t *testing.T)  {
	tokenstring :="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6InVzZXIxIiwibmFtZSI6Imx3ZCJ9.Dsvc1M50HdQ2VFnScy1wcjxcp4suzmjlyIe4j6oTF6E"
	userClaim :=new(userClaims)

	claims ,err :=jwt.ParseWithClaims(tokenstring,userClaim, func(token *jwt.Token) (interface{}, error) {
		return mykey,nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if claims.Valid {
		fmt.Println(userClaim)
	}
}
