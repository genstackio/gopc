package gopc

type Client struct {
	endpoint string
	env      string
	identity ClientIdentity
	options  ClientOptions
}

type ClientOptions struct {
	MinExpirationDelay int64
}
type ClientIdentity struct {
	ApiKey string
}

type FetchOptions struct {
	Method  string
	Body    interface{}
	Headers map[string]string
	Options HttpOptions
}

type HttpOptions struct {
	Timeout int64
}
