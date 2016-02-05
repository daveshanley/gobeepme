import {Component, Input}       from "angular2/core";
import {OnInit} from "angular2/core";

@Component({
    selector: "location",
    templateUrl: './app/location.component.html'
})
export class LocationComponent implements OnInit {

    apiKey:string = "AIzaSyBSlQ9m3DiaaiJB1l1ceSMAPZ_ZMD0jcUw";
    smallUrl:string;
    largeUrl:string;


    @Input() lon: number;
    @Input() lat: number;
    @Input() name: string;
    @Input() width:number;
    @Input() height: number;


    buildMarker() {
        var pipe = function() {
            return encodeURIComponent('|');
        }
        return "markers=size:medium" + pipe() + "color:red" + pipe() + "label:"+this.name+ pipe() + this.lat + "," + this.lon;
    }

    buildSmallUrl() {
        this.smallUrl = "https://maps.googleapis.com/maps/api/staticmap?center=" +
            this.lat + "," + this.lon + "&zoom=13&size="+ (this.width/3) + "x"+ (this.height/3)+ "&scale=1&"+this.buildMarker()+"&key=" + this.apiKey;
    }

    buildLargeUrl() {
        this.largeUrl = "https://maps.googleapis.com/maps/api/staticmap?center=" +
            this.lat + "," + this.lon + "&zoom=15&size="+ this.width + "x"+ this.height+ "&scale=3&"+this.buildMarker()+"&key=" + this.apiKey;
    }

    ngOnInit() {
        this.buildSmallUrl();
        this.buildLargeUrl();
    }
}

/*


    https://maps.googleapis.com/maps/api/staticmap?center=40.714728,-73.998672&zoom=14&size=400x400&key=YOUR_API_KEY

*/
