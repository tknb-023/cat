package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

type options struct {
	number   bool
	nonblank bool
	squeeze  bool
	table    bool
	help     bool
	args     []string
}

// オプションが指定された場合のフラグ立て
func parseArgs(args []string) (*options, error) {
	opts := &options{}
	flags := flag.NewFlagSet("ccat", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args[0])) }
	flags.BoolVarP(&opts.number, "number", "n", false, "行番号を表示する．")
	flags.BoolVarP(&opts.nonblank, "number-nonblank", "b", false, "行番号を表示する．ただし空白行には付けない．")
	flags.BoolVarP(&opts.squeeze, "squeeze", "s", false, "連続した空行を1行にする．")
	flags.BoolVarP(&opts.table, "table", "t", false, "2つ目以降のファイルでは先頭行を無視する．")
	flags.BoolVarP(&opts.help, "help", "h", false, "このメッセージを出力します．")
	if err := flags.Parse(args); err != nil {
		return nil, err
	}
	opts.args = flags.Args()[1:]
	return opts, nil
}
