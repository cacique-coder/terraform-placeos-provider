package placeos

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type System struct {
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	CreatedAt          int64    `json:"created_at"`
	UpdatedAt          int64    `json:"updated_at"`
	Description        string   `json:"description"`
	Features           []string `json:"features"`
	Email              string   `json:"email"`
	Bookable           bool     `json:"bookable"`
	DisplayName        string   `json:"display_name"`
	Code               string   `json:"code"`
	Type               string   `json:"type"`
	Capacity           int64    `json:"capacity"`
	MapId              string   `json:"map_id"`
	Images             []string `json:"images"`
	Timezone           string   `json:"timezone"`
	SupportUrl         string   `json:"support_url"`
	Version            int64    `json:"version"`
	InstalledUiDevices int64    `json:"installed_ui_devices"`
	Zones              []string `json:"zones"`
	Modules            []string `json:"modules"`
}

func (client *Client) GetSystem(id string) (System, error) {
	var system System
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/engine/v2/systems/%s", client.Host, id), nil)

	if err != nil {
		return system, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Token.AccessToken))
	c := &http.Client{Timeout: 10 * time.Second, Transport: tr}
	jsonString, err := getJsonString(req, c)

	if err != nil {
		return system, err
	}

	json.Unmarshal([]byte(jsonString), &system)

	return system, nil
}

// create driver with driver parameters

func (client *Client) CreateSystem(name string, zoneIds []string, email string, displayName string, supportUrl string, installedUiDevices int64, capacity int64, bookable bool, description string, features []string, mapId string, modules []string, timezone string, code string, version int64, images []string) (System, error) {
	var system = System{
		Name:               name,
		Zones:              zoneIds,
		Email:              email,
		DisplayName:        displayName,
		SupportUrl:         supportUrl,
		InstalledUiDevices: installedUiDevices,
		Capacity:           capacity,
		Bookable:           bookable,
		Description:        description,
		Features:           features,
		MapId:              mapId,
		Modules:            modules,
		Timezone:           timezone,
	}

	// get json from system struct
	postBody, _ := json.Marshal(system)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/engine/v2/systems", client.Host), bytes.NewBuffer(postBody))

	if err != nil {
		return system, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Token.AccessToken))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	c := &http.Client{Timeout: 10 * time.Second, Transport: tr}
	jsonString, err := getJsonString(req, c)

	if err != nil {
		return system, err
	}

	json.Unmarshal([]byte(jsonString), &system)

	return system, nil
}

// updates a driver in placeos when the parameter is the driver instance
func (client *Client) UpdateSystem(system System) (System, error) {
	// get json from system struct
	postBody, _ := json.Marshal(system)

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/api/engine/v2/systems/%s", client.Host, system.Id), bytes.NewBuffer(postBody))

	if err != nil {
		return system, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Token.AccessToken))
	c := &http.Client{Timeout: 10 * time.Second, Transport: tr}
	jsonString, err := getJsonString(req, c)

	if err != nil {
		return system, err
	}

	json.Unmarshal([]byte(jsonString), &system)

	return system, nil
}

// delete a systems in placeos
func (client *Client) DeleteSystem(id string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/engine/v2/systems/%s", client.Host, id), nil)

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
