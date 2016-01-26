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
    bf := flag.Bool("beep", true, "Send Beep?")
    sf := flag.Bool("server", false, "Run as http service")

    var un, pw, dn, msg string
    flag.Parse()
    ufVal := *uf
    pfVal := *pf
    nfVal := *nf
    mfVal := *mf
    bfVal := *bf
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
        pw = strings.TrimSpace(nfVal)
    }
    if nfVal == "" {
        dn = ""
    } else {
        pw = strings.TrimSpace(nfVal)
    }



    // print welcome!
    console.WelcomeBanner()

    /*
    if nfVal == "" {
        id = collectId()
    } else {
        id = strings.TrimSpace(nfVal)
    }
    */


    var creds = model.Creds{AppleID: username, Password: password}
    var cs model.CloudService
    cs, err := commands.Authenticate(creds, client)
    if err != nil {
        fmt.Printf("\nAuthentication Failed: %v", err)
        return
    }

    d, err := commands.RefreshDeviceList(&cs, client)
    if err != nil {
        fmt.Printf("\nCan't refresh devices: %v", err)
        return
    }
    console.PrintDevices(&d)
    dID := console.CollectDeviceSelection(len(d.Devices))
    message := console.CollectMessageSelection()
    play := console.CollectPlaySound()
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
