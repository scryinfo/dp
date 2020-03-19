// package: protobuf
// file: variflight.proto

var variflight_pb = require("./variflight_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var VariFlightDataService = (function () {
  function VariFlightDataService() {}
  VariFlightDataService.serviceName = "protobuf.VariFlightDataService";
  return VariFlightDataService;
}());

VariFlightDataService.GetFlightDataByFlightNumber = {
  methodName: "GetFlightDataByFlightNumber",
  service: VariFlightDataService,
  requestStream: false,
  responseStream: true,
  requestType: variflight_pb.GetFlightDataByFlightNumberRequest,
  responseType: variflight_pb.VariFlightData
};

VariFlightDataService.GetFlightDataBetweenTwoAirports = {
  methodName: "GetFlightDataBetweenTwoAirports",
  service: VariFlightDataService,
  requestStream: false,
  responseStream: true,
  requestType: variflight_pb.GetFlightDataBetweenTwoAirportsRequest,
  responseType: variflight_pb.VariFlightData
};

VariFlightDataService.GetFlightDataBetweenTwoCities = {
  methodName: "GetFlightDataBetweenTwoCities",
  service: VariFlightDataService,
  requestStream: false,
  responseStream: true,
  requestType: variflight_pb.GetFlightDataBetweenTwoCitiesRequest,
  responseType: variflight_pb.VariFlightData
};

VariFlightDataService.GetFlightDataByDepartureAndArrivalStatus = {
  methodName: "GetFlightDataByDepartureAndArrivalStatus",
  service: VariFlightDataService,
  requestStream: false,
  responseStream: true,
  requestType: variflight_pb.GetFlightDataAtOneAirportByStatusRequest,
  responseType: variflight_pb.VariFlightData
};

exports.VariFlightDataService = VariFlightDataService;

function VariFlightDataServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

VariFlightDataServiceClient.prototype.getFlightDataByFlightNumber = function getFlightDataByFlightNumber(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(VariFlightDataService.GetFlightDataByFlightNumber, {
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

VariFlightDataServiceClient.prototype.getFlightDataBetweenTwoAirports = function getFlightDataBetweenTwoAirports(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(VariFlightDataService.GetFlightDataBetweenTwoAirports, {
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

VariFlightDataServiceClient.prototype.getFlightDataBetweenTwoCities = function getFlightDataBetweenTwoCities(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(VariFlightDataService.GetFlightDataBetweenTwoCities, {
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

VariFlightDataServiceClient.prototype.getFlightDataByDepartureAndArrivalStatus = function getFlightDataByDepartureAndArrivalStatus(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(VariFlightDataService.GetFlightDataByDepartureAndArrivalStatus, {
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

exports.VariFlightDataServiceClient = VariFlightDataServiceClient;

