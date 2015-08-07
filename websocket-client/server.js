var fs = require('fs');

function handler(req, res) {
    'use strict';
    fs.readFile(
        __dirname + '/index.html',
        function (err, data) {
            if (err) {
                res.writeHead(500);
                return res.end('Error loading index.html');
            }

            res.writeHead(200);
            res.end(data);
        }
    );
}

var app = require('http').createServer(handler),
    io = require('socket.io').listen(app);

app.listen(8012);

io.sockets.on('connection', function (socket) {
    'use strict';


    setInterval(function () {
        socket.emit('alerts', '');
    }, 1000);

    setInterval(function () {
        socket.emit('news', {some: 'long', alert: 'clear'});
    }, 1000);

    socket.emit('a', 'sdjkjdskjdkj');
    socket.on('my other event', function (data) {
        console.log(data);
    });
});
