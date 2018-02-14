//
//   location /go {
//           proxy_set_header Host $host;
//	         proxy_set_header X-Real-IP $remote_addr;
//		     proxy_pass http://localhost:8080;
//			 }
//

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {

	var keys []string
	for k := range r.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, ":", r.Header[k], "</br>", r.URL.Path[1:])
		//fmt.Fprintln(w, k, ":", r.Header[k], "</br>", r.URL.Path[1:])
	}

	// get form
	r.ParseForm()
	param := r.Form
	log.Print(param["name"])
	name := fmt.Sprintf("hello, world %s!\n", param["name"][0])
	io.WriteString(w, name)
	//for _, e := range os.Environ() {
	//	fmt.Println(e)
	//}
}

func main() {
	http.HandleFunc("/go", HelloServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
