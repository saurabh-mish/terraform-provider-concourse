package main

import (
	"fmt"
	"github.com/saurabh-mish/terraform-provider-concourse/client"
)

func main() {
	username, password := client.CheckCredentials()
	//fmt.Printf("Concourse Email:    %s\n", *username)
	//fmt.Printf("Concourse Password: %s\n", *password)

	respData := client.GetAuthData(username, password)
	fmt.Println(respData)
}
