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

func Dummy() {

}

func main() {

    // configure console flags
    uf := flag.String("user", "", "Your iCloud ID / AppleID (normally an email)")
    pf := flag.String("pass", "", "Pretty sure this is self explanatory")
    nf := flag.String("name", "", "Name of the iOS device you want to beep")
    mf := flag.String("msg", model.DefaultMessage, "Message to be sent to iOS device")
    sf := flag.Bool("service", false, "Run as https service")
    portf := flag.Int("port", 9443, "(service only) Port to run https service on")
    certf := flag.String("key", "", "(service only) private server key")
    keyf := flag.String("cert", "", "(service only) certificate to use")

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
        console.PrintAuthFailed(err)
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
