var iv
var key

var request = new XMLHttpRequest();
var recv_cipher = document.getElementById('recv_cipher');
var send_cipher = document.getElementById('send_cipher');
recv_cipher.innerHTML = "send_cipher: "
send_cipher.innerHTML = "recv_cipher: "

document.addEventListener('DOMContentLoaded',function(e){
	document.getElementById('button').addEventListener('click',function(e){
		var cipher = document.form.command.value;
		
		request.open('POST', 'http://'+location.hostname+':8080/enc/echo');
		request.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');

		request.send(cipher);
		document.form.command.value = "";

		send_cipher.innerHTML = "send_cipher: " + cipher;
	});
});
	
request.onreadystatechange = function () {
	if (request.readyState != 4) {
		//
	} else if (request.status != 200) {
		//
	} else {
		var cipher = request.responseText;
		recv_cipher.innerHTML = "recv_cipher: " + cipher;
	}
};

