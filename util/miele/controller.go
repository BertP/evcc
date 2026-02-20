package miele

import (
	"context"
	"sync"

	"github.com/evcc-io/evcc/provider/miele"
	"github.com/evcc-io/evcc/server/db/settings"
	"github.com/evcc-io/evcc/util"
	"golang.org/x/oauth2"
)

var Instance *Controller

const mieleTokenKey = "miele_token"

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

	// load token from database
	var token oauth2.Token
	if err := settings.Json(mieleTokenKey, &token); err == nil {
		c.log.DEBUG.Println("loaded Miele token from database")
		c.SetToken(&token)
	}

	return c, nil
}

func (c *Controller) GetAuthURL(redirectURI, state string) string {
	// AccessTypeOffline requests a refresh token
	return c.client.AuthCodeURL(redirectURI, state, oauth2.AccessTypeOffline)
}

func (c *Controller) Exchange(ctx context.Context, redirectURI, code string) (*oauth2.Token, error) {
	token, err := c.client.Exchange(ctx, redirectURI, code)
	if err != nil {
		c.log.ERROR.Printf("failed to exchange code: %v", err)
		return nil, err
	}

	c.log.DEBUG.Println("code exchange successful")
	c.SetToken(token)
	return token, nil
}

func (c *Controller) SetToken(token *oauth2.Token) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.tokenSource = c.client.TokenSource(context.Background(), token)
	c.connected = true

	// persist token
	if err := settings.SetJson(mieleTokenKey, token); err != nil {
		c.log.ERROR.Printf("failed to persist token: %v", err)
	} else {
		c.log.DEBUG.Println("persisted Miele token to database")
	}
}

func (c *Controller) IsConnected() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.connected
}

func (c *Controller) Logout() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.tokenSource = nil
	c.connected = false
	if err := settings.Delete(mieleTokenKey); err != nil {
		c.log.ERROR.Printf("failed to delete token: %v", err)
	} else {
		c.log.DEBUG.Println("deleted Miele token from database")
	}
}
