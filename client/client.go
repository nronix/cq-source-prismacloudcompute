package client

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
)

type Client struct {
	logger   zerolog.Logger
	Spec     Spec
	Account  TwAccounts
	accounts []TwAccounts
}

func (c *Client) ID() string {
	// TODO: Change to either your plugin name or a unique dynamic identifier
	return fmt.Sprintf("Twistlock:%s", c.Account.Account)
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func New(ctx context.Context, logger zerolog.Logger, s *Spec) (Client, error) {
	logger.Printf("rao2")
	var pAcounts []TwAccounts
	for _, acc := range s.TWISTLOCK {

		logger.Printf("rao3")
		client := &TwClient{}
		client.Config = APIClientConfig{
			ConsoleURL: acc.ENDPOINT,
			Username:   acc.API_KEY,
			Password:   acc.API_SECRET,
			Project:    acc.ACCOUNT,
		}

		if err := client.Initialize(""); err != nil {
			logger.Err(err)
		}
		pAcounts = append(pAcounts, TwAccounts{acc.ACCOUNT, client})
	}

	c := Client{
		logger:   logger,
		Spec:     *s,
		accounts: pAcounts,
	}

	return c, nil
}

func (c *Client) WithAccount(account TwAccounts) *Client {
	newC := *c
	newC.logger = c.logger.With().Str("account", account.Account).Logger()
	newC.Account = account
	return &newC
}
