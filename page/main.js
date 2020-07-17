var iv
var key

function aesEnc(enc_in){
	var enc_out = enc_in
	//TODO: Encrypto
	//TODO: Base64 encode
	return enc_out
}

function aesDec(dec_in){
	var dec_out = dec_in
	//TODO: Base64 decode
	//TODO: Decrypted
	return dec_out
}

var request = new XMLHttpRequest();

document.addEventListener('DOMContentLoaded',function(e){
	document.getElementById('button').addEventListener('click',function(e){
		var plain = document.form.command.value;
		request.open('POST', 'http://'+location.hostname+':8080/enc/echo');
		request.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');

		var cipher = aesEnc(plain)
		request.send(cipher);
		document.form.command.value = "";

		var send_cipher = document.getElementById('send_msg_plain');
		send_cipher.innerHTML = "send_msg_cipher: " + plain;

		var send_cipher = document.getElementById('send_msg_cipher');
		send_cipher.innerHTML = "send_msg_cipher: " + cipher;
	});
});
	
request.onreadystatechange = function () {
	if (request.readyState != 4) {
		//
	} else if (request.status != 200) {
		//
	} else {
		var cipher = request.responseText;
		var recv_cipher = document.getElementById('recv_msg_cipher');
		recv_cipher.innerHTML = "recv_msg_cipher: " + cipher;

		var plain = aesDec(cipher);
		var recv_plain = document.getElementById('recv_msg_plain');
		recv_plain.innerHTML = "recv_msg_plain: " + plain;
	}
};

