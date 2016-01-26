package gobeepme

import (
	"fmt"
	"net/http"
	"bytes"
	"net/http/cookiejar"
	"errors"
	"encoding/json"
	"bufio"
	"os"
	"strings"
	"flag"
	"github.com/howeyc/gopass"
	"github.com/olekukonko/tablewriter"
	"strconv"
	"unicode/utf8"
)

type Creds struct {
	AppleID string `json:"apple_id"`
	Password string `json:"password"`
}

type DeviceResult struct {
	StatusCode string `json:"statusCode"`
	Devices []Device `json:"content"`
}

type Device struct {
	ID string `json:"id"`
	BatteryLevel float64 `json:"batteryLevel`
	BatteryStatus string `json:"batteryStatus`
	Class string `json:"deviceClass"`
	DisplayName string `json:"deviceDisplayName"`
	Location DeviceLocation `json:"location"`
	Model string `json:"deviceModel"`
	ModelDisplayName string `json:"modelDisplayName"`
	Name string `'json:"name"`
}

type DeviceLocation struct {
	Longitude float64 `json:"longitude"`
	Latitude float64 `json:"latitude"`
}

type ServerCommand struct {
	DeviceID string `json:"device"`
	Message string `json:"subject"`
}

type CloudService struct {
	Host string
	Scope string
	Creds Creds
}

func (d *DeviceResult) GetDevice(id string) (*Device,error) {
	for _,r := range d.Devices {
		if r.ID == id {
			return &r, nil
		}
	}
	return nil, errors.New("No device found")
}

func (d *DeviceResult) GetDeviceByIndex(index int) (*Device,error) {
	i := 0
	for _,d := range d.Devices {
		if i >= index {
			return &d, nil
		}
		i++
	}
	return nil, fmt.Errorf("No Device with index [%d] located", index)
}

func (d *DeviceResult) GetDeviceByName(name string) (*Device,error) {
	for _,d := range d.Devices {
		if strings.ToLower(d.DisplayName) == strings.ToLower(name) {
			return &d, nil
		}
	}
	return nil, fmt.Errorf("No Device with name [%s] located", name)
}


func (d *DeviceResult) GetDeviceByDisplayName(dn string) *Device {
	for _,r := range d.Devices {
		if r.DisplayName == dn {
			return &r
		}
	}
	return nil
}

var (
	client = &http.Client{
		Jar: cookieJar,
	}
	cookieJar, _ = cookiejar.New(nil)
)

func authenticate(c Creds) (CloudService, error){

	req, err := http.NewRequest("POST", "https://fmipmobile.icloud.com/fmipservice/device/" + c.AppleID + "/initClient", bytes.NewBufferString(""))
	req.Header.Set("Origin", "https://www.icloud.com")
	req.SetBasicAuth(c.AppleID, c.Password)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// validate response code
	if resp.StatusCode == http.StatusForbidden || resp.StatusCode == http.StatusUnauthorized  {
		return CloudService{}, fmt.Errorf("Unable to authenticate with ['%s'], Check credentials.\n\n", c.AppleID)
	}
	cs:=CloudService{resp.Header.Get("X-Apple-MMe-Host"),resp.Header.Get("X-Apple-MMe-Scope"), c}
	fmt.Printf("Authenticated: iCloud Host[%s] / Scope[%s]\n", cs.Host, cs.Scope)
	return cs,nil
}

func refreshDeviceList(cs *CloudService) (DeviceResult, error) {
	// make a new request for our most recent devices
	req, err := http.NewRequest("POST",
		"https://" + cs.Host + "/fmipservice/device/" +
		cs.Scope + "/initClient", bytes.NewBufferString(""))
	req.Header.Set("Origin", "https://www.icloud.com")
	req.SetBasicAuth(cs.Creds.AppleID, cs.Creds.Password)

	// print out an update.
	fmt.Println("\nRefreshing Device list...")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var devices DeviceResult
	if err := json.NewDecoder(resp.Body).Decode(&devices); err !=nil {
		return DeviceResult{}, fmt.Errorf("Error, unable to decode JSON : %v", err)
	}
	return devices, nil
}

func welcomeBanner() {
	fmt.Println("\nbeepme - page your iOS device")
	fmt.Println("-----------------------------\n")

}

func collectUsername() string {
	r := bufio.NewReader(os.Stdin)
	var username string

	for username == "" {
		fmt.Print("iCloud Username: ")
		u, _ := r.ReadString('\n')
		username = strings.TrimSpace(u)
	}
	return username
}

func collectPassword() string {
	fmt.Print("iCloud Password: ")
	pw := gopass.GetPasswd()
	return strings.TrimSpace(string(pw))
}

func collectDeviceSelection(size int) int {
	r := bufio.NewReader(os.Stdin)
	var s int
	var id string
	for id == "" {
		fmt.Print("Pick Target ID: ")
		i, _ := r.ReadString('\n')
		i = strings.TrimSpace(i)
		if a,b := strconv.Atoi(i); b != nil || a < 0 || a > size {
			fmt.Printf("\n\tInvalid input '%s', please try again, enter a number between 1-%d\n\n", i, size)
			continue
		} else {
			s = a
			id = i
		}
	}
	return s
}

func collectMessageSelection() string {
	r := bufio.NewReader(os.Stdin)
	var message string
	for message == "" {
		fmt.Print("Message: ")
		m, _ := r.ReadString('\n')
		message = strings.TrimSpace(m)
	}
	return message
}

func collectPlaySound() bool {
	r := bufio.NewReader(os.Stdin)
	var sound string
	var b bool = false
	for sound == "" {
		fmt.Print("Play Sound? [y/n]: ")
		s, _ := r.ReadString('\n')
		sound = strings.ToLower(strings.TrimSpace(s))
		rc := utf8.RuneCountInString(sound)
		if rc != 1 || (sound != "y" && sound != "n") {
			sound = ""
			continue
		} else {
			if sound == "y" {
				b = true
			} else {
				b = false
			}
		}
	}
	return b
}


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

func playSound (cs *CloudService, d *Device, msg string) bool {
	sc := ServerCommand{d.ID, msg}
	o,_ := json.Marshal(sc)
	req, err := http.NewRequest("POST", "https://"+cs.Host+"/fmipservice/device/" +
		cs.Scope + "/playSound", bytes.NewReader(o))
	req.Header.Set("Origin", "https://www.icloud.com")
	req.SetBasicAuth(cs.Creds.AppleID, cs.Creds.Password)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return true;
}

func sendMessage (cs *CloudService, d *Device, msg string) bool {
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
	return true;
}

func main() {
	uf := flag.String("user", "", "Your iCloud ID / AppleID (normally an email)")
	pf := flag.String("pass", "", "Pretty sure this is self explanatory")
	nf := flag.String("name", "", "Name of the iOS device you want to beep")

	// print welcome!
	welcomeBanner()

	var username, password, id, msg string
	flag.Parse()
	ufVal := *uf
	pfVal := *pf
	nfVal := *nf

	if ufVal == "" {
		username=collectUsername()
	} else {
		username = strings.TrimSpace(ufVal)
	}
	if pfVal == "" {
		password = collectPassword()
	} else {
		password = strings.TrimSpace(pfVal)
	}
	if nfVal == "" {
		id = collectId()
	} else {
		id = strings.TrimSpace(nfVal)
	}



	var creds = Creds{AppleID: username ,Password: password}
	var cs CloudService
	cs,err := authenticate(creds)
	if err!=nil {
		fmt.Printf("\nAuthentication Failed: %v", err)
		return
	}

	d,err := refreshDeviceList(&cs)
	if err != nil {
		fmt.Printf("\nCan't refresh devices: %v", err)
		return
	}
	printDevices(&d)
	dID := collectDeviceSelection(len(d.Devices))
	message := collectMessageSelection()
	play := collectPlaySound()
	fd,err :=d.GetDeviceByIndex(dID-1)
	if err!=nil {
		fmt.Printf("\nUnable to extract device: %v", err)
		return
	}

	if play {
		//playSound(&cs, fd, message);
	} else {
		//sendMessage(&cs, fd, message);
	}
	fmt.Printf("\nDelivered!\n")

}
