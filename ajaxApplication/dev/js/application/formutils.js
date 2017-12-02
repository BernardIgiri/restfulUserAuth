/*jshint esversion: 6 */
define(['jquery','mustache'], function ($) {
	return {
		serialize: function(keysArray, elsMap) {
			return keysArray.map(function (k) {
				let type = elsMap[k].attr('type');
				if (type === "checkbox" || type === "radio") {
					return elsMap[k].prop('checked');
				} else {
					return elsMap[k].val();
				}
			});
		},
	};
});
