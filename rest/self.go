package rest

import (
	"fmt"
	"net/http"
)

type User struct {
	ID                string `json:"id"`
	Type              string `json:"type"`
	Clock             int64  `json:"clock"`
	GlobalPermissions struct {
		A  bool `json:"a"`
		D  bool `json:"d"`
		Da bool `json:"da"`
		M  bool `json:"m"`
		R  bool `json:"r"`
	} `json:"globalPermissions"`
	UI struct {
	} `json:"ui"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	Auth        struct {
		Local string `json:"local"`
	} `json:"auth"`
	SmsNumber   string   `json:"smsNumber"`
	Tokens      []string `json:"tokens"`
	Permissions struct {
	} `json:"permissions"`
	Subscriptions struct {
	} `json:"subscriptions"`
}

// SelfService handles returns self endpoint
type SelfService service

func (s *SelfService) Get() (*User, *http.Response, error) {
	path := fmt.Sprintf("self")

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var u User
	resp, err := s.client.Do(req, &u)
	if err != nil {
		return nil, resp, err
	}

	return &u, resp, nil
}
