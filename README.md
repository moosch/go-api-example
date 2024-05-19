# Go API Example

This is a very small example of a simple API written in Go for a tutorial/walkthrough.

It demonstrates routing http requests, using auth middleware, defining types for requests and responses, some super basic caching for the auth middleware and a few other things.

The structure is very simple and should be easy to follow and build upon.

## Todo

- [ ] If we update a username, we need to then also update the Auth Token DB record, and clear the cache entry.

