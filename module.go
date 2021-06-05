package mythicbeasts

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/tombish/mythicbeasts-provider"
)

// Provider wraps the provider implementation as a Caddy module.
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
	repl := caddy.NewReplacer()
	p.Provider.KeyID = repl.ReplaceAll(p.Provider.KeyID, "")
	p.Provider.Secret = repl.ReplaceAll(p.Provider.Secret, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
// mythicbeasts {
//     key_id <string>
//     secret <string>
// }
//
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "key_id":
				if d.NextArg() {
					p.Provider.KeyID = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "secret":
				if d.NextArg() {
					p.Provider.Secret = d.Val()
				}
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
)
