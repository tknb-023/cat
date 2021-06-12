package main

import (
	"bufio"
	"fmt"
	"io"
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

func stin() {
	scanner := bufio.NewScanner(os.Stdin)

	// 標準入力から受け取ったテキストを出力
	for scanner.Scan() {
		fmt.Fprintf(os.Stdout, "%s\n", scanner.Text())
	}
}

func stout(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	rd := bufio.NewReader(file)
	for {
		s, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(s)
	}
	file.Close()
}

func stout_loop(filename []string) {
	for _, file := range filename {
		stout(file)
	}
}

func perform(opts *options) int {
	if len(opts.args) < 1 {
		stin()
	} else {
		stout_loop(opts.args)
	}
	return 0
}

func goMain(args []string) int {
	opts, err := parseArgs(args)
	if err != nil {
		fmt.Printf("parsing args fail: %s\n", err.Error())
		fmt.Println(helpMessage(filepath.Base(args[0])))
		return 1
	}
	if opts.help {
		fmt.Println(helpMessage(filepath.Base(args[0])))
		return 0
	}
	return perform(opts)
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
