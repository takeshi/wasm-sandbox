package main

import (
	"log"
	"syscall/js"

	. "github.com/takeshi/wasm-sandbox/src/util"
)

func main() {
	quit := make(chan struct{}, 0)
	html, err := Fetch("templates/editor.html")
	if err != nil {
		log.Fatal(err)
	}

	S("body").InsertAdjacentHTML("beforeend", html)

	S("#render").OnClick(func([]js.Value) {
		md := S("#markdown").Get("value").String()
		S("#preview").InnerHTML(md)
	})

	<-quit
}
