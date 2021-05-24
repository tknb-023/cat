# ccat

catコマンドの機能拡張

<img width="100" alt="ccat_icon.svg" src="./img/ccat_icon.svg">

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

[](
    -v, --show-nonprinting  TAB、改行、改ページ以外の非表示文字を表示する．\
    -t                      非表示文字を表示、TABを"^I"、用紙送りを"^L"とする．\
    -E, --show-ends         行の最後に"$"を表示する．\
    -A, --show-all          全ての非表示文字を表示する（-vETと同じ）．\
    -e                      TABを除く全ての非表示文字を表示する（-vEと同じ）．\
)