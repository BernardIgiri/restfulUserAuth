/*jshint esversion: 6 */
define(['jquery','mustache'], function ($, nustache) {
	return {
		renderHtml: function(selector, view) {
			return nustache.render($(selector).text(), view);
		},
	};
});
