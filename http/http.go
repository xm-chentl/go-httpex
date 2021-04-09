package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/xm-chentl/go-httpex"
)

// New 实例
func New() httpex.IHttp {
	return &httpex.BaseHTTP{
		HandleFunc: func(method, url string, requestData, responseData interface{}) error {
			var err error
			bodyBytes, err := json.Marshal(requestData)
			if err != nil {
				return fmt.Errorf("requestData to byte[] is fail, err(%v)", err)
			}

			req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyBytes))
			if err != nil {
				return err
			}

			req.Header.Add("content-type", httpex.ContentType)
			client := &http.Client{Timeout: 5 * time.Second}
			resp, err := client.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			respBodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			if err = json.Unmarshal(respBodyBytes, responseData); err != nil {
				return err
			}

			return nil
		},
	}
}
