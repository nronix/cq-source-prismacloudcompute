package client

import "github.com/cloudquery/plugin-sdk/v4/schema"

func AccountMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)

	for _, acc := range client.accounts {
		l = append(l, client.WithAccount(acc))
	}
	return l
}
