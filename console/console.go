// Copyright 2016 Dave Shanley <dave@quobix.com>
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

// Package console handles input and output to and from stdin/stdout
package console

import (
    "github.com/daveshanley/gobeepme/model"
    "github.com/olekukonko/tablewriter"
    "github.com/howeyc/gopass"
    "github.com/fatih/color"
    "fmt"
    "strconv"
    "os"
    "bufio"
    "strings"
)

// PrintDevices uses the tablewriter (github.com/olekukonko/tablewriter)
// package to pretty print a lovely grid of the devices registered on your
// iCloud account
func PrintDevices(dr *model.DeviceResult) {
    // to convert a float number to a string
    var fc = func(f float64) string {
        return fmt.Sprintf("%d%%", int(f * 100))
    }
    var br = func(x int, d model.Device) []string {
        return []string{strconv.Itoa(x + 1), d.Name, d.DisplayName,
            d.ModelDisplayName, d.BatteryStatus, fc(d.BatteryLevel)}
    }
    t := tablewriter.NewWriter(os.Stdout)
    t.SetHeader([]string{"ID", "Device Name", "Device Type", "Model",
        "Battery Status", "Battery Level"})
    var i int = 0
    for _, d := range dr.Devices {
        t.Append(br(i, d));
        i++
    }
    t.Render() // write table to stdout
}

// CollectMesageSelection asks the user to type a message to be sent to the
// iOS device you want to beep.
func CollectMessageSelection() string {
    r := bufio.NewReader(os.Stdin)
    var message string
    for message == "" {
        fmt.Print(model.BeepMessage)
        m, _ := r.ReadString('\n')
        message = strings.TrimSpace(m)
    }
    return message
}

// CollectUsername asks the user to type in their iCloud username. It then
// collects the input and returns it.
func CollectUsername() string {
    r := bufio.NewReader(os.Stdin)
    var username string
    for username == "" {
        fmt.Print("\n" + model.ICloudUsername)
        u, _ := r.ReadString('\n')
        username = strings.TrimSpace(u)
    }
    return username
}

// CollectPassword asks the user to type in their iCloud password. It then
// collects the input and returns it.
func CollectPassword() string {
    fmt.Print(model.ICloudPassword)
    pw,_ := gopass.GetPasswd()
    return strings.TrimSpace(string(pw))
}

// CollectDeviceSelection asks the user to select a device to beep. It then returns
// the index of the selected device
func CollectDeviceSelection(size int) int {
    r := bufio.NewReader(os.Stdin)
    var s int
    var id string
    for id == "" {
        fmt.Print(model.PickTargetID)
        i, _ := r.ReadString('\n')
        i = strings.TrimSpace(i)
        if a, b := strconv.Atoi(i); b != nil || a < 0 || a > size {
            fmt.Printf("\n\tInvalid input '%s', please try again, enter" +
                "a number between 1-%d\n\n", i, size)
            continue
        } else {
            s = a
            id = i
        }
    }
    return s
}

// printBlankLine does exactly what it says on the tin.
func printBlankLine() {
    fmt.Println("")
}

// PrintWelcomeBanner this one should be obvious.
func PrintWelcomeBanner() {
    color.Cyan("\n" + model.BeepHeader)
}

// PrintAuthFailed well, now we're just getting literal
func PrintAuthFailed(err error) {
    fmt.Printf(model.AuthFailedMessage, err)
    printBlankLine()
}

// PrintNoDeviceFound prints out the supplied device that wasn't found
// when looking for it against the users iCloud accountg
func PrintNoDeviceFound(d string) {
    fmt.Printf(model.NoDeviceName, d)
    printBlankLine()
}

// PrintPlayingSound prints out a message to the console, indicating that
// the beep is indeed beeping.
func PrintPlayingSound(name, msg string) {
    fmt.Printf(model.PlayingSound, name, msg)
    printBlankLine()
}

// PrintServiecMoode prints out a message informing the user that the app
// has started in a non interactive mode and is listening for web requests.
func PrintServiceMode(d int) {
    fmt.Printf(model.StartingService, d)
    printBlankLine()
}

// PrintKeyCertError simply prints out a message informing the user that
// they have failed to supply a private key and certificate. This is because
// gobeepme only runs over TLS.
// If you don't have a valid key/cert - you can generate a quality free one
// at https://letsencrypt.org
func PrintKeyCertError() {
    fmt.Println(model.ProvideCertificates)
}

// PrintKeyNotFoundError prints out a message to the console informing the user
// that the sky has indeed fallen. Or it may indicate that the key path/file
// simply does not exist.
func PrintKeyNotFoundError(k string) {
    fmt.Printf(model.KeyNotFoundError, k)
}

// PrintKeyNotFoundError prints out a message to the console informing the user
// that the sky has indeed fallen. Or it may indicate that the key path/file
// simply does not exist.
func PrintCertNotFoundError(c string) {
    fmt.Printf(model.CertNotFoundError, c)
}

// PrintPortInvalidError prints out a message to the console informing the user
// that they need to use a port higher than 1024 - because of reasons.
func PrintPortInvalidError(p int) {
    fmt.Printf(model.PortInvalidError, p)
    printBlankLine()
}

func Dummy() {

}
