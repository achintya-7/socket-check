const io = require('socket.io-client');
const socket = io('http://localhost/socket:8000')

console.log('Starting');

socket.on('connect', () => {
    console.log("Connected");
})

socket.on('message'), () => {
    console.log("Sent");
}

socket.emit('message', 'Hello World')