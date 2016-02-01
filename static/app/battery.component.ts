import {Component, Input}       from "angular2/core";
import {NgClass}                from "angular2/common";

@Component({
    selector: "battery",
    templateUrl:    './app/battery.component.html',
    directives: [NgClass]
})
export class BatteryComponent  {
    @Input() batteryLevel: number;
}
