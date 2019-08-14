import * as jspb from "google-protobuf"

export class HelloTask extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getMessage(): string;
  setMessage(value: string): void;

  getTimestamp(): number;
  setTimestamp(value: number): void;

  getResponse(): string;
  setResponse(value: string): void;

  getXmanList(): Array<HX>;
  setXmanList(value: Array<HX>): void;
  clearXmanList(): void;
  addXman(value?: HX, index?: number): HX;

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
    message: string,
    timestamp: number,
    response: string,
    xmanList: Array<HX.AsObject>,
  }
}

export class HX extends jspb.Message {
  getH(): string;
  setH(value: string): void;

  getX(): number;
  setX(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HX.AsObject;
  static toObject(includeInstance: boolean, msg: HX): HX.AsObject;
  static serializeBinaryToWriter(message: HX, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HX;
  static deserializeBinaryFromReader(message: HX, reader: jspb.BinaryReader): HX;
}

export namespace HX {
  export type AsObject = {
    h: string,
    x: number,
  }
}
