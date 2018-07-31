# Athenz Plugin
The Athenz plugin uses [Athenz](http://www.athenz.io) to authorize requests to the [yFuzz API](../..).

Requests are authenticated with mutual TLS, so the `todo` setting must be enabled in the [CLI config](../../../cmd/yfuzz-cli#settings).

## Configuration
Plugins in yFuzz are configured in the `config.yaml` file. The following options are available for the `athenz` plugin:
* `url`: URL for your Athenz server.
* `cert-file`: Path to the x509 certificate associated with your Athenz service.
* `key-file`: Path to the private key associated with the certificate.
* `ca-issuer-name`: The name of the CA used by your Athenz instance.
* `action` and `resource`: Athenz principals must be authorized to perform `action` on `resource` to be able to access the yFuzz API.

```yaml
plugins:
  # Information for Athenz (see http://athenz.io/)
  - athenz:
      url: https://your-athenz-server.com/zms/v1
      cert-file: /path/to/athenz/cert
      key-file: /path/to/athenz/key
      ca-issuer-name: Athenz CA Name
      action: access
      resource: yfuzz:yfuzz
```

## See Also
* The mutual TLS option in the [yFuzz CLI](../../../cmd/yfuzz-cli#settings).
* The [list of yFuzz plugins](../../../docs/plugins.md).