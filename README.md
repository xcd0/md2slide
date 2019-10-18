# md2slide

※ 開発中です \_(┐「ε:)\_

マークダウンからプレゼンテーション用の資料を生成します。

従来の[Marp](https://yhatt.github.io/marp/) はそれなりの可搬性がありました。  
しかし、新しく作り直されている[Marp Next](https://github.com/marp-team/marp/#readme)は現状CLIのものだけであり、可搬性が乏しいです。  
発表する際、直前になって修正というシチュエーションもあり、  
可搬性を重視してGolangで実装します。  
またスライドの比率変更やプレゼンテーション発表時に使用する  
全画面表示など実装されていないいくつかの機能があります。  
これらの実装されていない機能の実装も目指します。  
加えて、従来のMarpで問題となった任意のコード実行が可能という問題にも何らかの対策を行います。


## Requirement

### 書式

* 一般的なマークダウンの書式を拡張したものとする。

* 標準のマークダウンにはない、何らかの記述により以下の機能を実現する。

	* スライドのページ区切りを指定できる。

	* スライドの比率を指定できる。

	* フォント指定やフォントサイズ指定、フォントカラー指定ができる。

	* latex形式の数式を表示できる。

	* ページごとのタイトル表示を制御できる。

		* 印刷する文書としてはスライドの各ページにあるタイトルは不要。

		* 以下のどちらか、または両方を実装する。

			1. `<!-- previous -->` のように何らかの記述により、  
			前ページのタイトルをスライドで表示できる。

			2. `#`による見出しの記述の横に`<!-- no print -->`のような何らかの記述により、  
			印刷用の表示ではその見出しを表示しない。

	* ユーザーで独自にエイリアスを定義できる。
		* 例えば、`%%`や`$$`などの記述を何らかの記述に置き換えることができるようにする。  
		イメージとしてはC言語のマクロなど。

	* マークダウンの書式やmathjaxの書式に抵触しないものに制限したい。  
	そのため何らかの書式チェックを実装する。

* 使用方法として、2つのモードを実装する。  
どちらのモードにおいてもレンダリング時に、セキュリティ対策を行う。

	1. 生成するだけのCUIモード

	2. リアルタイムにレンダリングができるGUIモード


### 1. 単純なマークダウンからのプレゼンテーション資料作成

* バイナリに拡張子が`.md`のファイルを投げるとhtmlとpdfを生成する。

* pdfは印刷用とスライド表示用の2つを生成する。

	* スライド用のpdfでは各スライドに見出しを表示したい、  
	しかし、印刷用の文書としては同じ見出しが何度も表記されてしまうため、  
	それぞれ分けて生成する。

* htmlはブラウザを使ってそのままプレゼンテーションができるものにする。

	* pdfファイルよりいくつかの点で発表に使いやすいものとする。

		* キーボードやマウスでスライドの表示を操作できる。

		* 画像はbase64にエンコードして埋め込む。

		* 動画はネット上のもの、オフラインの外部ファイルどちらも埋め込めそのまま表示できる。

		* 音声は

			* 再生ボタン

			* 一時停止ボタン

			* 停止ボタン

			* 音量調節ボタン

			* シークバー

			を表示して、そのまま再生できる。

* コマンドラインからいくつかのリッチな機能を提供する。

	* CSSの指定

### 2. マークダウンを動的にレンダリングするGUIウィンドウを用いたプレゼンテーション資料作成、表示

* バイナリを起動するとウィンドウが開き、  
何らかの方法で拡張子が`.md`のファイルをユーザーに指定させる。

	* ドラッグ&ドロップ

	* メニューバーからのファイル指定

* ファイル指定をするとウィンドウが読み込んだマークダウンに合わせて変化し、  
また既定のテキストエディタで指定されたファイルが開く。

	* 指定されたマークダウンが更新されるたびに  
	htmlを生成しレンダリングしてウィンドウに表示する。

	* これにより、任意のテキストエディタを使用して、  
	リアルタイムに確認しながら作成することができる。
	操作感としては[Previm](https://github.com/previm/previm)をイメージしている。

* ウィンドウは全画面表示することができ、発表用のスライド表示として使用できる。

	* この時キーボードやマウスでスライドの表示を操作できる。

* ウィンドウの上部または下部に生成したhtmlのパスを表示し、  
付近のボタンを押下することで任意のブラウザで表示する。



