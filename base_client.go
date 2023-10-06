package gopc

import (
	"fmt"
	"net/http"
)

var envs = map[string]string{
	"production": "https://api.payconiq.com",
	"external":   "https://api.ext.payconiq.com",
}

func (c *Client) Init(apiKey string, env string) {
	c.identity.ApiKey = apiKey
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
	return len(c.identity.ApiKey) > 0
}

func (c *Client) request(uri string, method string, data interface{}, result interface{}) error {
	url := c.endpoint + "/v3" + uri
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.identity.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Printf("HTTP code %d\n", res.StatusCode)
		return err
	}
	return nil
}
