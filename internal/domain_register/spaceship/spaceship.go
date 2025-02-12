package spaceship

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gDDNS/internal/dns"
	"io"
	"log"
	"net/http"
	"time"
)

type SpaceShip struct {
	Key            string
	Secret         string
	TopLevelDomain string
	c              *http.Client
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

type getGroup struct {
	Type string `json:"type"`
}

type getItem struct {
	Address string   `json:"address"`
	Name    string   `json:"name"`
	Type    string   `json:"type"`
	TTL     int      `json:"ttl"`
	Group   getGroup `json:"group"`
}

type getResponse struct {
	Items []getItem `json:"items"`
	Total int       `json:"total"`
}

func NewSpaceShip() *SpaceShip {
	tr := &http.Transport{
		MaxIdleConns:       3,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{Transport: tr}
	return &SpaceShip{
		c: client,
	}
}

func (s SpaceShip) ListDomains() ([]dns.Record, error) {
	putURL := fmt.Sprintf("https://spaceship.dev/api/v1/dns/records/%s?take=100&skip=0", s.TopLevelDomain)
	putReq, _ := http.NewRequest("GET", putURL, nil)
	putReq.Header.Set("X-API-Key", s.Key)
	putReq.Header.Set("X-API-Secret", s.Secret)
	resp, err := s.c.Do(putReq)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response *getResponse
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
	_, err = s.c.Do(putReq)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
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
	_, err = s.c.Do(putReq)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
