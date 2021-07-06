---
title: "CCAT"
---

[![CI](https://github.com/tknb-023/ccat/actions/workflows/blank.yml/badge.svg)](https://github.com/tknb-023/ccat/actions/workflows/blank.yml)
[![Coverage Status](https://coveralls.io/repos/github/tknb-023/ccat/badge.svg?branch=main)](https://coveralls.io/github/tknb-023/ccat?branch=main)
[![codebeat badge](https://codebeat.co/badges/7baf5730-be98-43c2-b642-49e7887af865)](https://codebeat.co/projects/github-com-tknb-023-ccat-main)
[![License](https://img.shields.io/github/license/tknb-023/ccat)](https://github.com/tknb-023/ccat/blob/main/LICENSE)
[![Docker](https://img.shields.io/badge/Docker-saku2975%2Fccat%3A1.0.0-green?logo=docker)](https://hub.docker.com/r/saku2975/ccat)

catコマンドの機能拡張

## Description

catコマンドではファイルの結合が可能だが，結合の際はテキスト全てを結合する．
しかし，csvファイルなどで先頭行にカラム名がついている場合，2つ目以降のファイルではカラム名は必要ない．
そこでこのソフトウェアではcatコマンドの機能の一部に加え，2つ目以降のファイルでは先頭行を無視する機能を加える．

## Usage

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
## icon
<img width="100" alt="ccat_icon.svg" src="./img/ccat_icon.svg">
![Icon](https://raw.githubusercontent.com/tknb-023/ccat/main/docs/static/img/ccat_icon.svg)
アイコンは www.flaticon.com で公開されている画像を用いている．

