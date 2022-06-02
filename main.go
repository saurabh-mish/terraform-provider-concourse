package main

import (
	"fmt"
	"os"
	"github.com/saurabh-mish/terraform-provider-concourse/client"
)


//WORKS with overcloud.go
func main() {
	username := os.Getenv("CONCOURSE_USERNAME")
	password := os.Getenv("CONCOURSE_PASSWORD")
	respData, val := client.NewClient(&username, &password)
	if val != nil {
		fmt.Println(val)
	}
	fmt.Println(respData)
}


/*
func main() {

}
*/
