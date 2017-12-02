({
	baseUrl: '.',
	out: 'build/temp/require.es6.js',
	mainConfigFile: 'dev/js/application/config.js',
	wrap: true,
	include: ['node_modules/almond/almond.js'],
	name: "application/config",
	optimize: "none",
})
