import {Component, OnInit}              from '../../node_modules/angular2/core';
import {RouteConfig, ROUTER_DIRECTIVES} from '../../node_modules/angular2/router';
import {AuthComponent}                  from './auth.component'
import {Device}                         from './../model/device.model'
import {DeviceListComponent}            from "./device-list.component";
import {DataService}                    from "./../services/data.service";
import {ModelService}                   from "./../services/model.service";
import {Creds}                          from "./../model/creds.model";
import {Router}                         from "../../node_modules/angular2/router";
import {AuthEventService}               from "./../services/auth.service";
import {SpinnerComponent}               from "./../utils/spinner.component";
import {HeaderComponent}                from "./../utils/header.component";
import {BeepEventService}               from "./../services/beepevent.service";
import {ServiceResponse, BeepCommand}   from "./../model/beepcommand.model";

@Component({
    selector: 'gobeepme',
    templateUrl: './app/controllers/main.component.html',
    providers: [ModelService],
    directives: [DeviceListComponent, AuthComponent, HeaderComponent, ROUTER_DIRECTIVES]
})
@RouteConfig([
    {path:'/',      name: 'Authenticate',    component: AuthComponent, useAsDefault:true},
    {path:'/list',  name: 'ListDevices',     component: DeviceListComponent}

])
export class MainComponent implements OnInit{

    message:string;

    constructor(private _modelService: ModelService,
                private _creds: Creds,
                private _beepService: BeepEventService) {
    }

    ngOnInit() {
        this._beepService.subscribe(r => this.beepRequested(r));
    }

    messageChanged(msg) {
        this.message = msg;
    }

    beepSuccess(r) {
        this._beepService.beepEventResult(true);
    }

    beepFail(r) {
        this._beepService.beepEventResult(false);
    }

    beepRequested(device) {
        if(device != null) {

            var bc:BeepCommand =
                this._beepService.generateBeepCommand(this._creds, device, this.message);

            this._modelService.beep(
                bc,
                (r:ServiceResponse) => this.beepSuccess(r),
                (r:ServiceResponse) => this.beepFail(r));
        }
    }
}
