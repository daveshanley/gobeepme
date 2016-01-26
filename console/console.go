package console

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/daveshanley/gobeepme/model"
	"strconv"
	"os"
)

func PrintDevices(dr *model.DeviceResult) {
	// to convert a float number to a string
	var fc = func (f float64) string {
		return fmt.Sprintf("%d%%",int(f*100))
	}
	var br = func(x int, d model.Device) []string {
		return []string{strconv.Itoa(x+1), d.Name, d.DisplayName,
			d.ModelDisplayName, d.BatteryStatus, fc(d.BatteryLevel)}
	}
	t := tablewriter.NewWriter(os.Stdout)
	t.SetHeader([]string{"ID", "Device Name", "Device Type", "Model",
		"Battery Status", "Battery Level"})
	var i int = 0
	for _,d := range dr.Devices {
		t.Append(br(i,d));
		i++
	}
	t.Render() // write table to stdout
}
