package gopc

import (
	"errors"
	"net/url"
	"time"
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
	re.Set("vendor_token", c.identity.ApiKey)

	err := c.request("/payments/unknow", "GET", re, &response)
	if err != nil {
		return nil, err
	}
	if response.Error != "0" {
		return nil, errors.New(response.Message)
	}
	return &response, nil
}

type Payment struct {
	PaymentId   string    `json:"paymentId"`
	CreatedAt   time.Time `json:"createdAt"`
	ExpireAt    time.Time `json:"expireAt"`
	SucceededAt time.Time `json:"succeededAt"`
	Currency    string    `json:"currency"`
	Status      string    `json:"status"`
	Creditor    struct {
		ProfileId   string `json:"profileId"`
		MerchantId  string `json:"merchantId"`
		Name        string `json:"name"`
		Iban        string `json:"iban"`
		CallbackUrl string `json:"callbackUrl"`
	} `json:"creditor"`
	Debtor struct {
		Name string `json:"name"`
		Iban string `json:"iban"`
	} `json:"debtor"`
	Amount         int    `json:"amount"`
	TransferAmount int    `json:"transferAmount"`
	TippingAmount  int    `json:"tippingAmount"`
	TotalAmount    int    `json:"totalAmount"`
	Description    string `json:"description"`
	BulkId         string `json:"bulkId"`
	Links          struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Deeplink struct {
			Href string `json:"href"`
		} `json:"deeplink"`
		Qrcode struct {
			Href string `json:"href"`
		} `json:"qrcode"`
		Refund struct {
			Href string `json:"href"`
		} `json:"refund"`
	} `json:"_links"`
}

func (c *Client) GetPaymentDetails(paymentId string) (*Payment, error) {
	var result Payment
	err := c.request("/payments/"+paymentId, "get", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

//

func (c *Client) GetRequestState(id string) (*RequestState, error) {
	var res RequestState
	var re = url.Values{}
	re.Set("request_id", id)
	re.Set("vendor_token", c.identity.ApiKey)
	err := c.request("/request/state.json", "post", re, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
func (c *Client) GetRequestStateByOrderRef(orderRef string) (*RequestState, error) {
	var res RequestState
	var re = url.Values{}
	re.Set("order_ref", orderRef)
	re.Set("vendor_token", c.identity.ApiKey)
	err := c.request("/request/state.json", "post", re, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
func (c *Client) BuildSignature(data *map[string]string) (string, error) {
	if len(c.identity.ApiKey) == 0 {
		return "", errors.New("no api key specified")
	}
	return buildSignature(data, c.identity.ApiKey), nil
}
func (c *Client) GetB2CBalance() (*map[string]float64, error) {
	var res struct {
		Error   int                `json:"error,omitempty"`
		Balance map[string]float64 `json:"balance,omitempty"`
	}
	var re = url.Values{}
	re.Set("vendor_token", c.identity.ApiKey)
	signature, err := c.BuildSignature(&map[string]string{"vendor_token": c.identity.ApiKey})
	re.Set("signature", signature)
	err = c.request("/business/b2cbalance.json", "post", re, &res)
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
