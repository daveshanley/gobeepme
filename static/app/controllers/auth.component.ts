import {Component}   from "angular2/core";
import {Creds}                            from './../model/creds.model'
import {ModelService}                     from "./../services/model.service";
import {Router, ROUTER_DIRECTIVES}        from "angular2/router";
import {SpinnerComponent}                 from "./../utils/spinner.component";
import {AuthEventService}                 from "./../services/auth.service";

@Component({
    templateUrl: './app/controllers/auth.component.html',
    directives: [ROUTER_DIRECTIVES, SpinnerComponent],
    providers:  [ModelService]
})

export class AuthComponent {

    model:Creds;

    constructor(
                private _modelService: ModelService, creds: Creds,
                private _router: Router,
                private _authService:AuthEventService) {
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
        this._authService.authenicatedEvent(false);
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

        var cake = () => {
            this._modelService.auth(this.model,
                (r) => {
                    this.result()
                }, (r) => {
                    this.authFailed()
                }
            );
        }
        setTimeout(cake, 1000);
    }

    result() {
        this.authenticating = false;
        this.authenticated = true;
        this.model.authenticated=true;
        this._authService.authenicatedEvent(true);
        this._router.navigate( ["ListDevices", {}] );
    }
}
