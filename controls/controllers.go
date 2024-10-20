package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/AarizZafar/Nexus_protocol.git/bioMetrics"
	"github.com/AarizZafar/Nexus_protocol.git/netCredentials"
)

// the netCred variabel starts with a lower case hence it is private
var netCred = netCredentials.NetCred()

func LoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login_page.html",nil)
}

func NetAuthentication(ctx *gin.Context) {
	netName := ctx.PostForm("username")
	netPass := ctx.PostForm("Password")

	passWd, exists := netCred[netName]

	if exists && netPass == passWd {
		ctx.HTML(http.StatusOK, "success.html", nil)
	} else {
		ctx.HTML(http.StatusUnauthorized, "failure.html",nil)
    }
}

func GetSysBioMetrix(c *gin.Context) {
	// The BMstruct that we had declared in the bioMetrix file 
	BMmodel := bioMetrics.BMstruct
	c.JSON(http.StatusOK,BMmodel)
}