package Vmcontrols

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/AarizZafar/Nexus_protocol.git/BMmodel"
	"github.com/AarizZafar/Nexus_protocol.git/bioMetrics"
)

const NetVerifybaseURL			 = "http://13.90.73.228:8080/NetVerify"
const SysVerifybaseURL			 = "http://13.90.73.228:8080/SysVerify"
const GetAllNetRecbaseURL		 = "http://13.90.73.228:8080/GetSSIDS"
const GetAdminCredsURL			 = "http://13.90.73.228:8080/GetAdmincreds"
const GettestNetfromSSIDURL		 = "http://13.90.73.228:8080/GetTestNetsfromSSID"
const GetBioMetxtestNetURL		 = "http://13.90.73.228:8080/GetBioMetrixtestNet"
const CrtTestNetInSSIDURL		 = "http://13.90.73.228:8080/CrtTestNetInSSID"

func handleErr(err error) {
	if err != nil {
		log.Fatal("\033[97;41m", err, "\033[0m")
	}
}

func SendNetBioMetrix() {
	fmt.Print("\033[46m              SEND NET BIO METRIX                \033[0m")
	jsonData, err := json.Marshal(bioMetrics.NBMstruct)
	handleErr(err)

	resp, err := http.Post(NetVerifybaseURL, "application/json", bytes.NewBuffer(jsonData))
	handleErr(err)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Print("\033[42m ✔  \033[0m\n")
	fmt.Println("Response : ", string(body))
}

func SendSysBioMetrix() {
	fmt.Print("\033[46m            SEND SYS BIO METRIX               \033[0m")
	jsonData, err := json.Marshal(bioMetrics.BMstruct) // the device data that we have assigned to the struct in our local code
	handleErr(err)

	resp, err := http.Post(SysVerifybaseURL, "application/json", bytes.NewBuffer(jsonData))
	handleErr(err)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("\033[42m✔\033[0m")
	fmt.Println("Response : ", string(body))
}

func GetAdminCreds() map[string]string {
	fmt.Print("\033[46m               GET ADMIN CREDS                   \033[0m")
	resp, err := http.Get(GetAdminCredsURL)
	if err != nil {
		fmt.Println("\033[41m     ERROR MAKING REQUEST : ", err, "     \033[0m")
		return nil
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("\033[41m      ERROR NON-OK RESPONSE COULD NOT GET THE ADMIN CREDS : ", resp.Status, "     \033[0m")
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("\033[41m     ERROR READING RESPONSE BODY : ", err, "     \033[0m")
		return nil
	}

	var AdminCred map[string]string
	if err := json.Unmarshal(body, &AdminCred); err != nil {
		handleErr(err)
		return nil
	}
	fmt.Print("\033[42m ✔  \033[0m\n")
	return AdminCred
}

// the below function returns all the network that are registered in the Network db
func GetAllRecInNetDB() []string {
	fmt.Print("\033[46m             GET ALL REC IN NET DB               \033[0m")
	resp, err := http.Get(GetAllNetRecbaseURL)
	if err != nil {
		fmt.Println("\033[41m     ERROR MAKING REQUEST : ", err, "     \033[0m")
		return nil
	}

	defer resp.Body.Close() // free up resources, avoid potential memory leaks

	if resp.StatusCode != http.StatusOK {
		fmt.Println("\033[41m     ERROR NON-OK RESPONSE COULD NOT GET THE SSIDS : ", resp.Status, "     \033[0m")
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body) // reading all the data from resp.Body and returning it as a byte slice
	if err != nil {
		fmt.Println("\033[41m     ERROR READING RESPONSE BODY : ", err, "     \033[0m")
		return nil
	}

	var ssid []string
	if err := json.Unmarshal(body, &ssid); err != nil {
		handleErr(err)
		return nil
	}
	fmt.Print("\033[42m ✔  \033[0m\n")
	return ssid
}

// The below function returns all the testnets from the ssid
func GetTestNetsfromSSID(ssid string) []string {
	fmt.Print("\033[46m             GET TEST NETS FROM SSID             \033[0m")
	reqURL := fmt.Sprintf("%s?ssid=%s", GettestNetfromSSIDURL, ssid)

	resp, err := http.Get(reqURL)
	if err != nil {
		fmt.Println("\033[41m     ERROR MAKING REQUEST : ", err, "     \033[0m")
		return nil
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("\033[41m      ERROR NON-OK RESPONSE : ", resp.Status, "     \033[0m")
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("\033[41m     ERROR READING RESPONSE BODY : ", err, "     \033[0m")
		return nil
	}

	var result struct {
		Collection []string `json:"collection"` 
	}

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("\033[41m     ERROR UNMARSHALING RESPONSE : ", err, "     \033[0m")
		return nil
	}
	fmt.Print("\033[42m ✔  \033[0m\n")
	return result.Collection
}

// the below function return all the device biometrix from the testnet of the ssid
func GetBioMetxFromTestNet(ssid string, testNet string) []map[string]interface{} {
	fmt.Print("\033[46m           GET BIO METX FROM TEST NET                  \033[0m")
	// The return type in a []map[string]interface{} -> because the response from the server contains
	// multiple records and each record has a fiels with different data types (string, number) etc.
	// creates a URL by combining a base url with 2 query parameters
	reqURL := fmt.Sprintf("%s?ssid=%s&testNet=%s", GetBioMetxtestNetURL, ssid, testNet)

	resp, err := http.Get(reqURL)
	if err != nil {
		fmt.Println("\033[41m     ERROR MAKING REQUEST : ", err, "     \033[0m")
		return nil
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("\033[41m      ERROR NON-OK RESPONSE : ", resp.Status, "     \033[0m")
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("\033[41m     ERROR READING RESPONSE BODY : ", err, "     \033[0m")
		return nil
	}

	var result struct {
		Records []map[string]interface{} `json:"records"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("\033[41m     ERROR UNMARSHALING RESPONSE : ", err, "     \033[0m")
		return nil
	}
	fmt.Print("\033[42m ✔  \033[0m\n")
	return result.Records
}

func CreateTestNetInSSID(ssid string, testNet string, sysbioMetx BMmodel.SysBioMetrix) {
	fmt.Print("\033[46m            CREATE TEST NET IN SSID              \033[0m")
	CrtTestNetData := map[string]interface{}{
		"ssid":       ssid,
		"testNet":    testNet,
		"sysBioMetx": sysbioMetx,
	}

	jsonData, err := json.Marshal(CrtTestNetData)
	if err != nil {
		fmt.Println("\033[41mERROR MARSHALING JSON: ", err, "\033[0m")
		return
	}

	resp, err := http.Post(CrtTestNetInSSIDURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("\033[41mERROR MAKING POST REQUEST: ", err, "\033[0m")
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Print("\033[42m ✔  \033[0m\n")
		fmt.Println("POST REQUEST SENT SUCCESSFULLY")
	} else {
		fmt.Printf("\033[41mERROR RESPONSE: %s\033[0m\n", resp.Status)
	}
}