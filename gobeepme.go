package main

import (
    "fmt"
    "strings"
    "flag"
    "github.com/daveshanley/gobeepme/model"
    "github.com/daveshanley/gobeepme/console"
    "github.com/daveshanley/gobeepme/commands"
)

func Dummy() {

}

func main() {

    // configure console flags
    uf := flag.String("user", "", "Your iCloud ID / AppleID (normally an email)")
    pf := flag.String("pass", "", "Pretty sure this is self explanatory")
    nf := flag.String("name", "", "Name of the iOS device you want to beep")
    mf := flag.String("msg", "gobeepme is beeping you!", "Message to be sent to iOS device")
    sf := flag.Bool("server", false, "Run as http service")

    var un, pw, dn, msg string
    flag.Parse()
    ufVal := *uf
    pfVal := *pf
    nfVal := *nf
    mfVal := *mf
    sfVal := *sf

    // check for service mode
    if sfVal {
        console.PrintServiceMode()
        return
    }

    // print welcome!
    console.WelcomeBanner()

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
        fmt.Printf("\nAuthentication Failed: %v", err)
        return
    }

    d, err := commands.RefreshDeviceList(&cs)
    if err != nil {
        fmt.Printf("\nCan't refresh devices: %v", err)
        return
    }

    var dID int
    var dv *model.Device
    if nfVal == "" {
        console.PrintDevices(&d)
        dID = console.CollectDeviceSelection(len(d.Devices))
        dv, err = d.GetDeviceByIndex(dID-1)
        if err!=nil {
            fmt.Printf("Unable to extract device: %v\n\n", err)
            return
        }
    } else {
        dv, err = d.GetDeviceByName(dn)
        if err!=nil {
            fmt.Printf("Unable to locate iOS device: %v\n\n", err)
            return
        }
    }
    commands.PlaySound(&cs, dv, msg)
    //fmt.Printf("\nPlaying sound.. dummy: %s: %s\n\n", dv.Name, msg)
}
