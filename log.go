package yerro

import (
	"fmt"
	"log"
	"strings"
)

func (err *Error) Log() *Error {
	var valores = []string{}
	for l, v := range err.valores {
		valores = append(valores, fmt.Sprintf(`"%s":"%s"`, l, v))
	}

	json := fmt.Sprintf(`{"error":"%v","mensaje":"%s","código":%d,"valores":{%s}}`, err.original, err.mensaje, err.código, strings.Join(valores, ","))
	log.Println(json)

	return err
}
