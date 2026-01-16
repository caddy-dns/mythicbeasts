# Mythic Beasts module for Caddy

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with [Mythic Beasts](https://www.mythic-beasts.com).

## Caddy module name

```
dns.providers.mythicbeasts
```

## Building using xcaddy

To build Caddy with this module, use the following xcaddy command:

```
xcaddy build --with github.com/caddyserver/dnsproviders/mythicbeasts
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
                            "key_id": "{env.MYTHIC_KEY_ID}",
                            "secret": "{env.MYTHIC_SECRET}"
            		}
		}
	}
}
```

or with the Caddyfile:

```
{
	# Use the ACME DNS challenge with Mythic Beasts
	acme_dns mythicbeasts {
		# key_id "YOUR_KEY_ID"
		# secret "YOUR_SECRET"
		# Or better, use environment variables:
		key_id {$MYTHIC_KEY_ID}
		secret {$MYTHIC_SECRET}
	}
}

# Replace 'example.com' with a domain you own and want to manage via Mythic Beasts
example.com {
	respond "Hello, Caddy with Mythic Beasts!"
}
```

To run the Caddyfile test, set the environment variables and run:

```
export MYTHIC_KEY_ID="your-key-id"
export MYTHIC_SECRET="your-secret"
./caddy run --config caddyfile
```
