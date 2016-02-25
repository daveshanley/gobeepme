import {Component, Input, Output,
        EventEmitter, OnInit}       from "angular2/core";
import {MessageComponent}           from "../ui/message.component";
import {AuthEventService}           from "../services/auth.service";
import {BeepEventService}           from "../services/beepevent.service";

@Component({
    selector:       'app-header',
    templateUrl:    './app/utils/header.component.html',
    directives:      [MessageComponent]
})
export class HeaderComponent implements OnInit {

    @Input() shrunk:boolean = false;
    @Input() beeping:boolean = false;
    @Output() messageEvent:EventEmitter<String> = new EventEmitter();

    constructor(private _authService: AuthEventService,
                private _beepService: BeepEventService) { }

    messageChanged(msg) {
        this.messageEvent.emit(msg); // bubble up
    }

    ngOnInit() {
        this._authService.subscribe(r => this.authResult(r));
        this._beepService.subscribe(r => this.beepRequested(r));
    }

    authResult(r) {
        if(r) {
            this.shrunk = true;
        }
    }

    beepRequested(r) {
        if(r != null) {
            this.beeping = true;
            setTimeout(() => this.beeping = false, 1100);
        }
    }
}
