# MTLS Plugin
The mtls (mutual TLS) plugin uses mutual tls and self-signed certificates to authenticate requests to the [yFuzz API](../..).

Requests are authorized based on if the public key associated with the user's certificate is in a whitelist.

## Adding a New User
First, generate a user x509 certificate:
```bash
$ openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -nodes
$ openssl rsa -in key.pem -pubout
```

Then whitelist that certificate by adding the public key to the yFuzz configuration file.

## Configuration
Plugins in yFuzz are configured in the `config.yaml` file. The following options are available for the `mtls` plugin:

```yaml
plugins:
  - mtls:
      authorized-keys:
        - |
          -----BEGIN PUBLIC KEY-----
          Public key goes here.
          -----END PUBLIC KEY-----
        - |
          -----BEGIN PUBLIC KEY-----
          A second public key.
          -----END PUBLIC KEY-----
```

* `authorized-keys`: List of public keys corresponding to users authorized to access yFuzz.

## See Also
* Mutual TLS configuration in the [yFuzz CLI](../../../cmd/yfuzz-cli#settings).
* The [list of yFuzz plugins](../../../docs/plugins.md).