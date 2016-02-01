import {Component, Input}       from "angular2/core";
import {NgClass}                from "angular2/common";
import {Device}                 from "./model/device.model";
import {BatteryComponent}       from "./battery.component";
import {LocationComponent}      from "./location.component";

@Component({
    selector: "device-item",
    templateUrl:    './app/device-item.component.html',
    directives:     [BatteryComponent, LocationComponent]
})

export class DeviceItemComponent  {
    @Input() device: Device;
}
