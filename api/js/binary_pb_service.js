// package: api
// file: binary.proto

var binary_pb = require("./binary_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var BinaryService = (function () {
  function BinaryService() {}
  BinaryService.serviceName = "api.BinaryService";
  return BinaryService;
}());

BinaryService.SubscribeEvent = {
  methodName: "SubscribeEvent",
  service: BinaryService,
  requestStream: false,
  responseStream: false,
  requestType: binary_pb.SubscribeInfo,
  responseType: binary_pb.Result
};

BinaryService.UnSubscribeEvent = {
  methodName: "UnSubscribeEvent",
  service: BinaryService,
  requestStream: false,
  responseStream: false,
  requestType: binary_pb.SubscribeInfo,
  responseType: binary_pb.Result
};

BinaryService.RecvEvents = {
  methodName: "RecvEvents",
  service: BinaryService,
  requestStream: false,
  responseStream: true,
  requestType: binary_pb.ClientInfo,
  responseType: binary_pb.Event
};

BinaryService.Publish = {
  methodName: "Publish",
  service: BinaryService,
  requestStream: false,
  responseStream: false,
  requestType: binary_pb.PublishParams,
  responseType: binary_pb.PublishResult
};

BinaryService.PrepareToBuy = {
  methodName: "PrepareToBuy",
  service: BinaryService,
  requestStream: false,
  responseStream: false,
  requestType: binary_pb.PrepareParams,
  responseType: binary_pb.Result
};

BinaryService.BuyData = {
  methodName: "BuyData",
  service: BinaryService,
  requestStream: false,
  responseStream: false,
  requestType: binary_pb.BuyParams,
  responseType: binary_pb.Result
};

BinaryService.CancelTransaction = {
  methodName: "CancelTransaction",
  service: BinaryService,
  requestStream: false,
  responseStream: false,
  requestType: binary_pb.CancelTxParams,
  responseType: binary_pb.Result
};

BinaryService.SubmitMetaDataIdEncWithBuyer = {
  methodName: "SubmitMetaDataIdEncWithBuyer",
  service: BinaryService,
  requestStream: false,
  responseStream: false,
  requestType: binary_pb.SubmitMetaDataIdParams,
  responseType: binary_pb.Result
};

BinaryService.ConfirmDataTruth = {
  methodName: "ConfirmDataTruth",
  service: BinaryService,
  requestStream: false,
  responseStream: false,
  requestType: binary_pb.DataConfirmParams,
  responseType: binary_pb.Result
};

BinaryService.ApproveTransfer = {
  methodName: "ApproveTransfer",
  service: BinaryService,
  requestStream: false,
  responseStream: false,
  requestType: binary_pb.ApproveTransferParams,
  responseType: binary_pb.Result
};

BinaryService.Vote = {
  methodName: "Vote",
  service: BinaryService,
  requestStream: false,
  responseStream: false,
  requestType: binary_pb.VoteParams,
  responseType: binary_pb.Result
};

BinaryService.RegisterAsVerifier = {
  methodName: "RegisterAsVerifier",
  service: BinaryService,
  requestStream: false,
  responseStream: false,
  requestType: binary_pb.RegisterVerifierParams,
  responseType: binary_pb.Result
};

BinaryService.CreditsToVerifier = {
  methodName: "CreditsToVerifier",
  service: BinaryService,
  requestStream: false,
  responseStream: false,
  requestType: binary_pb.CreditVerifierParams,
  responseType: binary_pb.Result
};

BinaryService.TransferTokens = {
  methodName: "TransferTokens",
  service: BinaryService,
  requestStream: false,
  responseStream: false,
  requestType: binary_pb.TransferTokenParams,
  responseType: binary_pb.Result
};

BinaryService.GetTokenBalance = {
  methodName: "GetTokenBalance",
  service: BinaryService,
  requestStream: false,
  responseStream: false,
  requestType: binary_pb.TokenBalanceParams,
  responseType: binary_pb.TokenBalanceResult
};

BinaryService.CreateAccount = {
  methodName: "CreateAccount",
  service: BinaryService,
  requestStream: false,
  responseStream: false,
  requestType: binary_pb.CreateAccountParams,
  responseType: binary_pb.AccountResult
};

BinaryService.Authenticate = {
  methodName: "Authenticate",
  service: BinaryService,
  requestStream: false,
  responseStream: false,
  requestType: binary_pb.ClientInfo,
  responseType: binary_pb.Result
};

BinaryService.TransferEth = {
  methodName: "TransferEth",
  service: BinaryService,
  requestStream: false,
  responseStream: false,
  requestType: binary_pb.TransferEthParams,
  responseType: binary_pb.Result
};

BinaryService.GetEthBalance = {
  methodName: "GetEthBalance",
  service: BinaryService,
  requestStream: false,
  responseStream: false,
  requestType: binary_pb.EthBalanceParams,
  responseType: binary_pb.EthBalanceResult
};

exports.BinaryService = BinaryService;

function BinaryServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

BinaryServiceClient.prototype.subscribeEvent = function subscribeEvent(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(BinaryService.SubscribeEvent, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

BinaryServiceClient.prototype.unSubscribeEvent = function unSubscribeEvent(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(BinaryService.UnSubscribeEvent, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

BinaryServiceClient.prototype.recvEvents = function recvEvents(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(BinaryService.RecvEvents, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners.end.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

BinaryServiceClient.prototype.publish = function publish(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(BinaryService.Publish, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

BinaryServiceClient.prototype.prepareToBuy = function prepareToBuy(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(BinaryService.PrepareToBuy, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

BinaryServiceClient.prototype.buyData = function buyData(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(BinaryService.BuyData, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

BinaryServiceClient.prototype.cancelTransaction = function cancelTransaction(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(BinaryService.CancelTransaction, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

BinaryServiceClient.prototype.submitMetaDataIdEncWithBuyer = function submitMetaDataIdEncWithBuyer(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(BinaryService.SubmitMetaDataIdEncWithBuyer, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

BinaryServiceClient.prototype.confirmDataTruth = function confirmDataTruth(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(BinaryService.ConfirmDataTruth, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

BinaryServiceClient.prototype.approveTransfer = function approveTransfer(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(BinaryService.ApproveTransfer, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

BinaryServiceClient.prototype.vote = function vote(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(BinaryService.Vote, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

BinaryServiceClient.prototype.registerAsVerifier = function registerAsVerifier(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(BinaryService.RegisterAsVerifier, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

BinaryServiceClient.prototype.creditsToVerifier = function creditsToVerifier(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(BinaryService.CreditsToVerifier, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

BinaryServiceClient.prototype.transferTokens = function transferTokens(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(BinaryService.TransferTokens, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

BinaryServiceClient.prototype.getTokenBalance = function getTokenBalance(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(BinaryService.GetTokenBalance, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

BinaryServiceClient.prototype.createAccount = function createAccount(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(BinaryService.CreateAccount, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

BinaryServiceClient.prototype.authenticate = function authenticate(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(BinaryService.Authenticate, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

BinaryServiceClient.prototype.transferEth = function transferEth(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(BinaryService.TransferEth, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

BinaryServiceClient.prototype.getEthBalance = function getEthBalance(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(BinaryService.GetEthBalance, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.BinaryServiceClient = BinaryServiceClient;

