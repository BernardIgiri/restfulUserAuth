/*jshint esversion: 6 */
define(['jquery', 'progressbar', 'zxcvbn', 'application/validation', 'application/formutils'],
function ($, progressBar, zxcvbn, validation, formutils) {
	return function () {
		const weakColor = [252, 91, 63];  // Red
		const strongColor = [111, 213, 127];  // Green
		const defaultColor = [204, 204, 204];
		const passwordGrades = {
			0: 'Very weak',
			1: 'Weak',
			2: 'Average',
			3: 'Strong',
			4: 'Very strong'
		};
		let els = {};
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
		const getElements = function() {
			els = {
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
			return formutils.serialize([
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
			], els);
		};
		const showPasswordStrength = function(e) {
			els.strengthBarContainer.css("visibility", "visible");
		};
		const hidePasswordStrength = function(e) {
			els.strengthBarContainer.css("visibility", "hidden");
		};
		const getPasswordScore = ()=> zxcvbn(els.password.val()).score;
		const updateStrengthMeter = function() {
			const score = getPasswordScore();
			els.strengthLabel.text(passwordGrades[score]);
			let progress = score / 4;
			if (progress === 0 && els.password.val() > 0) {
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
					els.password.css({ "color": state.color });
					bar.path.setAttribute('stroke', state.color);
				}
			});
		};
		const submit = function(e) {
			e.stopPropagation();
			e.preventDefault();
			if (els.form.valid()) {
				let values = getValues();
				console.log(values);
				var request = {
					login: values[3],
					password: values[5],
					firstname: values[0],
					lastname: values[1],
					email: values[4],
					phonenumber: values[2],
					enable2fa: values[7],
					sendNewsLetter: values[8],
				};
				$.ajax({
					url: "user/register",
					method: "POST",
					data: request,
					dataType: 'json',
					contentType: "application/json",
					success: function(result, status, jqXHR) {
						console.log('success', result);
					},
					error: function(jqXHR, textStatus, errorThrown) {
						console.log('error', arguments);
					}
				});
			}
		};
		getElements();
		els.root.show();
		getElements();
		const strengthBar = new progressBar.Line(els.strengthBar[0], {
			color: '#ddd',
			trailColor: '#f7f7f7',
			duration: 1000,
			easing: 'easeOut',
			strokeWidth: 8
		});
		hidePasswordStrength();
		els.password.on('focus', showPasswordStrength);
		els.password.on('blur', hidePasswordStrength);
		els.password.on('input', updateStrengthMeter);
		els.form.on('submit', submit);
		els.form.validate();
		validation.enableLiveChecking(els.allFields, (isValid)=> els.submit.prop('disabled', !isValid));
	};
});
