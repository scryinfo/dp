// package: protobuf
// file: variflight.proto

import * as variflight_pb from "./variflight_pb";
import {grpc} from "@improbable-eng/grpc-web";

type VariFlightDataServiceGetFlightDataByFlightNumber = {
  readonly methodName: string;
  readonly service: typeof VariFlightDataService;
  readonly requestStream: true;
  readonly responseStream: true;
  readonly requestType: typeof variflight_pb.GetFlightDataByFlightNumberRequest;
  readonly responseType: typeof variflight_pb.VariFlightData;
};

type VariFlightDataServiceGetFlightDataBetweenTwoAirports = {
  readonly methodName: string;
  readonly service: typeof VariFlightDataService;
  readonly requestStream: true;
  readonly responseStream: true;
  readonly requestType: typeof variflight_pb.GetFlightDataBetweenTwoAirportsRequest;
  readonly responseType: typeof variflight_pb.VariFlightData;
};

type VariFlightDataServiceGetFlightDataBetweenTwoCities = {
  readonly methodName: string;
  readonly service: typeof VariFlightDataService;
  readonly requestStream: true;
  readonly responseStream: true;
  readonly requestType: typeof variflight_pb.GetFlightDataBetweenTwoCitiesRequest;
  readonly responseType: typeof variflight_pb.VariFlightData;
};

type VariFlightDataServiceGetFlightDataByDepartureAndArrivalStatus = {
  readonly methodName: string;
  readonly service: typeof VariFlightDataService;
  readonly requestStream: true;
  readonly responseStream: true;
  readonly requestType: typeof variflight_pb.GetFlightDataAtOneAirportByStatusRequest;
  readonly responseType: typeof variflight_pb.VariFlightData;
};

export class VariFlightDataService {
  static readonly serviceName: string;
  static readonly GetFlightDataByFlightNumber: VariFlightDataServiceGetFlightDataByFlightNumber;
  static readonly GetFlightDataBetweenTwoAirports: VariFlightDataServiceGetFlightDataBetweenTwoAirports;
  static readonly GetFlightDataBetweenTwoCities: VariFlightDataServiceGetFlightDataBetweenTwoCities;
  static readonly GetFlightDataByDepartureAndArrivalStatus: VariFlightDataServiceGetFlightDataByDepartureAndArrivalStatus;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class VariFlightDataServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  getFlightDataByFlightNumber(metadata?: grpc.Metadata): BidirectionalStream<variflight_pb.GetFlightDataByFlightNumberRequest, variflight_pb.VariFlightData>;
  getFlightDataBetweenTwoAirports(metadata?: grpc.Metadata): BidirectionalStream<variflight_pb.GetFlightDataBetweenTwoAirportsRequest, variflight_pb.VariFlightData>;
  getFlightDataBetweenTwoCities(metadata?: grpc.Metadata): BidirectionalStream<variflight_pb.GetFlightDataBetweenTwoCitiesRequest, variflight_pb.VariFlightData>;
  getFlightDataByDepartureAndArrivalStatus(metadata?: grpc.Metadata): BidirectionalStream<variflight_pb.GetFlightDataAtOneAirportByStatusRequest, variflight_pb.VariFlightData>;
}

