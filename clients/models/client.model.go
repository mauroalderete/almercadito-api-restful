package models

import (
	"fmt"
	"strings"

	hash "gitlab.com/vyra/almercadito/almercadito-api-restful/shared/hash"
)

type Client struct {
	ID            hash.Hash `json:"id"`
	Name          string    `json:"name"`
	Address       string    `json:"address"`
	Telephone     string    `json:"telephone"`
	Whatsapp      string    `json:"whatsapp"`
	Facebook      string    `json:"facebook"`
	DeliveryTimes string    `json:"deliveryTimes"`
	Comments      string    `json:"comments"`
}

func New(id hash.Hash,
	name string,
	address string,
	telephone string,
	whatsapp string,
	facebook string,
	deliveryTimes string,
	comments string) (*Client, error) {
	c := new(Client)

	if id.Empty() {
		return nil, fmt.Errorf("el ID no puede estar vacio: %v", id)
	}

	if strings.Trim(name, "\t \n") == "" {
		return nil, fmt.Errorf("el cliente debe tener un nombre válido: %v", name)
	}

	c.ID = id
	c.Name = name
	c.Address = address
	c.Telephone = telephone
	c.Whatsapp = whatsapp
	c.Facebook = facebook
	c.DeliveryTimes = deliveryTimes
	c.Comments = comments

	return c, nil
}
