# utfbom-remove

[![Build Status](https://travis-ci.org/alastairruhm/utfbom-remove.svg?branch=master)](https://travis-ci.org/alastairruhm/utfbom-remove)
[![Build status](https://ci.appveyor.com/api/projects/status/6lor56a2339hd8we?svg=true&passingText=windows%20build%20pass)](https://ci.appveyor.com/project/alastairruhm/utfbom-remove)
[![codecov](https://codecov.io/gh/alastairruhm/utfbom-remove/branch/master/graph/badge.svg)](https://codecov.io/gh/alastairruhm/utfbom-remove)
[![Go Walker](https://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/alastairruhm/utfbom-remove)
[![Go Report Card](https://goreportcard.com/badge/github.com/alastairruhm/utfbom-remove)](https://goreportcard.com/report/github.com/alastairruhm/utfbom-remove)
[![codebeat badge](https://codebeat.co/badges/dcefcf89-de89-4d8a-adfb-b542b025c067)](https://codebeat.co/projects/github-com-alastairruhm-utfbom-remove-master)

detect and remove BOM in utf-8 encoding files


## feature

* remove BOM in utf8 encoding files
* dry-run mode

## download and run

https://github.com/alastairruhm/utfbom-remove/releases

download binary according to your system

```bash

$ ./utfbom-remove -h

NAME:
   utfbom-remove - detect and remove BOM in utf-8 encoding files

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   v1.0.0

AUTHOR:
   alastairruhm <alastairruhm@gmail.com>

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --check-only   dry-run mode
   --path value   the path to scan (default: ".")
   --help, -h     show help
   --version, -v  print the version

COPYRIGHT:
   (c) 2017 alastairruhm
```

### macOS and linux

```bash
# dry-run mode
$ ./utfbom-remove --path=/path/to/directory --check-only

# repalce with bom free header
$ ./utfbom-remove --path=/path/to/directory
```

### windows

```bash
# dry-run mode
$ ./utfbom-remove.exe --path=/path/to/directory --check-only

# repalce with bom free header
$ ./utfbom-remove.exe --path=/path/to/directory --check-only
```


## Supported platforms

utfbom-remove is tested against multiple versions of Go on Linux, and against the latest released version of Go on OS X and Windows. For full details, see [.travis.yml](./.travis.yml) and [appveyor.yml](./appveyor.yml).