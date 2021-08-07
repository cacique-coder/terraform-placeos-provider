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

type Module struct {
	CreatedAt       int64  `json:"created_at"`
	UpdatedAt       int64  `json:"updated_at"`
	Ip              string `json:"ip"`
	Port            int    `json:"port"`
	Tls             bool   `json:"tls"`
	Udp             bool   `json:"udp"`
	Makebreak       bool   `json:"makebreak"`
	Uri             string `json:"uri"`
	Name            string `json:"name"`
	CustomName      string `json:"custom_name"`
	Role            int    `json:"role"`
	Connected       bool   `json:"connected"`
	Running         bool   `json:"running"`
	Notes           string `json:"notes"`
	IgnoreConnected bool   `json:"ignore_connected"`
	IgnoreStartStop bool   `json:"ignore_startstop"`
	DriverId        string `json:"driver_id"`
	Id              string `json:"id"`
}

func (client *Client) getModule(id string) (Module, error) {
	var module Module
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/engine/v2/modules/%s", client.Host, id), nil)

	if err != nil {
		return module, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Token.AccessToken))
	c := &http.Client{Timeout: 100 * time.Second, Transport: tr}
	jsonString, err := getJsonString(req, c)

	if err != nil {
		return module, err
	}

	json.Unmarshal([]byte(jsonString), &module)

	return module, nil
}

func (client *Client) createModule(ip string, driverId string, name string, uri string, port int, tlsModule bool, udp bool, makebreak bool, customName string, notes string, ignore_connected bool) (Module, error) {
	var module = Module{
		Name:       name,
		Uri:        uri,
		Port:       port,
		Tls:        tlsModule,
		Udp:        udp,
		Makebreak:  makebreak,
		CustomName: customName,
		Notes:      notes,
		DriverId:   driverId,
	}

	// get json from driver struct
	postBody, _ := json.Marshal(module)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/engine/v2/modules", client.Host), bytes.NewBuffer(postBody))
	if err != nil {
		return module, err
	}

	if err != nil {
		return module, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Token.AccessToken))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	c := &http.Client{Timeout: 100 * time.Second, Transport: tr}
	jsonString, err := getJsonString(req, c)

	if err != nil {
		return module, err
	}

	json.Unmarshal([]byte(jsonString), &module)

	return module, nil
}

// updates a driver in placeos when the parameter is the driver instance
func (client *Client) updateModule(moduleParams Module) (Module, error) {
	var module = Module{
		Name:            moduleParams.Name,
		Uri:             moduleParams.Uri,
		Port:            moduleParams.Port,
		Tls:             moduleParams.Tls,
		Udp:             moduleParams.Udp,
		Makebreak:       moduleParams.Makebreak,
		Notes:           moduleParams.Notes,
		IgnoreConnected: moduleParams.IgnoreConnected,
		DriverId:        moduleParams.DriverId,
		CustomName:      moduleParams.CustomName,
	}

	file, err := os.Create("/tmp/module.json")

	postBody, _ := json.Marshal(module)
	file.Write(postBody)

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/api/engine/v2/modules/%s", client.Host, moduleParams.Id), bytes.NewBuffer(postBody))

	if err != nil {
		return module, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.Token.AccessToken))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	c := &http.Client{Timeout: 100 * time.Second, Transport: tr}
	jsonString, err := getJsonString(req, c)
	file.Write(jsonString)

	file.Close()
	if err != nil {
		return module, err
	}

	json.Unmarshal([]byte(jsonString), &module)

	return module, nil
}

// deletes a driver from placeos
func (client *Client) deleteModule(id string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/engine/v2/modules/%s", client.Host, id), nil)

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
