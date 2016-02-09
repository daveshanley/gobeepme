import {Injectable, EventEmitter}           from 'angular2/core';
import {Device}                             from "../model/device.model";
import {Creds}                              from "../model/creds.model";
import {BeepCommand}                        from "../model/beepcommand.model";

export class BeepEventService {

    stream:EventEmitter<Device>;
    result:EventEmitter<boolean>;

    constructor(){
        this.stream = new EventEmitter();
        this.result = new EventEmitter();
    }

    beepEventRequest(result: Device) {
        this.stream.emit(result);
    }

    beepEventResult(result: boolean) {
        this.result.emit(result);
    }

    subscribe(s: Function, r?: Function) {
        if(s!=null) { this.stream.subscribe(s); }
        if(r!=null) { this.result.subscribe(r); }
    }

    generateBeepCommand(cr: Creds, device: Device, message: string) {
        return new BeepCommand(cr, device.Name, message)
    }
}
