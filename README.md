# ccat

[![CI](https://github.com/tknb-023/ccat/actions/workflows/blank.yml/badge.svg)](https://github.com/tknb-023/ccat/actions/workflows/blank.yml)
[![Coverage Status](https://coveralls.io/repos/github/tknb-023/ccat/badge.svg?branch=main)](https://coveralls.io/github/tknb-023/ccat?branch=main)
[![codebeat badge](https://codebeat.co/badges/7baf5730-be98-43c2-b642-49e7887af865)](https://codebeat.co/projects/github-com-tknb-023-ccat-main)

[![License](https://img.shields.io/github/license/tknb-023/ccat)](https://github.com/tknb-023/ccat/blob/main/LICENSE)
[![Version](https://img.shields.io/badge/Version-0.9.4-orange)](https://github.com/tknb-023/ccat/releases/tag/v0.9.4)
[![DOI](https://zenodo.org/badge/370349891.svg)](https://zenodo.org/badge/latestdoi/370349891)

[![Docker](https://img.shields.io/badge/Docker-saku2975%2Fccat%3A1.0.1-green?logo=docker)](https://hub.docker.com/r/saku2975/ccat)

catコマンドの機能拡張

## Description

catコマンドではファイルの結合が可能だが，結合の際はテキスト全てを結合する．
しかし，csvファイルなどで先頭行にカラム名がついている場合，2つ目以降のファイルではカラム名は必要ない．
そこでこのソフトウェアではcatコマンドの機能の一部に加え，2つ目以降のファイルでは先頭行を無視する機能を加える．

## Usage

### CLI help message

```sh
ccat [OPTIONS]  [FILEs...]
OPTIONS
    -n, --number            行番号を表示する．
    -b, --number-nonblank   行番号を表示する．ただし空白行には付けない．
    -s, --squeeze-blank     連続した空行を1行にする．
    -t, --table             2つ目以降のファイルでは先頭行を無視する．
    -h, --help              このメッセージを出力します．
ARGUMENTS
    FILEs...                中身の確認または結合を行うファイル．
```

### Examples

```sh
$ ccat testdata/csv1.csv                     
ID,Name,Mail
123456,神山太郎,123456@abc.com
234567,京産花子,234567@abc.com
345678,京都次郎,345678@abc.com
$ ccat testdata/csv2.csv -n
  1 ID,Name,Mail
  2 456789,神山三郎,456789@abc.com
  3 567890,京産町子,567890@abc.com
$ ccat testdata/csv1.csv testdata/csv2.csv -t 
ID,Name,Mail
123456,神山太郎,123456@abc.com
234567,京産花子,234567@abc.com
345678,京都次郎,345678@abc.com
456789,神山三郎,456789@abc.com
567890,京産町子,567890@abc.com
```

## Install

### Homebrew

```shell
brew tap tknb-023/brew
brew install ccat
```

### Go lang

```shell
go get github.com/tamada/scv
```

### Compiling yourself

```shell
git clone https://github.com/tknb-023/ccat.git
cd ccat
make
```

## About

### Version

```1.0.1``` , ```latest```
```1.0.0```

### License

MIT License

[![License](https://img.shields.io/github/license/tknb-023/ccat)](https://github.com/tknb-023/ccat/blob/main/LICENSE)

### icon

<img width="100" alt="ccat_icon.svg" src="./docs/static/images/ccat_icon.svg">
アイコンは www.flaticon.com で公開されている画像を用いている．
