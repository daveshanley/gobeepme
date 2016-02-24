// Copyright 2016 Dave Shanley <dave@quobix.com>
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

// Package model provides data models for server responses, devices and JSON mappings.
package model

import (
    "fmt"
    "strings"
)

type Creds struct {
    AppleID         string `json:"apple_id"`
    Password        string `json:"password"`
}

type ServiceCommand struct {
    Creds
    Message         string `json:"message"`
    Name            string `json:"name"`
}

type DeviceResult struct {
    StatusCode      string `json:"statusCode"`
    Devices         []Device `json:"content"`
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
    Longitude       float64 `json:"longitude"`
    Latitude        float64 `json:"latitude"`
}

type ServerCommand struct {
    DeviceID        string `json:"device"`
    Message         string `json:"subject"`
}

type CloudService struct {
    Host            string
    Scope           string
    Creds           Creds
}

type ServiceResponse struct {
    Error           bool `json:"error"`
    Message         string `json:"message"`
}

const (
    AuthFailedMessage   string = "Authentication failed: %v"
    ListDevicesFailed   string = "Unable to list iOS devices: %v"
    NoCredentials       string = "No authentication credentials were submitted"
    CommandMalformed    string = "Server command malformed"
    CommandMissingAttr  string = "Server command missing authentication, or iOS device name"
    NoDeviceName        string = "Can't Beep! No device with name [%s] located"
    NoDeviceIndex       string = "Can't Beep! No device with index [%d] located"
    NoDeviceID          string = "Can't Beep! No device with id [%s] located"
    DefaultMessage      string = "Beep Beep!"
    PlayingSound        string = "Playing sound on iOS Device [%s] with message: '%s'"
    StartingService     string = "Starting beepme as a service."
    ProvideCertificates string = "Please supply a port, private key and certficiate when starting service"
    KeyNotFoundError    string = "Unable to load key file '%s'"
    CertNotFoundError   string = "Unable to load cert file '%s'"
    PortInvalidError    string = "Port invalid [%s], must be higer than 1024"
    DeviceRefreshFailed string = "Can't refresh devices: %v"
    FlagAppleID         string = "Your iCloud ID / AppleID (normally an email)"
    FlagApplePass       string = "Pretty sure this is self explanatory"
    FlagDeviceName      string = "Name of the iOS device you want to beep"
    FlagDeviceMessage   string = "Message to be sent to iOS device"
    FlagRunService      string = "Run as https service"
    FlagServicePort     string = "(service only) Port to run https service on"
    FlagServiceCert     string = "(service only) certificate to use"
    FlagServiceKey      string = "(service only) private server key"
    PickTargetID        string = "Pick Target ID: "
    ICloudPassword      string = "iCloud Password: "
    ICloudUsername      string = "iCloud Username: "
    BeepMessage         string = "Message: "
    BeepHeader          string = "gobeepme - page your iOS device\n-----------------------------";
)

func (d *DeviceResult) GetDevice(id string) (*Device, error) {
    for _, r := range d.Devices {
        if r.ID == id {
            return &r, nil
        }
    }
    return nil, fmt.Errorf(NoDeviceID, id)
}

func (d *DeviceResult) GetDeviceByIndex(index int) (*Device, error) {
    i := 0
    for _, d := range d.Devices {
        if i >= index {
            return &d, nil
        }
        i++
    }
    return nil, fmt.Errorf(NoDeviceIndex, index)
}

func (d *DeviceResult) GetDeviceByName(name string) (*Device, error) {
    for _, d := range d.Devices {
        if strings.ToLower(d.Name) == strings.ToLower(name) {
            return &d, nil
        }
    }
    return nil, fmt.Errorf(NoDeviceName, name)
}

func Dummy() {

}
