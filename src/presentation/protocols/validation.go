package protocols

type Validation interface {
	Validate(input interface{}) error
}
