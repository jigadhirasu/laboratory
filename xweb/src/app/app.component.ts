import { Component } from '@angular/core';
import {HelloTask, HX} from './hello_pb';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'xweb';

  constructor(){
    const x = new HelloTask();
    x.setMessage('ABC');
    x.setTimestamp(123456789);
    x.setName('ffff');
    x.setXmanList([new HX(),new HX()]);

    const y = x.serializeBinary();
    console.log(y);

    const z = HelloTask.deserializeBinary(y);
    const hx = new HX();
    hx.setX(123456);
    z.addXman(hx,1);

    console.log(z.toObject());
  }
}
