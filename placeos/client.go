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

type Repository struct {
	CreatedAt   int64 `json:"created_at"`
	UpdatedAt   int64 `json:"updated_at"`
	Id          string
	Name        string
	Description string
	FolderName  string `json:"folder_name"`
	Uri         string
	CommitHash  string `json:"commit_hash"`
	Branch      string
	Username    string
	Password    string
	RepoType    string `json:"repo_type"`
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

	c := &http.Client{Timeout: 10 * time.Second, Transport: tr}

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

func (client *Client) getRepositories() ([]Repository, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/engine/v2/repositories", client.Host), nil)

	if err != nil {
		return nil, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Token.AccessToken))
	c := &http.Client{Timeout: 10 * time.Second, Transport: tr}
	jsonString, err := getJsonString(req, c)

	if err != nil {
		return nil, err
	}

	var repositories []Repository
	json.Unmarshal([]byte(jsonString), &repositories)

	return repositories, nil
}

func (client *Client) getRepository(id string) (Repository, error) {
	var repository Repository
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/engine/v2/repositories/%s", client.Host, id), nil)

	if err != nil {
		return repository, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Token.AccessToken))
	c := &http.Client{Timeout: 10 * time.Second, Transport: tr}
	jsonString, err := getJsonString(req, c)

	if err != nil {
		return repository, err
	}

	json.Unmarshal([]byte(jsonString), &repository)

	return repository, nil
}

func (client *Client) createRepository(name string, folder_name string, uri string, repo_type string, description string, branch string, username string, password string) (Repository, error) {
	var repository Repository

	postBody, _ := json.Marshal(map[string]string{
		"name":        name,
		"folder_name": folder_name,
		"uri":         uri,
		"repo_type":   repo_type,
		"description": description,
		"branch":      branch,
		"username":    username,
		"password":    password,
	})

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/engine/v2/repositories", client.Host), bytes.NewBuffer(postBody))

	if err != nil {
		return repository, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Token.AccessToken))
	c := &http.Client{Timeout: 10 * time.Second, Transport: tr}
	jsonString, err := getJsonString(req, c)

	if err != nil {
		return repository, err
	}

	json.Unmarshal([]byte(jsonString), &repository)

	return repository, nil
}

func (client *Client) deleteRepository(id string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/engine/v2/repositories/%s", client.Host, id), nil)

	if err != nil {
		return err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Token.AccessToken))
	c := &http.Client{Timeout: 10 * time.Second, Transport: tr}
	r, err := c.Do(req)

	if err != nil {
		return err
	}
	defer r.Body.Close()

	return nil
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
