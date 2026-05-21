package mythicbeasts

import (
	"context"
	"testing"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/mythicbeasts"
)

func TestUnmarshalCaddyfile(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedKeyID string
		expectedSec   string
		expectErr     bool
	}{
		{
			name: "valid block configuration",
			input: `mythicbeasts {
				key_id my_key_id
				secret my_secret
			}`,
			expectedKeyID: "my_key_id",
			expectedSec:   "my_secret",
			expectErr:     false,
		},
		{
			name: "invalid directive",
			input: `mythicbeasts {
				invalid_key my_key_id
			}`,
			expectErr: true,
		},
		{
			name: "missing key_id value",
			input: `mythicbeasts {
				key_id
			}`,
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := new(Provider)
			d := caddyfile.NewTestDispenser(tc.input)
			err := p.UnmarshalCaddyfile(d)
			if (err != nil) != tc.expectErr {
				t.Fatalf("expected error: %v, got: %v", tc.expectErr, err)
			}
			if !tc.expectErr {
				if p.Provider.KeyID != tc.expectedKeyID {
					t.Errorf("expected KeyID %s, got %s", tc.expectedKeyID, p.Provider.KeyID)
				}
				if p.Provider.Secret != tc.expectedSec {
					t.Errorf("expected Secret %s, got %s", tc.expectedSec, p.Provider.Secret)
				}
			}
		})
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name      string
		provider  *Provider
		expectErr bool
	}{
		{
			name: "valid provider",
			provider: &Provider{
				Provider: &mythicbeasts.Provider{
					KeyID:  "my_key_id",
					Secret: "my_secret",
				},
			},
			expectErr: false,
		},
		{
			name: "missing secret",
			provider: &Provider{
				Provider: &mythicbeasts.Provider{
					KeyID: "my_key_id",
				},
			},
			expectErr: true,
		},
		{
			name: "missing key_id",
			provider: &Provider{
				Provider: &mythicbeasts.Provider{
					Secret: "my_secret",
				},
			},
			expectErr: true,
		},
		{
			name:      "nil provider struct",
			provider:  &Provider{},
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.provider.Validate()
			if (err != nil) != tc.expectErr {
				t.Fatalf("expected error: %v, got: %v", tc.expectErr, err)
			}
		})
	}
}

func TestProvision(t *testing.T) {
	t.Setenv("MYTHIC_KEY_ID", "env_key_id")
	t.Setenv("MYTHIC_SECRET", "env_secret")

	p := &Provider{
		Provider: &mythicbeasts.Provider{
			KeyID:  "{env.MYTHIC_KEY_ID}",
			Secret: "{env.MYTHIC_SECRET}",
		},
	}

	ctx, cancel := caddy.NewContext(caddy.Context{Context: context.Background()})
	defer cancel()

	err := p.Provision(ctx)
	if err != nil {
		t.Fatalf("expected no error from Provision, got: %v", err)
	}

	if p.Provider.KeyID != "env_key_id" {
		t.Errorf("expected KeyID env_key_id, got: %s", p.Provider.KeyID)
	}
	if p.Provider.Secret != "env_secret" {
		t.Errorf("expected Secret env_secret, got: %s", p.Provider.Secret)
	}
}
