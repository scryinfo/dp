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
  requestStream: true,
  responseStream: true,
  requestType: variflight_pb.GetFlightDataByFlightNumberRequest,
  responseType: variflight_pb.VariFlightData
};

VariFlightDataService.GetFlightDataBetweenTwoAirports = {
  methodName: "GetFlightDataBetweenTwoAirports",
  service: VariFlightDataService,
  requestStream: true,
  responseStream: true,
  requestType: variflight_pb.GetFlightDataBetweenTwoAirportsRequest,
  responseType: variflight_pb.VariFlightData
};

VariFlightDataService.GetFlightDataBetweenTwoCities = {
  methodName: "GetFlightDataBetweenTwoCities",
  service: VariFlightDataService,
  requestStream: true,
  responseStream: true,
  requestType: variflight_pb.GetFlightDataBetweenTwoCitiesRequest,
  responseType: variflight_pb.VariFlightData
};

VariFlightDataService.GetFlightDataByDepartureAndArrivalStatus = {
  methodName: "GetFlightDataByDepartureAndArrivalStatus",
  service: VariFlightDataService,
  requestStream: true,
  responseStream: true,
  requestType: variflight_pb.GetFlightDataAtOneAirportByStatusRequest,
  responseType: variflight_pb.VariFlightData
};

exports.VariFlightDataService = VariFlightDataService;

function VariFlightDataServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

VariFlightDataServiceClient.prototype.getFlightDataByFlightNumber = function getFlightDataByFlightNumber(metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.client(VariFlightDataService.GetFlightDataByFlightNumber, {
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport
  });
  client.onEnd(function (status, statusMessage, trailers) {
    listeners.status.forEach(function (handler) {
      handler({ code: status, details: statusMessage, metadata: trailers });
    });
    listeners.end.forEach(function (handler) {
      handler({ code: status, details: statusMessage, metadata: trailers });
    });
    listeners = null;
  });
  client.onMessage(function (message) {
    listeners.data.forEach(function (handler) {
      handler(message);
    })
  });
  client.start(metadata);
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    write: function (requestMessage) {
      client.send(requestMessage);
      return this;
    },
    end: function () {
      client.finishSend();
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

VariFlightDataServiceClient.prototype.getFlightDataBetweenTwoAirports = function getFlightDataBetweenTwoAirports(metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.client(VariFlightDataService.GetFlightDataBetweenTwoAirports, {
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport
  });
  client.onEnd(function (status, statusMessage, trailers) {
    listeners.status.forEach(function (handler) {
      handler({ code: status, details: statusMessage, metadata: trailers });
    });
    listeners.end.forEach(function (handler) {
      handler({ code: status, details: statusMessage, metadata: trailers });
    });
    listeners = null;
  });
  client.onMessage(function (message) {
    listeners.data.forEach(function (handler) {
      handler(message);
    })
  });
  client.start(metadata);
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    write: function (requestMessage) {
      client.send(requestMessage);
      return this;
    },
    end: function () {
      client.finishSend();
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

VariFlightDataServiceClient.prototype.getFlightDataBetweenTwoCities = function getFlightDataBetweenTwoCities(metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.client(VariFlightDataService.GetFlightDataBetweenTwoCities, {
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport
  });
  client.onEnd(function (status, statusMessage, trailers) {
    listeners.status.forEach(function (handler) {
      handler({ code: status, details: statusMessage, metadata: trailers });
    });
    listeners.end.forEach(function (handler) {
      handler({ code: status, details: statusMessage, metadata: trailers });
    });
    listeners = null;
  });
  client.onMessage(function (message) {
    listeners.data.forEach(function (handler) {
      handler(message);
    })
  });
  client.start(metadata);
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    write: function (requestMessage) {
      client.send(requestMessage);
      return this;
    },
    end: function () {
      client.finishSend();
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

VariFlightDataServiceClient.prototype.getFlightDataByDepartureAndArrivalStatus = function getFlightDataByDepartureAndArrivalStatus(metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.client(VariFlightDataService.GetFlightDataByDepartureAndArrivalStatus, {
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport
  });
  client.onEnd(function (status, statusMessage, trailers) {
    listeners.status.forEach(function (handler) {
      handler({ code: status, details: statusMessage, metadata: trailers });
    });
    listeners.end.forEach(function (handler) {
      handler({ code: status, details: statusMessage, metadata: trailers });
    });
    listeners = null;
  });
  client.onMessage(function (message) {
    listeners.data.forEach(function (handler) {
      handler(message);
    })
  });
  client.start(metadata);
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    write: function (requestMessage) {
      client.send(requestMessage);
      return this;
    },
    end: function () {
      client.finishSend();
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

exports.VariFlightDataServiceClient = VariFlightDataServiceClient;

