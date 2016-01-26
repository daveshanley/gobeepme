package gobeepme

import (
	"fmt"
	//"net/http"
	//"bytes"
	//"net/http/cookiejar"
	//"errors"
	//"encoding/json"
	//"bufio"
	//"os"
	//"strings"
	//"flag"
	//"github.com/howeyc/gopass"
	//"github.com/olekukonko/tablewriter"
	//"strconv"
	//"unicode/utf8"
)

func printDevices(dr *DeviceResult) {
	var floatConv = func (f float64) string {
		// to convert a float number to a string
		return fmt.Sprintf("%d%%",int(f*100))
	}
	var buildRow = func(x int, d Device) []string {
		return []string{strconv.Itoa(x+1), d.Name, d.DisplayName,
			d.ModelDisplayName, d.BatteryStatus, floatConv(d.BatteryLevel)}
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Device Name", "Device Type", "Model",
		"Battery Status", "Battery Level"})
	var i int = 0
	for _,d := range dr.Devices {
		table.Append(buildRow(i,d));
		i++
	}
	table.Render() // write table to stdout
}
