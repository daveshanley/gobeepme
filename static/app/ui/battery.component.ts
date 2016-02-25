import {Component, Input}       from "../../node_modules/angular2/core";
import {NgClass}                from "../../node_modules/angular2/common";

@Component({
    selector: "battery",
    templateUrl:    './app/ui/battery.component.html',
    directives: [NgClass]
})
export class BatteryComponent  {
    @Input() batteryLevel: number;
}
