package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tsuchi77777/go-api-sample/handler"
	"github.com/tsuchi77777/go-api-sample/model"
)

// ServeHTTP(http.ResponseWriter, *http.Request) を実装すれば string も Handler になれる
// 同一 package 内ならば 小文字("myString")開始OK
type myString string

func (s myString) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", s)
}

func main() {
	//
	// Handle() の第2引数 に Handler を渡す
	//
	http.Handle("/hello", &handler.HelloHandler{})
	http.Handle("/world", handler.WorldHandler{})

	http.Handle("/mystring", myString("ServeHTTP() を実装した string"))

	// 静的ファイルを特定のディレクトリから提供
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//
	// HandleFunc() の第2引数 に HandlerFunc を渡す
	//
	http.HandleFunc("/hello2", handler.Hello2)
	// Middleware 使用
	http.HandleFunc("/hello2-log", handler.Logging(handler.Hello2))

	// 無名関数 を指定
	http.HandleFunc("/item", func(w http.ResponseWriter, req *http.Request) {
		item := model.GetItem(1)
		fmt.Fprintln(w, item)
	})

	http.HandleFunc("/items", handler.Items) // URL の末尾 が `/ 以外` の場合は 完全一致

	http.HandleFunc("/items/", handler.Item) // URL の末尾 が `/` の場合は 前方一致

	//
	// 8080ポートで起動
	//
	http.ListenAndServe(":8080", nil)
}

func init() {
	// 日付 | 時刻 | 時刻のマイクロ秒 を出力
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
}
