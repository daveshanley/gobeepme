package service

import (
    "net/http"
    "log"
    "encoding/json"
    "fmt"
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
    //commands.PlaySound(cs, dv, sc.Message)
    return
}

func StartService(port int, key, cert string) {

   // if key == "" || cert =="" {
   //     console.PrintKeyCertError()
   //     return
   // }
    console.PrintServiceMode()
    router := NewRouter()
    //log.Fatal(http.ListenAndServeTLS(":9443", cert, key, router))
    log.Fatal(http.ListenAndServeTLS(":9443", "server.pem", "server.key", router))
}

func Dummy() {

}
