import { Component } from '@angular/core';
import {HelloTask, HX} from '../pb/hello_pb';
import { HelloServiceClient } from 'src/pb/HelloServiceClientPb';

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

    const h1=new HX();
    h1.setH('128');
    const h2=new HX();
    h2.setX(128);
    x.setXmanList([h1,h2]);

    const y = x.serializeBinary();
    console.log(y);

    const z = HelloTask.deserializeBinary(y);
    const hx = new HX();
    hx.setX(123456);
    z.addXman(hx,1);

    console.log(z.toObject());

    const service = new HelloServiceClient('http://localhost:8080');
    service.xHello(new HelloTask(),null,(err,response)=>{
      console.log(response);
    });

    const stream = service.sayHello(new HelloTask());
    stream.on('status',status=>{console.log(status)});
    stream.on('error',err=>{console.log(err)});
    stream.on('end',()=>{console.log('end')});
    stream.on('data',response=>{console.log(response)});

    setTimeout(()=>{
      stream.cancel();
    },5000)
  }
}
