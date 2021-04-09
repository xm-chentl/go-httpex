package httpex

import (
	"fmt"
	"net/http"
	"reflect"
)

// HandleFunc 处理函数
type HandleFunc func(url string, requestData, responseData interface{}) error

// BaseHTTP 基础请求类
type BaseHTTP struct {
	url         string
	method      string
	requestData interface{}
	HandleFunc  HandleFunc
}

// SetURL 设置请求地址
func (b *BaseHTTP) SetURL(url string) IHttp {
	b.url = url
	return b
}

// SetBody 设置请求的参数
func (b *BaseHTTP) SetBody(requestData interface{}) IHttp {
	b.requestData = requestData
	return b
}

// SetMethod 设置请求方式
func (b *BaseHTTP) SetMethod(method string) IHttp {
	b.method = method
	return b
}

// Send 发送请求
func (b *BaseHTTP) Send(responseData interface{}) error {
	defer b.Reset()
	if reflect.TypeOf(responseData).Kind() != reflect.Ptr {
		return fmt.Errorf("receive parameter responseData must ptr")
	}
	if b.method == "" {
		// 默认post
		b.method = http.MethodPost
	}

	return b.HandleFunc(b.method, b.url, b.requestData, responseData)
}

// Reset 重置参数
func (b *BaseHTTP) Reset() {
	b.url = ""
	b.requestData = nil
}
