package Vmcontrols

import (
	"fmt"
	"log"
	"bytes"
	"io/ioutil"
	"net/http"
	"encoding/json"

	"github.com/AarizZafar/Nexus_protocol.git/bioMetrics"
)

const NetVerifybaseURL = "http://13.90.73.228:8080/NetVerify"
const SysVerifybaseURL = "http://13.90.73.228:8080/SysVerify"

func handleErr(err error) {
	if err != nil {
		log.Fatal("\033[97;41m", err, "\033[0m")
	}
}

func SendNetBioMetrix() {
	// the current network that we are in our system will send the network biometrix also
	jsonData, err := json.Marshal(bioMetrics.NBMstruct)
	handleErr(err)

	resp, err := http.Post(NetVerifybaseURL, "application/json", bytes.NewBuffer(jsonData))
	handleErr(err)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response : ", string(body))
}

func SendSysBioMetrix() {
	// the device data that we have assigned to the struct in our local code
	jsonData, err := json.Marshal(bioMetrics.BMstruct)
	handleErr(err)

	resp, err := http.Post(SysVerifybaseURL, "application/json", bytes.NewBuffer(jsonData))
	handleErr(err)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response : ", string(body))
}

