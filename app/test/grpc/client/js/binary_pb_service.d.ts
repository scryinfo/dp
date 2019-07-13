// package: api
// file: binary.proto

import * as binary_pb from "./binary_pb";
import {grpc} from "@improbable-eng/grpc-web";

type BinaryServiceSubscribeEvent = {
  readonly methodName: string;
  readonly service: typeof BinaryService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof binary_pb.SubscribeInfo;
  readonly responseType: typeof binary_pb.Result;
};

type BinaryServiceUnSubscribeEvent = {
  readonly methodName: string;
  readonly service: typeof BinaryService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof binary_pb.SubscribeInfo;
  readonly responseType: typeof binary_pb.Result;
};

type BinaryServiceRecvEvents = {
  readonly methodName: string;
  readonly service: typeof BinaryService;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof binary_pb.ClientInfo;
  readonly responseType: typeof binary_pb.Event;
};

type BinaryServicePublish = {
  readonly methodName: string;
  readonly service: typeof BinaryService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof binary_pb.PublishParams;
  readonly responseType: typeof binary_pb.PublishResult;
};

type BinaryServicePrepareToBuy = {
  readonly methodName: string;
  readonly service: typeof BinaryService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof binary_pb.PrepareParams;
  readonly responseType: typeof binary_pb.Result;
};

type BinaryServiceBuyData = {
  readonly methodName: string;
  readonly service: typeof BinaryService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof binary_pb.BuyParams;
  readonly responseType: typeof binary_pb.Result;
};

type BinaryServiceCancelTransaction = {
  readonly methodName: string;
  readonly service: typeof BinaryService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof binary_pb.CancelTxParams;
  readonly responseType: typeof binary_pb.Result;
};

type BinaryServiceSubmitMetaDataIdEncWithBuyer = {
  readonly methodName: string;
  readonly service: typeof BinaryService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof binary_pb.SubmitMetaDataIdParams;
  readonly responseType: typeof binary_pb.Result;
};

type BinaryServiceConfirmDataTruth = {
  readonly methodName: string;
  readonly service: typeof BinaryService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof binary_pb.DataConfirmParams;
  readonly responseType: typeof binary_pb.Result;
};

type BinaryServiceApproveTransfer = {
  readonly methodName: string;
  readonly service: typeof BinaryService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof binary_pb.ApproveTransferParams;
  readonly responseType: typeof binary_pb.Result;
};

type BinaryServiceVote = {
  readonly methodName: string;
  readonly service: typeof BinaryService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof binary_pb.VoteParams;
  readonly responseType: typeof binary_pb.Result;
};

type BinaryServiceRegisterAsVerifier = {
  readonly methodName: string;
  readonly service: typeof BinaryService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof binary_pb.RegisterVerifierParams;
  readonly responseType: typeof binary_pb.Result;
};

type BinaryServiceCreditsToVerifier = {
  readonly methodName: string;
  readonly service: typeof BinaryService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof binary_pb.CreditVerifierParams;
  readonly responseType: typeof binary_pb.Result;
};

type BinaryServiceTransferTokens = {
  readonly methodName: string;
  readonly service: typeof BinaryService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof binary_pb.TransferTokenParams;
  readonly responseType: typeof binary_pb.Result;
};

type BinaryServiceGetTokenBalance = {
  readonly methodName: string;
  readonly service: typeof BinaryService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof binary_pb.TokenBalanceParams;
  readonly responseType: typeof binary_pb.TokenBalanceResult;
};

type BinaryServiceCreateAccount = {
  readonly methodName: string;
  readonly service: typeof BinaryService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof binary_pb.CreateAccountParams;
  readonly responseType: typeof binary_pb.AccountResult;
};

type BinaryServiceAuthenticate = {
  readonly methodName: string;
  readonly service: typeof BinaryService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof binary_pb.ClientInfo;
  readonly responseType: typeof binary_pb.Result;
};

type BinaryServiceTransferEth = {
  readonly methodName: string;
  readonly service: typeof BinaryService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof binary_pb.TransferEthParams;
  readonly responseType: typeof binary_pb.Result;
};

type BinaryServiceGetEthBalance = {
  readonly methodName: string;
  readonly service: typeof BinaryService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof binary_pb.EthBalanceParams;
  readonly responseType: typeof binary_pb.EthBalanceResult;
};

export class BinaryService {
  static readonly serviceName: string;
  static readonly SubscribeEvent: BinaryServiceSubscribeEvent;
  static readonly UnSubscribeEvent: BinaryServiceUnSubscribeEvent;
  static readonly RecvEvents: BinaryServiceRecvEvents;
  static readonly Publish: BinaryServicePublish;
  static readonly PrepareToBuy: BinaryServicePrepareToBuy;
  static readonly BuyData: BinaryServiceBuyData;
  static readonly CancelTransaction: BinaryServiceCancelTransaction;
  static readonly SubmitMetaDataIdEncWithBuyer: BinaryServiceSubmitMetaDataIdEncWithBuyer;
  static readonly ConfirmDataTruth: BinaryServiceConfirmDataTruth;
  static readonly ApproveTransfer: BinaryServiceApproveTransfer;
  static readonly Vote: BinaryServiceVote;
  static readonly RegisterAsVerifier: BinaryServiceRegisterAsVerifier;
  static readonly CreditsToVerifier: BinaryServiceCreditsToVerifier;
  static readonly TransferTokens: BinaryServiceTransferTokens;
  static readonly GetTokenBalance: BinaryServiceGetTokenBalance;
  static readonly CreateAccount: BinaryServiceCreateAccount;
  static readonly Authenticate: BinaryServiceAuthenticate;
  static readonly TransferEth: BinaryServiceTransferEth;
  static readonly GetEthBalance: BinaryServiceGetEthBalance;
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

export class BinaryServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  subscribeEvent(
    requestMessage: binary_pb.SubscribeInfo,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  subscribeEvent(
    requestMessage: binary_pb.SubscribeInfo,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  unSubscribeEvent(
    requestMessage: binary_pb.SubscribeInfo,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  unSubscribeEvent(
    requestMessage: binary_pb.SubscribeInfo,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  recvEvents(requestMessage: binary_pb.ClientInfo, metadata?: grpc.Metadata): ResponseStream<binary_pb.Event>;
  publish(
    requestMessage: binary_pb.PublishParams,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: binary_pb.PublishResult|null) => void
  ): UnaryResponse;
  publish(
    requestMessage: binary_pb.PublishParams,
    callback: (error: ServiceError|null, responseMessage: binary_pb.PublishResult|null) => void
  ): UnaryResponse;
  prepareToBuy(
    requestMessage: binary_pb.PrepareParams,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  prepareToBuy(
    requestMessage: binary_pb.PrepareParams,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  buyData(
    requestMessage: binary_pb.BuyParams,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  buyData(
    requestMessage: binary_pb.BuyParams,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  cancelTransaction(
    requestMessage: binary_pb.CancelTxParams,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  cancelTransaction(
    requestMessage: binary_pb.CancelTxParams,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  submitMetaDataIdEncWithBuyer(
    requestMessage: binary_pb.SubmitMetaDataIdParams,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  submitMetaDataIdEncWithBuyer(
    requestMessage: binary_pb.SubmitMetaDataIdParams,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  confirmDataTruth(
    requestMessage: binary_pb.DataConfirmParams,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  confirmDataTruth(
    requestMessage: binary_pb.DataConfirmParams,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  approveTransfer(
    requestMessage: binary_pb.ApproveTransferParams,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  approveTransfer(
    requestMessage: binary_pb.ApproveTransferParams,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  vote(
    requestMessage: binary_pb.VoteParams,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  vote(
    requestMessage: binary_pb.VoteParams,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  registerAsVerifier(
    requestMessage: binary_pb.RegisterVerifierParams,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  registerAsVerifier(
    requestMessage: binary_pb.RegisterVerifierParams,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  creditsToVerifier(
    requestMessage: binary_pb.CreditVerifierParams,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  creditsToVerifier(
    requestMessage: binary_pb.CreditVerifierParams,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  transferTokens(
    requestMessage: binary_pb.TransferTokenParams,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  transferTokens(
    requestMessage: binary_pb.TransferTokenParams,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  getTokenBalance(
    requestMessage: binary_pb.TokenBalanceParams,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: binary_pb.TokenBalanceResult|null) => void
  ): UnaryResponse;
  getTokenBalance(
    requestMessage: binary_pb.TokenBalanceParams,
    callback: (error: ServiceError|null, responseMessage: binary_pb.TokenBalanceResult|null) => void
  ): UnaryResponse;
  createAccount(
    requestMessage: binary_pb.CreateAccountParams,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: binary_pb.AccountResult|null) => void
  ): UnaryResponse;
  createAccount(
    requestMessage: binary_pb.CreateAccountParams,
    callback: (error: ServiceError|null, responseMessage: binary_pb.AccountResult|null) => void
  ): UnaryResponse;
  authenticate(
    requestMessage: binary_pb.ClientInfo,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  authenticate(
    requestMessage: binary_pb.ClientInfo,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  transferEth(
    requestMessage: binary_pb.TransferEthParams,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  transferEth(
    requestMessage: binary_pb.TransferEthParams,
    callback: (error: ServiceError|null, responseMessage: binary_pb.Result|null) => void
  ): UnaryResponse;
  getEthBalance(
    requestMessage: binary_pb.EthBalanceParams,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: binary_pb.EthBalanceResult|null) => void
  ): UnaryResponse;
  getEthBalance(
    requestMessage: binary_pb.EthBalanceParams,
    callback: (error: ServiceError|null, responseMessage: binary_pb.EthBalanceResult|null) => void
  ): UnaryResponse;
}

