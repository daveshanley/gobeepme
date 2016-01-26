package main

import (
    "fmt"
    "net/http"
    "net/http/cookiejar"
    "strings"
    "flag"
    "github.com/daveshanley/gobeepme/model"
    "github.com/daveshanley/gobeepme/console"
    "github.com/daveshanley/gobeepme/commands"
)

var (
    client = &http.Client{
        Jar: cookieJar,
    }
    cookieJar, _ = cookiejar.New(nil)
)

func main() {

    // configure console flags
    uf := flag.String("user", "", "Your iCloud ID / AppleID (normally an email)")
    pf := flag.String("pass", "", "Pretty sure this is self explanatory")
    nf := flag.String("name", "", "Name of the iOS device you want to beep")
    mf := flag.String("msg", "", "Message to be sent to iOS device")
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
    if mfVal == "" {
        msg = "gobeepme is beeping you!" // message
    } else {
        msg = strings.TrimSpace(mfVal)
    }

    // print welcome!
    console.WelcomeBanner()

    var cr = model.Creds{AppleID: un, Password: pw}
    var cs model.CloudService

    cs, err := commands.Authenticate(cr, client)
    if err != nil {
        fmt.Printf("\nAuthentication Failed: %v", err)
        return
    }

    d, err := commands.RefreshDeviceList(&cs, client)
    if err != nil {
        fmt.Printf("\nCan't refresh devices: %v", err)
        return
    }

    var dID int
    if dn == "" {
        console.PrintDevices(&d)
        dID = console.CollectDeviceSelection(len(d.Devices))
    } else {

    }
    message := console.CollectMessageSelection()

    fd, err := d.GetDeviceByIndex(dID - 1)
    if err != nil {
        fmt.Printf("\nUnable to extract device: %v", err)
        return
    }

    if play {
        commands.PlaySound(&cs, fd, message);
    } else {
        commands.SendMessage(&cs, fd, message);
    }
    fmt.Printf("\nDelivered!\n")

}
