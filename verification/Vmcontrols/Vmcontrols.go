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

const NetVerifybaseURL        = "http://13.90.73.228:8080/NetVerify"
const SysVerifybaseURL        = "http://13.90.73.228:8080/SysVerify"
const GetAllNetRecbaseURL     = "http://13.90.73.228:8080/GetSSIDS"
const GetAdminCredsURL        = "http://13.90.73.228:8080/GetAdmincreds"

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

func GetAllRecInNetDB() []string {
	resp, err := http.Get(GetAllNetRecbaseURL)
	if err != nil {
		fmt.Println("\033[41m     Error making request : ", err, "     \033[0m")
		return nil
	}

	defer resp.Body.Close()  // free up resources, avoid potential memory leaks
	
	if resp.StatusCode != http.StatusOK {
		fmt.Println("\033[41m     Error non-OK response could not get the SSIDS : ", resp.Status, "     \033[0m")
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)     // reading all the data from resp.Body and returning it as a byte slice
	if err != nil {
		fmt.Println("\033[41m     Error reading response body : ", err, "     \033[0m")
		return nil
	}
	
	var ssid []string
	if err := json.Unmarshal(body, &ssid); err != nil {
		handleErr(err)
		return nil
	}
	return ssid
}

func GetAdminCreds() map[string]string {
	resp, err := http.Get(GetAdminCredsURL)
	if err != nil {
		fmt.Println("Error in making the request : ", err)
		return nil
	}
	
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		fmt.Println("\033[41m  Error non-OK response could not get the ADMIN CREDS : ", resp.Status, "     \033[0m")
		return nil
	}
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("\033[41m     Error reading response body : ", err, "     \033[0m")
		return nil
	}

	var AdminCred map[string]string 
	if err := json.Unmarshal(body, &AdminCred); err != nil {
		handleErr(err)
		return nil
	}
	return AdminCred
}


