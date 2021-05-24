# ccat

PDF結合ができるcatコマンド

![ccat_icon.svg](./img/ccat_icon.svg)

## Description

catコマンドはファイルの結合を行うことができるがpdfは結合できない．
このソフトウェアでは通常のcatコマンドの振る舞いに加えPDFの結合も行うことができる．

## Usage

```sh
ccat [OPTIONS]  [FILEs...]
OPTIONS
    -n, --number            行番号を表示する．
    -b, --number-nonblank   行番号を表示する．ただし空白行には付けない．
    -s, --squeeze-blank     連続した空行を1行にする．
    -v, --show-nonprinting  TAB、改行、改ページ以外の非表示文字を表示する．
    -t                      非表示文字を表示、TABを"^I"、用紙送りを"^L"とする．
    -E, --show-ends         行の最後に"$"を表示する．
    -A, --show-all          全ての非表示文字を表示する（-vETと同じ）．
    -e                      TABを除く全ての非表示文字を表示する（-vEと同じ）．
    -p, --pdf               PDFの結合を行う．
    -h, --help              このメッセージを出力します．
ARGUMENTS
    FILEs...                中身の確認または結合を行うファイル．
```
