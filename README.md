# Nodelab API
![Actions Status](https://github.com/no-de-lab/nodelab-server/workflows/CICD/badge.svg?branch=dev)
![golangci-lint](https://github.com/no-de-lab/nodelab-server/workflows/golangci-lint/badge.svg?branch=dev)
[![codecov](https://codecov.io/gh/no-de-lab/nodelab-server/branch/dev/graph/badge.svg?token=BAEElqEtoc)](https://codecov.io/gh/no-de-lab/nodelab-server)

## Tech stack
- Golang
- MySQL 
- AWS ECS
- GraphQL
## Core dependency
- [echo](https://github.com/labstack/echo) 
  - Web framework
- [sqlx](https://github.com/jmoiron/sqlx)
  - Database entity mapping
- [wire](https://github.com/google/wire)
  - Dependency injection
- [viper](https://github.com/spf13/viper)
  - Manage configuration
- model
  - DTO & entity mapping

## Structure
- Inspired from https://github.com/bxcodec/go-clean-arch

## pre-commit
Install pre-commit (https://pre-commit.com/)
```
brew install pre-commit
```

Run install command (project root)
```
pre-commit install
```

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

## Appendix
## Design and Planning Docs
- [Figma](https://www.figma.com/file/wSDzlnpDbM5B3yigiVYbgX/1127_nodelab_wireframe_%EC%8A%A4%ED%81%AC%EB%9F%BC?node-id=275%3A887)
- [Google Docs](https://docs.google.com/spreadsheets/d/1tkgqKZP7wX2VGBmsYFr--c4LLsIzV3mR5X0P80Gb9nc/edit#gid=0)
### Workflow
- Choose an issue to work on from [server project board](https://github.com/no-de-lab/nodelab-server/projects/1)
- Create a branch `#issue/short_description_of_issue` from `dev` branch
- Work on branch and make a PR (PR review is required)
- If no review is received within 2 days, you are free to merge to dev

### Semantic Commit messages
- feat: :zap: 새로운 기능 (issue 번호 달 것)
- fix: :bug: 버그 수정
- refactor: :hammer: 기능을 추가하지 않는 코드 변경
- chore: :package: src 또는 테스트 파일을 수정하지 않는 기타 변경 사항
- docs: :books: 문서만 변경
- style: :shirt: 코드의 의미에 영향을 미치지 않는 변경 사항 (공백, 서식, 누락 된 세미콜론 등)
- test: :rotating_light: 테스트 케이스 추가
- deploy: :rocket:
