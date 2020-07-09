var iv
var key

function aesEnc(enc_in){
	var enc_out = enc_in
	//TODO: Encrypto
	return enc_out
}

function aesDec(dec_in){
	var dec_out = dec_in
	//TODO: Decrypted
	return dec_out
}

var request = new XMLHttpRequest();

document.addEventListener('DOMContentLoaded',function(e){
	document.getElementById('button').addEventListener('click',function(e){
		var msg = document.form.command.value;
		request.open('POST', 'http://'+location.hostname+':8080/enc/echo');
		request.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
		request.send(aesEnc(msg));
		document.form.command.value = "";
	});
});
	
request.onreadystatechange = function () {
	if (request.readyState != 4) {
		//
	} else if (request.status != 200) {
		//
	} else {
		var ct = request.responseText;
		target = document.getElementById('cipher');
		target.innerHTML = "cipher text: " + ct;

		var dt = aesDec(ct);
		c = document.getElementById('dec');
		c.innerHTML = "decrypted text: " + dt;
	}
};

