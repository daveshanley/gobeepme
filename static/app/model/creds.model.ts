import {Injectable} from "angular2/core";

@Injectable()
export class Creds {
    public  appleid: string;
    public password: string;
    public authenticated: boolean
    public toJSON = function(){
        return { apple_id : this.appleid, password: this.password}
    };
}
