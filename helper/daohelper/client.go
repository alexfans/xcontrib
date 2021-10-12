package daohelper

import (
	"xcontrib/rename/ent"
)

var (
	entClient *ent.Client
)

func EntClient() *ent.Client {
	return entClient
}

func SetEntClient(client *ent.Client) {
	entClient = client
}
