package main

import "testing"

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

func TestParseArguments(t *testing.T) {
	testdata := []struct {
		giveArgs     []string
		wontNumber   bool
		wontNonblank bool
		wontSqueeze  bool
		wontTable    bool
		wontHelp     bool
		wontError    bool
		wantArgs     []string
		message      string
	}{
		{[]string{"ccat"}, false, false, false, false, false, false, []string{}, "success with stdin"},
		{[]string{"ccat", "-h"}, false, false, false, false, true, false, []string{}, "success with help"},
		{[]string{"ccat", "-h", "--background"}, false, false, false, false, true, false, []string{}, "success with help"},
		{[]string{"ccat", "./testdata/hotarunohikari.txt"}, false, false, false, false, false, false, []string{"../../testdata/hotarunohikari.txt"}, "success"},
		// {[]string{"ccat", "10"}, false, false, false, false, false, false, "success"},
		// {[]string{"ccat", "-H", "10"}, false, false, false, false, false, false, "success with header"},
		// {[]string{"ccat", "10", "--", "echo", "hoge"}, false, false, false, false, false, false, "success with command"},
		// {[]string{"ccat", "-u", "unknown_unit", "10"}, false, false, false, false, false, false, "parsing fail: the unknown unit"},
		// {[]string{"ccat", "10", "--"}, false, false, false, false, false, true, "parsing fail: no commands given"},
		// {[]string{"ccat", "-unknown-flag"}, false, false, false, false, false, true, "parsing fail: the unknown flag"},
	}
	for _, td := range testdata {
		opts, err := parseArgs(td.giveArgs)
		if (err == nil) && td.wontError {
			t.Errorf("parseArgs(%v) wont error, but got no error: %s", td.giveArgs, td.message)
		}
		if err != nil && !td.wontError {
			t.Errorf("parseArgs(%v) wont no error, but got error: %s (%s)", td.giveArgs, err.Error(), td.message)
		}
		if err != nil {
			continue
		}
		// if opts.number != td.wontNumber {
		// 	t.Errorf("parseArgs(%v) background did not match, wont %v, but got %v", td.giveArgs, td.wontNumber, opts.number)
		// }
		// if opts.nonblank != td.wontNonblank {
		// 	t.Errorf("parseArgs(%v) header did not match, wont %v, but got %v", td.giveArgs, td.wontNonblank, opts.nonblank)
		// }
		if opts.help != td.wontHelp {
			t.Errorf("parseArgs(%v) help did not match, wont %v, but got %v", td.giveArgs, td.wontHelp, opts.help)
		}
	}
}
