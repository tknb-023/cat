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

// func Example_ccat7() {
// 	goMain([]string{"ccat", "./testdata/hotarunohikari.txt", "-n", "-s"})
// 	// Output:
// 	// 	1 ほたるの光、窓(まど)の雪。
// 	// 	2
// 	// 	3 書(ふみ)よむ月日、重ねつつ。
// 	// 	4
// 	// 	5 いつしか年も、すぎの戸を、
// 	// 	6
// 	// 	7 明けてぞ、けさは、別れゆく。
// 	// 	8
// 	// 	9 とまるも行くも、限りとて、
// 	//    10
// 	//    11 かたみに思う、ちよろずの、
// 	//    12
// 	//    13 心のはしを、一言(ひとこと)に、
// 	//    14
// 	//    15 さきくとばかり、歌うなり。
// 	//    16
// 	//    17 筑紫(つくし)のきわみ、みちのおく、
// 	//    18
// 	//    19 海山(うみやま)とおく、へだつとも、
// 	//    20
// 	//    21 その真心(まごころ)は、へだてなく、
// 	//    22
// 	//    23 ひとつに尽くせ、国のため。
// 	//    24
// 	//    25 千島(ちしま)のおくも、沖縄(おきなわ)も、
// 	//    26
// 	//    27 八洲(やしま)のうちの、守りなり。
// 	//    28
// 	//    29 至らんくにに、いさお　しく。
// 	//    30
// 	//    31 つとめよ　わがせ、つつがなく。
// }

func Example_ccat6() {
	goMain([]string{"ccat", "./testdata/hoge.txt", "-n"})
	// Output:
	// cat: ./testdata/hoge.txt : No such file or directory
}

func Example_ccat5() {
	goMain([]string{"ccat", "./testdata/aogebatotoshi.txt", "-b", "-n"})
	// Output:
	// 	1 仰げば　とうとし、わが師の恩。

	// 	2 教(おしえ)の庭にも、はや　幾年(いくとせ)。

	// 	3 思えば　いと疾(と)し、この年月(としつき)。

	// 	4 今こそ　別れめ、いざさらば。

	// 	5 互(たがい)にむつみし、日ごろの恩。

	// 	6 別るる後(のち)にも、やよ　忘るな。

	// 	7 身を立て　名をあげ、やよ　はげめよ。

	// 	8 今こそ　別れめ、いざさらば。

	// 	9 朝夕　馴(なれ)にし、まなびの窓。

	//    10 螢の　ともし火、積む白雪(しらゆき)。

	//    11 忘るる　間(ま)ぞなき、ゆく年月。

	//    12 今こそ　別れめ、いざさらば。
}

func Example_ccat4() {
	goMain([]string{"ccat", "./testdata/csv1.csv", "./testdata/csv2.csv", "-n"})
	// Output:
	// 	1 ID,Name,Mail
	//   2 123456,神山太郎,123456@abc.com
	//   3 234567,京産花子,234567@abc.com
	//   4 345678,京都次郎,345678@abc.com
	//   1 ID,Name,Mail
	//   2 456789,神山三郎,456789@abc.com
	//   3 567890,京産町子,567890@abc.com
}

func Example_ccat3() {
	goMain([]string{"ccat", "./testdata/csv1.csv", "./testdata/csv2.csv", "-t"})
	// Output:
	// 	ID,Name,Mail
	// 123456,神山太郎,123456@abc.com
	// 234567,京産花子,234567@abc.com
	// 345678,京都次郎,345678@abc.com
	// 456789,神山三郎,456789@abc.com
	// 567890,京産町子,567890@abc.com
}

func Example_ccat2() {
	goMain([]string{"ccat", "./testdata/aogebatotoshi.txt", "-s"})
	// Output:
	// 	仰げば　とうとし、わが師の恩。

	// 教(おしえ)の庭にも、はや　幾年(いくとせ)。

	// 思えば　いと疾(と)し、この年月(としつき)。

	// 今こそ　別れめ、いざさらば。

	// 互(たがい)にむつみし、日ごろの恩。

	// 別るる後(のち)にも、やよ　忘るな。

	// 身を立て　名をあげ、やよ　はげめよ。

	// 今こそ　別れめ、いざさらば。

	// 朝夕　馴(なれ)にし、まなびの窓。

	// 螢の　ともし火、積む白雪(しらゆき)。

	// 忘るる　間(ま)ぞなき、ゆく年月。

	// 今こそ　別れめ、いざさらば。
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
