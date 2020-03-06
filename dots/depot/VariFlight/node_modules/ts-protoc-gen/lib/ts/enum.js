"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var Printer_1 = require("../Printer");
function printEnum(enumDescriptor, indentLevel) {
    var printer = new Printer_1.Printer(indentLevel);
    var enumInterfaceName = enumDescriptor.getName() + "Map";
    printer.printEmptyLn();
    printer.printLn("export interface " + enumInterfaceName + " {");
    enumDescriptor.getValueList().forEach(function (value) {
        printer.printIndentedLn(value.getName().toUpperCase() + ": " + value.getNumber() + ";");
    });
    printer.printLn("}");
    printer.printEmptyLn();
    printer.printLn("export const " + enumDescriptor.getName() + ": " + enumInterfaceName + ";");
    return printer.getOutput();
}
exports.printEnum = printEnum;
//# sourceMappingURL=enum.js.map