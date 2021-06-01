package main

func Example_help() {
	goMain([]string{"/some/path/of/ccat", "-h"})
	// Output:
	// ccat [OPTIONS]  [FILEs...]
	// OPTIONS
	// 	-n, --number            行番号を表示する．
	// 	-b, --number-nonblank   行番号を表示する．ただし空白行には付けない．
	// 	-s, --squeeze-blank     連続した空行を1行にする．
	// 	-t, --table             2つ目以降のファイルでは先頭行を無視する．
	// 	-h, --help              このメッセージを出力します．
	// ARGUMENTS
	// 	FILEs...                中身の確認または結合を行うファイル．
}
