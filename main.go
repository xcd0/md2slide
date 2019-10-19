package main

import ( // {{{
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"./md2html/md2html"
) // }}}

func main() {
	flag.Parse()
	// 第一引数にマークダウンのファイルを受け取る
	fi := md2html.Argparse(flag.Arg(0))
	make_html_by_schooL(fi)
}

func make_html_by_schooL(fi md2html.Fileinfo) {

	fi.Flavor = "gfm"
	// htmlを作成する
	html, err := md2html.Makehtml(fi)
	if err != nil {
		fmt.Println(err)
		return
	}

	htmlpath := fi.Dpath + "/" + fi.Basename + ".html"

	err = ioutil.WriteFile(htmlpath, []byte(html), 0644)
	if err != nil {
		// Openエラー処理
		fmt.Fprintf(os.Stderr, "File %s could not open : %v\n", fi.Htmlpath, err)
		fmt.Println(err)
		return
	}
}
