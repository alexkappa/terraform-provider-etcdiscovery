# Terraform Provider for the [coreos/etcd](https://github.com/coreos/etcd) [discovery](https://discovery.etcd.io) service.

Terraform provider for the etcd discovery service.

# Discovery

Use the discovery service to create a new token for your etcd cluster.

```
resource "etcd_discovery" "token" {
  size = 3
}

output "etcd_discovery_url" {
  value = "${etcd_discovery.token.id}"
}
```
