// package: _proto
// file: VariFlightDataService.proto

import * as VariFlightDataService_pb from "./VariFlightDataService_pb";
import {grpc} from "@improbable-eng/grpc-web";

type VariFlightDataServiceGetFlightDataByFlightNumber = {
  readonly methodName: string;
  readonly service: typeof VariFlightDataService;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof VariFlightDataService_pb.GetFlightDataByFlightNumberRequest;
  readonly responseType: typeof VariFlightDataService_pb.VariFlightData;
};

type VariFlightDataServiceGetFlightDataBetweenTwoAirports = {
  readonly methodName: string;
  readonly service: typeof VariFlightDataService;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof VariFlightDataService_pb.GetFlightDataBetweenTwoAirportsRequest;
  readonly responseType: typeof VariFlightDataService_pb.VariFlightData;
};

type VariFlightDataServiceGetFlightDataBetweenTwoCities = {
  readonly methodName: string;
  readonly service: typeof VariFlightDataService;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof VariFlightDataService_pb.GetFlightDataBetweenTwoCitiesRequest;
  readonly responseType: typeof VariFlightDataService_pb.VariFlightData;
};

type VariFlightDataServiceGetFlightDataByDepartureAndArrivalStatus = {
  readonly methodName: string;
  readonly service: typeof VariFlightDataService;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof VariFlightDataService_pb.GetFlightDataAtOneAirportByStatusRequest;
  readonly responseType: typeof VariFlightDataService_pb.VariFlightData;
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
  getFlightDataByFlightNumber(requestMessage: VariFlightDataService_pb.GetFlightDataByFlightNumberRequest, metadata?: grpc.Metadata): ResponseStream<VariFlightDataService_pb.VariFlightData>;
  getFlightDataBetweenTwoAirports(requestMessage: VariFlightDataService_pb.GetFlightDataBetweenTwoAirportsRequest, metadata?: grpc.Metadata): ResponseStream<VariFlightDataService_pb.VariFlightData>;
  getFlightDataBetweenTwoCities(requestMessage: VariFlightDataService_pb.GetFlightDataBetweenTwoCitiesRequest, metadata?: grpc.Metadata): ResponseStream<VariFlightDataService_pb.VariFlightData>;
  getFlightDataByDepartureAndArrivalStatus(requestMessage: VariFlightDataService_pb.GetFlightDataAtOneAirportByStatusRequest, metadata?: grpc.Metadata): ResponseStream<VariFlightDataService_pb.VariFlightData>;
}

