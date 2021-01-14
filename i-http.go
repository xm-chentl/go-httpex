package httpex

// IHttp http 统一接口
type IHttp interface {
	SetURL(string) IHttp
	SetBody(interface{}) IHttp
	Send(interface{}) error
}
