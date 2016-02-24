// Copyright 2016 Dave Shanley <dave@quobix.com>
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

// Package commands provides iCloud server call functionality
package commands

import (
    "github.com/daveshanley/gobeepme/model"
    "net/http"
    "net/http/cookiejar"
    "encoding/json"
    "bytes"
    "fmt"
    //"io/ioutil"
)

const (
    FPIServiceURL   string = "https://fmipmobile.icloud.com/fmipservice/device/"
    Referrer        string = "https://www.icloud.com"
    ServiceHost     string = "X-Apple-MMe-Host"
    ServiceScope    string = "X-Apple-MMe-Scope"
    ServiceEndpoint string = "/fmipservice/device/"
    InitCommand     string = "/initClient"
    RefreshCommand  string = "/refreshClient"
    SoundCommand    string = "/playSound"
    MessageCommand  string = "/sendMessage"
)

var (
    cl = &http.Client{
        Jar: cookieJar,
    }
    cookieJar, _ = cookiejar.New(nil)
)

func setOriginHeader(req *http.Request) {
    req.Header.Set("Origin",Referrer)
}

func setBasicAuth(req *http.Request, c model.Creds) {
    req.SetBasicAuth(c.AppleID, c.Password)
}

func createRequest(cmd string, cs *model.CloudService, d *bytes.Reader) (*http.Request, error) {
    return http.NewRequest("POST", "https://" + cs.Host + ServiceEndpoint + cs.Scope + cmd, d)
}

func prepareRequest(cmd string, cs *model.CloudService, c model.Creds, d *bytes.Reader) *http.Request {
    req, err := createRequest(cmd, cs, d)
    if err != nil {
        panic(err)
    }
    setOriginHeader(req)
    setBasicAuth(req, c)
    return req
}

func executeCommand(req *http.Request) (*http.Response, error) {
    resp, err := cl.Do(req)
    if err != nil {
        panic(err)
    }
    return resp, err
}

// Play a sound (and send a message) to iOS Device
func PlaySound(cs *model.CloudService, d *model.Device, msg string) bool {
    sc := model.ServerCommand{d.ID, msg}
    o,_ := json.Marshal(sc)
    req := prepareRequest(SoundCommand,cs, cs.Creds, bytes.NewReader(o))
    if _,err :=executeCommand(req); err!=nil {
        return false
    }
    return true
}

// Send a message to iOS device
func SendMessage(cs *model.CloudService, d *model.Device, msg string) bool {
    sc := model.ServerCommand{d.ID, msg}
    o,_ := json.Marshal(sc)
    req := prepareRequest(MessageCommand,cs, cs.Creds, bytes.NewReader(o))
    if _,err :=executeCommand(req); err!=nil {
        return false
    }
    return true
}

// Authenticate user
func Authenticate(c model.Creds) (model.CloudService, error) {
    req, _ := http.NewRequest("POST", FPIServiceURL + c.AppleID + InitCommand, bytes.NewBufferString(""))
    setOriginHeader(req)
    setBasicAuth(req, c)
    resp,_ :=executeCommand(req)
    if resp.StatusCode == http.StatusForbidden ||
        resp.StatusCode == http.StatusUnauthorized {
        return model.CloudService{},
        fmt.Errorf("Your credentials were rejected, try again.")
    }
    return model.CloudService{resp.Header.Get(ServiceHost),
        resp.Header.Get(ServiceScope), c}, nil
}

// Make a new request for our most recent devices
func RefreshDeviceList(cs *model.CloudService) (model.DeviceResult, error) {
    req := prepareRequest(RefreshCommand,cs, cs.Creds, bytes.NewReader([]byte("")))
    resp,_ :=executeCommand(req)
    var dv model.DeviceResult
    if err := json.NewDecoder(resp.Body).Decode(&dv); err != nil {
        return model.DeviceResult{},
        fmt.Errorf("unable to decode JSON: %v", err)
    }
    return dv, nil
}

func Dummy() {

}
