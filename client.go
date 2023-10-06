package goly

import (
	"errors"
	"net/url"
)

func (c *Client) CreateRequest(data url.Values) (*Request, error) {
	var response Request
	re := url.Values{}
	for k, v := range data {
		if len(v) > 0 {
			re.Set(k, v[0])
		} else {
			re.Set(k, "")
		}
	}
	re.Set("vendor_token", c.identity.PublicVendorToken)

	err := c.request("/request/do.json", re, &response)
	if err != nil {
		return nil, err
	}
	if response.Error != "0" {
		return nil, errors.New(response.Message)
	}
	return &response, nil
}

func (c *Client) GetRequestState(id string) (*RequestState, error) {
	var res RequestState
	var re = url.Values{}
	re.Set("request_id", id)
	re.Set("vendor_token", c.identity.PublicVendorToken)
	err := c.request("/request/state.json", re, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
func (c *Client) GetRequestStateByOrderRef(orderRef string) (*RequestState, error) {
	var res RequestState
	var re = url.Values{}
	re.Set("order_ref", orderRef)
	re.Set("vendor_token", c.identity.PublicVendorToken)
	err := c.request("/request/state.json", re, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
func (c *Client) BuildSignature(data *map[string]string) (string, error) {
	if len(c.identity.PrivateVendorToken) == 0 {
		return "", errors.New("no private vendor token specified")
	}
	return buildSignature(data, c.identity.PrivateVendorToken), nil
}
func (c *Client) GetB2CBalance() (*map[string]float64, error) {
	var res struct {
		Error   int                `json:"error,omitempty"`
		Balance map[string]float64 `json:"balance,omitempty"`
	}
	var re = url.Values{}
	re.Set("vendor_token", c.identity.PublicVendorToken)
	signature, err := c.BuildSignature(&map[string]string{"vendor_token": c.identity.PublicVendorToken})
	re.Set("signature", signature)
	err = c.request("/business/b2cbalance.json", re, &res)
	if err != nil {
		return nil, err
	}
	if res.Error != 0 {
		return nil, errors.New("unable to retrieve b2c balance")
	}
	if nil == res.Balance || len(res.Balance) == 0 {
		return nil, errors.New("unable to retrieve b2c balance")
	}

	return &res.Balance, nil
}
