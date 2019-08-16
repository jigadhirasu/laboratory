import { Component } from '@angular/core';
import {HelloTask} from '../pb/hello_pb';
import { HelloServiceClient } from 'src/pb/HelloServiceClientPb';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'xweb';

  constructor(){
    const service = new HelloServiceClient("http://grpc.qqegg.me");


    const h1 = new HelloTask();
    h1.setName('angular');
    service.hello(h1,{},(err,response)=>{
      console.log(h1);
      console.log(err,response);
    });
  }
}
