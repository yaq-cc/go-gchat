package objects

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type ChatRequest struct {
	Type                      string    `json:"type"`
	EventTime                 time.Time `json:"eventTime"`
	Message                   Message   `json:"message"`
	User                      User      `json:"user"`
	Space                     Space     `json:"space"`
	ConfigCompleteRedirectURL string    `json:"configCompleteRedirectUrl"`
}

func (cr *ChatRequest) FromJSONReader(r io.Reader) error {
	err := json.NewDecoder(r).Decode(cr)
	if err != nil {
		return err
	}
	return nil
}

func (cr *ChatRequest) FromHTTPRequest(r *http.Request) error {
	err := cr.FromJSONReader(r.Body)
	if err != nil {
		return err
	}
	return nil
}

func FromHTTPRequest(r *http.Request) (*ChatRequest, error) {
	var cr *ChatRequest
	err := cr.FromHTTPRequest(r)
	if err != nil {
		return nil, err
	}
	return cr, nil
}

type Sender struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	AvatarURL   string `json:"avatarUrl"`
	Email       string `json:"email"`
	Type        string `json:"type"`
	DomainID    string `json:"domainId"`
}

type UserMention struct {
	User User   `json:"user"`
	Type string `json:"type"`
}

type Annotations struct {
	Type         string       `json:"type"`
	StartIndex   int          `json:"startIndex"`
	Length       int          `json:"length"`
	UserMention  UserMention  `json:"userMention"`
	SlashCommand SlashCommand `json:"slashCommand,omitempty"`
}

type SlashCommand struct {
	Type      string `json:"type,omitempty"`
	User      User   `json:"user,omitempty"`
	CommandId string `json:"commandId,omitempty"`
}

type RetentionSettings struct {
	State string `json:"state"`
}

type Thread struct {
	Name              string            `json:"name"`
	RetentionSettings RetentionSettings `json:"retentionSettings"`
}

type Space struct {
	Name                string `json:"name"`
	Type                string `json:"type"`
	DisplayName         string `json:"displayName"`
	ExternalUserAllowed bool   `json:"externalUserAllowed"`
	SpaceThreadingState string `json:"spaceThreadingState"`
	SpaceType           string `json:"spaceType"`
	SpaceHistoryState   string `json:"spaceHistoryState"`
}

type Message struct {
	Name           string        `json:"name"`
	Sender         Sender        `json:"sender"`
	CreateTime     time.Time     `json:"createTime"`
	Text           string        `json:"text"`
	Annotations    []Annotations `json:"annotations"`
	Thread         Thread        `json:"thread"`
	Space          Space         `json:"space"`
	ArgumentText   string        `json:"argumentText"`
	LastUpdateTime time.Time     `json:"lastUpdateTime"`
	SlashCommand   SlashCommand  `json:"slashCommand"`
}

type User struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	AvatarURL   string `json:"avatarUrl"`
	Email       string `json:"email"`
	Type        string `json:"type"`
	DomainID    string `json:"domainId"`
}
