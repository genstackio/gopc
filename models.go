package gopc

type Request struct {
	Error       string `json:"error"`
	RequestId   string `json:"request_id,omitempty"`
	RequestUuid string `json:"request_uuid,omitempty"`
	Message     string `json:"message"`
	MobileUrl   string `json:"mobile_url,omitempty"`
	OrderRef    string `json:"order_ref,omitempty"`
}
type RequestState struct {
	State             string `json:"state"`
	UsedEaseOfPayment string `json:"used_ease_of_payment"`
}

func (o *RequestState) IsPaid() bool {
	return o.State == "1"
}
func (o *RequestState) IsNotPaid() bool {
	return !o.IsPaid()
}
