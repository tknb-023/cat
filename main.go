package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func helpMessage(originalProgramName string) string {
	programName := filepath.Base(originalProgramName)
	return fmt.Sprintf(`%s [OPTIONS]  [FILEs...]
OPTIONS
	-n, --number            行番号を表示する．
	-b, --number-nonblank   行番号を表示する．ただし空白行には付けない．
	-s, --squeeze-blank     連続した空行を1行にする．
	-t, --table             2つ目以降のファイルでは先頭行を無視する．
	-h, --help              このメッセージを出力します．
ARGUMENTS
	FILEs...                中身の確認または結合を行うファイル．`, programName)
}

func goMain(args []string) int {
	fmt.Println(helpMessage(args[0]))
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
