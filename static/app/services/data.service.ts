import {Injectable} from 'angular2/core';
import {Http}       from 'angular2/http';
import {Observable} from 'rxjs/Observable';

import {Device}     from '../model/device.model'
import {Creds}      from '../model/creds.model'
import 'rxjs/add/operator/map';


@Injectable()
export class DataService {


    constructor (private http: Http) {}

    private _authUrl = 'test.json';

    auth (creds: Creds) {
        var cr = creds.toJSON()
        //return this.http.post('https://localhost:9443', JSON.stringify(cr))
        //    .map(res => <Device[]> res.json())

        return this.http.get('test.json')
            .map(res => <Device[]> res.json())


    }


    getDevices () {
        return this.http.get(this._authUrl)
            .map(res => <Device[]> res.json().data)

    }
}
