// package: api
// file: binary.proto

import * as jspb from "google-protobuf";

export class CreateAccountParams extends jspb.Message {
  getPassword(): string;
  setPassword(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateAccountParams.AsObject;
  static toObject(includeInstance: boolean, msg: CreateAccountParams): CreateAccountParams.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateAccountParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateAccountParams;
  static deserializeBinaryFromReader(message: CreateAccountParams, reader: jspb.BinaryReader): CreateAccountParams;
}

export namespace CreateAccountParams {
  export type AsObject = {
    password: string,
  }
}

export class AccountResult extends jspb.Message {
  hasResult(): boolean;
  clearResult(): void;
  getResult(): Result | undefined;
  setResult(value?: Result): void;

  getAccountid(): string;
  setAccountid(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AccountResult.AsObject;
  static toObject(includeInstance: boolean, msg: AccountResult): AccountResult.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AccountResult, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AccountResult;
  static deserializeBinaryFromReader(message: AccountResult, reader: jspb.BinaryReader): AccountResult;
}

export namespace AccountResult {
  export type AsObject = {
    result?: Result.AsObject,
    accountid: string,
  }
}

export class TransferEthParams extends jspb.Message {
  getFrom(): string;
  setFrom(value: string): void;

  getPassword(): string;
  setPassword(value: string): void;

  getTo(): string;
  setTo(value: string): void;

  getValue(): number;
  setValue(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TransferEthParams.AsObject;
  static toObject(includeInstance: boolean, msg: TransferEthParams): TransferEthParams.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TransferEthParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TransferEthParams;
  static deserializeBinaryFromReader(message: TransferEthParams, reader: jspb.BinaryReader): TransferEthParams;
}

export namespace TransferEthParams {
  export type AsObject = {
    from: string,
    password: string,
    to: string,
    value: number,
  }
}

export class EthBalanceParams extends jspb.Message {
  getOwner(): string;
  setOwner(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EthBalanceParams.AsObject;
  static toObject(includeInstance: boolean, msg: EthBalanceParams): EthBalanceParams.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: EthBalanceParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EthBalanceParams;
  static deserializeBinaryFromReader(message: EthBalanceParams, reader: jspb.BinaryReader): EthBalanceParams;
}

export namespace EthBalanceParams {
  export type AsObject = {
    owner: string,
  }
}

export class EthBalanceResult extends jspb.Message {
  hasResult(): boolean;
  clearResult(): void;
  getResult(): Result | undefined;
  setResult(value?: Result): void;

  getBalance(): number;
  setBalance(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EthBalanceResult.AsObject;
  static toObject(includeInstance: boolean, msg: EthBalanceResult): EthBalanceResult.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: EthBalanceResult, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EthBalanceResult;
  static deserializeBinaryFromReader(message: EthBalanceResult, reader: jspb.BinaryReader): EthBalanceResult;
}

export namespace EthBalanceResult {
  export type AsObject = {
    result?: Result.AsObject,
    balance: number,
  }
}

export class ClientInfo extends jspb.Message {
  getAddress(): string;
  setAddress(value: string): void;

  getPassword(): string;
  setPassword(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ClientInfo.AsObject;
  static toObject(includeInstance: boolean, msg: ClientInfo): ClientInfo.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ClientInfo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ClientInfo;
  static deserializeBinaryFromReader(message: ClientInfo, reader: jspb.BinaryReader): ClientInfo;
}

export namespace ClientInfo {
  export type AsObject = {
    address: string,
    password: string,
  }
}

export class Event extends jspb.Message {
  getTime(): number;
  setTime(value: number): void;

  getJsondata(): string;
  setJsondata(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Event.AsObject;
  static toObject(includeInstance: boolean, msg: Event): Event.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Event, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Event;
  static deserializeBinaryFromReader(message: Event, reader: jspb.BinaryReader): Event;
}

export namespace Event {
  export type AsObject = {
    time: number,
    jsondata: string,
  }
}

export class TxParams extends jspb.Message {
  getFrom(): string;
  setFrom(value: string): void;

  getPassword(): string;
  setPassword(value: string): void;

  getValue(): number;
  setValue(value: number): void;

  getPending(): boolean;
  setPending(value: boolean): void;

  getGasprice(): number;
  setGasprice(value: number): void;

  getGaslimit(): number;
  setGaslimit(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TxParams.AsObject;
  static toObject(includeInstance: boolean, msg: TxParams): TxParams.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TxParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TxParams;
  static deserializeBinaryFromReader(message: TxParams, reader: jspb.BinaryReader): TxParams;
}

export namespace TxParams {
  export type AsObject = {
    from: string,
    password: string,
    value: number,
    pending: boolean,
    gasprice: number,
    gaslimit: number,
  }
}

export class PublishParams extends jspb.Message {
  hasTxparam(): boolean;
  clearTxparam(): void;
  getTxparam(): TxParams | undefined;
  setTxparam(value?: TxParams): void;

  getPrice(): number;
  setPrice(value: number): void;

  getMetadataid(): Uint8Array | string;
  getMetadataid_asU8(): Uint8Array;
  getMetadataid_asB64(): string;
  setMetadataid(value: Uint8Array | string): void;

  clearProofdataidsList(): void;
  getProofdataidsList(): Array<string>;
  setProofdataidsList(value: Array<string>): void;
  addProofdataids(value: string, index?: number): string;

  getProofnum(): number;
  setProofnum(value: number): void;

  getDetailsid(): string;
  setDetailsid(value: string): void;

  getSupportverify(): boolean;
  setSupportverify(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PublishParams.AsObject;
  static toObject(includeInstance: boolean, msg: PublishParams): PublishParams.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PublishParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PublishParams;
  static deserializeBinaryFromReader(message: PublishParams, reader: jspb.BinaryReader): PublishParams;
}

export namespace PublishParams {
  export type AsObject = {
    txparam?: TxParams.AsObject,
    price: number,
    metadataid: Uint8Array | string,
    proofdataidsList: Array<string>,
    proofnum: number,
    detailsid: string,
    supportverify: boolean,
  }
}

export class PublishResult extends jspb.Message {
  getPublishid(): string;
  setPublishid(value: string): void;

  hasResult(): boolean;
  clearResult(): void;
  getResult(): Result | undefined;
  setResult(value?: Result): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PublishResult.AsObject;
  static toObject(includeInstance: boolean, msg: PublishResult): PublishResult.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PublishResult, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PublishResult;
  static deserializeBinaryFromReader(message: PublishResult, reader: jspb.BinaryReader): PublishResult;
}

export namespace PublishResult {
  export type AsObject = {
    publishid: string,
    result?: Result.AsObject,
  }
}

export class PrepareParams extends jspb.Message {
  hasTxparam(): boolean;
  clearTxparam(): void;
  getTxparam(): TxParams | undefined;
  setTxparam(value?: TxParams): void;

  getPublishid(): string;
  setPublishid(value: string): void;

  getStartverify(): boolean;
  setStartverify(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PrepareParams.AsObject;
  static toObject(includeInstance: boolean, msg: PrepareParams): PrepareParams.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PrepareParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PrepareParams;
  static deserializeBinaryFromReader(message: PrepareParams, reader: jspb.BinaryReader): PrepareParams;
}

export namespace PrepareParams {
  export type AsObject = {
    txparam?: TxParams.AsObject,
    publishid: string,
    startverify: boolean,
  }
}

export class Result extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): void;

  getErrmsg(): string;
  setErrmsg(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Result.AsObject;
  static toObject(includeInstance: boolean, msg: Result): Result.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Result, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Result;
  static deserializeBinaryFromReader(message: Result, reader: jspb.BinaryReader): Result;
}

export namespace Result {
  export type AsObject = {
    success: boolean,
    errmsg: string,
  }
}

export class BuyParams extends jspb.Message {
  hasTxparam(): boolean;
  clearTxparam(): void;
  getTxparam(): TxParams | undefined;
  setTxparam(value?: TxParams): void;

  getTxid(): number;
  setTxid(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BuyParams.AsObject;
  static toObject(includeInstance: boolean, msg: BuyParams): BuyParams.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BuyParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BuyParams;
  static deserializeBinaryFromReader(message: BuyParams, reader: jspb.BinaryReader): BuyParams;
}

export namespace BuyParams {
  export type AsObject = {
    txparam?: TxParams.AsObject,
    txid: number,
  }
}

export class CancelTxParams extends jspb.Message {
  hasTxparam(): boolean;
  clearTxparam(): void;
  getTxparam(): TxParams | undefined;
  setTxparam(value?: TxParams): void;

  getTxid(): number;
  setTxid(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CancelTxParams.AsObject;
  static toObject(includeInstance: boolean, msg: CancelTxParams): CancelTxParams.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CancelTxParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CancelTxParams;
  static deserializeBinaryFromReader(message: CancelTxParams, reader: jspb.BinaryReader): CancelTxParams;
}

export namespace CancelTxParams {
  export type AsObject = {
    txparam?: TxParams.AsObject,
    txid: number,
  }
}

export class ReEncryptDataParams extends jspb.Message {
  hasTxparam(): boolean;
  clearTxparam(): void;
  getTxparam(): TxParams | undefined;
  setTxparam(value?: TxParams): void;

  getTxid(): number;
  setTxid(value: number): void;

  getEncodeddatawithseller(): Uint8Array | string;
  getEncodeddatawithseller_asU8(): Uint8Array;
  getEncodeddatawithseller_asB64(): string;
  setEncodeddatawithseller(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReEncryptDataParams.AsObject;
  static toObject(includeInstance: boolean, msg: ReEncryptDataParams): ReEncryptDataParams.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ReEncryptDataParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReEncryptDataParams;
  static deserializeBinaryFromReader(message: ReEncryptDataParams, reader: jspb.BinaryReader): ReEncryptDataParams;
}

export namespace ReEncryptDataParams {
  export type AsObject = {
    txparam?: TxParams.AsObject,
    txid: number,
    encodeddatawithseller: Uint8Array | string,
  }
}

export class DataConfirmParams extends jspb.Message {
  hasTxparam(): boolean;
  clearTxparam(): void;
  getTxparam(): TxParams | undefined;
  setTxparam(value?: TxParams): void;

  getTxid(): number;
  setTxid(value: number): void;

  getTruth(): boolean;
  setTruth(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DataConfirmParams.AsObject;
  static toObject(includeInstance: boolean, msg: DataConfirmParams): DataConfirmParams.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DataConfirmParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DataConfirmParams;
  static deserializeBinaryFromReader(message: DataConfirmParams, reader: jspb.BinaryReader): DataConfirmParams;
}

export namespace DataConfirmParams {
  export type AsObject = {
    txparam?: TxParams.AsObject,
    txid: number,
    truth: boolean,
  }
}

export class ApproveTransferParams extends jspb.Message {
  hasTxparam(): boolean;
  clearTxparam(): void;
  getTxparam(): TxParams | undefined;
  setTxparam(value?: TxParams): void;

  getSpenderaddr(): string;
  setSpenderaddr(value: string): void;

  getValue(): number;
  setValue(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ApproveTransferParams.AsObject;
  static toObject(includeInstance: boolean, msg: ApproveTransferParams): ApproveTransferParams.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ApproveTransferParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ApproveTransferParams;
  static deserializeBinaryFromReader(message: ApproveTransferParams, reader: jspb.BinaryReader): ApproveTransferParams;
}

export namespace ApproveTransferParams {
  export type AsObject = {
    txparam?: TxParams.AsObject,
    spenderaddr: string,
    value: number,
  }
}

export class VoteParams extends jspb.Message {
  hasTxparam(): boolean;
  clearTxparam(): void;
  getTxparam(): TxParams | undefined;
  setTxparam(value?: TxParams): void;

  getTxid(): number;
  setTxid(value: number): void;

  getJudge(): boolean;
  setJudge(value: boolean): void;

  getComments(): string;
  setComments(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): VoteParams.AsObject;
  static toObject(includeInstance: boolean, msg: VoteParams): VoteParams.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: VoteParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): VoteParams;
  static deserializeBinaryFromReader(message: VoteParams, reader: jspb.BinaryReader): VoteParams;
}

export namespace VoteParams {
  export type AsObject = {
    txparam?: TxParams.AsObject,
    txid: number,
    judge: boolean,
    comments: string,
  }
}

export class RegisterVerifierParams extends jspb.Message {
  hasTxparam(): boolean;
  clearTxparam(): void;
  getTxparam(): TxParams | undefined;
  setTxparam(value?: TxParams): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterVerifierParams.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterVerifierParams): RegisterVerifierParams.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RegisterVerifierParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterVerifierParams;
  static deserializeBinaryFromReader(message: RegisterVerifierParams, reader: jspb.BinaryReader): RegisterVerifierParams;
}

export namespace RegisterVerifierParams {
  export type AsObject = {
    txparam?: TxParams.AsObject,
  }
}

export class CreditVerifierParams extends jspb.Message {
  hasTxparam(): boolean;
  clearTxparam(): void;
  getTxparam(): TxParams | undefined;
  setTxparam(value?: TxParams): void;

  getTxid(): number;
  setTxid(value: number): void;

  getIndex(): number;
  setIndex(value: number): void;

  getCredit(): number;
  setCredit(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreditVerifierParams.AsObject;
  static toObject(includeInstance: boolean, msg: CreditVerifierParams): CreditVerifierParams.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreditVerifierParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreditVerifierParams;
  static deserializeBinaryFromReader(message: CreditVerifierParams, reader: jspb.BinaryReader): CreditVerifierParams;
}

export namespace CreditVerifierParams {
  export type AsObject = {
    txparam?: TxParams.AsObject,
    txid: number,
    index: number,
    credit: number,
  }
}

export class TransferTokenParams extends jspb.Message {
  hasTxparam(): boolean;
  clearTxparam(): void;
  getTxparam(): TxParams | undefined;
  setTxparam(value?: TxParams): void;

  getTo(): string;
  setTo(value: string): void;

  getValue(): number;
  setValue(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TransferTokenParams.AsObject;
  static toObject(includeInstance: boolean, msg: TransferTokenParams): TransferTokenParams.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TransferTokenParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TransferTokenParams;
  static deserializeBinaryFromReader(message: TransferTokenParams, reader: jspb.BinaryReader): TransferTokenParams;
}

export namespace TransferTokenParams {
  export type AsObject = {
    txparam?: TxParams.AsObject,
    to: string,
    value: number,
  }
}

export class TokenBalanceParams extends jspb.Message {
  hasTxparam(): boolean;
  clearTxparam(): void;
  getTxparam(): TxParams | undefined;
  setTxparam(value?: TxParams): void;

  getOwner(): string;
  setOwner(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TokenBalanceParams.AsObject;
  static toObject(includeInstance: boolean, msg: TokenBalanceParams): TokenBalanceParams.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TokenBalanceParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TokenBalanceParams;
  static deserializeBinaryFromReader(message: TokenBalanceParams, reader: jspb.BinaryReader): TokenBalanceParams;
}

export namespace TokenBalanceParams {
  export type AsObject = {
    txparam?: TxParams.AsObject,
    owner: string,
  }
}

export class TokenBalanceResult extends jspb.Message {
  getBalance(): number;
  setBalance(value: number): void;

  hasResult(): boolean;
  clearResult(): void;
  getResult(): Result | undefined;
  setResult(value?: Result): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TokenBalanceResult.AsObject;
  static toObject(includeInstance: boolean, msg: TokenBalanceResult): TokenBalanceResult.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TokenBalanceResult, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TokenBalanceResult;
  static deserializeBinaryFromReader(message: TokenBalanceResult, reader: jspb.BinaryReader): TokenBalanceResult;
}

export namespace TokenBalanceResult {
  export type AsObject = {
    balance: number,
    result?: Result.AsObject,
  }
}

export class SubscribeInfo extends jspb.Message {
  getAddress(): string;
  setAddress(value: string): void;

  clearEventList(): void;
  getEventList(): Array<string>;
  setEventList(value: Array<string>): void;
  addEvent(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SubscribeInfo.AsObject;
  static toObject(includeInstance: boolean, msg: SubscribeInfo): SubscribeInfo.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SubscribeInfo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SubscribeInfo;
  static deserializeBinaryFromReader(message: SubscribeInfo, reader: jspb.BinaryReader): SubscribeInfo;
}

export namespace SubscribeInfo {
  export type AsObject = {
    address: string,
    eventList: Array<string>,
  }
}

