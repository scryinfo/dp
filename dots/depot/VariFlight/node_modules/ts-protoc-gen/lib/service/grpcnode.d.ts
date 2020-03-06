import { ExportMap } from "../ExportMap";
import { FileDescriptorProto } from "google-protobuf/google/protobuf/descriptor_pb";
import { CodeGeneratorResponse } from "google-protobuf/google/protobuf/compiler/plugin_pb";
export declare function generateGrpcNodeService(filename: string, descriptor: FileDescriptorProto, exportMap: ExportMap): CodeGeneratorResponse.File;
