package placeos

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Zone struct {
	Name string `json:"name"`

	Description string `json:"description"`
	Tags        []string
	Location    string `json:"location"`
	DisplayName string `json:"display_name"`
	Code        string `json:"code"`
	Type        string `json:"type"`
	Count       int    `json:"count"`
	Capacity    int    `json:"capacity"`
	MapId       string `json:"map_id"`
	ParentId    string `json:"parent_id"`

	Id        string `json:"id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (client *Client) GetZone(id string) (Zone, error) {
	var zone Zone
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/engine/v2/zones/%s", client.Host, id), nil)

	if err != nil {
		return zone, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Token.AccessToken))
	c := &http.Client{Timeout: 10 * time.Second, Transport: tr}
	jsonString, err := getJsonString(req, c)

	if err != nil {
		return zone, err
	}

	json.Unmarshal([]byte(jsonString), &zone)

	return zone, nil
}

// Create zone with following parameters
// Name        string `json:"name"`
// Description string `json:"description"`
// Tags        []string
// Location    string `json:"location"`
// DisplayName string `json:"display_name"`
// Code        string `json:"code"`
// Type        string `json:"type"`
// Count       int    `json:"count"`
// Capacity    int    `json:"capacity"`
// MapId       string `json:"map_id"`
// ParentId    string `json:"parent_id"`

func (client *Client) CreateZone(name string, description string, tags []string, location string, displayName string, code string, typeZone string, count int, capacity int, mapId string, parentId string) (Zone, error) {
	var zone Zone
	postBody := map[string]interface{}{
		"name":         name,
		"description":  description,
		"tags":         tags,
		"location":     location,
		"display_name": displayName,
		"code":         code,
		"type":         typeZone,
		"count":        count,
		"capacity":     capacity,
		"map_id":       mapId,
		"parent_id":    parentId,
	}
	postBodyJson, _ := json.Marshal(postBody)
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/engine/v2/zones", client.Host), bytes.NewBuffer(postBodyJson))

	if err != nil {
		return zone, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Token.AccessToken))
	c := &http.Client{Timeout: 10 * time.Minute, Transport: tr}
	jsonString, err := getJsonString(req, c)

	if err != nil {
		return zone, err
	}

	json.Unmarshal([]byte(jsonString), &zone)

	return zone, nil
}

// updates a zone in placeos when the parameter is the zone instance
func (client *Client) UpdateZone(zone Zone) (Zone, error) {
	// get json from zone struct
	postBody, _ := json.Marshal(zone)

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/api/engine/v2/zones/%s", client.Host, zone.Id), bytes.NewBuffer(postBody))

	if err != nil {
		return zone, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Token.AccessToken))
	c := &http.Client{Timeout: 10 * time.Second, Transport: tr}
	jsonString, err := getJsonString(req, c)

	if err != nil {
		return zone, err
	}

	json.Unmarshal([]byte(jsonString), &zone)

	return zone, nil
}

// delete a zones in placeos
func (client *Client) deleteZone(id string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/engine/v2/zones/%s", client.Host, id), nil)

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
