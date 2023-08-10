package extension

import (
	"fmt"
	"io"
	"net/http"
)

func FooBarBaz(w http.ResponseWriter) {
	for _, str := range []string{"foo", "bar", "baz"} {
		io.WriteString(w, fmt.Sprintf("%s\n", str))
	}
}
