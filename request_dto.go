package nyarakadb

type Request struct {
	Endpoint string
	QueryParams map[string]string
	Headers  map[string]string
	Body     map[string]interface{}
}