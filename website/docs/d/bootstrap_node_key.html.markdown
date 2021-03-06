---
layout: "quorum"
page_title: "Quorum: quorum_bootstrap_node_key"
sidebar_current: "docs-quorum-bootstrap-node-key"
description: |-
   Use this data source to obtain various information parsed from an existing node key in hex format.
   
   Node key encodes a private key that defines an identity of a Quorum node in the network. It is primarily used in P2P networking.
---

# quorum_bootstrap_node_key

Use this data source to obtain various information parsed from an existing node key in hex format.

Node key encodes a private key that defines an identity of a Quorum node in the network. It is primarily used in P2P networking.

## Example Usage

```hcl
resource "quorum_bootstrap_node_key" "test" {
}

data "quorum_bootstrap_node_key" "test" {
  node_key_hex = quorum_bootstrap_node_key.test.node_key_hex
}
```

## Argument Reference

- `node_key_hex` - (Required) This hex value encodes the private key of a node

## Attributes Reference

- `hex_node_id` - 64-byte hex value represents node ID which is seen being encoded in the username portion of enode URL. E.g.: `enode:/[<hex node id]@localhost:22000`
- `istanbul_address` - Address representing public key for the newly created node key. This is mainly used to construct the initial validators set encoded in `extradata` field of the genesis file
- `node_id` - 32-byte hex value represents the unique identifier for a node
