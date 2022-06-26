# Ledger
## Server
### Installation and Setup
1. Dependencies
    - `go 1.17` (for local development without Docker)
    - `Docker`
    - `docker-compose`
2. Setup
    - Add `dev.env` and `prod.env` files to `/server` directory
    - Run `make dev` in the root dir to spin up a development server
    - Run `make prod` in the root dir to spin up a development server
    - `8080` should be only exposed port on container

## Client
1. Dependencies
    - `yarn`
    - `node.js >=v17.6`
2. Setup
    - Run `yarn install` to install packages
    - Run `yarn start` to start client on port `3000`