package placeos

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// Role values
// 99 : Logic
// 2  : service
// 3  : websocket
// 1  : device
type Driver struct {
	CreatedAt        int64  `json:"created_at"`
	UpdatedAt        int64  `json:"updated_at"`
	Id               string `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	FileName         string `json:"file_name"`
	DefaultUri       string `json:"default_uri"`
	Commit           string `json:"commit"`
	Role             int    `json:"role"`
	ModuleName       string `json:"module_name"`
	RepositoryId     string `json:"repository_id"`
	IgnoredConnected bool   `json:"ignore_connected"`
}

func (client *Client) getDriver(id string) (Driver, error) {
	var driver Driver
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/engine/v2/drivers/%s", client.Host, id), nil)

	if err != nil {
		return driver, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Token.AccessToken))
	c := &http.Client{Timeout: 100 * time.Second, Transport: tr}
	jsonString, err := getJsonString(req, c)

	if err != nil {
		return driver, err
	}

	json.Unmarshal([]byte(jsonString), &driver)

	return driver, nil
}

// create driver with driver parameters

func (client *Client) createDriver(name string, description string, file_name string, default_uri string, module_name string, repository_id string, commit string, role int, ignore_connected bool) (Driver, error) {
	var driver = Driver{
		Name:             name,
		Description:      description,
		FileName:         file_name,
		DefaultUri:       default_uri,
		Role:             role,
		ModuleName:       module_name,
		RepositoryId:     repository_id,
		Commit:           commit,
		IgnoredConnected: ignore_connected,
	}

	// get json from driver struct
	postBody, _ := json.Marshal(driver)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/engine/v2/drivers", client.Host), bytes.NewBuffer(postBody))

	if err != nil {
		return driver, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Token.AccessToken))
	c := &http.Client{Timeout: 300 * time.Second, Transport: tr}
	jsonString, err := getJsonString(req, c)

	if err != nil {
		return driver, err
	}

	json.Unmarshal([]byte(jsonString), &driver)

	return driver, nil
}

// updates a driver in placeos when the parameter is the driver instance
func (client *Client) updateDriver(id string, name string, description string, file_name string, default_uri string, module_name string, repository_id string, commit string, role int, ignore_connected bool, created_at int64, updated_at int64) error {
	var driver = Driver{
		Id:               id,
		Name:             name,
		Description:      description,
		FileName:         file_name,
		DefaultUri:       default_uri,
		Role:             role,
		ModuleName:       module_name,
		RepositoryId:     repository_id,
		Commit:           commit,
		IgnoredConnected: ignore_connected,
		CreatedAt:        created_at,
		UpdatedAt:        updated_at,
	}

	postBody, _ := json.Marshal(driver)

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/api/engine/v2/drivers/%s", client.Host, driver.Id), bytes.NewBuffer(postBody))

	if err != nil {
		return err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Token.AccessToken))
	c := &http.Client{Timeout: 100 * time.Second, Transport: tr}
	jsonString, err := getJsonString(req, c)

	// print json to a file
	file, err := os.Create("/tmp/patch.json")
	if err != nil {
		return err
	}
	file.Write(postBody)
	file.Close()

	if err != nil {
		return err
	}

	json.Unmarshal([]byte(jsonString), &driver)

	return nil
}

// deletes a driver from placeos
func (client *Client) deleteDriver(id string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/engine/v2/drivers/%s", client.Host, id), nil)

	if err != nil {
		return err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Token.AccessToken))
	c := &http.Client{Timeout: 100 * time.Second, Transport: tr}
	_, err = c.Do(req)

	if err != nil {
		return err
	}

	return nil
}
