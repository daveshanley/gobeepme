// Copyright 2016 Dave Shanley <dave@quobix.com>
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

// gobeepme is both a console app and webservice to allow you to ping your iOS device from anywhere.
package main

import (
    "fmt"
    "strings"
    "flag"
    "github.com/daveshanley/gobeepme/model"
    "github.com/daveshanley/gobeepme/console"
    "github.com/daveshanley/gobeepme/commands"
    "github.com/daveshanley/gobeepme/service"
)

// main can either accept every arg as a flag, or you can step through in
// an interactive manner.
func main() {
    // console flags
    uf := flag.String("user", "", model.FlagAppleID)
    pf := flag.String("pass", "", model.FlagApplePass)
    nf := flag.String("name", "", model.FlagDeviceName)
    mf := flag.String("msg", model.DefaultMessage, model.FlagDeviceMessage)
    sf := flag.Bool("service", false, model.FlagRunService)
    portf := flag.Int("port", 9443, model.FlagServicePort)
    certf := flag.String("key", "", model.FlagServiceKey)
    keyf := flag.String("cert", "", model.FlagServiceCert)

    var un, pw, dn, msg string
    flag.Parse()
    ufVal := *uf
    pfVal := *pf
    nfVal := *nf
    mfVal := *mf
    sfVal := *sf
    portVal := *portf
    keyVal := *keyf
    certVal := *certf

    // check for service mode
    if sfVal {
        service.StartService(portVal, certVal, keyVal)
        return
    }

    // print welcome!
    console.PrintWelcomeBanner()

    // defaults
    if ufVal == "" {
        un = console.CollectUsername()
    } else {
        un = strings.TrimSpace(ufVal)
    }
    if pfVal == "" {
        pw = console.CollectPassword()
    } else {
        pw = strings.TrimSpace(pfVal)
    }
    if nfVal == "" {
        dn = ""    // device name
    } else {
        dn = strings.TrimSpace(nfVal)
    }
    msg = strings.TrimSpace(mfVal)

    var cr = model.Creds{AppleID: un, Password: pw}
    var cs model.CloudService

    cs, err := commands.Authenticate(cr)
    if err != nil {
        console.PrintAuthFailed(err)
        return
    }

    d, err := commands.RefreshDeviceList(&cs)
    if err != nil {
        fmt.Printf("\n" + model.DeviceRefreshFailed, err)
        return
    }

    var dID int
    var dv *model.Device
    if nfVal == "" {
        console.PrintDevices(&d)
        dID = console.CollectDeviceSelection(len(d.Devices))
        dv, err = d.GetDeviceByIndex(dID-1)
        if err!=nil {
            console.PrintNoDeviceFound(string(dID))
            return
        }
    } else {
        dv, err = d.GetDeviceByName(dn)
        if err!=nil {
            console.PrintNoDeviceFound(dn)
            return
        }
    }
    console.PrintPlayingSound(dv.Name, msg)
    commands.PlaySound(&cs, dv, msg)
}

func Dummy() {

}

