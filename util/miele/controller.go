package miele

import (
	"context"
	"sync"

	"github.com/evcc-io/evcc/provider/miele"
	"github.com/evcc-io/evcc/util"
	"golang.org/x/oauth2"
)

var Instance *Controller

type Controller struct {
	mu          sync.Mutex
	log         *util.Logger
	client      *miele.Client
	tokenSource oauth2.TokenSource
	connected   bool
}

func NewController(path, redirectURI string) (*Controller, error) {
	creds, err := miele.LoadCredentials(path)
	if err != nil {
		return nil, err
	}

	if creds.RedirectURI != "" {
		redirectURI = creds.RedirectURI
	}

	client := miele.NewClient(creds.ClientID, creds.ClientSecret, redirectURI)

	c := &Controller{
		log:    util.NewLogger("miele"),
		client: client,
	}

	return c, nil
}

func (c *Controller) GetAuthURL(state string) string {
	// AccessTypeOffline requests a refresh token
	return c.client.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (c *Controller) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := c.client.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	c.SetToken(token)
	return token, nil
}

func (c *Controller) SetToken(token *oauth2.Token) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.tokenSource = c.client.TokenSource(context.Background(), token)
	c.connected = true
}

func (c *Controller) IsConnected() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.connected
}
