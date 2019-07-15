// package: api
// file: auth.proto

import * as jspb from "google-protobuf";

export class ImportParameter extends jspb.Message {
  getContentPassword(): string;
  setContentPassword(value: string): void;

  getImportPsd(): string;
  setImportPsd(value: string): void;

  getContent(): Uint8Array | string;
  getContent_asU8(): Uint8Array;
  getContent_asB64(): string;
  setContent(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ImportParameter.AsObject;
  static toObject(includeInstance: boolean, msg: ImportParameter): ImportParameter.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ImportParameter, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ImportParameter;
  static deserializeBinaryFromReader(message: ImportParameter, reader: jspb.BinaryReader): ImportParameter;
}

export namespace ImportParameter {
  export type AsObject = {
    contentPassword: string,
    importPsd: string,
    content: Uint8Array | string,
  }
}

export class AddressParameter extends jspb.Message {
  getPassword(): string;
  setPassword(value: string): void;

  getAddress(): string;
  setAddress(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddressParameter.AsObject;
  static toObject(includeInstance: boolean, msg: AddressParameter): AddressParameter.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AddressParameter, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddressParameter;
  static deserializeBinaryFromReader(message: AddressParameter, reader: jspb.BinaryReader): AddressParameter;
}

export namespace AddressParameter {
  export type AsObject = {
    password: string,
    address: string,
  }
}

export class AddressInfo extends jspb.Message {
  getStatus(): StatusMap[keyof StatusMap];
  setStatus(value: StatusMap[keyof StatusMap]): void;

  getAddress(): string;
  setAddress(value: string): void;

  getMsg(): string;
  setMsg(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddressInfo.AsObject;
  static toObject(includeInstance: boolean, msg: AddressInfo): AddressInfo.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AddressInfo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddressInfo;
  static deserializeBinaryFromReader(message: AddressInfo, reader: jspb.BinaryReader): AddressInfo;
}

export namespace AddressInfo {
  export type AsObject = {
    status: StatusMap[keyof StatusMap],
    address: string,
    msg: string,
  }
}

export class CipherParameter extends jspb.Message {
  getPassword(): string;
  setPassword(value: string): void;

  getAddress(): string;
  setAddress(value: string): void;

  getMessage(): Uint8Array | string;
  getMessage_asU8(): Uint8Array;
  getMessage_asB64(): string;
  setMessage(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CipherParameter.AsObject;
  static toObject(includeInstance: boolean, msg: CipherParameter): CipherParameter.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CipherParameter, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CipherParameter;
  static deserializeBinaryFromReader(message: CipherParameter, reader: jspb.BinaryReader): CipherParameter;
}

export namespace CipherParameter {
  export type AsObject = {
    password: string,
    address: string,
    message: Uint8Array | string,
  }
}

export class CipherText extends jspb.Message {
  getStatus(): StatusMap[keyof StatusMap];
  setStatus(value: StatusMap[keyof StatusMap]): void;

  getData(): Uint8Array | string;
  getData_asU8(): Uint8Array;
  getData_asB64(): string;
  setData(value: Uint8Array | string): void;

  getMsg(): string;
  setMsg(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CipherText.AsObject;
  static toObject(includeInstance: boolean, msg: CipherText): CipherText.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CipherText, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CipherText;
  static deserializeBinaryFromReader(message: CipherText, reader: jspb.BinaryReader): CipherText;
}

export namespace CipherText {
  export type AsObject = {
    status: StatusMap[keyof StatusMap],
    data: Uint8Array | string,
    msg: string,
  }
}

export interface StatusMap {
  OK: 0;
  ERROR: 1;
}

export const Status: StatusMap;

