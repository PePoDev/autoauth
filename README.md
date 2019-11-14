# AutoAuth [![Go Report Card](https://goreportcard.com/badge/github.com/pepodev/autoauth)](https://goreportcard.com/report/github.com/pepodev/autoauth) [![GoDoc](https://godoc.org/github.com/PePoDev/autoauth?status.svg)](http://godoc.org/github.com/PePoDev/autoauth) [![codebeat badge](https://codebeat.co/badges/b7d3c2af-ac18-457e-9ff0-4976f11061d3)](https://codebeat.co/projects/github-com-pepodev-autoauth-master) [![GolangCI](https://golangci.com/badges/github.com/PePoDev/autoauth.svg)](https://golangci.com) [![fuzzit](https://app.fuzzit.dev/badge?org_id=pepodev-gh)](https://app.fuzzit.dev/orgs/pepodev-gh/dashboard) [![Build Status](https://travis-ci.com/PePoDev/autoauth.svg?branch=master)](https://travis-ci.com/PePoDev/autoauth) [![CircleCI](https://circleci.com/gh/PePoDev/autoauth/tree/master.svg?style=svg)](https://circleci.com/gh/PePoDev/autoauth/tree/master) [![codecov](https://codecov.io/gh/PePoDev/autoauth/branch/master/graph/badge.svg)](https://codecov.io/gh/PePoDev/autoauth) [![Sourcegraph](https://sourcegraph.com/github.com/PePoDev/autoauth/-/badge.svg)](https://sourcegraph.com/github.com/PePoDev/autoauth?badge) [![GitHub tag](https://img.shields.io/github/tag/PePoDev/autoauth.svg)]()

![Banner](https://raw.githubusercontent.com/PePoDev/pepodev.github.io/master/doc-assets/autoauth/banner.png)

Automatic Internet Portal Authentication CLI written in Go 🛰🛰

Current supported on Windows, Maybe Linux and MacOS also support too (untested)

## 📚 Table of contents

- [Installation](https://github.com/PePoDev/autoauth#installation)

- [Example Config](https://github.com/PePoDev/autoauth#example-preset-file)

- [Usage](https://github.com/PePoDev/autoauth#cli-usage)

- [Todos](https://github.com/PePoDev/autoauth#todos)

## ⛏ Installation

### go get

> go get -u github.com/pepodev/autoauth

### docker (coming soon)

> docker run --name AutoAuth -d -v ./autoauth.yml:./autoauth.yml coming_soon

Note: docker network need ability to access external internet

### Complied binary (coming soon too)

Check [release](https://github.com/PePoDev/autoauth/releases) page to download binary file.

## 📃 Example Preset file

```yml
# Example preset file for univercity KMITL wifi portal 
autoauth:
  name: KMITL
  encrypted: false

  login:
    endpoint: https://connect.kmitl.ac.th:8445/PortalServer/Webauth/webAuthAction!login.action
    method: POST
    header:
      - User-Agent:AutoAuth
    body:
      - userName=59050xxx@kmitl.ac.th
      - password=s3cr3t_p@ssw0rd
      - authLan=en
      - validCode=200
      - hasValidateNextUpdatePassword=true
      - browserFlag=en

  logout:
    endpoint: https://connect.kmitl.ac.th:8445/PortalServer/Webauth/webAuthAction!logout.action
    method: POST
    header:
      - User-Agent:AutoAuth
      - X-XSRF-TOKEN:{token}

  heartbeat:
    endpoint: http://clients3.google.com/generate_204
    method: GET
    header:
      - User-Agent:AutoAuth
    interval: 3
    timeout: 3
    retry: 3

save:
  - token
```

## 📕 CLI Usage

Help command

> autoauth --help

## 📝 Todos

- Save key from response for use in another purpose

- Add retries rule

- Implement timeout from preset file

- Write unit tests

- Write benchmark

- Write documents

## 🕵️‍♀️ Contributing

Yes, Thanks for all contributor.
