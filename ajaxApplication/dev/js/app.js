requirejs.config({
	baseUrl: 'js/vendor',
	paths: {
		application: '../application'
	}
});

requirejs(['application/main']);
