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
import {MessageComponent}                                   from "./message.component";
import {HeaderComponent}                                    from "./utils/header.component";
import {AuthEventService}                                   from "./services/auth.service";
import {BeepEventService}                                   from "./services/beepevent.service";

@Component({
    templateUrl:    './app/device-list.component.html',
    providers:      [ModelService],
    directives:     [DeviceItemComponent, BatteryComponent,
                        LocationComponent, MessageComponent, HeaderComponent]
})
export class DeviceListComponent implements OnInit, OnActivate {

    devices: Device[];
    message:string;
    spinner:boolean = true;
    beeping:boolean = false;
    beepingError:boolean = false;
    beepThrottle = null;
    beepThrottleEnd = null;

    beepCount:number = 0;
    throttled:boolean = false;
    throttleClose:boolean = false;
    throttleTimeout:number = 10000; // 10 seconds

    constructor(private _creds: Creds,
                private _router: Router,
                private _modelService: ModelService,
                private _authService: AuthEventService,
                private _beepService: BeepEventService) {
    }

    ngOnInit() {
        if(!this._creds.authenticated) {
            this._router.navigate( ["Authenticate", {}] );
            return
        }
        this._authService.authenicatedEvent(true); // tell the header to shrink
        this._beepService.subscribe(null, r => this.beepResponse(r));
    }

    beepResponse(success) {
        if(!success) {

            this.beepingError = true;
            setTimeout(() => this.beepingError = false, 2500);

            return;
        }

        if(this.beepCount>=2) { // two device beeps is enough.
            this.startThrottle();
            return;
        }
        this.beepCount++;
        this.beeping = true;
        setTimeout(() => this.beeping = false, 1200);
    }

    routerOnActivate() {
       this._modelService.getDevices(
           this._creds,
           (r:Device[]) => this.updateList(r),
           (r:Device[]) => this.listFailed(r));
    }

    listFailed(err) {
        console.log('error!', err)
    }

    updateList(result) {
       this.devices = result;
        setTimeout(() => {
            this.spinner = false;}, 1000); // give the animations a second
    }

    startThrottle() {
        this.throttled = true;
        clearTimeout(this.beepThrottle)
        clearTimeout(this.beepThrottleEnd)
        this.beepThrottle = setTimeout(
            () => { this.throttleClose = false;
                    this.beepCount = 0;
                    this.throttled = false },
            this.throttleTimeout);

        this.beepThrottleEnd = setTimeout(
            () => this.throttleClose = true,
            this.throttleTimeout-500); // close anim before dom removal;
    }

    selectDevice(device) {
        if(!this.beeping && !this.throttled) {
            this._beepService.beepEventRequest(device);
        }
    }
}
