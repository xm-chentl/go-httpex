package httpex

import (
	"fmt"
	"reflect"
)

// HandleFunc 处理函数
type HandleFunc func(url string, requestData, responseData interface{}) error

// BaseHTTP 基础请求类
type BaseHTTP struct {
	url         string
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

// Send 发送请求
func (b *BaseHTTP) Send(responseData interface{}) error {
	defer b.Reset()
	if reflect.TypeOf(responseData).Kind() != reflect.Ptr {
		return fmt.Errorf("receive parameter responseData must ptr")
	}

	return b.HandleFunc(b.url, b.requestData, responseData)
}

// Reset 重置参数
func (b *BaseHTTP) Reset() {
	b.url = ""
	b.requestData = nil
}
