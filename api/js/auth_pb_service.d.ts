// package: api
// file: auth.proto

import * as auth_pb from "./auth_pb";
import {grpc} from "@improbable-eng/grpc-web";

type KeyServiceGenerateAddress = {
  readonly methodName: string;
  readonly service: typeof KeyService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof auth_pb.AddressParameter;
  readonly responseType: typeof auth_pb.AddressInfo;
};

type KeyServiceVerifyAddress = {
  readonly methodName: string;
  readonly service: typeof KeyService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof auth_pb.AddressParameter;
  readonly responseType: typeof auth_pb.AddressInfo;
};

type KeyServiceContentEncrypt = {
  readonly methodName: string;
  readonly service: typeof KeyService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof auth_pb.CipherParameter;
  readonly responseType: typeof auth_pb.CipherText;
};

type KeyServiceContentDecrypt = {
  readonly methodName: string;
  readonly service: typeof KeyService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof auth_pb.CipherParameter;
  readonly responseType: typeof auth_pb.CipherText;
};

type KeyServiceSignature = {
  readonly methodName: string;
  readonly service: typeof KeyService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof auth_pb.CipherParameter;
  readonly responseType: typeof auth_pb.CipherText;
};

type KeyServiceimport_keystore = {
  readonly methodName: string;
  readonly service: typeof KeyService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof auth_pb.ImportParameter;
  readonly responseType: typeof auth_pb.AddressInfo;
};

export class KeyService {
  static readonly serviceName: string;
  static readonly GenerateAddress: KeyServiceGenerateAddress;
  static readonly VerifyAddress: KeyServiceVerifyAddress;
  static readonly ContentEncrypt: KeyServiceContentEncrypt;
  static readonly ContentDecrypt: KeyServiceContentDecrypt;
  static readonly Signature: KeyServiceSignature;
  static readonly import_keystore: KeyServiceimport_keystore;
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

export class KeyServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  generateAddress(
    requestMessage: auth_pb.AddressParameter,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: auth_pb.AddressInfo|null) => void
  ): UnaryResponse;
  generateAddress(
    requestMessage: auth_pb.AddressParameter,
    callback: (error: ServiceError|null, responseMessage: auth_pb.AddressInfo|null) => void
  ): UnaryResponse;
  verifyAddress(
    requestMessage: auth_pb.AddressParameter,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: auth_pb.AddressInfo|null) => void
  ): UnaryResponse;
  verifyAddress(
    requestMessage: auth_pb.AddressParameter,
    callback: (error: ServiceError|null, responseMessage: auth_pb.AddressInfo|null) => void
  ): UnaryResponse;
  contentEncrypt(
    requestMessage: auth_pb.CipherParameter,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: auth_pb.CipherText|null) => void
  ): UnaryResponse;
  contentEncrypt(
    requestMessage: auth_pb.CipherParameter,
    callback: (error: ServiceError|null, responseMessage: auth_pb.CipherText|null) => void
  ): UnaryResponse;
  contentDecrypt(
    requestMessage: auth_pb.CipherParameter,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: auth_pb.CipherText|null) => void
  ): UnaryResponse;
  contentDecrypt(
    requestMessage: auth_pb.CipherParameter,
    callback: (error: ServiceError|null, responseMessage: auth_pb.CipherText|null) => void
  ): UnaryResponse;
  signature(
    requestMessage: auth_pb.CipherParameter,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: auth_pb.CipherText|null) => void
  ): UnaryResponse;
  signature(
    requestMessage: auth_pb.CipherParameter,
    callback: (error: ServiceError|null, responseMessage: auth_pb.CipherText|null) => void
  ): UnaryResponse;
  import_keystore(
    requestMessage: auth_pb.ImportParameter,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: auth_pb.AddressInfo|null) => void
  ): UnaryResponse;
  import_keystore(
    requestMessage: auth_pb.ImportParameter,
    callback: (error: ServiceError|null, responseMessage: auth_pb.AddressInfo|null) => void
  ): UnaryResponse;
}

