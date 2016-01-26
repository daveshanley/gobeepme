package commands

import (
	"github.com/daveshanley/gobeepme/model"
)


func PlaySound (cs *model.CloudService, d *model.Device, msg string) bool {
	//sc := ServerCommand{d.ID, msg}
	//o,_ := json.Marshal(sc)

	/*
	req, err := http.NewRequest("POST", "https://"+cs.Host+"/fmipservice/device/" +
		cs.Scope + "/playSound", bytes.NewReader(o))
	req.Header.Set("Origin", "https://www.icloud.com")
	req.SetBasicAuth(cs.Creds.AppleID, cs.Creds.Password)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	*/
	return true;
}

func SendMessage (cs *model.CloudService, d *model.Device, msg string) bool {
	/*
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
	*/
	return true;
}
