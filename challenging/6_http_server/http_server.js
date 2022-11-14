var http = require('http');
var url = require('url');

http.createServer(function (req, res) {
	var q = url.parse(req.url, true);
	var qdata = q.query;
	var txt = qdata.year + " " + qdata.month;
	res.writeHead(200, {'Content-Type': 'text/html'});
	res.end(txt);
}).listen(8080);
