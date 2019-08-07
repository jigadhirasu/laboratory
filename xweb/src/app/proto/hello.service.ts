import { Injectable } from '@angular/core';
import { HelloServiceClient } from '../pb/HelloServiceClientPb';
import { HelloTask } from '../pb/hello_pb';


@Injectable({
  providedIn: 'root'
})
export class HelloService {

  constructor() { 
    const service = new HelloServiceClient('http://localhost:17887',null,null);
    const request = new HelloTask();
    request.setName('TMD');
    request.setMessage('Message');

    const stream = service.sayHello(request);

    stream.on('data',(response)=>{
      console.log(response.getResponse());
    });

    stream.on('status',(status)=>{
      console.log('status : ',status);
    });
  }
}
