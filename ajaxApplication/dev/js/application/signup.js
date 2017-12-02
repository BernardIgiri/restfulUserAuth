/*jshint esversion: 6 */
define(['jquery', 'progressbar', 'zxcvbn', 'validate', 'application/validation'],
function ($, progressBar, zxcvbn, jqvalidate, validation) {
	return function () {
		const weakColor = [252, 91, 63];  // Red
		const strongColor = [111, 213, 127];  // Green
		const defaultColor = [204, 204, 204];
		const minimumPasswordGrade = 3;
		const passwordGrades = {
			0: 'Very weak',
			1: 'Weak',
			2: 'Average',
			3: 'Strong',
			4: 'Very strong'
		};
		let parts = {};
		const interpolateColor = function(rgbA, rgbB, value) {
			let rDiff = rgbA[0] - rgbB[0];
			let gDiff = rgbA[1] - rgbB[1];
			let bDiff = rgbA[2] - rgbB[2];
			value = 1 - value;
			return [
				rgbB[0] + rDiff * value,
				rgbB[1] + gDiff * value,
				rgbB[2] + bDiff * value
			];
		};
 		const rgbArrayToString = function(rgb) {
			return 'rgb(' + rgb[0] + ',' + rgb[1] + ',' + rgb[2] + ')';
		};
		const barColor = function(progress) {
			return interpolateColor(weakColor, strongColor, progress);
		};
		const getParts = function() {
			parts = {
				root: $("div.page.signup"),
				form: $("div.page.signup form"),
				firstname: $("#signup_firstname"),
				lastname: $("#signup_lastname"),
				phonenumber: $("#signup_phonenumber"),
				login: $("#signup_login"),
				email: $("#signup_email"),
				password: $("#signup_password"),
				confirm: $("#signup_confirm"),
				enable2fa: $("#signup_enable2fa"),
				sendnewsletter: $("#signup_sendnewsletter"),
				terms: $("#signup_terms"),
				strengthBarContainer: $("div.page.signup .strengthBarContainer"),
				strengthBar: $("div.page.signup .strengthBar"),
				strengthLabel: $("div.page.signup .strengthLabel"),
				submit: $("div.page.signup [type='submit']"),
				allFields: $("div.page.signup input"),
			};
		};
		const getValues = function() {
			return [
				'firstname',
				'lastname',
				'phonenumber',
				'login',
				'email',
				'password',
				'confirm',
				'enable2fa',
				'sendnewsletter',
				'terms'
				].map((k) => parts[k].val());
		};
		const showPasswordStrength = function(e) {
			parts.strengthBarContainer.css("visibility", "visible");
		};
		const hidePasswordStrength = function(e) {
			parts.strengthBarContainer.css("visibility", "hidden");
		};
		const updateStrengthMeter = function() {
			const result = zxcvbn(parts.password.val());
			parts.strengthLabel.text(passwordGrades[result.score]);
			let progress = result.score / 4;
			if (progress === 0 && parts.password.val() > 0) {
				progress = 0.1;
			}
			let startColor = +strengthBar.value().toFixed(3) === 0 ?
				rgbArrayToString(defaultColor) :
				rgbArrayToString(barColor(strengthBar.value()));
			let endColor = progress === 0 ?
				rgbArrayToString(defaultColor) :
				rgbArrayToString(barColor(progress));
			strengthBar.animate(progress, {
				from: { color: startColor },
				to: { color: endColor },
				step: function(state, bar) {
					//input.style.color = state.color;
					bar.path.setAttribute('stroke', state.color);
				}
			});
		};
		/*const validate = function(e) {
			updateStrengthMeter();
			const values = getValues();
			isValid = isPasswordValid &&
				values[1].length > 8 &&
				values[1]===values[2] &&
				values[0].match(validation.email) !== null;
			parts.submit.prop('disabled', !isValid);
			console.log("isValid", isValid);
		};*/
		const submit = function(e) {
			e.stopPropagation();
			e.preventDefault();
			//if (isValid) {
				let values = getValues();
				console.log(values);
			//}
		};
		getParts();
		const strengthBar = new progressBar.Line(parts.strengthBar[0], {
			color: '#ddd',
			trailColor: '#f7f7f7',
			duration: 1000,
			easing: 'easeOut',
			strokeWidth: 8
		});
		hidePasswordStrength();
		parts.password.on('focus', showPasswordStrength);
		parts.password.on('blur', hidePasswordStrength);
		parts.password.on('input', updateStrengthMeter);
		parts.form.on('submit', submit);
		parts.form.validate();
		parts.root.show();
	};
});
