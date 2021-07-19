---
title: "Usage"
---

## CLI help message

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

## Examples

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