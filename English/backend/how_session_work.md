#### How Session Work?
>   When using an application，the computer knows who you are and how much time you spend using that application, However,
in HTTP web servers, things are a bit difference. An HTTP server does not known you and the time you spend interacting
with the server. This is because HTTP servers are stateless and don't store user information. they only get request from
the user and back a response. This means each request your browser make is unaware of the actions of the previous request.
Session are used to make the HTTP server stateful. Session variables are used to store user information to be used across
the Web application. It helps to keep track of the state between the server and the client/browser. This way the browser
keeps session variables with the information of previous request made to the server. This is made possibly by allowing
the server to crate the session ID. A cookie value is saved on the client's browser when the session is created. This way,
the request will be send along with the value of this cookie.The server will then check if the cookie value matches the
session store value.The difference between a cookie and a session is that a cookie is saved on the client/browser. while
the session is saved on the server. Unlike the session,a cookie don't store any user credentials.