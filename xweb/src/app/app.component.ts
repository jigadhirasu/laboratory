import { Component } from '@angular/core';
import { HelloTask } from '../pb/hello_pb';
import { HelloServiceClient } from 'src/pb/HelloServiceClientPb';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'xweb';

  constructor() {
    const service = new HelloServiceClient('https://grpc.qqegg.me', {}, {});


    const h1 = new HelloTask();
    // h1.setName('angular');
    // const s1 = service.hello(h1, {}, (err, response) => {
    //   console.log(err, response.toObject());
    // });
    // s1.on('status', (status) => { console.log(status); });
    // s1.on('data', (response) => { console.log(response.toObject()); });
    // s1.on('end', () => { console.log('request end'); });

    const stream = service.helloStream(h1);
    stream.on('data', (response) => {
      console.log(response);
    });
    stream.on('status', (status) => {
      console.log(status);
    });
    stream.on('error', (err) => {
      console.log(err);
    });

  }
}
