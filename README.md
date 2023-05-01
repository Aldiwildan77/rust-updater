# Rust Updater

This repository is for update notifications from the rust game which helps server admins to get official game updates

## Game Plugin
To make this API works, please install this [Rust Plugin](https://github.com/Psystec/RUST-Update-Notice) into your Rust server.

## How to build
### Native build:
1. Please check your current OS and Architecture, run `go env`.
2. Build this project using command `env GOOS=<YOUR_OS> GOARCH=<YOUR_ARCH> go build -v -o bin/api`.
3. Give output permission to execute.
4. Run the server, depending on your operating system.

### Docker build:
1. Make sure you have installed docker on your machine.
2. Build the project using command `docker build . `
3. Start the project using command `docker run -h localhost -p 8000:8000 -e PORT=8000 rust-updater-api_app:latest`

### Docker Compose:
1. Make sure you have installed docker and docker-compose on your machine.
2. Start the container (detached) using command `docker-compose up -d --build`

## Contact
If you need more information please join the discord: https://discord.chroma-gaming.xyz/

## Maintenance
- Developer: Aldiwildan77
