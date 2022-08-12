package tmio

import nadeo "github.com/codecat/gonadeo"

type client struct {
	nadeoClient nadeo.Nadeo
}

func NewClient(nadeoClient nadeo.Nadeo) *client {
	return &client{
		nadeoClient: nadeoClient,
	}
}
