import {Component, Output, OnInit, EventEmitter}       from "angular2/core";

@Component({
    selector:       "message-input",
    templateUrl:    './app/message.component.html',
})

export class MessageComponent implements OnInit {
    message:string;
    @Output()  messageUpdate:EventEmitter<String> = new EventEmitter();

    ngOnInit() {
        this.message = "Beep!";
        this.messageChange(); // fire the default up the stack
    }

    messageChange() {
        this.messageUpdate.emit(this.message);

    }
}
