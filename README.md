# AutoAuth

[![Go Report Card](https://goreportcard.com/badge/github.com/pepodev/autoauth)](https://goreportcard.com/report/github.com/pepodev/autoauth)

## Installation

### go get

> go get -u github.com/pepodev/autoauth

### docker

> docker run --name AutoAuth -d -v ./autoauth.yml:./autoauth.yml pepodev/autoauth

## Preset file

```yml
autoauth:
    name: simple-name
    enc: true

    login:
        endpoint: https://simple.com/login
```
