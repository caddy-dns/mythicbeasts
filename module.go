package mythicbeasts

import (
	"fmt"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/mythicbeasts"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *mythicbeasts.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.mythicbeasts",
		New: func() caddy.Module { return &Provider{new(mythicbeasts.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	if p.Provider == nil {
		p.Provider = new(mythicbeasts.Provider)
	}
	repl := caddy.NewReplacer()
	p.Provider.KeyID = repl.ReplaceAll(p.Provider.KeyID, "")
	p.Provider.Secret = repl.ReplaceAll(p.Provider.Secret, "")
	return nil
}

// UnmarshalCaddyfile() sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	mythicbeasts {
//	    key_id <string>
//	    secret <string>
//	}
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	if p.Provider == nil {
		p.Provider = new(mythicbeasts.Provider)
	}
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "key_id":
				if !d.NextArg() {
					return d.ArgErr()
				}
				p.Provider.KeyID = d.Val()
				if d.NextArg() {
					return d.ArgErr()
				}
			case "secret":
				if !d.NextArg() {
					return d.ArgErr()
				}
				p.Provider.Secret = d.Val()
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized directive '%s'", d.Val())
			}
		}
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
	_ caddy.Validator       = (*Provider)(nil)
)

// Validate implements caddy.Validator.
func (p *Provider) Validate() error {
	if p.Provider == nil {
		return fmt.Errorf("mythicbeasts: provider is not initialized")
	}
	if p.Provider.KeyID == "" || p.Provider.Secret == "" {
		return fmt.Errorf("mythicbeasts: key_id and secret are required")
	}
	return nil
}
