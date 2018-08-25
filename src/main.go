package main

import (
	u "app/wasm/src/util"
	"log"
	"syscall/js"
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
