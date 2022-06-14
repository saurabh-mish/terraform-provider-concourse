package concourse

import (
	"context"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"attrtag": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	got := Provider().InternalValidate()
	if got != nil {
		t.Errorf("got %v; wanted %v", got, nil)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if os.Getenv("CONCOURSE_USERNAME") == "" || os.Getenv("CONCOURSE_PASSWORD") == "" {
		t.Errorf("Environment variables CONCOURSE_USERNAME and CONCOURSE_PASSWORD must be set")
	}

	diags := Provider().Configure(context.Background(), &terraform.ResourceConfig{})
	if diags.HasError() {
		for _, d := range diags {
			if d.Severity == diag.Error {
				t.Errorf("Error %v", d.Summary)
			}
		}
	}
}
