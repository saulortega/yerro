package yerro

func Nuevo(e error) *Error {
	err, ok := e.(*Error)
	if ok {
		return err
	}

	return &Error{
		original: e,
		valores:  map[string]interface{}{},
	}
}

func Mensaje(m string) *Error {
	return &Error{
		mensaje: m,
		valores: map[string]interface{}{},
	}
}

func Código(c int) *Error {
	return &Error{
		código:  c,
		valores: map[string]interface{}{},
	}
}

func (err *Error) Mensaje(m string) *Error {
	err.mensaje = m
	return err
}

func (err *Error) Código(c int) *Error {
	err.código = c
	return err
}

func (err *Error) Valor(llave string, valor interface{}) *Error {
	err.valores[llave] = valor
	return err
}
