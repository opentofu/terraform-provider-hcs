---
page_title: "hcs_federation_token Data Source - terraform-provider-hcs"
subcategory: ""
description: |-
  The federation token data source can be used during HCS cluster creation to join the cluster to a federation.
---

# Data Source `hcs_federation_token`

The federation token data source can be used during HCS cluster creation to join the cluster to a federation.

## Example Usage

```terraform
data "hcs_federation_token" "default" {
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

- **token** (String, Sensitive) The federation token.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- **default** (String)

