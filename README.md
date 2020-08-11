# HashID Tool

A command line tool that encodes & decodes [hashids](https://hashids.org/).

## Installation

Make sure your Go Version is greater than 1.13. And simply run:

```shell
$ GO111MODULE=on go install github.com/gg-tools/hashid
```

Make sure your `PATH` includes the `$GOPATH/bin` directory so commands can be easily used:

`export PATH=$PATH:$GOPATH/bin`

## Usage

Configuration:

```shell
$ export HASHID_SALT=!@#!@F234t23g2354rr2tg35
$ export HASHID_MIN_LENGTH=8
$ export HASHID_ALPHABET=qwertyuiopasdfghjklzxcvbnm1234567890
```

Get hashed string:

```shell
$ hashid encode 1988
25yodre1
$ hashid encode 1988 1990
25yodre1
re7og60y
```

Get ID from hashed string:

```shell
$ hashid decode 25yodre1
$ hashid decode 25yodre1 re7og60y
```