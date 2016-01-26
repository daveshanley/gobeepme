package commands

import (
    "github.com/daveshanley/gobeepme/model"
    "net/http"
    "encoding/json"
    "bytes"
    "fmt"
)


func PlaySound(cs *model.CloudService, d *model.Device, msg string) bool {
    //sc := ServerCommand{d.ID, msg}
    //o,_ := json.Marshal(sc)

    /*
    req, err := http.NewRequest("POST", "https://"+cs.Host+"/fmipservice/device/" +
        cs.Scope + "/playSound", bytes.NewReader(o))
    req.Header.Set("Origin", "https://www.icloud.com")
    req.SetBasicAuth(cs.Creds.AppleID, cs.Creds.Password)

    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    */
    return true;
}

func SendMessage(cs *model.CloudService, d *model.Device, msg string) bool {
    /*
    sc := ServerCommand{d.ID, msg}

    o,_ := json.Marshal(sc)
    req, err := http.NewRequest("POST", "https://"+cs.Host+"/fmipservice/device/" +
    cs.Scope + "/sendMessage", bytes.NewReader(o))
    req.Header.Set("Origin", "https://www.icloud.com")
    req.SetBasicAuth(cs.Creds.AppleID, cs.Creds.Password)

    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()
    */
    return true;
}

func Authenticate(c model.Creds, cl *http.Client) (model.CloudService, error) {

    req, err := http.NewRequest("POST",
        "https://fmipmobile.icloud.com/fmipservice/device/" +
        c.AppleID + "/initClient", bytes.NewBufferString(""))
    req.Header.Set("Origin", "https://www.icloud.com")
    req.SetBasicAuth(c.AppleID, c.Password)

    resp, err := cl.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    // validate response code
    if resp.StatusCode == http.StatusForbidden ||
    resp.StatusCode == http.StatusUnauthorized {
        return model.CloudService{},
        fmt.Errorf("Unable to authenticate with ['%s'], Check credentials.\n\n", c.AppleID)
    }
    return model.CloudService{resp.Header.Get("X-Apple-MMe-Host"),
        resp.Header.Get("X-Apple-MMe-Scope"), c}, nil
}

func RefreshDeviceList(cs *model.CloudService, cl *http.Client) (model.DeviceResult, error) {
    // make a new request for our most recent devices
    req, err := http.NewRequest("POST",
        "https://" + cs.Host + "/fmipservice/device/" +
        cs.Scope + "/initClient", bytes.NewBufferString(""))
    req.Header.Set("Origin", "https://www.icloud.com")
    req.SetBasicAuth(cs.Creds.AppleID, cs.Creds.Password)

    // print out an update.
    fmt.Println("\nRefreshing Device list...")
    resp, err := cl.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    var devices model.DeviceResult
    if err := json.NewDecoder(resp.Body).Decode(&devices); err != nil {
        return model.DeviceResult{},
        fmt.Errorf("Error, unable to decode JSON : %v", err)
    }
    return devices, nil
}
