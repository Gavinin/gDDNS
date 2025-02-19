package _default

import (
	"errors"
	"fmt"
	"gDDNS/internal/local"
	"gDDNS/internal/log"
	"gDDNS/internal/util"
	"io"
	"net/http"
	"strings"
)

type Default struct {
	Addr string
}

func NewDefault(addr string) *Default {
	if !strings.HasPrefix(addr, "http") {
		addr = fmt.Sprintf("http://%s", addr)
	}
	return &Default{
		Addr: addr,
	}
}

func (d Default) Query() (*local.IP, error) {
	request, err := http.NewRequest("GET", d.Addr, nil)
	request.Header.Set("Content-Type", "text/plain")
	request.Header.Set("Content-Length", "16")
	request.Header.Set("User-Agent", "curl/8.7.1")
	if err != nil {
		log.Log.Errorf("[Default] create request address (%s) error: %v", d.Addr, err)
		return nil, err
	}
	response, err := util.GetClient().Do(request)
	if err != nil {
		log.Log.Errorf("request err: %v", err)
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Log.Errorf("response status code: %d", response.StatusCode)
		return nil, errors.New(response.Status)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Log.Errorf("read response body fail: %v", err)
		return nil, err
	}
	ip := string(body)
	ip = strings.TrimSuffix(ip, "\n")
	log.Log.Debugf("Get IP (%s) from %s ", ip, d.Addr)
	return &local.IP{IPv4: ip}, nil

}
