---
page_title: "hcs_agent_kubernetes_secret Data Source - terraform-provider-hcs"
subcategory: ""
description: |-
  The agent config Kubernetes secret data source provides Consul agents running in Kubernetes the configuration needed to connect to the Consul cluster.
---

# Data Source `hcs_agent_kubernetes_secret`

The agent config Kubernetes secret data source provides Consul agents running in Kubernetes the configuration needed to connect to the Consul cluster.

## Example Usage

```terraform
data "hcs_agent_kubernetes_secret" "default" {
  resource_group_name      = var.resource_group_name
  managed_application_name = var.managed_application_name
}
```

## Schema

### Required

- **managed_application_name** (String) The name of the HCS Azure Managed Application.
- **resource_group_name** (String) The name of the Resource Group in which the HCS Azure Managed Application belongs.

### Optional

- **id** (String) The ID of this resource.
- **timeouts** (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-only

- **secret** (String, Sensitive) The Consul agent configuration in the format of a Kubernetes secret (YAML).

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- **default** (String)

