import {Component, Input,
            Output, EventEmitter}       from "angular2/core";
import {Device}                         from "./model/device.model";
import {BatteryComponent}               from "./battery.component";
import {LocationComponent}              from "./location.component";

@Component({
    selector:       "device-item",
    templateUrl:    './app/device-item.component.html',
    directives:     [BatteryComponent, LocationComponent]
})

export class DeviceItemComponent  {

    @Input() device: Device;
    @Output() selectedDevice:EventEmitter<Device> = new EventEmitter();

    selectDevice(device) {
        this.selectedDevice.emit(device);
    }
}
