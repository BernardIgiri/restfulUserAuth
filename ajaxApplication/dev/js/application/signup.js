/*jshint esversion: 6 */
define(function (require) {
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
		const validation = require('application/validation');
		const $ = jQuery;
		let parts = {};
		let isPasswordValid = false;
		let isValid = false;
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
				email: $("div.page.signup [name='email']"),
				password: $("div.page.signup [name='password']"),
				strengthBarContainer: $("div.page.signup .strengthBarContainer"),
				strengthBar: $("div.page.signup .strengthBar"),
				strengthLabel: $("div.page.signup .strengthLabel"),
				confirm: $("div.page.signup [name='confirm']"),
				submit: $("div.page.signup [type='submit']"),
				allFields: $("div.page.signup input"),
			};
		};
		const getValues = function() {
			return ['email','password','confirm'].map((k) => parts[k].val());
		};
		const showPasswordStrength = function(e) {
			parts.strengthBarContainer.attr("visibility", "visible");
		};
		const hidePasswordStrength = function(e) {
			parts.strengthBarContainer.attr("visibility", "hidden");
		};
		const updateStrengthMeter = function() {
			const result = zxcvbn(parts.password.val());
			parts.strengthLabel.text(passwordGrades[result.score]);
			let progress = result.score / 4;
			inputLabel.dataset.info = passwordGrades[result.score];
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
			isPasswordValid = result.score >= minimumPasswordGrade;
		};
		const validate = function(e) {
			updateStrengthMeter();
			const values = getValues();
			isValid = isPasswordValid &&
				values[1].length > 8 &&
				values[1]===values[2] &&
				values[0].match(validation.email) !== null;
			parts.submit.prop('disabled', !isValid);
			console.log("isValid", isValid);
		};
		const submit = function(e) {
			e.stopPropagation();
			e.preventDefault();
			if (isValid) {
				let values = getValues();
				console.log(values);
			}
		};
		getParts();
		validate();
		const strengthBar = new ProgressBar.Line(parts.strengthBar, {
			color: '#ddd',
			trailColor: '#f7f7f7',
			duration: 1000,
			easing: 'easeOut',
			strokeWidth: 5
		});
		hidePasswordStrength();
		parts.password.on('focus', showPasswordStrength);
		parts.password.on('focus', showPasswordStrength);
		parts.allFields.on('input', validate);
		parts.form.on('submit', submit);
		parts.root.show();
	};
});
