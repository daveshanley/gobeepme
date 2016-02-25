import {bootstrap}          from 'angular2/platform/browser'
import {MainComponent}      from './controllers/main.component.ts'
import {ROUTER_PROVIDERS}   from 'angular2/router';
import {HTTP_PROVIDERS}     from 'angular2/http';
import {DataService}        from './services/data.service';
import {Creds}              from "./model/creds.model";
import {ModelService}       from "./services/model.service";
import {AuthEventService}   from "./services/auth.service";
import {enableProdMode}     from 'angular2/core';
import {BeepEventService}   from "./services/beepevent.service";

//enableProdMode();

bootstrap(MainComponent, [
    ROUTER_PROVIDERS, HTTP_PROVIDERS, DataService, Creds, ModelService, AuthEventService, BeepEventService
]);

