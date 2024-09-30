package controllers

import (
	"fmt"
	"net/http"

	"github.com/AarizZafar/Nexus_protocol.git/netCredentials"
	"github.com/gin-gonic/gin"
)

// the netCred variabel starts with a lower case hence it is private
var netCred = netCredentials.NetCred()

func LoginPage(ctx *gin.Context) {
	fmt.Println("------------------------------------------------")
	ctx.HTML(http.StatusOK, "login_page.html",nil)
}

func NetAuthentication(ctx *gin.Context) {
	netName := ctx.PostForm("username")
	netPass := ctx.PostForm("Password")

	passWd, exists := netCred[netName]

	if exists && netPass == passWd {
		fmt.Println("CORRECT CREDENTIALS ")
		fmt.Println("netName : ", netName)
		fmt.Println("netPass : ", netPass)
		ctx.HTML(http.StatusOK, "success.html", nil)


		} else {
		fmt.Println("WRONG CREDENTIALS ")
		fmt.Println("netName : ", netName)
		fmt.Println("netPass : Does not exist")


		// // Failed authentication: render the login page with an error message
		// ctx.HTML(http.StatusUnauthorized, "login_page.html", gin.H{
		// 	"error": "Invalid network name or password",
		// })
	}
}


