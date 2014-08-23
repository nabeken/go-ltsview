# go-ltsview [![build status](https://travis-ci.org/nabeken/go-ltsview.svg?branch=master)](https://travis-ci.org/nabeken/go-ltsview)

A [LTSV](http://ltsv.org) viewer written in Go.

# Installation

```sh
$ go get -u github.com/nabeken/go-ltsview/ltsview
```

# Usage

```sh
$ ltsview --help
Usage of ltsview:
  -i="": Specify comma-separated keys to ignore
  -k="": Specify comma-separated keys to show
  -n=false: Set to true to disable colorize outputs
```

# Examples

```sh
$ cat example.ltsv | ltsview
---
host: 192.168.0.3
req: GET /nagios/ HTTP/1.1
status: 200
time: 17/Aug/2014:06:27:16 +0000
ua: Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)
---
host: 192.168.0.4
req: POST /nagios/ HTTP/1.1
status: 302
time: 18/Aug/2014:06:27:16 +0000
ua: Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)

$ cat example.ltsv | ltsview -k host,status
---
host: 192.168.0.3
status: 200
---
host: 192.168.0.4
status: 302

$ cat example.ltsv | ltsview -i req,ua
---
host: 192.168.0.3
status: 200
time: 17/Aug/2014:06:27:16 +0000
---
host: 192.168.0.4
status: 302
time: 18/Aug/2014:06:27:16 +0000
```

# Inspired by

- https://github.com/naoya/perl-Text-LTSV
- http://d.hatena.ne.jp/naoya/20130207/1360229220 (written in Japanese)

# Author

TANABE Ken-ichi

# LICENSE

See [LICENSE](LICENSE).
