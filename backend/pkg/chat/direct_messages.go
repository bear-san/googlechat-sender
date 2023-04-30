package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bear-san/googlechat-sender/backend/ent"
	"io"
	"net/http"
)

func CreateDirectMessage(token *ent.GoogleApiKey, googleUserId string) (*Space, error) {
	spaceType := "DIRECT_MESSAGE"
	spaceThreadingState := "UNTHREADED_MESSAGES"

	memberships := []Membership{
		{
			Member: &Member{
				Name: fmt.Sprintf("users/%s", googleUserId),
				Type: "HUMAN",
			},
		},
	}

	reqBody := CreateDirectMessageRequest{
		Space: Space{
			SpaceType:           &spaceType,
			SpaceThreadingState: &spaceThreadingState,
		},
		Memberships: &memberships,
	}

	reqTxt, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://chat.googleapis.com/v1/spaces:setup", bytes.NewBuffer(reqTxt))
	if err != nil {
		return nil, err
	}

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	resTxt, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var space Space
	err = json.Unmarshal(resTxt, &space)

	return &space, err
}

func FindDirectMessage(token *ent.GoogleApiKey, googleUserId string) (*Space, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("https://chat.googleapis.com/v1/spaces:findDirectMessage?name=users/%s", googleUserId),
		nil,
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	client := new(http.Client)
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	resTxt, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var space Space
	err = json.Unmarshal(resTxt, &space)

	return &space, err
}

type Membership struct {
	Name   *string `json:"name,omitempty"`
	Member *Member `json:"member,omitempty"`
}

type Member struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type CreateDirectMessageRequest struct {
	Space       Space         `json:"space"`
	Memberships *[]Membership `json:"memberships,omitempty"`
}
