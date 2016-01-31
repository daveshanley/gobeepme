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

/*
\id: "ACb1s3rZJele31CoOFdq39NTFNDKX53n9cdOQR7ApDV7+mMBzJyImw==",
 BatteryLevel: 0,
 BatteryStatus: "Unknown",
 deviceClass: "MacBookPro",
 deviceDisplayName: "MacBook Pro 15"",
 location: {
 longitude: 0.03796026181537588,
 latitude: 51.51237741338369
 },
 deviceModel: "MacBookPro8_2",
 modelDisplayName: "MacBook Pro",
 Name: "ShanBook Pro"
 */
