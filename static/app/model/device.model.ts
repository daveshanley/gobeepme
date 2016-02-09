import {DeviceLocation} from "./devicelocation.model"

export class Device {
    public id:                  string;
    public BatteryLevel:        number;
    public BatteryStatus:       string;
    public deviceClass:         string;
    public deviceDisplayName:   string;
    public location:            DeviceLocation;
    public deviceModel:         string;
    public modelDisplayName:    string;
    public Name:                string;
}
