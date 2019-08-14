/**
 * @fileoverview gRPC-Web generated client stub for hello
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import {HelloTask} from './hello_pb';

export class HelloServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'binary';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoSayHello = new grpcWeb.AbstractClientBase.MethodInfo(
    HelloTask,
    (request: HelloTask) => {
      return request.serializeBinary();
    },
    HelloTask.deserializeBinary
  );

  sayHello(
    request: HelloTask,
    metadata?: grpcWeb.Metadata) {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/hello.HelloService/SayHello',
      request,
      metadata || {},
      this.methodInfoSayHello);
  }

  methodInfoXHello = new grpcWeb.AbstractClientBase.MethodInfo(
    HelloTask,
    (request: HelloTask) => {
      return request.serializeBinary();
    },
    HelloTask.deserializeBinary
  );

  xHello(
    request: HelloTask,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: HelloTask) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/hello.HelloService/XHello',
      request,
      metadata || {},
      this.methodInfoXHello,
      callback);
  }

}

