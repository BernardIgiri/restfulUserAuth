/*jshint esversion: 6 */
define([], function ($, mustache) {
	return {
		renderHtml: function(selector, view) {
			return mustache.render($(selector).text(), view);
		},
	};
});
