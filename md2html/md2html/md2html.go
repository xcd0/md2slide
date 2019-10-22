package md2html

import ( // {{{
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/google/go-github/github"
	"github.com/russross/blackfriday"
	gfm "github.com/shurcooL/github_flavored_markdown"
	"github.com/xcd0/go-nkf"
) // }}}

type Fileinfo struct { // {{{
	Apath    string // 入力mdファイルの絶対パス
	Dpath    string // 入力mdファイルのあるディレクトリのパス
	Filename string // 入力mdファイルのファイル名
	Basename string // 入力mdファイルのベースネーム 拡張子抜きの名前
	Ext      string // 入力mdファイルの拡張子
	Htmlpath string // 生成されるhtmlファイルの出力先パス
	Flavor   string // 生成に用いるmarkdownの方言
	Html     string // 生成したhtml本体が入る
	Pdfpath  string // 生成されるpdfファイルの出力先パス
} // }}}

func Makehtml(fi Fileinfo) (string, error) { // {{{

	header := Makeheader()
	body, err := Makebody(fi)
	footer := Makefooter()

	fi.Html = header + body + footer

	return fi.Html, err
} // }}}

func Argparse(arg string) Fileinfo { // {{{

	fi := Fileinfo{}

	// 絶対パスを得る
	fi.Apath, _ = filepath.Abs(arg)
	// ファイルパスをディレクトリパスとファイル名に分割する
	fi.Dpath, fi.Filename = filepath.Split(fi.Apath)
	// 拡張子を得る
	fi.Ext = filepath.Ext(fi.Filename)
	// 拡張子なしの名前を得る
	fi.Basename = fi.Filename[:len(fi.Filename)-len(fi.Ext)]
	// 出力するhtmlのパスを得る
	fi.Htmlpath = fi.Dpath + fi.Basename + ".html"

	return fi
} // }}}

func Makeheader() string { // {{{
	header := `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" />
<style type="text/css"><!--
body{font-family:Helvetica,arial,sans-serif;font-size:14px;line-height:1.8;padding:30px;background-color:#fff;color:#333}body>:first-child{margin-top:0!important}body>:last-child{margin-bottom:0!important}a{color:#4183c4;tExt-decoration:none}a.absent{color:#c00}a.anchor{display:block;padding-left:30px;margin-left:-30px;cursor:pointer;position:absolute;top:0;left:0;bottom:0x}h1,h2,h3,h4,h5,h6{margin:20px 0 10px;padding:0;font-weight:700;-webkit-font-smoothing:antialiased;cursor:tExt;position:relative}h1:first-child,h1:first-child+h2,h2:first-child,h3:first-child,h4:first-child,h5:first-child,h6:first-child{margin-top:0;padding-top:0}h1:hover a.anchor,h2:hover a.anchor,h3:hover a.anchor,h4:hover a.anchor,h5:hover a.anchor,h6:hover a.anchor{tExt-decoration:none}h1 code,h1 tt,h2 code,h2 tt,h3 code,h3 tt,h4 code,h4 tt,h5 code,h5 tt,h6 code,h6 tt{font-size:inherit}h1{font-size:34px;margin-bottom:40px;padding-bottom:0}h1,h2{color:#000}h2{padding-top:20px;font-size:30px;border-bottom:2px solid #ccc}h3{font-size:24px;border-bottom:1px solid #ddd}h4{font-size:20px}h5{font-size:18px}h6{font-size:16px;color:#777}blockquote,dl,li,ol,p,pre,table,ul{margin:15px 0}hr{border:0 0 0;color:#ccc;height:4px;padding:0}a:first-child h1,a:first-child h2,a:first-child h3,a:first-child h4,a:first-child h5,a:first-child h6,body>h1:first-child,body>h1:first-child+h2,body>h2:first-child,body>h3:first-child,body>h4:first-child,body>h5:first-child,body>h6:first-child{margin-top:0;padding-top:0}h1 p,h2 p,h3 p,h4 p,h5 p,h6 p{margin-top:0}li p.first{display:inline-block}ol,ul{padding-left:30px}ol:first-child,ul:first-child{margin-top:0}ol:last-child,ul:last-child{margin-bottom:0}dl,dl dt{padding:0}dl dt{font-size:14px;font-weight:700;font-style:italic;margin:15px 0 5px}dl dt:first-child{padding:0}dl dt>:first-child{margin-top:0}dl dt>:last-child{margin-bottom:0}dl dd{margin:0 0 15px;padding:0 15px}dl dd>:first-child{margin-top:0}dl dd>:last-child{margin-bottom:0}blockquote{border-left:4px solid #ddd;padding:0 15px;color:#777}blockquote>:first-child{margin-top:0}blockquote>:last-child{margin-bottom:0}table{padding:0;border-spacing:2px;border-collapse:collapse;width:80%;margin:auto}table,td,th{border:1px solid #ccc}td,th{padding:0;margin:0}table tr{background-color:#fff;border-top:1px solid #c6cbd1;margin:0;padding:0}table tr:nth-child(2n){background-color:#f6f8fa}table tr th{font-weight:700}table tr td,table tr th{border:1px solid #ccc;tExt-align:center;margin:0;padding:6px 13px}table tr td:first-child,table tr th:first-child{margin-top:0}table tr td:last-child,table tr th:last-child{margin-bottom:0}img{max-width:100%}span.frame,span.frame>span{display:block;overflow:hidden}span.frame>span{border:1px solid #ddd;float:left;margin:13px 0 0;padding:7px;width:auto}span.frame span img{display:block;float:left}span.frame span span{clear:both;color:#333;display:block;padding:5px 0 0}span.align-center{display:block;overflow:hidden;clear:both}span.align-center>span{display:block;overflow:hidden;margin:13px auto 0;tExt-align:center}span.align-center span img{margin:0 auto;tExt-align:center}span.align-right{display:block;overflow:hidden;clear:both}span.align-right>span{display:block;overflow:hidden;margin:13px 0 0;tExt-align:right}span.align-right span img{margin:0;tExt-align:right}span.float-left{display:block;margin-right:13px;overflow:hidden;float:left}span.float-left span{margin:13px 0 0}span.float-right{display:block;margin-left:13px;overflow:hidden;float:right}span.float-right>span{display:block;overflow:hidden;margin:13px auto 0;tExt-align:right}code,tt{margin:0 2px;padding:0 5px;white-space:nowrap;border:1px solid #eaeaea;background-color:#f8f8f8;border-radius:3px}pre code{margin:0;padding:0;white-space:pre;border:0;background:0 0}.highlight pre,pre{background-color:#f8f8f8;border:1px solid #ccc;font-size:13px;line-height:19px;overflow:auto;padding:6px 10px;border-radius:3px}pre code,pre tt{background-color:transparent;border:0}.main-content{max-width:50pc;margin:auto;padding-bottom:50px}hr{page-break-after:always;border:0!important;color:#fff;height:4px;padding:0}
--></style>
<script type="text/x-mathjax-config">
	MathJax.Hub.Config({
		tex2jax: { inlineMath: [['$','$'], ['\\(','\\)']], processEscapes: true },
		CommonHTML: { matchFontHeight: false }
	});
</script>

<!-- オフラインの時様に取ってきたからいらない気がする -->
<!--
<script src="https://polyfill.io/v3/polyfill.min.js?features=es6"></script>
<script id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-mml-chtml.js"></script>
-->

<script id="MathJax-script" async src="./md2html/md2html/MathJax-3.0.0/es5/tex-mml-chtml.js"></script>

</head>
<body>
`
	return header
} // }}}

func Makebody(fi Fileinfo) (string, error) { //{{{

	if fi.Ext != ".md" {
		fmt.Println("拡張子が.mdではありません")
		fmt.Fprintf(
			os.Stderr,
			"%s は拡張子が.mdではありません。\n"+
				"拡張子が.mdのマークダウンのファイルを指定してください。\n",
			fi.Filename)
		return "", errors.New("The Extention of this file is not .md")
	}

	bytemd, err := ioutil.ReadFile(fi.Apath)
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"ファイル%vが読み込めません\n",
			fi.Filename)
		log.Fatal(err)
		return "", err
	}

	// ファイルの文字コード変換
	charset, err := nkf.CharDet(bytemd)
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"文字コード変換に失敗しました\n"+
				"utf8を使用してください\n")
		log.Fatal(err)
		return "", err
	}
	stringmd, err := nkf.ToUtf8(string(bytemd), charset)
	bytemd = []byte(stringmd)
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"文字コード変換に失敗しました\n"+
				"utf8を使用してください\n")
		log.Fatal(err)
		return "", err
	}

	var bytebody []byte
	if fi.Flavor == "github" {
		bytebody, err = gitHubAPI(bytemd)
	} else if fi.Flavor == "gfm" {
		bytebody, err = shurcooL_GFM(bytemd)
	} else {
		bytebody = blackfriday.MarkdownBasic(bytemd)
	}

	body := string(bytebody)

	return body, err

} //}}}

func Makefooter() string { // {{{
	footer := "</body>\n</html>"
	return footer
} // }}}

func gitHubAPI(md []byte) ([]byte, error) { // {{{
	client := github.NewClient(nil)
	opt := &github.MarkdownOptions{Mode: "gfm", Context: "google/go-github"}
	body, _, err := client.Markdown(context.Background(), string(md), opt)
	return []byte(body), err
} // }}}

func shurcooL_GFM(md []byte) ([]byte, error) { // {{{

	bytehtml := gfm.Markdown(md)

	return bytehtml, nil
}

// }}}