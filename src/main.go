package main

import (
	"log"
	"syscall/js"

	u "github.com/takeshi/wasm-sandbox/src/util"
)

func main() {
	quit := make(chan struct{}, 0)
	html, err := u.Fetch("templates/editor.html")
	if err != nil {
		log.Fatal(err)
	}

	u.S("body").InsertAdjacentHTML("beforeend", html)

	u.S("#render").OnClick(func([]js.Value) {
		md := u.S("#markdown").Get("value").String()
		u.S("#preview").InnerHTML(md)
	})

	<-quit
}
