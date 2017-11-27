/*jshint esversion: 6 */
define(function (require) {
	require('mustache');
	return {
		renderHtml: function(selector, view) {
			return Mustache.render($(selector).text(), view);
		},
	};
});
