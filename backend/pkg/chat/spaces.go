package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bear-san/googlechat-sender/backend/ent"
	"io"
	"net/http"
	"net/url"
)

func GetOneSpace(token *ent.GoogleApiKey, name string) (*Space, error) {
	var result Space

	u, _ := url.Parse(fmt.Sprintf("https://chat.googleapis.com/v1/%s", name))
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	cl := new(http.Client)

	res, err := cl.Do(req)
	if err != nil {
		return nil, err
	}

	t, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(t, &result)
	return &result, err
}

func GetSpaces(token *ent.GoogleApiKey) (*[]Space, error) {
	lst := make([]Space, 0)
	u, _ := url.Parse("https://chat.googleapis.com/v1/spaces")

	pageToken := ""
	for {
		if pageToken != "" {
			q := u.Query()
			q.Set("pageToken", pageToken)

			u.RawQuery = q.Encode()
		}

		req, err := http.NewRequest("GET", u.String(), nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
		client := new(http.Client)
		res, err := client.Do(req)

		resTxt, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var result SpaceList
		err = json.Unmarshal(resTxt, &result)

		lst = append(lst, result.Spaces...)
		if result.NextPageToken != nil {
			pageToken = *result.NextPageToken
			continue
		}

		break
	}

	return &lst, nil
}

func (s *Space) Post(cred *ent.GoogleApiKey, msg Message) (*Message, error) {
	u, err := url.Parse(fmt.Sprintf("https://chat.googleapis.com/v1/%s/messages", s.Name))
	if err != nil {
		return nil, err
	}

	body, _ := json.Marshal(msg)
	req, _ := http.NewRequest("POST", u.String(), bytes.NewBuffer(body))

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cred.AccessToken))
	client := new(http.Client)
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	var result Message
	resTxt, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resTxt, &result)

	return &result, err
}

type Message struct {
	Text string `json:"text"`
}

type Space struct {
	Name                string `json:"name"`
	Type                string `json:"type"`
	DisplayName         string `json:"displayName,omitempty"`
	SpaceThreadingState string `json:"spaceThreadingState"`
	SpaceType           string `json:"spaceType"`
	SpaceHistoryState   string `json:"spaceHistoryState"`
	SpaceDetails        struct {
		Guidelines  string `json:"guidelines,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"spaceDetails,omitempty"`
	SingleUserBotDm bool `json:"singleUserBotDm,omitempty"`
	Threaded        bool `json:"threaded,omitempty"`
}

type SpaceList struct {
	Spaces        []Space `json:"spaces"`
	NextPageToken *string `json:"nextPageToken,omitempty"`
}
