package vsslib

import (
	"errors"
	"log"
	"os"

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

func NewOPSession() (OpHandler, error) {
	var err error

	var opURL string = ""
	opURLValue, opURLPresent := os.LookupEnv("OP_URL")
	if opURLPresent {
		opURL = opURLValue
	} else {
		panic("Missing ENV Variable OP_URL")
	}

	var opToken string = ""
	opTokenValue, opTokenPresent := os.LookupEnv("OP_TOKEN")
	if opTokenPresent {
		opToken = opTokenValue
	} else {
		panic("Missing ENV Variable OP_TOKEN")
	}

	var opVault string = ""
	opVaultValue, opVaultPresent := os.LookupEnv("OP_VAULT")
	if opVaultPresent {
		opVault = opVaultValue
	} else {
		panic("Missing ENV Variable OP_VAULT")
	}

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
