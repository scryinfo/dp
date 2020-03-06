"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var plugin_pb_1 = require("google-protobuf/google/protobuf/compiler/plugin_pb");
var WellKnown_1 = require("../WellKnown");
var FieldTypes_1 = require("../ts/FieldTypes");
var util_1 = require("../util");
function createFile(output, filename) {
    var file = new plugin_pb_1.CodeGeneratorResponse.File();
    file.setName(filename);
    file.setContent(output);
    return file;
}
exports.createFile = createFile;
function getCallingTypes(method, exportMap) {
    return {
        requestType: FieldTypes_1.getFieldType(FieldTypes_1.MESSAGE_TYPE, method.getInputType().slice(1), "", exportMap),
        responseType: FieldTypes_1.getFieldType(FieldTypes_1.MESSAGE_TYPE, method.getOutputType().slice(1), "", exportMap),
    };
}
function isUsed(fileDescriptor, pseudoNamespace, exportMap) {
    return fileDescriptor.getServiceList().some(function (service) {
        return service.getMethodList().some(function (method) {
            var callingTypes = getCallingTypes(method, exportMap);
            var namespacePackage = pseudoNamespace + ".";
            return (callingTypes.requestType.indexOf(namespacePackage) === 0 ||
                callingTypes.responseType.indexOf(namespacePackage) === 0);
        });
    });
}
var RPCDescriptor = (function () {
    function RPCDescriptor(grpcService, protoService, exportMap) {
        this.grpcService = grpcService;
        this.protoService = protoService;
        this.exportMap = exportMap;
    }
    Object.defineProperty(RPCDescriptor.prototype, "name", {
        get: function () {
            return this.protoService.getName();
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(RPCDescriptor.prototype, "qualifiedName", {
        get: function () {
            return (this.grpcService.packageName ? this.grpcService.packageName + "." : "") + this.name;
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(RPCDescriptor.prototype, "methods", {
        get: function () {
            var _this = this;
            return this.protoService.getMethodList()
                .map(function (method) {
                var callingTypes = getCallingTypes(method, _this.exportMap);
                var nameAsCamelCase = method.getName()[0].toLowerCase() + method.getName().substr(1);
                return {
                    nameAsPascalCase: method.getName(),
                    nameAsCamelCase: nameAsCamelCase,
                    functionName: util_1.normaliseFieldObjectName(nameAsCamelCase),
                    serviceName: _this.name,
                    requestStream: method.getClientStreaming(),
                    responseStream: method.getServerStreaming(),
                    requestType: callingTypes.requestType,
                    responseType: callingTypes.responseType,
                };
            });
        },
        enumerable: true,
        configurable: true
    });
    return RPCDescriptor;
}());
exports.RPCDescriptor = RPCDescriptor;
var GrpcServiceDescriptor = (function () {
    function GrpcServiceDescriptor(fileDescriptor, exportMap) {
        this.fileDescriptor = fileDescriptor;
        this.exportMap = exportMap;
        this.pathToRoot = util_1.getPathToRoot(fileDescriptor.getName());
    }
    Object.defineProperty(GrpcServiceDescriptor.prototype, "filename", {
        get: function () {
            return this.fileDescriptor.getName();
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(GrpcServiceDescriptor.prototype, "packageName", {
        get: function () {
            return this.fileDescriptor.getPackage();
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(GrpcServiceDescriptor.prototype, "imports", {
        get: function () {
            var _this = this;
            var dependencies = this.fileDescriptor.getDependencyList()
                .filter(function (dependency) { return isUsed(_this.fileDescriptor, util_1.filePathToPseudoNamespace(dependency), _this.exportMap); })
                .map(function (dependency) {
                var namespace = util_1.filePathToPseudoNamespace(dependency);
                if (dependency in WellKnown_1.WellKnownTypesMap) {
                    return {
                        namespace: namespace,
                        path: WellKnown_1.WellKnownTypesMap[dependency],
                    };
                }
                else {
                    return {
                        namespace: namespace,
                        path: "" + _this.pathToRoot + util_1.replaceProtoSuffix(util_1.replaceProtoSuffix(dependency))
                    };
                }
            });
            var hostProto = {
                namespace: util_1.filePathToPseudoNamespace(this.filename),
                path: "" + this.pathToRoot + util_1.replaceProtoSuffix(this.filename),
            };
            return [hostProto].concat(dependencies);
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(GrpcServiceDescriptor.prototype, "services", {
        get: function () {
            var _this = this;
            return this.fileDescriptor.getServiceList()
                .map(function (service) {
                return new RPCDescriptor(_this, service, _this.exportMap);
            });
        },
        enumerable: true,
        configurable: true
    });
    return GrpcServiceDescriptor;
}());
exports.GrpcServiceDescriptor = GrpcServiceDescriptor;
//# sourceMappingURL=common.js.map