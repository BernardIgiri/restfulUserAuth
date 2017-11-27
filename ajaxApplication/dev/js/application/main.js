/*jshint esversion: 6 */
define(function (require) {
	require('jquery');
	let signup = require('application/signup');
	let $ = jQuery;
	$("body > h1").text("Hello, world!");
	signup();
});
