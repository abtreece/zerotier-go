package rest

import (
	"fmt"
	"net/http"
)

type Network struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Clock int64  `json:"clock"`
	UI    struct {
		Properties struct{} `json:"properties"`
	} `json:"ui"`
	Config struct {
		ActiveMemberCount     int64    `json:"activeMemberCount"`
		AuthTokens            []string `json:"authTokens"`
		AuthorizedMemberCount int64    `json:"authorizedMemberCount"`
		Capabilities          []string `json:"capabilities"`
		Clock                 int64    `json:"clock"`
		CreationTime          int64    `json:"creationTime"`
		ID                    string   `json:"id"`
		LastModified          int64    `json:"lastModified"`
		MulticastLimit        int64    `json:"multicastLimit"`
		Name                  string   `json:"name"`
		Nwid                  string   `json:"nwid"`
		Objtype               string   `json:"objtype"`
		Private               bool     `json:"private"`
		Revision              int64    `json:"revision"`
		Routes                []struct {
			Target string `json:"target"`
			Via    string `json:"via"`
		} `json:"routes"`
		Rules []struct {
			EtherType int64  `json:"ethertype"`
			Not       bool   `json:"not"`
			Or        bool   `json:"or"`
			Type      string `json:"type"`
		} `json:"rules"`
		Tags             []string `json:"tags"`
		TotalMemberCount int64    `json:"totalMemberCount"`
		V4AssignMode     struct {
			Properties struct{} `json:"properties"`
		} `json:"v4AssignMode"`
		V6AssignMode struct {
			Properties struct{} `json:"properties"`
		} `json:"v6AssignMode"`
	} `json:"config"`
	Description       string `json:"description"`
	OnlineMemberCount int64  `json:onlineMemberCount`
	Permissions       struct {
		ID struct {
			A bool `json:"a"`
			D bool `json:"d"`
			M bool `json:"m"`
			O bool `json:"o"`
			R bool `json:"r"`
			T bool `json:"t"`
		} `json:"{id}"`
	} `json:"permissions"`
	RulesSource string `json:"rulesSource"`
	TagsByName  struct {
		Properties struct{} `json:"properties"`
	} `json:"tagsByName"`
}

// NetworkService handles network endpoint
type NetworkService service

func (ns *NetworkService) List() ([]*Network, *http.Response, error) {
	path := fmt.Sprintf("network")

	req, err := ns.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	nl := []*Network{}
	resp, err := ns.client.Do(req, &nl)
	if err != nil {
		return nil, resp, err
	}

	return nl, resp, nil
}

func (ns *NetworkService) Get(id string) (*Network, *http.Response, error) {
	path := fmt.Sprintf("network/%s", id)

	req, err := ns.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var n Network
	resp, err := ns.client.Do(req, &n)
	if err != nil {
		return nil, resp, err
	}

	return &n, resp, nil
}
