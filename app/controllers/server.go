package controllers

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"todo_app1/app/models"
	"todo_app1/config"
)

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

func session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

// urlの型を指定
// |(または)で複数url宣言
var validPath = regexp.MustCompile("^/todos/(edit|save|update|delete)/([0-9]+)$")

// /todos/edit/1
// 最後の1にIDが入る

// リクエストがあったら，そのurlからIDを取得の関数
// ハンドラファンク型を返す
func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// URL.PathとvalidPathの一致部分
		// マッチする部分をスライスで取得
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		// IDが数値で取得できているはず
		// int型に変換
		id, _ := strconv.Atoi(q[2])
		fmt.Println(id)
		// idを渡してtodoの編集などに渡す
		fn(w, r, id)
	}
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// httpで指定するためにハンドラのurlを登録
	// 指定したurlの末尾に/ない場合は完全一致を求める
	// /がある場合は，先頭が一致すれば良い(末尾にurlの続きがあっていい)
	http.HandleFunc("/", top) //top
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))

	port := os.Getenv("PORT")
	return http.ListenAndServe(":"+port, nil)
}
