// Copyright 2016 Dave Shanley <dave@quobix.com>
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

// service package handles all web service calls and provides a simple RESTful API for the bundled UI.
package service

import (
    "os"
    "net/http"
    "log"
    "encoding/json"
    "fmt"
    "strconv"
    "github.com/daveshanley/gobeepme/model"
    "github.com/daveshanley/gobeepme/commands"
    "github.com/daveshanley/gobeepme/console"
)

func getDevices(w http.ResponseWriter, cs *model.CloudService) (*model.DeviceResult, error) {
    d, err := commands.RefreshDeviceList(cs)
    if err != nil {
       writeError(w,fmt.Sprintf(model.ListDevicesFailed, err))
        return nil, fmt.Errorf(model.ListDevicesFailed, err)
    }
    return &d, nil
}

func authenticate(w http.ResponseWriter, cr model.Creds) (*model.CloudService, error) {
    cs, err := commands.Authenticate(cr)
    if err != nil {
        writeError(w,fmt.Sprintf(model.AuthFailedMessage, err))
        return nil, fmt.Errorf(model.AuthFailedMessage, err)

    }
    return &cs, nil
}

func setHeaders(w http.ResponseWriter) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Method", "POST, GET")
}

func checkServiceCommand(sc model.ServiceCommand) bool {
    if sc.Creds.AppleID     == "" { return false }
    if sc.Creds.Password    == "" { return false }
    if sc.Name              == "" { return false }
    return true
}

func writeError(w http.ResponseWriter, msg string) {
    w.WriteHeader(http.StatusForbidden)
    json.NewEncoder(w).Encode(model.ServiceResponse{true,msg})
}

func ListDevices(w http.ResponseWriter, req *http.Request) {
    setHeaders(w)
    var cr model.Creds
    if err := json.NewDecoder(req.Body).Decode(&cr); err != nil {
        writeError(w,model.NoCredentials)
        return
    }

    cs,err := authenticate(w, cr)
    if err != nil {
        return
    }

    d,err := getDevices(w, cs)
    if err != nil {
       return
    }

    w.WriteHeader(http.StatusOK)
    if err :=  json.NewEncoder(w).Encode(d.Devices); err != nil {
        panic(err)
    }
}

func BeepDevice(w http.ResponseWriter, req *http.Request) {
    setHeaders(w)
    var sc model.ServiceCommand
    if err := json.NewDecoder(req.Body).Decode(&sc); err != nil {
        writeError(w,model.CommandMalformed)
        return
    }
    if !checkServiceCommand(sc) {

        w.WriteHeader(http.StatusForbidden)
        json.NewEncoder(w).Encode(sc)
        //writeError(w,model.CommandMissingAttr)
        return
    }

    cs,err := authenticate(w, sc.Creds)
    if err != nil {
        return
    }

    d,err := getDevices(w, cs)
    if err != nil {
        return
    }

    dv, err := d.GetDeviceByName(sc.Name)
    if err!=nil {
        writeError(w, fmt.Sprintf(model.NoDeviceName, err))
        return
    }

    if sc.Message == "" { sc.Message = model.DefaultMessage }
    r := model.ServiceResponse{false,fmt.Sprintf(model.PlayingSound, dv.Name, sc.Message)}
    w.WriteHeader(http.StatusOK)
    if err :=  json.NewEncoder(w).Encode(r); err != nil {
        panic(err)
    }
    commands.PlaySound(cs, dv, sc.Message)
    return
}

func StartService(port int, key, cert string) {
    if key == "" || cert =="" {
        console.PrintKeyCertError()
        return
    }
    if(port <= 1024) {
        console.PrintPortInvalidError(port);
        return;
    }
    if _, err := os.Stat(key); os.IsNotExist(err) {
        console.PrintKeyNotFoundError(key)
        return
    }
    if _, err := os.Stat(cert); os.IsNotExist(err) {
        console.PrintCertNotFoundError(cert)
        return
    }

    console.PrintServiceMode(port)
    router := NewRouter()
    log.Fatal(http.ListenAndServeTLS(":" + strconv.Itoa(port), cert, key, router))
}

func Dummy() {

}
