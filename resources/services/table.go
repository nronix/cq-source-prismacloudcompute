package services

import (
	"context"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/nronix/cq-source-twistlock/client"
	"github.com/nronix/cq-source-twistlock/internal/defender"
	"net/http"
	"strconv"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func DefenderTable() *schema.Table {
	return &schema.Table{
		Name:      "twistlock_defenders",
		Resolver:  fetchDefendersTables,
		Transform: transformers.TransformWithStruct(&defender.Defenders{}),
		Multiplex: client.AccountMultiplex,
	}
}

func CloudDiscoveryTable() *schema.Table {
	return &schema.Table{
		Name:      "twistlock_cloud_vms",
		Resolver:  fetchCloudVmsTables,
		Transform: transformers.TransformWithStruct(&defender.Defenders{}),
		Multiplex: client.AccountMultiplex,
	}
}

func fetchDefendersTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	//account := cl.Account

	var lastSize = 0

	for {
		var resp = []defender.Defenders{}
		var query = map[string]string{"offset": strconv.Itoa(lastSize), "limit": "100"}

		if err := cl.Account.TwClient.Request(http.MethodGet, "api/v1/defenders", query, nil, &resp, cl.Logger()); err != nil {
			return err
		}

		if len(resp) == 0 {
			break
		}
		lastSize += len(resp)
		res <- resp
	}
	//response, err := defender.GetDefendersNew(cl)

	return nil
}

func fetchCloudVmsTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	//account := cl.Account

	var lastSize = 0

	for {
		var resp = []defender.VmsDefender{}
		var query = map[string]string{"offset": strconv.Itoa(lastSize), "limit": "100"}

		if err := cl.Account.TwClient.Request(http.MethodGet, "api/v1/cloud/discovery/vms", query, nil, &resp, cl.Logger()); err != nil {
			return err
		}

		if len(resp) == 0 {
			break
		}
		lastSize += len(resp)
		res <- resp
	}
	//response, err := defender.GetDefendersNew(cl)

	return nil
}
