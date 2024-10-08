package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/AarizZafar/Nexus_protocol.git/bioMetrics"
	"github.com/AarizZafar/Nexus_protocol.git/netCredentials"
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
		ctx.HTML(http.StatusUnauthorized, "failure.html",nil)
    }
}

func GetSysBioMetrix(c *gin.Context) {
	// The BMstruct that we had declared in the bioMetrix file 
	BMmodel := bioMetrics.BMstruct
	c.JSON(http.StatusOK,BMmodel)
}

