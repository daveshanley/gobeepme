import {Creds} from "./creds.model";

export class BeepCommand {
    public creds: Creds;
    public message: string;
    public name: string;
    constructor(creds: Creds, name: string, message: string) {
        this.creds = creds;
        this.name = name;
        this.message = message;
    }
    public toJSON = function(){
        return {apple_id: this.creds.appleid,
                password: this.creds.password,
                message: this.message,
                name: this.name}
    };
}

export class ServiceResponse {
    public error:boolean;
    public message:string;
}
