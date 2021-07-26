package yerro

import "fmt"

type Error struct {
	original error
	mensaje  string
	código   int
	valores  map[string]interface{}
}

func (err *Error) Error() string {
	if err.original != nil {
		return err.original.Error()
	}

	if len(err.mensaje) > 0 {
		return fmt.Sprintf("msj: %s", err.mensaje)
	}

	if err.código != 0 {
		return fmt.Sprintf("cdg: %s", err.mensaje)
	}

	return "yerro"
}
