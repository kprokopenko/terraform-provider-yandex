---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "yandex_kubernetes_cluster_iam_member Resource - yandex"
subcategory: ""
description: |-
  Allows creation and management of a single member for a single binding within the IAM policy for an existing Yandex Managed Service for Kubernetes cluster.
  ~> Roles controlled by yandex_kubernetes_cluster_iam_binding should not be assigned using yandex_kubernetes_cluster_iam_member.
---

# yandex_kubernetes_cluster_iam_member (Resource)

Allows creation and management of a single member for a single binding within the IAM policy for an existing Yandex Managed Service for Kubernetes cluster.

~> Roles controlled by `yandex_kubernetes_cluster_iam_binding` should not be assigned using `yandex_kubernetes_cluster_iam_member`.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cluster_id` (String) The [Yandex Managed Service for Kubernetes](https://yandex.cloud/docs/managed-kubernetes/) cluster ID to apply a binding to.
- `member` (String) An array of identities that will be granted the privilege in the `role`. Each entry can have one of the following values:
  * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
  * **serviceAccount:{service_account_id}**: A unique service account ID.
  * **federatedUser:{federated_user_id}**: A unique federated user ID.
  * **federatedUser:{federated_user_id}:**: A unique SAML federation user account ID.
  * **group:{group_id}**: A unique group ID.
  * **system:group:federation:{federation_id}:users**: All users in federation.
  * **system:group:organization:{organization_id}:users**: All users in organization.
  * **system:allAuthenticatedUsers**: All authenticated users.
  * **system:allUsers**: All users, including unauthenticated ones.

~> for more information about system groups, see [Cloud Documentation](https://yandex.cloud/docs/iam/concepts/access-control/system-group).
- `role` (String) The role that should be applied. See [roles catalog](https://yandex.cloud/docs/iam/roles-reference).

### Optional

- `sleep_after` (Number)
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `default` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
