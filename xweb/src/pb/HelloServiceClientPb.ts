/**
 * @fileoverview gRPC-Web generated client stub for hello
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import {
  Empty,
  HelloTask} from './hello_pb';

export class HelloServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials: null | { [index: string]: string; },
               options: null | { [index: string]: string; }) {
    if (!options) options = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoHelloStream = new grpcWeb.AbstractClientBase.MethodInfo(
    HelloTask,
    (request: HelloTask) => {
      return request.serializeBinary();
    },
    HelloTask.deserializeBinary
  );

  helloStream(
    request: HelloTask,
    metadata?: grpcWeb.Metadata) {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/hello.HelloService/HelloStream',
      request,
      metadata || {},
      this.methodInfoHelloStream);
  }

  methodInfoHello = new grpcWeb.AbstractClientBase.MethodInfo(
    HelloTask,
    (request: Empty) => {
      return request.serializeBinary();
    },
    HelloTask.deserializeBinary
  );

  hello(
    request: Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: HelloTask) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/hello.HelloService/Hello',
      request,
      metadata || {},
      this.methodInfoHello,
      callback);
  }

}

