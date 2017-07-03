package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Config struct {
	ActiveMemberCount     int64        `json:"activeMemberCount,omitempty"`
	AuthTokens            []string     `json:"authTokens,omitempty"`
	AuthorizedMemberCount int64        `json:"authorizedMemberCount,omitempty"`
	Capabilities          []string     `json:"capabilities,omitempty"`
	Clock                 int64        `json:"clock,omitempty"`
	CreationTime          int64        `json:"creationTime,omitempty"`
	ID                    string       `json:"id,omitempty"`
	LastModified          int64        `json:"lastModified,omitempty"`
	MulticastLimit        int64        `json:"multicastLimit,omitempty"`
	Name                  string       `json:"name,omitempty"`
	Nwid                  string       `json:"nwid,omitempty"`
	Objtype               string       `json:"objtype,omitempty"`
	Private               bool         `json:"private,omitempty"`
	Revision              int64        `json:"revision,omitempty"`
	Routes                []Routes     `json:"routes,omitempty"`
	Rules                 []Rules      `json:"rules,omitempty"`
	Tags                  []string     `json:"tags,omitempty"`
	TotalMemberCount      int64        `json:"totalMemberCount,omitempty"`
	V4AssignMode          V4AssignMode `json:"v4AssignMode,omitempty"`
	V6AssignMode          V6AssignMode `json:"v6AssignMode,omitempty"`
}

type Routes struct {
	Target string `json:"target,omitempty"`
	Via    string `json:"via,omitempty"`
}

type Rules struct {
	EtherType int64  `json:"ethertype,omitempty"`
	Not       bool   `json:"not,omitempty"`
	Or        bool   `json:"or,omitempty"`
	Type      string `json:"type,omitempty"`
}

type V4AssignMode struct {
	Properties Properties `json:"properties,omitempty"`
}

type V6AssignMode struct {
	Properties Properties `json:"properties,omitempty"`
}

type Network struct {
	ID                string      `json:"id,omitempty"`
	Type              string      `json:"type,omitempty"`
	Clock             int64       `json:"clock,omitempty"`
	UI                UI          `json:"ui,omitempty"`
	Config            Config      `json:"config,omitempty"`
	Description       string      `json:"description,omitempty"`
	OnlineMemberCount int64       `json:"onlineMemberCount,omitempty"`
	Permissions       Permissions `json:"permissions,omitempty"`
	RulesSource       string      `json:"rulesSource,omitempty"`
	TagsByName        TagsByName  `json:"tagsByName,omitempty"`
}

type UI struct {
	Properties Properties `json:"properties,omitempty"`
}

type TagsByName struct {
	Properties Properties `json:"properties,omitempty"`
}

type Properties struct{}

type Permissions struct {
	ID ID `json:"{id},omitempty"`
}

type ID struct {
	A bool `json:"a,omitempty"`
	D bool `json:"d,omitempty"`
	M bool `json:"m,omitempty"`
	O bool `json:"o,omitempty"`
	R bool `json:"r,omitempty"`
	T bool `json:"t,omitempty"`
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

// Updates a network or creates a new network if no network ID is provided
func (ns *NetworkService) Update(b []byte) (*Network, *http.Response, error) {
	var d map[string]interface{}

	if err := json.Unmarshal(b, &d); err != nil {
		panic(err)
	}

	path := fmt.Sprintf("network/%s", d["id"])

	req, err := ns.client.NewRequest("POST", path, &d)
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

func (ns *NetworkService) Delete(id string) (*http.Response, error) {
	path := fmt.Sprintf("network/%s", id)

	req, err := ns.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := ns.client.Do(req, nil)
	if err != nil {
		switch err.(type) {
		case *Error:
			if err.(*Error).Message == "network not found" {
				return resp, err
			}
		}
		return resp, err
	}
	return resp, nil
}
