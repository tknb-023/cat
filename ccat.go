package main

import (
	"bufio"
	"fmt"
	"os"
)

func printLine(line int, opts *options, s string) {
	if opts.number || opts.nonblank {
		fmt.Printf("%3d %s\n", line, s)
	} else {
		fmt.Printf("%s\n", s)
	}
}

func scanOut(opts *options, scanner *bufio.Scanner, fc int) {
	line := 1
	squeeze := 0
	for scanner.Scan() {
		s := scanner.Text()
		if opts.table && fc > 1 && line == 1 {
			line++
			continue
		}
		if scanner.Text() == "" {
			squeeze++
		} else {
			squeeze = 0
		}
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

func fileDataOut(opts *options, filename string, fc int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanOut(opts, scanner, fc)
	file.Close()
}

func existFile(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func fileLoop(opts *options) {
	fc := 1
	for _, file := range opts.args {
		if !existFile(file) {
			fmt.Printf("cat: %s : No such file or directory\n", file)
		} else {
			fileDataOut(opts, file, fc)
		}
		fc++
	}
}
