requirejs.config({
	baseUrl: '../',
	paths: {
		application: 'dev/js/application',
		mustache: "node_modules/mustache/mustache",
		jquery: "node_modules/jquery/dist/jquery",
		zxcvbn: "node_modules/zxcvbn/dist/zxcvbn",
		progressbar: "node_modules/progressbar.js/dist/progressbar",
	}
});

requirejs(['application/main']);

//                                                 ajaxApplication/node_modules/jquery/dist/jquery.js
// file:///home/bigiri/Documents/Projects/Personal/restfulUserAuth/node_modules/jquery/dist/jquery.js.js
// file:///home/bigiri/Documents/Projects/Personal/restfulUserAuth/ajaxApplication/node_modules/jquery/dist/jquery.js
