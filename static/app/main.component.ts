import {Component}                      from 'angular2/core';
import {RouteConfig, ROUTER_DIRECTIVES} from 'angular2/router';
import {AuthComponent}                  from './auth.component'
import {Device}                         from './model/device.model'
import {DeviceListComponent}            from "./device-list.component";
import {DataService}                    from "./services/data.service";
import {ModelService}                   from "./services/model.service";
import {Creds}                          from "./model/creds.model";
import {Router}                         from "angular2/router";

@Component({
    selector: 'gobeepme',
    templateUrl: './app/main.component.html',
    providers: [DataService, ModelService],
    directives: [DeviceListComponent,AuthComponent, ROUTER_DIRECTIVES]
})
@RouteConfig([
    {path:'/',      name: 'Authenticate',    component: AuthComponent, useAsDefault:true},
    {path:'/list',  name: 'ListDevices',     component: DeviceListComponent}

])

export class MainComponent {
    modelService:modelService;

    constructor(modelService: ModelService, private _creds: Creds, private _router: Router) {
        this.modelService = modelService;
    }

    authenticatedEvent() {
        console.log('AUTHENTICATED');
        this._router.navigate( ["ListDevices", {}] );

    }
}
