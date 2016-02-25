import {Component, Output, OnInit, EventEmitter} from "../../node_modules/angular2/core";

@Component({
    selector:       "message-input",
    templateUrl:    './app/ui/message.component.html',
})

export class MessageComponent implements OnInit {
    message:string;
    @Output()  messageEvent:EventEmitter<String> = new EventEmitter();

    ngOnInit() {
        this.message = "Beep!";
        this.messageChange(); // fire the default up the stack
    }

    messageChange() {
        this.messageEvent.emit(this.message);

    }
}
