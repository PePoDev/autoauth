# AutoAuth

![Banner](https://raw.githubusercontent.com/PePoDev/pepodev.github.io/master/doc-assets/autoauth/banner.png)

[![Go Report Card](https://goreportcard.com/badge/github.com/pepodev/autoauth)](https://goreportcard.com/report/github.com/pepodev/autoauth) [![GoDoc](https://godoc.org/github.com/PePoDev/autoauth?status.svg)](http://godoc.org/github.com/PePoDev/autoauth) [![codebeat badge](https://codebeat.co/badges/b7d3c2af-ac18-457e-9ff0-4976f11061d3)](https://codebeat.co/projects/github-com-pepodev-autoauth-master) [![GolangCI](https://golangci.com/badges/github.com/PePoDev/autoauth.svg)](https://golangci.com) [![fuzzit](https://app.fuzzit.dev/badge?org_id=pepodev-gh)](https://app.fuzzit.dev/orgs/pepodev-gh/dashboard) [![Build Status](https://travis-ci.com/PePoDev/autoauth.svg?branch=master)](https://travis-ci.com/PePoDev/autoauth) [![codecov](https://codecov.io/gh/PePoDev/autoauth/branch/master/graph/badge.svg)](https://codecov.io/gh/PePoDev/autoauth) [![Sourcegraph](https://sourcegraph.com/github.com/PePoDev/autoauth/-/badge.svg)](https://sourcegraph.com/github.com/PePoDev/autoauth?badge) [![GitHub tag](https://img.shields.io/github/tag/PePoDev/autoauth.svg)]() [![Docker Repository on Quay](https://quay.io/repository/pepodev/autoauth/status "Docker Repository on Quay")](https://quay.io/repository/pepodev/autoauth) ![GitHub](https://img.shields.io/github/license/pepodev/autoauth)

**AutoAuth** is Automatic Internet Portal Authentication CLI written in Golang 🛰🛰

## 📚 Table of contents

- [Installation](#-installation)

- [Example Preset](#-example-preset-file)

- [Usage](#-cli-usage)

- [Todos](#-todos)

- [Dependencies](#-dependencies)

- [Contributing](#️️-contributing)

## ⛏ Installation

### Go get

```console
~$ go get -u github.com/pepodev/autoauth
```

Supported on Windows, Maybe Linux and MacOS also support too (untested)

### Docker

```console
~$ docker run --name autoauth -d -v ${pwd}/autoauth.yml:/autoauth.yml quay.io/pepodev/autoauth
```

Note: docker network need ability to access external internet

### Complied binary (coming soon too)

Check [release](https://github.com/PePoDev/autoauth/releases) page to see available binary file.

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
    timeout: 5

  logout:
    endpoint: https://connect.kmitl.ac.th:8445/PortalServer/Webauth/webAuthAction!logout.action
    method: POST
    header:
      - User-Agent:AutoAuth
      - X-XSRF-TOKEN:{token}
    timeout: 5

  heartbeat:
    endpoint: http://clients3.google.com/generate_204
    method: GET
    header:
      - User-Agent:AutoAuth
    interval: 5
    timeout: 5
    retry: 3

  save:
  - token
```

Note: also support in various format include json, toml, hcl, envfile. It's powered by [Viper](https://github.com/spf13/viper)

## 📕 CLI Usage

Basic usage run this command to **Start** Autoauth

```console
~$ autoauth start -f ./autoauth.yml
```

![Start Command](https://raw.githubusercontent.com/PePoDev/pepodev.github.io/master/doc-assets/autoauth/screenshots/screenshot-2.png)

Or use **Help** command to see, What's Autoauth CLI can do for you.

```console
~$ autoauth --help
```

![Banner](https://raw.githubusercontent.com/PePoDev/pepodev.github.io/master/doc-assets/autoauth/screenshots/screenshot-1.png)

## 📝 Todos

- ~~Add docker image~~

- Add complied binary file

- Save key from response for use in another purpose

- Add retries rule

- Create preset file from CLI

- Decrypt with key for read secret variables from preset file

- ~~Implement timeout from preset file~~

- Write unit tests

- Write benchmark

- Write documents

- Create GUI

## 🛒 Dependencies

- [Cobra](https://github.com/spf13/cobra) A Commander for modern Go CLI interactions

- [Viper](https://github.com/spf13/viper) Go configuration with fangs

- [Fasthttp](https://github.com/valyala/fasthttp) Used for httpclient. Zero memory allocations. Up to 10x faster than net/http

- [Go-ps](https://github.com/mitchellh/go-ps) Find, list, and inspect processes from Go (golang).

## 🕵️‍♀️ Contributing

Yes, Thanks for all contributor.

If you have any question or issues also check [this](https://github.com/PePoDev/autoauth/issues/new) page.

## Specials Thanks

- [Wandee](https://github.com/userkmitl) for internet portal endpoint at KMITL university.
