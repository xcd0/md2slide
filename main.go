package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/russross/blackfriday"
)

func main() {
	flag.Parse()

	// 第一引数にマークダウンのファイルを受け取る

	arg1 := flag.Arg(0)
	// 第一引数から絶対パスを得る
	apath, _ := filepath.Abs(arg1)
	// ファイルパスをディレクトリパスとファイル名に分割する
	dpath, filename := filepath.Split(apath)
	// 拡張子を得る
	ext := filepath.Ext(filename)
	// 拡張子なしの名前を得る
	basename := filename[:len(filename)-len(ext)]
	// 出力するhtmlのパスを得る
	htmlpath := dpath + "/" + basename + ".html"

	// htmlに変換 ヘッダーなどは含まれていない
	body, err := makebody(filename, ext)
	if err != nil {
		os.Exit(1)
	}

	// ヘッダーを作成する
	header := `<!DOCTYPE html>
<html>
<head>
<link rel="stylesheet" type="text/css" href="./markdown.css">
</head>
<body>
`
	footer := "</body>\n</html>"

	html := header + body + footer

	err = ioutil.WriteFile(htmlpath, []byte(html), 644)
	if err != nil {
		// Openエラー処理
		fmt.Fprintf(os.Stderr, "File %s could not open : %v\n", htmlpath, err)
		os.Exit(1)
	}
}

func makebody(filename string, ext string) (string, error) { //{{{

	if ext != ".md" {
		fmt.Println("拡張子が.mdではありません")
		fmt.Fprintf(
			os.Stderr,
			"%s は拡張子が.mdではありません。\n"+
				"拡張子が.mdのマークダウンのファイルを指定してください。\n",
			filename)
		return "", errors.New("The extention of this file is not .md")
	}

	md, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"ファイル%vが読み込めません\n",
			filename)
		log.Fatal(err)
		return "", err
	}

	body := string(blackfriday.MarkdownBasic(md))

	return body, nil

} //}}}
