# Explanation of the OAuth process
1. The sign-in button html code on the frontend contains the client-id
    >The client ID is not a secret but domain validation occurs at google cloud and redirec_uri (see below) is always specified to our own backend domain
2. The sign-in button on the frontend makes the request to gauth server
3. google oauth server sends a PUT request to redirect_uri with 'credential' containing a token
4. Backend processes the said token and extracts required info from it
5. Ideally the backend would sign and send back its own jwt token which would henceforth be included in all frontend requests. But currently the same token is sent back to frontend
6. Once the user is created / verified the backend redirects to frontend/login and sets a cookie named `auth` with value as jwt
7. henceforth an Authorization header is sent with every frontend request with `Bearer {token}`