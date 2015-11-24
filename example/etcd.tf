resource "etcdiscovery_token" "example" {
    size = 3
}

output "etcd_discovery_url" {
	value = "${etcdiscovery_token.example.id}"
}
