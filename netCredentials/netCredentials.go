package netCredentials

// Defining the map  outside the function as a global variable 
var cred = map[string]string{}

// these values will be assigned later on by the vm in azure later on 
var netName = "testNet"
var Password = "testnet123"

// adding the values to the map
func init() {
	cred[netName] = Password
}

// return the credentials map
func NetCred() map[string]string {
	return cred
}