package gopc

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

var envs = map[string]string{
	"production":   "https://payconiq-app.com",
	"homologation": "https://homologation.payconiq-app.com",
}

func (c *Client) Init(publicVendorToken string, privateVendorToken string, env string) {
	c.identity.PublicVendorToken = publicVendorToken
	c.identity.PrivateVendorToken = privateVendorToken
	c.env = env
	endpoint, ok := envs[env]
	if !ok {
		endpoint = envs["production"]
	}
	c.endpoint = endpoint
	c.options = ClientOptions{
		MinExpirationDelay: 1,
	}
}

func (c *Client) isIdentified() bool {
	return len(c.identity.PublicVendorToken) > 0
}

func (c *Client) request(uri string, data url.Values, result interface{}) error {
	res, err := http.Post(c.endpoint+"/api"+uri, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return err
	}
	return nil
}
