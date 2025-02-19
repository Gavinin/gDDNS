package spaceship

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gDDNS/enum"
	"gDDNS/internal/dns"
	"gDDNS/internal/util"
	"io"
	"log"
	"net/http"
)

type SpaceShip struct {
	Key            string
	Secret         string
	TopLevelDomain string
}

type updateReq struct {
	Force bool            `json:"force"`
	Items []updateReqItme `json:"items"`
}
type updateReqItme struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Address string `json:"address"`
	TTL     int    `json:"ttl"`
}

type deleteReq struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type group struct {
	Type string `json:"type"`
}

type item struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	TTL     int    `json:"ttl"`
	Group   group  `json:"group"`
}

type listDomainsResponse struct {
	Items []item `json:"items"`
	Total int    `json:"total"`
}

func NewSpaceShip() *SpaceShip {
	return &SpaceShip{}
}

func (s SpaceShip) ListDomains() ([]dns.Record, error) {
	putURL := fmt.Sprintf("https://spaceship.dev/api/v1/dns/records/%s?take=100&skip=0", s.TopLevelDomain)
	putReq, _ := http.NewRequest("GET", putURL, nil)
	putReq.Header.Set("X-API-Key", s.Key)
	putReq.Header.Set("X-API-Secret", s.Secret)
	resp, err := util.GetClient().Do(putReq)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response *listDomainsResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, err
	}
	records := make([]dns.Record, len(response.Items))
	for _, item := range response.Items {
		records = append(records, dns.Record{
			IP:         item.Address,
			Domain:     item.Name,
			RecordType: item.Type,
		})
	}
	return records, nil
}

func (s SpaceShip) Update(record dns.Record) error {
	return enum.UnSupportErr
}

func (s SpaceShip) Delete(record dns.Record) error {
	putURL := "https://spaceship.dev/api/v1/dns/records/" + s.TopLevelDomain
	req := []deleteReq{
		{
			Type:    record.RecordType,
			Name:    record.Domain,
			Address: record.IP,
		},
	}
	reqJSON, err := json.Marshal(req)
	if err != nil {
		return err
	}
	putReq, _ := http.NewRequest("DELETE", putURL, bytes.NewBuffer(reqJSON))
	putReq.Header.Set("X-API-Key", s.Key)
	putReq.Header.Set("X-API-Secret", s.Secret)
	putReq.Header.Set("Content-Type", "application/json")
	_, err = util.GetClient().Do(putReq)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (s SpaceShip) PUT(record dns.Record) error {
	putURL := "https://spaceship.dev/api/v1/dns/records/" + s.TopLevelDomain
	req := updateReq{
		Force: true,
		Items: []updateReqItme{
			{
				Type:    record.RecordType,
				Name:    record.Domain,
				Address: record.IP,
				TTL:     60,
			},
		},
	}
	reqJSON, err := json.Marshal(req)
	if err != nil {
		return err
	}
	putReq, _ := http.NewRequest("PUT", putURL, bytes.NewBuffer(reqJSON))
	putReq.Header.Set("X-API-Key", s.Key)
	putReq.Header.Set("X-API-Secret", s.Secret)
	putReq.Header.Set("Content-Type", "application/json")
	_, err = util.GetClient().Do(putReq)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (s SpaceShip) Name() string {
	return string(enum.SpaceShip)
}
