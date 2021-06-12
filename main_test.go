package main

import (
	"io/ioutil"
	"os"
	"testing"
)

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

func TestStdin(t *testing.T) {
	temp, _ := ioutil.TempFile("", "temp")
	origStdin := os.Stdin
	os.Stdin = temp
	defer func() {
		os.Stdin = origStdin
		os.Remove(temp.Name())
	}()
	temp.Write([]byte(`./testdata/hotarunohikari.txt
	./testdata/aogebatotoshi.txt`))
	temp.Seek(0, 0)

	goMain([]string{"ccat"})
	// Output:
	// ./testdata/hotarunohikari.txt
	// ./testdata/aogebatotoshi.txt
}

func Example_ccat() {
	goMain([]string{"ccat", "./testdata/hotarunohikari.txt", "./testdata/csv1.csv"})
	// Output:
	// ほたるの光、窓(まど)の雪。

	// 書(ふみ)よむ月日、重ねつつ。

	// いつしか年も、すぎの戸を、

	// 明けてぞ、けさは、別れゆく。

	// とまるも行くも、限りとて、

	// かたみに思う、ちよろずの、

	// 心のはしを、一言(ひとこと)に、

	// さきくとばかり、歌うなり。

	// 筑紫(つくし)のきわみ、みちのおく、

	// 海山(うみやま)とおく、へだつとも、

	// その真心(まごころ)は、へだてなく、

	// ひとつに尽くせ、国のため。

	// 千島(ちしま)のおくも、沖縄(おきなわ)も、

	// 八洲(やしま)のうちの、守りなり。

	// 至らんくにに、いさお　しく。

	// つとめよ　わがせ、つつがなく。
	// ID,Name,Mail
	// 123456,神山太郎,123456@abc.com
	// 234567,京産花子,234567@abc.com
	// 345678,京都次郎,345678@abc.com

}

func Test_main(t *testing.T) {
	status := goMain([]string{"ccat", "-h"})
	if status != 0 {
		t.Errorf("status code wont 0, but got %d", status)
	}
}
