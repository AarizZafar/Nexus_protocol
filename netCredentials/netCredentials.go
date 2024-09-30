package netCredentials

// Defining the map  outside the function as a global variable
var cred = map[string]string{}

// these values will be assigned later on by the vm in azure later on 
var netName = "testnet"
var newPass = "testnet123"

// adding the values to the map
func init() {
	setCred(netName, newPass)
}

func setCred(newName, newPass string) {
	cred[newName] = newPass
}

// return the credentials map
func NetCred() map[string]string {
	return cred
}

