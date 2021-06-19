# KrakenD API Gateway

Ultra-High performance API Gateway with middlewares

**KrakenD Community Edition** (or *KrakenD-CE*) is the binary distribution of [KrakenD](http://www.krakend.io).

This application will run at port `8080` by default.

For more details, pls read at [KrakenD's Docs](https://www.krakend.io/docs/overview/introduction/)

## ------ Getting started ------

## Clone repository
```bash
git clone https://github.com/phucps89/krakend-playground.git --recursive krakend-api-gateway
```

----------

### Build KrakenD
* Install `Go` environment. More details at [https://golang.org/doc/install ](https://golang.org/doc/install)
* Navigate to `krakend-api-gateway` folder
```bash
./build-krakend.sh
```

### Setup env for configuration files
Navigate to `krakend-api-gateway/src` folder
1. Create `.env` file based on `.env.example` 
2. Change to the right values in `.env` file

## Compile configuration files
Run one of the following commands depended on your OS
```bash
./bin/env-linux
./bin/env-linux-arm
./bin/env-mac
./bin/env-windows
``` 

## Run on real machine

Pls add the parameter `run` after above commands if you want to compile & run.
For example:
```bash
./bin/env-linux run
```
Or you can navigate to `krakend-api-gateway/dist` and run
```bash
./run.sh
```

----------

## Using docker compose
Navigate to `krakend-api-gateway` folder
### Setup env
1. Create `.env` file based on `.env.example` 
2. Change to the right values in `.env` file

## Run command
```bash
docker-compose up krakend_ce
```