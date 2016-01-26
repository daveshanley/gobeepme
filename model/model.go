// Copyright 2015 Dave Shanley <dave@quobix.com>
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

// Package model provides data models for server resonses, devices and JSON mappings.
package model

import (
    "fmt"
    "errors"
    "strings"
)

type Creds struct {
    AppleID  string `json:"apple_id"`
    Password string `json:"password"`
}

type DeviceResult struct {
    StatusCode string `json:"statusCode"`
    Devices    []Device `json:"content"`
}

type Device struct {
    ID               string `json:"id"`
    BatteryLevel     float64 `json:"batteryLevel`
    BatteryStatus    string `json:"batteryStatus`
    Class            string `json:"deviceClass"`
    DisplayName      string `json:"deviceDisplayName"`
    Location         DeviceLocation `json:"location"`
    Model            string `json:"deviceModel"`
    ModelDisplayName string `json:"modelDisplayName"`
    Name             string `'json:"name"`
}

type DeviceLocation struct {
    Longitude float64 `json:"longitude"`
    Latitude  float64 `json:"latitude"`
}

type ServerCommand struct {
    DeviceID string `json:"device"`
    Message  string `json:"subject"`
}

type CloudService struct {
    Host  string
    Scope string
    Creds Creds
}

func (d *DeviceResult) GetDevice(id string) (*Device, error) {
    for _, r := range d.Devices {
        if r.ID == id {
            return &r, nil
        }
    }
    return nil, errors.New("No device found")
}

func (d *DeviceResult) GetDeviceByIndex(index int) (*Device, error) {
    i := 0
    for _, d := range d.Devices {
        if i >= index {
            return &d, nil
        }
        i++
    }
    return nil, fmt.Errorf("No Device with index [%d] located", index)
}

func (d *DeviceResult) GetDeviceByName(name string) (*Device, error) {
    for _, d := range d.Devices {
        if strings.ToLower(d.DisplayName) == strings.ToLower(name) {
            return &d, nil
        }
    }
    return nil, fmt.Errorf("No Device with name [%s] located", name)
}

func (d *DeviceResult) GetDeviceByDisplayName(dn string) *Device {
    for _, r := range d.Devices {
        if r.DisplayName == dn {
            return &r
        }
    }
    return nil
}
