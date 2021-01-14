package mock

import (
	"encoding/json"

	"github.com/xm-chentl/go-httpex"
)

// New 实例一个mock实例
func New(callback func(url string, requestData interface{}) (interface{}, error)) httpex.IHttp {
	return &httpex.BaseHTTP{
		HandleFunc: func(url string, requestData, responseData interface{}) error {
			respData, err := callback(url, requestData)
			if err != nil {
				return err
			}

			// 序列化，使用mock者无需要关注
			respDataByte, _ := json.Marshal(respData)
			json.Unmarshal(respDataByte, responseData)
			return nil
		},
	}
}
