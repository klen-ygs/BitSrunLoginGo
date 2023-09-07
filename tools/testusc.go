package tools

import (
	"net/http"
	"strings"
	"time"
)

func TestInUSCButLogout() bool {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	client.CheckRedirect = func(_ *http.Request, _ []*http.Request) error {
		return http.ErrUseLastResponse
	}

	resp, err := client.Get("http://bilibili.com")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode == 302 && strings.HasPrefix(resp.Header.Get("location"), "http://210.43.112.9") {
		return true
	}

	return false
}
