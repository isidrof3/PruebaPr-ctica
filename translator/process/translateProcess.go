package process

import (
	"fmt"
	"translator/endpoint/translate"
	"translator/factory"
)

type Service interface {
	GetTranslateProcess(request interface{}) (interface{}, error)
}

type processTranslate struct {
}

func NewProcessTranslate() *processTranslate {
	return &processTranslate{}
}

func (pt *processTranslate) GetTranslateProcess(request interface{}) (interface{}, error) {

	var req translate.Request
	req = request.(translate.Request)

	translate := factory.Factory(req.FormatoDestino)

	if translate == nil {
		fmt.Println("Formato no reconocido")
		return nil, nil
	}

	textoTraducido, err := translate.Translate(req.FormatoDestino)

	if err != nil {

		fmt.Println("Error %v", err)
		return nil, err
	}

	return textoTraducido, nil
}
