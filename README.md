# Hide n Seek

## What is the aim of the project?

This enables products to roll out a feature to different environments strategically. Some of the features are —

1. Target a feature to a specific user id.
2. Target a feature to specific user ids.
3. Target to x% of users.
4. Setup different environment strategies like staging, production, and QA.
5. Ability to turn on and off the feature flag without having to restart the services.
6. Console to view and control all the feature flags of the different environments in a single place.

## How to work on this project?

1. All the Backend services use Golang
    * https://golang.org/doc/install
2. We are currently using Postgres as our data store  
    * https://www.postgresql.org/download/
    * you can also use the corresponding package manager based on your OS.
3. We also use Redis for caching, and the feature flag for faster access since the value is frequently accessed
4. Create `hide_seek.env` and drop the following values to the environmental variables required for the project. Here is a link to the list.
5. `~/go/bin/gin main.go` to run the HTTP server
6. To build the executable file use `go build` and run `./hide_seek`. 

## Client SDK

Client SDK to access the feature flag in the code without having to hassle around building/updating API on your own. The target is to launch the SDK for ruby as the first mile stone and support other languages soon after.