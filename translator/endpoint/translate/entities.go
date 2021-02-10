package translate

type Request struct {
	TextoATraducir  string `json:"textoATraducir"`
	FormatoOrigen   string `json:"formatoOrigen"`
	FormatoDestino   string `json:"formatoDestino "`
}

type Response struct {
	TextoTraducido string `json:"textoTraducido "`
} 
