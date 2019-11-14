# AutoAuth [![Go Report Card](https://goreportcard.com/badge/github.com/pepodev/autoauth)](https://goreportcard.com/report/github.com/pepodev/autoauth) [![GoDoc](https://godoc.org/github.com/PePoDev/autoauth?status.svg)](http://godoc.org/github.com/PePoDev/autoauth) [![codebeat badge](https://codebeat.co/badges/b7d3c2af-ac18-457e-9ff0-4976f11061d3)](https://codebeat.co/projects/github-com-pepodev-autoauth-master) [![GolangCI](https://golangci.com/badges/github.com/PePoDev/autoauth.svg)](https://golangci.com) [![fuzzit](https://app.fuzzit.dev/badge?org_id=pepodev-gh)](https://app.fuzzit.dev/orgs/pepodev-gh/dashboard) [![Build Status](https://travis-ci.com/PePoDev/autoauth.svg?branch=master)](https://travis-ci.com/PePoDev/autoauth) [![CircleCI](https://circleci.com/gh/PePoDev/autoauth/tree/master.svg?style=svg)](https://circleci.com/gh/PePoDev/autoauth/tree/master) [![codecov](https://codecov.io/gh/PePoDev/autoauth/branch/master/graph/badge.svg)](https://codecov.io/gh/PePoDev/autoauth) [![Sourcegraph](https://sourcegraph.com/github.com/PePoDev/autoauth/-/badge.svg)](https://sourcegraph.com/github.com/PePoDev/autoauth?badge) [![GitHub tag](https://img.shields.io/github/tag/PePoDev/autoauth.svg)]()

![Banner](https://raw.githubusercontent.com/PePoDev/pepodev.github.io/master/doc-assets/autoauth/banner.png)

Automatic Internet Portal Authentication CLI written in Go 🛰🛰

## Table of contents

[Installation](google.com)

## Installation

### go get

> go get -u github.com/pepodev/autoauth

### docker

> docker run --name AutoAuth -d -v ./autoauth.yml:./autoauth.yml pepodev/autoauth

Note: docker network need ability to access external internet

## Example Preset file

```yml
# Example preset file for univercity KMITL wifi portal
autoauth:
  name: KMITL
  encrypted: false

  login:
    endpoint: https://connect.kmitl.ac.th:8445/PortalServer/Webauth/webAuthAction!login.action
    method: POST
    encoding: utf-8
    header:
      - User-Agent:AutoAuth
    body:
      - userName:59050xxx@kmitl.ac.th
      - password:s3cr3t_p@ssw0rd
      - authLan:en
      - validCode:200
      - hasValidateNextUpdatePassword:true
      - browserFlag:en
      - ClientIp:127.0.0.1

  logout:
    endpoint: https://connect.kmitl.ac.th:8445/PortalServer/Webauth/webAuthAction!logout.action
    method: POST
    encoding: utf-8
    header:
      - User-Agent:AutoAuth
      - X-XSRF-TOKEN:{token}

  heartbeat:
    endpoint: https://connect.kmitl.ac.th:8445/PortalServer/Webauth/webAuthAction!hearbeat.action
    method: POST
    encoding: utf-8
    header:
      - User-Agent:AutoAuth
      - X-XSRF-TOKEN:{token}
    body:
      - userName:59050xxx@kmitl.ac.th
    interval: 5
    timeout: 5
    retry: 2

save:
  token: null
```

## CLI Usage

Help command

> autoauth --help



## asd
