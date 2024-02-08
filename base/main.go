package base

import (
	"fmt"
	"net/http"
	thttp "talex-gee/base/gee"
)

func Main() {

	gee := thttp.New()

	gee.GET("/", indexHandler)
	gee.GET("/hello", helloHandler)

	gee.Run(":9999")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "URL.Path = %q\n", r.URL.Path)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Println(w, "Header[%q] = %q\n", k, v)
	}
}
