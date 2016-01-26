package main

import (
	"fmt"
	"net/http"
	"bytes"
	"net/http/cookiejar"
	//"errors"
	"encoding/json"
	"bufio"
	"os"
	"strings"
	"flag"
	"github.com/howeyc/gopass"
	"github.com/olekukonko/tablewriter"
	"strconv"
	"unicode/utf8"
	"github.com/daveshanley/gobeepme/model"
	"github.com/daveshanley/gobeepme/console"

)


var (
	client = &http.Client{
		Jar: cookieJar,
	}
	cookieJar, _ = cookiejar.New(nil)
)

func authenticate(c model.Creds) (model.CloudService, error){

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
		return model.CloudService{}, fmt.Errorf("Unable to authenticate with ['%s'], Check credentials.\n\n", c.AppleID)
	}
	cs:=model.CloudService{resp.Header.Get("X-Apple-MMe-Host"),resp.Header.Get("X-Apple-MMe-Scope"), c}
	fmt.Printf("Authenticated: iCloud Host[%s] / Scope[%s]\n", cs.Host, cs.Scope)
	return cs,nil
}

func refreshDeviceList(cs *model.CloudService) (model.DeviceResult, error) {
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
	var devices model.DeviceResult
	if err := json.NewDecoder(resp.Body).Decode(&devices); err !=nil {
		return model.DeviceResult{}, fmt.Errorf("Error, unable to decode JSON : %v", err)
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


func main() {
	uf := flag.String("user", "", "Your iCloud ID / AppleID (normally an email)")
	pf := flag.String("pass", "", "Pretty sure this is self explanatory")
	//nf := flag.String("name", "", "Name of the iOS device you want to beep")

	// print welcome!
	welcomeBanner()

	var username, password string
	flag.Parse()
	ufVal := *uf
	pfVal := *pf
	//nfVal := *nf

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

	/*
	if nfVal == "" {
		id = collectId()
	} else {
		id = strings.TrimSpace(nfVal)
	}
	*/


	var creds = model.Creds{AppleID: username ,Password: password}
	var cs model.CloudService
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
	console.PrintDevices(&d)
	dID := collectDeviceSelection(len(d.Devices))
	message := collectMessageSelection()
	play := collectPlaySound()
	fd,err :=d.GetDeviceByIndex(dID-1)
	if err!=nil {
		fmt.Printf("\nUnable to extract device: %v", err)
		return
	}

	if play {
		playSound(&cs, fd, message);
	} else {
		sendMessage(&cs, fd, message);
	}
	fmt.Printf("\nDelivered!\n")

}
