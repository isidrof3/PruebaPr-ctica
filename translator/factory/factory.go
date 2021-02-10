package factory

import "translator/implementations"

func Factory(tipo string) implementations.ITranslate {

	// se construye de acuerdo al tipo solicitado
	switch tipo {
	case "TEXT":
		return &implementations.Texto{}

	case "BINARY":
		return &implementations.Binary{}

	case "MORSE":
		return &implementations.Morse{}

	default:
		return nil

	}
}
