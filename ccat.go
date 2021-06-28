package main

import (
	"bufio"
	"fmt"
	"os"
)

// 画面出力
func printLine(line int, opts *options, s string) {
	if opts.number || opts.nonblank {
		fmt.Printf("%3d %s\n", line, s)
	} else {
		fmt.Printf("%s\n", s)
	}
}

// 1行ずつ読み込み　オプション処理
func scanOut(opts *options, scanner *bufio.Scanner, fc int) {
	line := 1
	squeeze := 0
	for scanner.Scan() {
		s := scanner.Text()
		// -tの処理
		if opts.table && fc > 1 && line == 1 {
			line++
			continue
		}
		// sが改行のみであればsqueezeをプラス1 そうでなければ0に
		if scanner.Text() == "" {
			squeeze++
		} else {
			squeeze = 0
		}
		// -sの処理
		if opts.squeeze {
			if squeeze == 1 {
				if opts.number {
					fmt.Printf("%3d\n", line)
					line++
				} else {
					fmt.Println()
				}
				continue
			} else if squeeze > 1 {
				continue
			}
		}
		// -bの処理
		if opts.nonblank {
			if s == "" {
				fmt.Println()
				continue
			}
		}
		printLine(line, opts, s)
		line++
		squeeze = 0
	}
}

// 1ファイルずつの読み込み
func fileDataOut(opts *options, filename string, fc int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanOut(opts, scanner, fc)
	file.Close()
}

// 複数ファイル指定された場合のループ
func fileLoop(opts *options) {
	fc := 1
	for _, file := range opts.args {
		ford, err := os.Stat(file)
		if err != nil {
			fmt.Printf("ccat: %s : No such file or directory\n", file)
		} else if ford.IsDir() {
			fmt.Printf("ccat: %s : Is a directory\n", file)
		} else {
			fileDataOut(opts, file, fc)
		}
		fc++
	}
}
