"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const http_1 = __importDefault(require("http"));
// Create a new HTTP server and wrap the Express app
const server = http_1.default.createServer();
// Set up the server to listen on port 3000
server.listen(4123, () => {
    console.log('listening on *:4123');
});
