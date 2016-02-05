import {Component, Input, Output, EventEmitter} from "angular2/core";
import {MessageComponent}                       from "../message.component";

@Component({
    selector:       'app-header',
    templateUrl:    './app/utils/header.component.html',
    directives:      [MessageComponent]
})
export class HeaderComponent {

    @Input() shrunk:boolean = false;
    @Output() messageEvent:EventEmitter<String> = new EventEmitter();

    messageChanged(msg) {
        this.messageEvent.emit(msg);
    }
}
