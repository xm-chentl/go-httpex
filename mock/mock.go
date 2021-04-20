package mock

import (
	"encoding/json"

	"github.com/xm-chentl/go-httpex"
)

// New 实例一个mock实例
func New(callback func(method, url string, requestData interface{}, header map[string]string) (interface{}, error)) httpex.IHttp {
	return &httpex.BaseHTTP{
		HandleFunc: func(method, url string, requestData, responseData interface{}, header map[string]string) error {
			respData, err := callback(method, url, requestData, header)
			if err != nil {
				return err
			}

			// 序列化，使用mock者无需要关注
			respDataByte, _ := json.Marshal(respData)
			err = json.Unmarshal(respDataByte, responseData)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
