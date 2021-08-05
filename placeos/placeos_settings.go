package placeos

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Role values
// 99 : Logic
// 2  : service
// 3  : websocket
// 1  : device

// // {
//   "id": "sets-HjI~xeEP_Md",
//   "name": "",
//   "created_at": 1626781157,
//   "updated_at": 1626781157,
//   "version": 0,
//   "parent_id": "driver-HAvHh-3XZUs",
//   "settings_string": "q: 2",
//   "encryption_level": 0,
//   "keys": [
//     "q"
//   ]
// }
type Setting struct {
	Id              string   `json:"id"`
	Name            string   `json:"name"`
	CreatedAt       int64    `json:"created_at"`
	UpdatedAt       int64    `json:"updated_at"`
	Version         int64    `json:"version"`
	ParentId        string   `json:"parent_id"`
	ParentType      string   `json:"parent_type"`
	SettingsString  string   `json:"settings_string"`
	EncryptionLevel int      `json:"encryption_level"`
	Keys            []string `json:"keys"`
}

func (client *Client) getSetting(id string) (Setting, error) {
	var setting Setting
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/engine/v2/settings/%s", client.Host, id), nil)

	if err != nil {
		return setting, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Token.AccessToken))
	c := &http.Client{Timeout: 10 * time.Second, Transport: tr}
	jsonString, err := getJsonString(req, c)

	if err != nil {
		return setting, err
	}

	json.Unmarshal([]byte(jsonString), &setting)

	return setting, nil
}

// create driver with driver parameters

func (client *Client) createSetting(name string, parent_id string, parent_type string, settings_string string, encryption_level int, keys []string) (Setting, error) {
	var setting = Setting{
		Name:            name,
		ParentId:        parent_id,
		ParentType:      parent_type,
		SettingsString:  settings_string,
		EncryptionLevel: encryption_level,
		Keys:            keys,
	}

	// get json from setting struct
	postBody, _ := json.Marshal(setting)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/engine/v2/settings", client.Host), bytes.NewBuffer(postBody))

	if err != nil {
		return setting, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Token.AccessToken))
	c := &http.Client{Timeout: 10 * time.Second, Transport: tr}
	jsonString, err := getJsonString(req, c)

	if err != nil {
		return setting, err
	}

	json.Unmarshal([]byte(jsonString), &setting)

	return setting, nil
}

// updates a driver in placeos when the parameter is the driver instance
func (client *Client) updateSetting(setting Setting) (Setting, error) {
	// get json from setting struct
	postBody, _ := json.Marshal(setting)

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/api/engine/v2/settings/%s", client.Host, setting.Id), bytes.NewBuffer(postBody))

	if err != nil {
		return setting, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Token.AccessToken))
	c := &http.Client{Timeout: 10 * time.Second, Transport: tr}
	jsonString, err := getJsonString(req, c)

	if err != nil {
		return setting, err
	}

	json.Unmarshal([]byte(jsonString), &setting)

	return setting, nil
}

// delete a settings in placeos
func (client *Client) deleteSetting(id string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/engine/v2/settings/%s", client.Host, id), nil)

	if err != nil {
		return err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Token.AccessToken))
	c := &http.Client{Timeout: 10 * time.Second, Transport: tr}
	_, err = getJsonString(req, c)

	if err != nil {
		return err
	}

	return nil
}
