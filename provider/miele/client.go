package miele

import (
	"context"
	"net/http"

	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/request"
	"golang.org/x/oauth2"
)

// Endpoints
const (
	AuthURL  = "https://auth.domestic.miele-iot.com/partner/realms/mcs/protocol/openid-connect/auth"
	TokenURL = "https://auth.domestic.miele-iot.com/partner/realms/mcs/protocol/openid-connect/token"
	ApiURL   = "https://api.mcs3.miele.com/v1"
)

type Client struct {
	*request.Helper
	conf *oauth2.Config
}

func NewClient(clientID, clientSecret, redirectURI string) *Client {
	return &Client{
		Helper: request.NewHelper(util.NewLogger("miele")),
		conf: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURI,
			Endpoint: oauth2.Endpoint{
				AuthURL:  AuthURL,
				TokenURL: TokenURL,
			},
			Scopes: []string{"openid", "mcs_thirdparty_read", "mcs_thirdparty_write"},
		},
	}
}

func (c *Client) SetRedirectURL(url string) {
	c.conf.RedirectURL = url
}

func (c *Client) config(redirectURI string) *oauth2.Config {
	conf := *c.conf
	if redirectURI != "" {
		conf.RedirectURL = redirectURI
	}
	return &conf
}

func (c *Client) Exchange(ctx context.Context, redirectURI, code string) (*oauth2.Token, error) {
	return c.config(redirectURI).Exchange(ctx, code)
}

func (c *Client) TokenSource(ctx context.Context, token *oauth2.Token) oauth2.TokenSource {
	return c.conf.TokenSource(ctx, token)
}

// Client returns an authenticated http client
func (c *Client) Client(ctx context.Context, token *oauth2.Token) *http.Client {
	return c.conf.Client(ctx, token)
}

func (c *Client) AuthCodeURL(redirectURI, state string, opts ...oauth2.AuthCodeOption) string {
	return c.config(redirectURI).AuthCodeURL(state, opts...)
}
