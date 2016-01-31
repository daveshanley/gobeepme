import {bootstrap}          from 'angular2/platform/browser'
import {MainComponent}      from './main.component'
import {ROUTER_PROVIDERS}   from 'angular2/router';
import {HTTP_PROVIDERS}     from 'angular2/http';
import {DataService}        from './services/data.service';
import {Creds}              from "./model/creds.model";
import {ModelService}       from "./services/model.service";

bootstrap(MainComponent, [
    ROUTER_PROVIDERS, HTTP_PROVIDERS, DataService, Creds, ModelService
]);

