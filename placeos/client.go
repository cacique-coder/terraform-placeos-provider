package placeos

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type AccessToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string
	CreatedAt    int64 `json:"created_at"`
}

type Client struct {
	Username     string
	Password     string
	Host         string
	InsecureSsl  bool
	Token        AccessToken
	ClientId     string
	ClientSecret string
}

func NewBasicAuthClient(username string, password string, host string, insecureSsl bool, clientId string, clientSecret string) *Client {
	client := Client{
		Username:     username,
		Password:     password,
		Host:         host,
		InsecureSsl:  insecureSsl,
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}
	client.authorize()
	return &client
}

func (client *Client) authorize() (bool, error) {

	postBody, _ := json.Marshal(map[string]string{
		"grant_type": "password",
		"username":   client.Username,
		"password":   client.Password,
		"scope":      "public",
	})

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/auth/oauth/token", client.Host), bytes.NewBuffer(postBody))
	if err != nil {
		return false, err
	}

	headerAuthorization := fmt.Sprintf("%s:%s", client.ClientId, client.ClientSecret)
	encoded := base64.StdEncoding.EncodeToString([]byte(headerAuthorization))
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encoded))
	req.Header.Add("Content-Type", "application/json")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	c := &http.Client{Timeout: 100 * time.Second, Transport: tr}

	r, err := c.Do(req)
	if err != nil {
		return false, err
	}

	var accessToken AccessToken
	w, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return false, err
	}

	json.Unmarshal(w, &accessToken)
	client.Token = accessToken
	return true, nil
}

func getJsonString(req *http.Request, c *http.Client) ([]byte, error) {
	r, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	w, err := ioutil.ReadAll(r.Body)

	return w, nil
}
