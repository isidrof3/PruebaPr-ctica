package implementations

type ITranslate interface {
	Translate(request interface{}) (interface{}, error)
}
