import {Component, OnInit}              from 'angular2/core';
import {RouteConfig, ROUTER_DIRECTIVES} from 'angular2/router';
import {AuthComponent}                  from './auth.component'
import {DeviceListComponent}            from "./device-list.component";
import {ModelService}                   from "./../services/model.service";
import {Creds}                          from "./../model/creds.model";
import {HeaderComponent}                from "./../utils/header.component";
import {BeepEventService}               from "./../services/beepevent.service";
import {ServiceResponse, BeepCommand}   from "./../model/beepcommand.model";

@Component({
    selector:   'gobeepme',
    templateUrl: './app/controllers/main.component.html',
    providers:  [ModelService],
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

    beepSuccess() {
        this._beepService.beepEventResult(true);
    }

    beepFail() {
        this._beepService.beepEventResult(false);
    }

    beepRequested(device) {
        if(device != null) {

            var bc:BeepCommand =
                this._beepService.generateBeepCommand(this._creds, device, this.message);

            this._modelService.beep(
                bc,
                (r:ServiceResponse) => this.beepSuccess(),
                (r:ServiceResponse) => this.beepFail());
        }
    }
}
