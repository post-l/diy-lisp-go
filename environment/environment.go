package environment

type Environment map[string]interface{}

func New() *Environment {
	return &Environment{}
}
