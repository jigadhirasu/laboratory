import * as jspb from "google-protobuf"

export class HelloTask extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getTimestamp(): number;
  setTimestamp(value: number): void;

  getResponse(): string;
  setResponse(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HelloTask.AsObject;
  static toObject(includeInstance: boolean, msg: HelloTask): HelloTask.AsObject;
  static serializeBinaryToWriter(message: HelloTask, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HelloTask;
  static deserializeBinaryFromReader(message: HelloTask, reader: jspb.BinaryReader): HelloTask;
}

export namespace HelloTask {
  export type AsObject = {
    name: string,
    timestamp: number,
    response: string,
  }
}

export class Empty extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Empty.AsObject;
  static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
  static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Empty;
  static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
  export type AsObject = {
  }
}

