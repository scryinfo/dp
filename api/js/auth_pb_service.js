// package: api
// file: auth.proto

var auth_pb = require("./auth_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var KeyService = (function () {
  function KeyService() {}
  KeyService.serviceName = "api.KeyService";
  return KeyService;
}());

KeyService.GenerateAddress = {
  methodName: "GenerateAddress",
  service: KeyService,
  requestStream: false,
  responseStream: false,
  requestType: auth_pb.AddressParameter,
  responseType: auth_pb.AddressInfo
};

KeyService.VerifyAddress = {
  methodName: "VerifyAddress",
  service: KeyService,
  requestStream: false,
  responseStream: false,
  requestType: auth_pb.AddressParameter,
  responseType: auth_pb.AddressInfo
};

KeyService.ContentEncrypt = {
  methodName: "ContentEncrypt",
  service: KeyService,
  requestStream: false,
  responseStream: false,
  requestType: auth_pb.CipherParameter,
  responseType: auth_pb.CipherText
};

KeyService.ContentDecrypt = {
  methodName: "ContentDecrypt",
  service: KeyService,
  requestStream: false,
  responseStream: false,
  requestType: auth_pb.CipherParameter,
  responseType: auth_pb.CipherText
};

KeyService.Signature = {
  methodName: "Signature",
  service: KeyService,
  requestStream: false,
  responseStream: false,
  requestType: auth_pb.CipherParameter,
  responseType: auth_pb.CipherText
};

KeyService.import_keystore = {
  methodName: "import_keystore",
  service: KeyService,
  requestStream: false,
  responseStream: false,
  requestType: auth_pb.ImportParameter,
  responseType: auth_pb.AddressInfo
};

exports.KeyService = KeyService;

function KeyServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

KeyServiceClient.prototype.generateAddress = function generateAddress(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(KeyService.GenerateAddress, {
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

KeyServiceClient.prototype.verifyAddress = function verifyAddress(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(KeyService.VerifyAddress, {
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

KeyServiceClient.prototype.contentEncrypt = function contentEncrypt(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(KeyService.ContentEncrypt, {
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

KeyServiceClient.prototype.contentDecrypt = function contentDecrypt(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(KeyService.ContentDecrypt, {
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

KeyServiceClient.prototype.signature = function signature(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(KeyService.Signature, {
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

KeyServiceClient.prototype.import_keystore = function import_keystore(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(KeyService.import_keystore, {
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

exports.KeyServiceClient = KeyServiceClient;

