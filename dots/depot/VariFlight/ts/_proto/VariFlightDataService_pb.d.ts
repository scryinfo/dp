// package: _proto
// file: VariFlightDataService.proto

import * as jspb from "google-protobuf";

export class VariFlightData extends jspb.Message {
  getFcategory(): string;
  setFcategory(value: string): void;

  getFlightno(): string;
  setFlightno(value: string): void;

  getFlightcompany(): string;
  setFlightcompany(value: string): void;

  getFlightdepcode(): string;
  setFlightdepcode(value: string): void;

  getFlightarrcode(): string;
  setFlightarrcode(value: string): void;

  getFlightdeptimeplandate(): string;
  setFlightdeptimeplandate(value: string): void;

  getFlightarrtimeplandate(): string;
  setFlightarrtimeplandate(value: string): void;

  getFlightdeptimedate(): string;
  setFlightdeptimedate(value: string): void;

  getFlightarrtimedate(): string;
  setFlightarrtimedate(value: string): void;

  getFlightstate(): string;
  setFlightstate(value: string): void;

  getFlighthterminal(): string;
  setFlighthterminal(value: string): void;

  getFlightterminal(): string;
  setFlightterminal(value: string): void;

  getOrgTimezone(): string;
  setOrgTimezone(value: string): void;

  getDstTimezone(): string;
  setDstTimezone(value: string): void;

  getShareflightno(): string;
  setShareflightno(value: string): void;

  getStopflag(): string;
  setStopflag(value: string): void;

  getShareflag(): string;
  setShareflag(value: string): void;

  getVirtualflag(): string;
  setVirtualflag(value: string): void;

  getLegflag(): string;
  setLegflag(value: string): void;

  getFlightdep(): string;
  setFlightdep(value: string): void;

  getFlightarr(): string;
  setFlightarr(value: string): void;

  getFlightdepairport(): string;
  setFlightdepairport(value: string): void;

  getFlightarrairport(): string;
  setFlightarrairport(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): VariFlightData.AsObject;
  static toObject(includeInstance: boolean, msg: VariFlightData): VariFlightData.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: VariFlightData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): VariFlightData;
  static deserializeBinaryFromReader(message: VariFlightData, reader: jspb.BinaryReader): VariFlightData;
}

export namespace VariFlightData {
  export type AsObject = {
    fcategory: string,
    flightno: string,
    flightcompany: string,
    flightdepcode: string,
    flightarrcode: string,
    flightdeptimeplandate: string,
    flightarrtimeplandate: string,
    flightdeptimedate: string,
    flightarrtimedate: string,
    flightstate: string,
    flighthterminal: string,
    flightterminal: string,
    orgTimezone: string,
    dstTimezone: string,
    shareflightno: string,
    stopflag: string,
    shareflag: string,
    virtualflag: string,
    legflag: string,
    flightdep: string,
    flightarr: string,
    flightdepairport: string,
    flightarrairport: string,
  }
}

export class GetFlightDataByFlightNumberRequest extends jspb.Message {
  getFlightnumber(): string;
  setFlightnumber(value: string): void;

  getDate(): string;
  setDate(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetFlightDataByFlightNumberRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetFlightDataByFlightNumberRequest): GetFlightDataByFlightNumberRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetFlightDataByFlightNumberRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetFlightDataByFlightNumberRequest;
  static deserializeBinaryFromReader(message: GetFlightDataByFlightNumberRequest, reader: jspb.BinaryReader): GetFlightDataByFlightNumberRequest;
}

export namespace GetFlightDataByFlightNumberRequest {
  export type AsObject = {
    flightnumber: string,
    date: string,
  }
}

export class GetFlightDataBetweenTwoAirportsRequest extends jspb.Message {
  getDepartureairport(): string;
  setDepartureairport(value: string): void;

  getArrivalairport(): string;
  setArrivalairport(value: string): void;

  getDate(): string;
  setDate(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetFlightDataBetweenTwoAirportsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetFlightDataBetweenTwoAirportsRequest): GetFlightDataBetweenTwoAirportsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetFlightDataBetweenTwoAirportsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetFlightDataBetweenTwoAirportsRequest;
  static deserializeBinaryFromReader(message: GetFlightDataBetweenTwoAirportsRequest, reader: jspb.BinaryReader): GetFlightDataBetweenTwoAirportsRequest;
}

export namespace GetFlightDataBetweenTwoAirportsRequest {
  export type AsObject = {
    departureairport: string,
    arrivalairport: string,
    date: string,
  }
}

export class GetFlightDataBetweenTwoCitiesRequest extends jspb.Message {
  getDeparturecity(): string;
  setDeparturecity(value: string): void;

  getArrivalcity(): string;
  setArrivalcity(value: string): void;

  getDate(): string;
  setDate(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetFlightDataBetweenTwoCitiesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetFlightDataBetweenTwoCitiesRequest): GetFlightDataBetweenTwoCitiesRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetFlightDataBetweenTwoCitiesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetFlightDataBetweenTwoCitiesRequest;
  static deserializeBinaryFromReader(message: GetFlightDataBetweenTwoCitiesRequest, reader: jspb.BinaryReader): GetFlightDataBetweenTwoCitiesRequest;
}

export namespace GetFlightDataBetweenTwoCitiesRequest {
  export type AsObject = {
    departurecity: string,
    arrivalcity: string,
    date: string,
  }
}

export class GetFlightDataAtOneAirportByStatusRequest extends jspb.Message {
  getAirport(): string;
  setAirport(value: string): void;

  getStatus(): string;
  setStatus(value: string): void;

  getDate(): string;
  setDate(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetFlightDataAtOneAirportByStatusRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetFlightDataAtOneAirportByStatusRequest): GetFlightDataAtOneAirportByStatusRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetFlightDataAtOneAirportByStatusRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetFlightDataAtOneAirportByStatusRequest;
  static deserializeBinaryFromReader(message: GetFlightDataAtOneAirportByStatusRequest, reader: jspb.BinaryReader): GetFlightDataAtOneAirportByStatusRequest;
}

export namespace GetFlightDataAtOneAirportByStatusRequest {
  export type AsObject = {
    airport: string,
    status: string,
    date: string,
  }
}

