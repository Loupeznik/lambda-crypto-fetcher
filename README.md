# Lambda Cryptocurrency Fetcher
A simple script to fetch cryptocurrency rates from [messari.io](https://messari.io) API and store these into a MongoDB database
meant to be used on AWS Lambda.

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
[![License](https://img.shields.io/github/license/Loupeznik/ServerStatusChecker?style=for-the-badge)](./LICENSE)

## Build (for publishing to Lambda)
```bash
git clone https://github.com/Loupeznik/lambda-messario.git
cd lambda-messario
go get .
GOOS=linux GOARCH=amd64 go build -o main .
zip main.zip main
```

## License
This project is MIT licensed.
