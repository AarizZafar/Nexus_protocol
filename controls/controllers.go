package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AarizZafar/Nexus_protocol.git/bioMetrics"
	"github.com/AarizZafar/Nexus_protocol.git/verification/Vmcontrols"
)

func LoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login_page.html",nil)
}

func NetAuthentication(ctx *gin.Context) {
	// userName         := ctx.PostForm("username")
	// userPass         := ctx.PostForm("Password")
	netAccess        := ctx.PostForm("networkname")

	var adminCred       = Vmcontrols.GetAdminCreds()
	fmt.Println(adminCred)
	var currNetwork     = bioMetrics.BMstruct.SSID
	var ssids           = Vmcontrols.GetAllRecInNetDB()                  // this has all the network that are registered in the db
	fmt.Println(ssids)
	var testNets        = Vmcontrols.GetTestNetsfromSSID(currNetwork)
	fmt.Println(testNets)
	var testNetBioMetx  = Vmcontrols.GetBioMetxFromTestNet(currNetwork, netAccess)    // GET ALL THE BIOMETRIX FROM THE TEST NET 
	fmt.Println(testNetBioMetx)

	// fmt.Println(testNetBioMetx)
	
	// var regSSID bool = false
	// for _, ssid := range ssids {
	// 	if currNetwork == ssid {
	// 		regSSID = true
	// 	}
	// }

	// var regTestNet bool = false
	// for _, testNet := range testNets {
	// 	if testNet == netAccess {
	// 		regTestNet = true
	// 	}
	// }
	
	// adminPass, exists := adminCred[userName]
	
	// if exists && userPass == adminPass {
	// 	if regSSID {
	// 		if regTestNet {
	// 			ctx.HTML(http.StatusOK, "success.html", nil)
	// 		}
	// 	}
		
	// } else {
	// 	ctx.HTML(http.StatusUnauthorized, "failure.html",nil)
    // }
}

func GetSysBioMetrix(c *gin.Context) {
	// The BMstruct that we had declared in the bioMetrix file 
	BMmodel := bioMetrics.BMstruct
	c.JSON(http.StatusOK,BMmodel)
}
