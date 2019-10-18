package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/russross/blackfriday"
)

func main() {
	flag.Parse()

	// 第一引数にマークダウンのファイルを受け取る
	// 第一引数から絶対パスを得る
	apath, _ := filepath.Abs(flag.Arg[0])
	// ファイルパスをディレクトリパスとファイル名に分割する
	dpath, filename := filepath.Split(apath)
	// 拡張子を得る
	ext, _ := filepath.Ext(filename)
	// 拡張子なしの名前を得る
	basename, _ := filename[:len(filename)-len(ext)]

	if ext != ".md" {
		fmt.Println("拡張子が.mdではありません")
		fmt.Fprintf(
			os.Stderr,
			"%s は拡張子が.mdではありません。\n"+
				"拡張子が.mdのマークダウンのファイルを指定してください。\n",
			filename)
		os.Exit(1)
	}

	md, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	position := strings.Index(flag.Arg(0), ".md")
	filename := flag.Arg(0)[:position]
	fmt.Println(filename)

	html := blackfriday.MarkdownBasic(md)

	htmlpath := dpath + "/" + basename + ".html"
	err := ioutil.WriteFile(htmlpath, html, 644)
	if err != nil {
		// Openエラー処理
		fmt.Fprintf(os.Stderr, "File %s could not open : %v\n", htmlpath, err)
		os.Exit(1)
	}
}
