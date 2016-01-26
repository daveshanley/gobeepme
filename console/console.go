package console

import (
    "github.com/daveshanley/gobeepme/model"
    "github.com/olekukonko/tablewriter"
    "github.com/howeyc/gopass"
    "fmt"
    "strconv"
    "os"
    "bufio"
    "strings"
    "unicode/utf8"
)

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

func CollectMessageSelection() string {
    r := bufio.NewReader(os.Stdin)
    var message string
    for message == "" {
        fmt.Print("Message: ")
        m, _ := r.ReadString('\n')
        message = strings.TrimSpace(m)
    }
    return message
}

func CollectUsername() string {
    r := bufio.NewReader(os.Stdin)
    var username string
    for username == "" {
        fmt.Print("iCloud Username: ")
        u, _ := r.ReadString('\n')
        username = strings.TrimSpace(u)
    }
    return username
}

func CollectPassword() string {
    fmt.Print("iCloud Password: ")
    pw := gopass.GetPasswd()
    return strings.TrimSpace(string(pw))
}

func CollectDeviceSelection(size int) int {
    r := bufio.NewReader(os.Stdin)
    var s int
    var id string
    for id == "" {
        fmt.Print("Pick Target ID: ")
        i, _ := r.ReadString('\n')
        i = strings.TrimSpace(i)
        if a, b := strconv.Atoi(i); b != nil || a < 0 || a > size {
            fmt.Printf("\n\tInvalid input '%s', please try again, enter a number between 1-%d\n\n", i, size)
            continue
        } else {
            s = a
            id = i
        }
    }
    return s
}

func WelcomeBanner() {
    fmt.Println("\nbeepme - page your iOS device")
    fmt.Println("-----------------------------\n")
}

func PrintServiceMode() {
    fmt.Println("Starting beepme as a service, listening on port 37556")
}
