package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/tsuchi77777/go-api-sample/model"
)

// "/hello" 用 Handler
type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

// "/world" 用 Handler
type WorldHandler struct{}

func (h WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

// "/hello2" 用 HandlerFunc
func Hello2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello 2!")
}

// Middleware
func Logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("前処理: %s\n", r.URL.Path)
		f(w, r)
		log.Printf("後処理: %s\n", r.URL.Path)
	}
}

// Item 一覧 を json で返す HandlerFunc
func Items(w http.ResponseWriter, r *http.Request) {
	startId := getId(r, "start_id", 1)
	limit := getId(r, "limit", 10)

	items := model.GetItems(startId, limit)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(items)
}

// Item 詳細 を html で返す HandlerFunc
func Item(w http.ResponseWriter, r *http.Request) {
	// items/{id} の {id} 取得
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "ID 未指定: %s", r.URL.Path)
		return
	}

	// HTTPメソッドをチェック (GET のみ許可)
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		fmt.Fprintf(w, "GET Method のみ対応. %s Method 未対応.", r.Method)
		return
	}

	item := model.GetItem(id)
	tmpl.Execute(w, item)
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("view/layout.html", "view/header.html"))
}

func getId(r *http.Request, key string, defaultId int) int {
	id, err := strconv.Atoi(r.FormValue(key))
	if err != nil {
		return defaultId
	}
	return id
}
