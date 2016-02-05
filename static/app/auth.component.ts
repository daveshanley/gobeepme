import {Input, Output, EventEmitter, Host,Component }   from "angular2/core";
import {NgClass}                                        from 'angular2/common';
import {Creds}                                          from './model/creds.model'
import {DataService}                                    from './services/data.service'
import {ModelService}                                   from "./services/model.service";
import {Router, ROUTER_DIRECTIVES}                      from "angular2/router";
import {SpinnerComponent}                               from "./utils/spinner.component";
import {HeaderComponent}                                from "./utils/header.component";

@Component({
    templateUrl: './app/auth.component.html',
    directives: [NgClass, ROUTER_DIRECTIVES, SpinnerComponent, HeaderComponent],
    providers:  [ModelService]
})

export class AuthComponent {

    @Output authenticatedEvent = new EventEmitter();

    dataService:DataService;
    modelService:ModelService;
    model:Creds;

    constructor(dataService: DataService,
                modelService: ModelService, creds: Creds,
                private _router: Router) {
        this.dataService = dataService;
        this.modelService = modelService;
        this.model = creds;
    }

    authenticating = false;
    authenticated = false;
    error = false;
    errorMessage = ""
    fadeOut=false;

    authFailed(err) {
        this.model.password=""; // reset password
        this.fadeOut=false;
        switch (err.status) {
            case 403:
                var resp = JSON.parse(err._body);
                this.errorMessage = resp.message;
                setTimeout( () => {
                    this.authenticating = false; this.error = true;
                    setTimeout( () => {
                        this.fadeOut=true; // no animations yet!
                    }, 2000)
                }, 500);
                break;
            default:
                this.errorMessage = "Can't connect to gobeepme service!";
                this.authenticating = false; this.error = true;
                setTimeout( () => {
                    this.fadeOut=true; // no animations yet!
                }, 2000)
        }
    }
    auth() {
        this.authenticating = true;
        this.modelService.auth(this.model,
            (r) => { this.result(r)}, (r) => { this.authFailed(r)}
        );
    }

    result(result) {
        this.authenticating = false;
        this.authenticated = true;
        this.model.authenticated=true;
        this.authenticatedEvent.emit(true);
        this._router.navigate( ["ListDevices", {}] );
    }
}
