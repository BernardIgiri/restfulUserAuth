/*jshint esversion: 6 */
define(['jquery','validate', 'zxcvbn'], function ($, validate, zxcvbn) {
	const getPasswordScore = (value)=> zxcvbn(value).score;
	$.validator.addMethod('password-complexity', function (value, el, minimumPasswordScore) {
		return getPasswordScore(value) >= parseInt(minimumPasswordScore);
	}, 'Please enter a password with sufficient complexity. Try using numbers, symbols, and mixing lower and uppercase letters.');
	return {
		enableLiveChecking: function (allFieldElements, validationHandler) {
			const checkForm = function () {
				let isValid = allFieldElements.toArray().
					every(function(el) {
						const $el = $(el);
						if ($el.attr('type')==='checkbox') {
							console.log('checkbox', $el.attr('id'), $el.attr('required'), $el.prop('checked'), typeof $el.attr('required') === 'undefined' || $el.prop('checked'));
							return typeof $el.attr('required') === 'undefined' || $el.prop('checked');
						} else {
							return $el.attr('aria-invalid') === "false";
						}
					});
				validationHandler(isValid);
			};
			allFieldElements.on('input change', checkForm);
			checkForm();
		},
		val: validate
	};
});
