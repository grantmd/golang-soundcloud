package soundcloud

import (
	"fmt"
	"testing"
)

var doAuthorizedRequests bool
var api *Api
var ladygaga_id uint64 = 55213175
var regain_control_id uint64 = 120702493
var shiroban_id uint64 = 448757
var joneisen_id uint64 = 557633

// -- helpers --

func init() {
	doAuthorizedRequests = (TestConfig["access_token"] != "" || (TestConfig["username"] != "" && TestConfig["password"] != ""))
	if !doAuthorizedRequests {
		fmt.Println("*** Authorized requests will not be performed because no access_token or username/password was specified in config_test.go")
	}
	api = createApi()
}

func authorizedRequest(t *testing.T) {
	if !doAuthorizedRequests {
		t.Skip("access_token or username/password not provided.")
	}
}

func createApi() *Api {
	var api *Api
	if TestConfig["username"] != "" && TestConfig["password"] != "" {
		fmt.Println("*** Using username/password auth")
		api = &Api{
			ClientId:     TestConfig["client_id"].(string),
			ClientSecret: TestConfig["client_secret"].(string),
			Username:     TestConfig["username"].(string),
			Password:     TestConfig["password"].(string),
		}
	} else {
		fmt.Println("*** Using access token auth")
		api = &Api{
			ClientId:     TestConfig["client_id"].(string),
			ClientSecret: TestConfig["client_secret"].(string),
			AccessToken:  TestConfig["access_token"].(string),
			RefreshToken: TestConfig["refresh_token"].(string),
		}
	}

	_ = api.Refresh()
	return api
}
