import http from 'http';
import { Server as SocketIO } from 'socket.io';
import { EventEmitter } from 'node:events';
import express from 'express';
import path from 'path';

// Initialize an Event Emitter.
const ee = new EventEmitter();

// Listen for the startingServer event.
ee.on('startingServer', port => {
    console.log('Roger. Youre starting the server on port:', port);
});

// Initialize the Express app.
const app = express();
const port = 4123;

// Serve static files from the 'public' directory.
app.use(express.static(path.join(__dirname, 'public')));

// Serve index.html for the root route.
app.get('/', (req, res) => {
    res.sendFile(path.join(__dirname, 'public', 'index.html'));
});

// Create a new HTTP server and wrap the Express app.
const server = http.createServer(app);

// Initialize Socket.IO and attach it to the HTTP server.
const io = new SocketIO(server);

// Handling Socket.IO connections.
io.on('connection', (socket) => {
    console.log('A user connected');

    // You can setup your socket.io events here
    // For example:
    // socket.on('chat message', (msg) => {
    //     io.emit('chat message', msg);
    // });

    socket.on('disconnect', () => {
        console.log('User disconnected');
    });
});

// Set up the server to listen on port 4123, and emit an event on startup.
server.listen(port, () => {
    ee.emit('startingServer', port);
    console.log('Listening on *:', port);
});
