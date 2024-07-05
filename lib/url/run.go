package url

import (
	"net/http"
	"strings"
	"time"
)

// RunFor 尝试三次
func RunFor(c http.Client, req *http.Request) (resp *http.Response, err error) {
	deadline := time.Now().Add(30 * time.Second)
	for i := 0; i < 3; i++ {
		if time.Now().After(deadline) {
			break
		}
		if resp, err = c.Do(req); err == nil {
			return resp, nil
		}
		if strings.Contains(err.Error(), "no such host") {
			break
		}
		if strings.Contains(err.Error(), "timeout") {
			break
		}
	}
	// log.Println(err)
	return nil, err
}
