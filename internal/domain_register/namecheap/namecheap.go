package namecheap

import (
	"gDDNS/enum"
	"gDDNS/internal/dns"
	"gDDNS/internal/log"
	"gDDNS/internal/util"
	"io"
	"net/http"
	"net/url"
)

type NameCheap struct {
	Key            string
	TopLevelDomain string
}

func NewNameCheap(key string) *NameCheap {
	return &NameCheap{
		Key: key,
	}
}

func (n NameCheap) ListDomains() ([]dns.Record, error) {
	return nil, enum.UnSupportErr
}

func (n NameCheap) Update(record dns.Record) error {

	baseUrl, err := url.Parse("https://dynamicdns.park-your-domain.com/update")
	if err != nil {
		return err
	}
	params := url.Values{}
	params.Set("domain", n.TopLevelDomain)
	params.Set("password", n.Key)
	params.Set("host", record.Domain)
	params.Set("ip", record.IP.IPv4)
	baseUrl.RawQuery = params.Encode()

	putReq, _ := http.NewRequest("GET", baseUrl.String(), nil)

	resp, err := util.GetClient().Do(putReq)
	if err != nil {
		log.Log.Error(err)
		return err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Log.Error(err)
		return err
	}
	respStr := string(respBody)
	log.Log.Debug(respStr)
	return nil
}

func (n NameCheap) PUT(record dns.Record) error {
	return enum.UnSupportErr
}

func (n NameCheap) Delete(record dns.Record) error {
	return enum.UnSupportErr
}
