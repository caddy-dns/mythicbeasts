# Mythic Beasts module for Caddy

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with [Mythic Beasts](https://www.mythic-beasts.com).

## Caddy module name

```
dns.providers.mythicbeasts
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
    "module": "acme",
    "challenges": {
        "dns": {
            "provider": {
                            "name": "mythicbeasts",
                            "key_id": "{env.MYTHICBEASTS_KEYID}",
                            "secret": "{env.MYTHICBEASTS_SECRET}"
            		}
		}
	}
}
```

or with the Caddyfile:

```
tls {
	dns mythicbeasts {
            key_id {$MYTHICBEASTS_KEYID}
            secret {$MYTHICBEASTS_SECRET}
    }
}
```
