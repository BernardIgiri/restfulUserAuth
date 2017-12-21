requirejs.config({
	baseUrl: '../',
	paths: {
		application: 'dev/js/application',
		mustache: "node_modules/mustache/mustache",
		jquery: "node_modules/jquery/dist/jquery",
		zxcvbn: "node_modules/zxcvbn/dist/zxcvbn",
		progressbar: "node_modules/progressbar.js/dist/progressbar",
		validate: "node_modules/jquery-validation/dist/jquery.validate",
	}
});

requirejs(['application/main']);
