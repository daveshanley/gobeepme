import {Injectable, EventEmitter}         from 'angular2/core';

export class AuthEventService {

    Stream:EventEmitter<boolean>;

    constructor(){
        this.Stream = new EventEmitter();
    }

    authenicatedEvent(result: boolean) {
        this.Stream.emit(result);
    }

    subscribe(s: Function) {
        this.Stream.subscribe(s);
    }
}
