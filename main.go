package main

import ( // {{{
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"./md2html"
) // }}}

func main() {
	flag.Parse()
	// 第一引数にマークダウンのファイルを受け取る
	fi := md2html.Argparse(flag.Arg(0))

	{
		// htmlを作成する
		html, err := md2html.Makehtml(fi)
		if err != nil {
			os.Exit(1)
		}

		err = ioutil.WriteFile(fi.Htmlpath, []byte(html), 0644)
		if err != nil {
			// Openエラー処理
			fmt.Fprintf(os.Stderr, "File %s could not open : %v\n", fi.Htmlpath, err)
			os.Exit(1)
		}
	}

	{
		fi.Flavor = "gm"
		// htmlを作成する
		html, err := md2html.Makehtml(fi)
		if err != nil {
			os.Exit(1)
		}

		htmlpath := fi.Dpath + "/" + fi.Basename + "_goldmark.html"

		err = ioutil.WriteFile(htmlpath, []byte(html), 0644)
		if err != nil {
			// Openエラー処理
			fmt.Fprintf(os.Stderr, "File %s could not open : %v\n", fi.Htmlpath, err)
			os.Exit(1)
		}
	}
	{
		fi.Flavor = "github"
		// htmlを作成する
		html, err := md2html.Makehtml(fi)
		if err != nil {
			os.Exit(1)
		}

		htmlpath := fi.Dpath + "/" + fi.Basename + "_github.html"

		err = ioutil.WriteFile(htmlpath, []byte(html), 0644)
		if err != nil {
			// Openエラー処理
			fmt.Fprintf(os.Stderr, "File %s could not open : %v\n", fi.Htmlpath, err)
			os.Exit(1)
		}
	}
}
