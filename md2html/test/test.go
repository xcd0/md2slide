package main

import ( // {{{
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"../md2html"
) // }}}

func main() {
	flag.Parse()
	// 第一引数にマークダウンのファイルを受け取る
	fi := md2html.Argparse(flag.Arg(0))

	black(fi)
	schooL(fi)
	githubapi(fi)

}

func black(fi md2html.Fileinfo) {
	// htmlを作成する
	html, err := md2html.Makehtml(fi)
	if err != nil {
		fmt.Println("生成に失敗しました")
		fmt.Println(err)
		return
	}

	htmlpath := fi.Dpath + "/" + "test_" + fi.Basename + "_black.html"

	err = ioutil.WriteFile(htmlpath, []byte(html), 0644)
	if err != nil {
		// Openエラー処理
		fmt.Fprintf(os.Stderr, "File %s could not open : %v\n", htmlpath, err)
		fmt.Println(err)
		return
	}
}

func schooL(fi md2html.Fileinfo) {
	fi.Flavor = "gfm"
	// htmlを作成する
	html, err := md2html.Makehtml(fi)
	if err != nil {
		fmt.Println("生成に失敗しました")
		fmt.Println(err)
		return
	}

	htmlpath := fi.Dpath + "/" + "test_" + fi.Basename + "_schooL.html"

	err = ioutil.WriteFile(htmlpath, []byte(html), 0644)
	if err != nil {
		// Openエラー処理
		fmt.Fprintf(os.Stderr, "File %s could not open : %v\n", fi.Htmlpath, err)
		fmt.Println(err)
		return
	}
}
func githubapi(fi md2html.Fileinfo) {
	fi.Flavor = "github"
	// htmlを作成する
	html, err := md2html.Makehtml(fi)
	if err != nil {
		fmt.Println("生成に失敗しました")
		fmt.Println("(注意)Github API はオフラインでは使用できません。")
		fmt.Println(err)
		return
	}

	htmlpath := fi.Dpath + "/" + "test_" + fi.Basename + "_github.html"

	err = ioutil.WriteFile(htmlpath, []byte(html), 0644)
	if err != nil {
		// Openエラー処理
		fmt.Fprintf(os.Stderr, "File %s could not open : %v\n", fi.Htmlpath, err)
		fmt.Println(err)
		return
	}
}
