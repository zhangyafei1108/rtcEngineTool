package main

import "C"
import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "golang.org/x/net/route"
	"net/http"

	//"time"
)

func main()  {

	gin.SetMode(gin.DebugMode)

	httpserver := gin.New()


	httpserver.GET("/parse", func(c *gin.Context) {

		srcToken := c.DefaultQuery("token","")



		//-APPSECRET := "65aa3ffe992148a480bbf813f3f72fee"

		//-APPID := "f19aa9bc123f4e81af71ff53beebb5a4"

		//claims := Claims{
		//	appid: APPID,
		//	timestamp: time.Now().Unix(),
		//}
		//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		//auth, err := token.SignedString([]byte(APPSECERT))

		parser:=jwt.Parser{SkipClaimsValidation:true}
		token, parts, err := parser.ParseUnverified(srcToken, jwt.MapClaims{})

		room := token.Claims.(jwt.MapClaims)["room"];
		user := token.Claims.(jwt.MapClaims)["user"];


		if err != nil {
			panic(err)
		}else {
			c.JSON(http.StatusOK,gin.H{"room":room,"user":user})
		}

		fmt.Println(token)
		fmt.Println(parts)


	})

	port:= "8080"
	_ = httpserver.Run(":" + port)

}
