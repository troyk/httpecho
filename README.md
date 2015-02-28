HttpEcho
========

Simple go http server that echos any query or form values as a json response.
The `delay` query value is special in that it will cause the server to delay 
the response x milliseconds.

We use this to find out which ruby http and background processing libs can 
haul butt with twillio (to send 1000's of txt messages per minute)
	
