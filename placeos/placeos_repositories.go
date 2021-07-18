package placeos

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

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

type Commit struct {
	Commit  string `json:"commit"`
	Date    string `json:"date"`
	Author  string `json:"author"`
	Subject string `json:"subject"`
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

func (client *Client) updateRepository(repository Repository) (Repository, error) {
	var repositoryNew Repository

	postBody, _ := json.Marshal(map[string]string{
		"name":        repository.Name,
		"folder_name": repository.FolderName,
		"uri":         repository.Uri,
		"repo_type":   repository.RepoType,
		"description": repository.Description,
		"branch":      repository.Branch,
		"username":    repository.Username,
		"password":    repository.Password,
	})

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/api/engine/v2/repositories/%s", client.Host, repository.Id), bytes.NewBuffer(postBody))

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

	json.Unmarshal([]byte(jsonString), &repositoryNew)

	return repositoryNew, nil
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

// Pulling last commit hash from a repository
func (client *Client) getLastCommitHash(repository_id string, driver_file_name string) (string, error) {
	// make a get request to commits in a repository with a query string
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/engine/v2/repositories/%s/commits?driver=%s", client.Host, repository_id, driver_file_name), nil)

	if err != nil {
		return "", err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Token.AccessToken))
	c := &http.Client{Timeout: 10 * time.Second, Transport: tr}
	jsonString, err := getJsonString(req, c)

	if err != nil {
		return "", err
	}

	// Parsing json into and array of commits
	var commits []Commit
	json.Unmarshal([]byte(jsonString), &commits)
	// Get last commit hash
	return commits[0].Commit, nil
}
