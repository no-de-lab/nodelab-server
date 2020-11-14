# Nodelab API 
Nodelab API project

## Core dependency
- gorilla mux
  - web framework
- sqlx
  - database entity mapping
- wire 
  - dependency injection
- viper
  - manage configuration 
- model
  - dto & entity mapping

## Structure
https://github.com/bxcodec/go-clean-arch


## Setup
```bash
# install go1.15
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)

# install go 1.7 (binary install)
gvm install go1.7 -B
gvm use go1.7

# install go 1.15
gvm install go1.15

# run main
make run
```

Run with docker-compose 
```
# local database create & run with air
$ make up

# cleanup container & local data
$ make down
```


## Make command
- test
  - run all test
- run
  - run main
- vendor
  - install dependencies
- up
  - run with docker-compose for development
- down
  - remove all container
  - remove with docker named volume (local data)
- wire
  - make wire_gen
- build
  - build for production
- build-air
  - build for air


## App Configuration
```toml
# for db connection
[database]
host = "mysql"
database = "nodelab"
username = "nodelab"
password = "test"

# context timeout
[context]
timeout = 2

# etc 
# ...
```
