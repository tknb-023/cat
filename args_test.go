package main

import (
	"testing"
)

func TestParseArguments(t *testing.T) {
	testdata := []struct {
		giveArgs     []string
		wontNumber   bool
		wontNonblank bool
		wontSqueeze  bool
		wontTable    bool
		wontHelp     bool
		wontError    bool
		message      string
	}{
		{[]string{"ccat"}, false, false, false, false, false, false, "success with stdin"},
		{[]string{"ccat", "-h"}, false, false, false, false, true, false, "success with help"},
		{[]string{"ccat", "-n"}, true, false, false, false, false, false, "success with number"},
		{[]string{"ccat", "-b"}, false, true, false, false, false, false, "success with number-nonblank"},
		{[]string{"ccat", "-s"}, false, false, true, false, false, false, "success with squeeze"},
		{[]string{"ccat", "-t"}, false, false, false, true, false, false, "success with table"},
		{[]string{"ccat", "./testdata/hotarunohikari.txt"}, false, false, false, false, false, false, "success"},
		{[]string{"ccat", "-e"}, false, false, false, false, false, true, "parsing fail: the unknown flag"},
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
		if opts.number != td.wontNumber {
			t.Errorf("parseArgs(%v) number did not match, wont %v, but got %v", td.giveArgs, td.wontNumber, opts.number)
		}
		if opts.nonblank != td.wontNonblank {
			t.Errorf("parseArgs(%v) number-nonblank did not match, wont %v, but got %v", td.giveArgs, td.wontNonblank, opts.nonblank)
		}
		if opts.squeeze != td.wontSqueeze {
			t.Errorf("parseArgs(%v) nsqueeze did not match, wont %v, but got %v", td.giveArgs, td.wontSqueeze, opts.squeeze)
		}
		if opts.table != td.wontTable {
			t.Errorf("parseArgs(%v) table did not match, wont %v, but got %v", td.giveArgs, td.wontTable, opts.table)
		}
		if opts.help != td.wontHelp {
			t.Errorf("parseArgs(%v) help did not match, wont %v, but got %v", td.giveArgs, td.wontHelp, opts.help)
		}
	}
}
