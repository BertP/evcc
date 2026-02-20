package whitegood

import (
	"context"
	"fmt"
	"strings"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/util"
	reg "github.com/evcc-io/evcc/util/registry"
)

var registry = reg.New[api.Whitegood]("whitegood")

// Types returns the list of types
func Types() []string {
	return registry.Types()
}

// NewFromConfig creates whitegood from configuration
func NewFromConfig(ctx context.Context, typ string, other map[string]any) (api.Whitegood, error) {
	factory, err := registry.Get(strings.ToLower(typ))
	if err != nil {
		return nil, err
	}

	v, err := factory(ctx, other)
	if err != nil {
		return nil, fmt.Errorf("cannot create whitegood type '%s': %w", util.TypeWithTemplateName(typ, other), err)
	}

	return v, nil
}
