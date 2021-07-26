package yerro

import (
	"fmt"
	"net/http"
	"strings"
)

func ResponderHTTP(w http.ResponseWriter, err error) {
	Nuevo(err).ResponderHTTP(w)
}

func (err *Error) ResponderHTTP(w http.ResponseWriter) {
	var código = err.código
	if código <= 0 {
		código = http.StatusInternalServerError
	}

	var mensaje string
	if len(strings.TrimSpace(err.mensaje)) > 0 {
		mensaje = fmt.Sprintf(`{"error":"%s"}`, err.mensaje)
	}

	http.Error(w, mensaje, código)
}
