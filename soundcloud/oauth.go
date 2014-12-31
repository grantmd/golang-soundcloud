package soundcloud

import (
	"fmt"
	"net/url"
)

func (api *Api) Refresh() error {
	if api.ClientId == "" || api.ClientSecret == "" {
		return fmt.Errorf("ClientId and ClientSecret must all be specified")
	}

	var values url.Values
	if api.RefreshToken != "" {
		values = Values(
			"client_id", api.ClientId,
			"client_secret", api.ClientSecret,
			"refresh_token", api.RefreshToken,
			"grant_type", "refresh_token",
		)
	} else if api.Username != "" && api.Password != "" {
		values = Values(
			"client_id", api.ClientId,
			"client_secret", api.ClientSecret,
			"username", api.Username,
			"password", api.Password,
			"grant_type", "password",
		)
	}

	ret := new(AuthResponse)
	err := api.post("/oauth2/token", values, ret)

	if err == nil {
		api.AccessToken = ret.AccessToken
		api.RefreshToken = ret.RefreshToken
	}

	return err
}
