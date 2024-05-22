# Go API Example

This is a very small example of a simple API written in Go for a tutorial/walkthrough.

It demonstrates routing http requests, using auth middleware, defining types for requests and responses, some super basic caching for the auth middleware and a few other things.

The structure is very simple and should be easy to follow and build upon.

## Endpoints

```
POST /account/login?username=<username>
POST /account/logout


GET    /account?username=<username>
    Authorization: <authToken>

POST   /account?username=<username>
    Authorization: <authToken>

DELETE /account?username=<username>
    Authorization: <authToken>
```

