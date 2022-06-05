/*
package main

import (
	"fmt"
	"os"
	"github.com/saurabh-mish/terraform-provider-concourse/client"
)


func main() {
	username := os.Getenv("CONCOURSE_USERNAME")
	password := os.Getenv("CONCOURSE_PASSWORD")
	hosturl  := "https://auth.prod.concourselabs.io/api/v1/oauth"
	respData, val := client.NewClient(&hosturl, &username, &password)
	if val != nil {
		fmt.Println(val)
	}
	fmt.Println(respData)
}
*/

package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/saurabh-mish/terraform-provider-concourse/concourse"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return concourse.Provider()
		},
	})
}
