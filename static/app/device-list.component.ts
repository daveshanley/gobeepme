import {Input, Output, EventEmitter, Component, OnInit}     from "angular2/core";
import {DataService}                                        from "./services/data.service";
import {Creds}                                              from "./model/creds.model";
import {Router}                                             from "angular2/router";
import {ModelService}                                       from "./services/model.service";
import {Device}                                             from "./model/device.model";
import {OnActivate}                                         from "angular2/router";
import {BatteryComponent}                                   from "./battery.component";
import {LocationComponent}                                  from "./location.component";
import {DeviceItemComponent}                                from "./device-item.component";

@Component({
    templateUrl:    './app/device-list.component.html',
    providers:      [ModelService],
    directives:     [DeviceItemComponent, BatteryComponent, LocationComponent]

})

export class DeviceListComponent implements OnInit, OnActivate {

    devices: Device[];

    constructor(dataService: DataService,
                private _creds: Creds,
                private _router: Router,
                private _modelService: ModelService) {

    }

    ngOnInit() {
        if(!this._creds.authenticated) {
            this._router.navigate( ["Authenticate", {}] );
        }
    }

    routerOnActivate() {
       this._modelService.getDevices(this._creds,
           (r) => { this.updateList(r)}, (r) => { this.listFailed(r)});
    }

    listFailed(err) {
        console.log('error!', err)
    }

    updateList(result) {
        this.devices = result;
    }


}
