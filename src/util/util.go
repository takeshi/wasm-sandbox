package util

import (
	"io/ioutil"
	"net/http"
	"syscall/js"
)

type HtmlElement struct {
	value *js.Value
}

type Value interface {
	String() string
}

func (elem *HtmlElement) Get(attr string) Value {
	return elem.value.Get(attr)
}

func (elem *HtmlElement) IsNull() bool {
	return *elem.value == js.Null()
}

func (elem *HtmlElement) InnerHTML(html string) *HtmlElement {

	if elem.IsNull() {
		return elem
	}
	elem.value.Set("innerHTML", html)
	return elem
}

func (elem *HtmlElement) InsertAdjacentHTML(position string, html string) *HtmlElement {

	if elem.IsNull() {
		return elem
	}
	// args := []interface{
	// 	position,
	// 	html,
	// };
	elem.value.Call("insertAdjacentHTML", position, html)
	return elem
}

func (elem *HtmlElement) OnClick(fn func([]js.Value)) *HtmlElement {
	if elem.IsNull() {
		return elem
	}
	elem.value.Set("onclick", js.NewCallback(fn))
	return elem
}

func S(selector string) *HtmlElement {
	elem := QuerySelector(selector)
	return &HtmlElement{
		value: &elem,
	}

}

var document = js.Global().Get("document")

func Fetch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func GetElementByID(id string) js.Value {
	return document.Call("getElementById", id)
}

func QuerySelector(selector string) js.Value {
	return document.Call("querySelector", selector)
}

func RenderHtml(selector string, html string) bool {
	elem := QuerySelector(selector)
	if elem == js.Null() {
		return false
	}
	elem.Set("innerHTML", html)
	return true
}
