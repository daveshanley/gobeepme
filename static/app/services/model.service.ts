import {Injectable}             from 'angular2/core';
import {Http}                   from 'angular2/http';
import {Observable}             from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import {Device}                 from '../model/device.model'
import {Creds}                  from '../model/creds.model'
import {DataService}            from "./data.service";


@Injectable()
export class ModelService {

    dataService:DataService;

    constructor (private http: Http, dataService: DataService) {
        this.dataService = dataService;
    }


    getDevices (creds: Creds, success, fail) {
        this.dataService.getDevices(creds)
            .subscribe(
                result  => success(result),
                error => fail(error));


    }

    auth (creds: Creds, success, fail) {
        this.dataService.auth(creds)
            .subscribe(
                result  => success(result),
                error => fail(error));


    }

}
