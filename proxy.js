var express = require('express');
var request = require('request');
var app = express();

app.use('/', function (req, res) {
	var url = req.query.url;
	console.log('real url=' + req.query.url);
	
	res.setHeader('Access-Control-Allow-Origin', '*');
	res.setHeader('Access-Control-Allow-Methods', 'GET, POST, OPTIONS, PUT, PATCH, DELETE');
	res.setHeader('Access-Control-Allow-Headers', 'X-Requested-With,content-type');
	res.setHeader('Access-Control-Allow-Credentials', true);
	res.setHeader('Content-Type', 'application/json;charset=UTF-8');

	if (req.method === 'OPTIONS') {
		console.log('\nOPTIONS method.');
		res.send('');
	} else {
		var resProxy = res;
		var options = {
			url: url
		};
		req.pipe(request(options)).pipe(resProxy);
	}

});
console.log('running at 8080...');
app.listen(process.env.PORT || 8080);


// npm install express
// npm install request
// node proxy.js
// http://localhost:8080?url=realUrl