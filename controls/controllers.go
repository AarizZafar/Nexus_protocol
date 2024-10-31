package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AarizZafar/Nexus_protocol.git/bioMetrics"
	"github.com/AarizZafar/Nexus_protocol.git/verification/Vmcontrols"
)

func LoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login_page.html",nil)
}

func NetAuthentication(ctx *gin.Context) {
	userName           := ctx.PostForm("username")
	userPass           := ctx.PostForm("Password")
	testNetAccess      := ctx.PostForm("networkname")

	var adminCred       = Vmcontrols.GetAdminCreds()
	var currNetwork     = bioMetrics.BMstruct.SSID
	var ssids           = Vmcontrols.GetAllRecInNetDB()                  // this has all the network that are registered in the db
	var testNets        = Vmcontrols.GetTestNetsfromSSID(currNetwork)
	// var testNetBioMetx  = Vmcontrols.GetBioMetxFromTestNet(currNetwork, netAccess)    // GET ALL THE BIOMETRIX FROM THE TEST NET 
	
	var regSSID bool = false
	for _, ssid := range ssids {
		if currNetwork == ssid {
			regSSID = true
		}
	}

	var regTestNet bool = false
	for _, testNet := range testNets {
		if testNet == testNetAccess {
			regTestNet = true
		}
	}
	
	adminPass, exists := adminCred[userName]
	
	if exists && userPass == adminPass {     					// the person trying to enter is he an admin or not 
		if regSSID {                         					// the network is he in is it registered or not 
			if regTestNet {                  					// the test network he trying to access is it registered or not
				ctx.HTML(http.StatusOK, "success.html", nil)
				} else {                                        // the test net the admin trying to access does not exist hence create it
					Vmcontrols.CreateTestNetInSSID(currNetwork, testNetAccess, bioMetrics.BMstruct)
				    ctx.HTML(http.StatusOK, "success.html", nil)
			}
		} else {                                                // the network the admin is in is not registered 
			Vmcontrols.SendNetBioMetrix()
			Vmcontrols.CreateTestNetInSSID(currNetwork, testNetAccess, bioMetrics.BMstruct)
			ctx.HTML(http.StatusOK, "success.html", nil)
		}
	} else {
		ctx.HTML(http.StatusUnauthorized, "failure.html",nil)
    }
}

func GetSysBioMetrix(c *gin.Context) {
	BMmodel := bioMetrics.BMstruct
	c.JSON(http.StatusOK,BMmodel)
}