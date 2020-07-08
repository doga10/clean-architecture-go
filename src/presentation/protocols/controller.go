package protocols

type Controller interface {
	Handle(req Request) Response
}

