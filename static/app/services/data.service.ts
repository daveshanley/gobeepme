import {Injectable}     from 'angular2/core';
import {Http}           from 'angular2/http';
import {Observable}     from 'rxjs/Observable';
import {Device}         from '../model/device.model'
import {Creds}          from '../model/creds.model'
import {BeepCommand}    from "../model/beepcommand.model";
import                       'rxjs/add/operator/map';



@Injectable()
export class DataService {

    constructor (private http: Http) {}

    getDevices (creds: Creds) {
        return this.auth(creds);
    }

    auth (creds: Creds) {
        var cr = creds.toJSON()
        return this.http.post('https://127.0.0.1:9443', JSON.stringify(cr))
            .map(res => <Device[]> res.json())
    }

    beep (bc) {
     return this.http.post('https://127.0.0.1:9443/beep', JSON.stringify(bc))
          .map(res => res.json())
    }
}
