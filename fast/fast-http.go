package fast

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/valyala/fasthttp"
	"github.com/xm-chentl/go-httpex"
)

// New 新建一个http实例
func New() httpex.IHttp {
	return &httpex.BaseHTTP{
		HandleFunc: func(method, url string, requestData, responseData interface{}) error {
			var err error
			req := fasthttp.AcquireRequest()
			defer fasthttp.ReleaseRequest(req)

			req.Header.SetContentType(httpex.ContentType)
			req.Header.SetMethod(method)
			req.SetRequestURI(url)
			requestDataOfByte, err := json.Marshal(requestData)
			if err != nil {
				return fmt.Errorf("fasthttp requestData to []byte fail err: %v", err)
			}
			req.SetBody(requestDataOfByte)

			resp := fasthttp.AcquireResponse()
			fasthttp.ReleaseResponse(resp)

			if err = fasthttp.Do(req, resp); err != nil {
				return fmt.Errorf("fasthttp post request fail err: %v", err)
			}
			if resp.StatusCode() != http.StatusOK {
				return fmt.Errorf("fasthttp request status fail")
			}

			if err = json.Unmarshal(resp.Body(), responseData); err != nil {
				return fmt.Errorf("[]byte to responseData fail, err: %v", err)
			}

			return nil
		},
	}
}
