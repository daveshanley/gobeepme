import {Input, Output, EventEmitter, Component, OnInit}     from "angular2/core";
import {DataService}                                        from "./services/data.service";
import {Creds}                                              from "./model/creds.model";
import {Router}                                             from "angular2/router";

@Component({
    selector: 'device-list',
    template: `<h1>devices</h1>`

})

export class DeviceListComponent implements  OnInit {

    constructor(dataService: DataService, private _creds: Creds, private _router: Router) {

    }

    ngOnInit() {
        if(!this._creds.authenticated) {
            this._router.navigate( ["Authenticate", {}] );
        }
    }
}
