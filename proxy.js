var express = require('express');
var request = require('request');
var app = express();

var desHost = 'http://localhost:3001';

app.use('/', function (req, res) {
	var url = req.url;
	var newURL = desHost + url;
	console.log(url);
	res.setHeader('Access-Control-Allow-Origin', '*');
	res.setHeader('Access-Control-Allow-Methods', 'GET, POST, OPTIONS, PUT, PATCH, DELETE');
	res.setHeader('Access-Control-Allow-Headers', 'X-Requested-With,content-type');
	res.setHeader('Access-Control-Allow-Credentials', true);
	res.setHeader('Content-Type', 'application/json;charset=UTF-8');
	// skip OPTIONS
	if (req.method === 'OPTIONS') {
		console.log('\nOPTIONS method.');
		res.send('');
	} else {
		var resProxy = res;
		var options = {
			url: newURL,
			// proxy: 'http://cn-proxy.cn.oracle.com:80'
		};

		req.pipe(request(options)).pipe(resProxy);
	}

});
app.listen(process.env.PORT || 8888);
