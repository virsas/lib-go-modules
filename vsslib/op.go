package vsslib

import (
	"errors"
	"log"

	"github.com/1Password/connect-sdk-go/connect"
	"github.com/1Password/connect-sdk-go/onepassword"
)

type OpHandler interface {
	Get(item string, label string) (string, error)
}

type op struct {
	client connect.Client
	vault  string
}

func NewOPSession(opURL string, opToken string, opVault string) (OpHandler, error) {
	var err error
	client := connect.NewClient(opURL, opToken)

	vaults, err := client.GetVaults()
	if err != nil {
		return nil, err
	}

	var vault string = ""
	for _, v := range vaults {
		if v.Name == opVault {
			vault = v.ID
			break
		}
	}

	if vault == "" {
		err = errors.New("unable to find vault")
		return nil, err
	}

	o := &op{client: client, vault: vault}

	return o, nil
}

func (o *op) Get(item_name string, label string) (string, error) {
	var err error
	var item *onepassword.Item
	var value string

	item, err = o.client.GetItemByTitle(item_name, o.vault)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range item.Fields {
		if f.Label == label {
			value = f.Value
			break
		}
	}

	return value, nil
}
