---
subcategory: "Managed Service for Trino"
page_title: "Yandex: yandex_trino_catalog"
description: |-
  Manages Trino catalog within Yandex Cloud.
---

# yandex_trino_catalog (Resource)

Catalog for Manage Trino cluster.

## Example usage

```terraform
//
// Create a new Trino catalog
//

resource "yandex_trino_catalog" "catalog" {
  name        = "name"
  description = "descriptionr"
  cluster_id  = yandex_trino_cluster.trino.id
  postgresql = {
    connection_manager = {
      connection_id = "<connection_id>"
      database      = "database-name"
      connection_properties = {
        "targetServerType" = "primary"
      }
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cluster_id` (String) ID of the Trino cluster. Provided by the client when the Catalog is created.
- `name` (String) The resource name.

### Optional

- `clickhouse` (Attributes) Configuration for Clickhouse connector. (see [below for nested schema](#nestedatt--clickhouse))
- `delta_lake` (Attributes) Configuration for DeltaLake connector. (see [below for nested schema](#nestedatt--delta_lake))
- `description` (String) The resource description.
- `hive` (Attributes) Configuration for Hive connector. (see [below for nested schema](#nestedatt--hive))
- `iceberg` (Attributes) Configuration for Iceberg connector. (see [below for nested schema](#nestedatt--iceberg))
- `labels` (Map of String) A set of key/value label pairs which assigned to resource.
- `oracle` (Attributes) Configuration for Oracle connector. (see [below for nested schema](#nestedatt--oracle))
- `postgresql` (Attributes) Configuration for Postgresql connector. (see [below for nested schema](#nestedatt--postgresql))
- `sqlserver` (Attributes) Configuration for SQLServer connector. (see [below for nested schema](#nestedatt--sqlserver))
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `tpcds` (Attributes) Configuration for TPCDS connector. (see [below for nested schema](#nestedatt--tpcds))
- `tpch` (Attributes) Configuration for TPCH connector. (see [below for nested schema](#nestedatt--tpch))

### Read-Only

- `id` (String) The resource identifier.

<a id="nestedatt--clickhouse"></a>
### Nested Schema for `clickhouse`

Optional:

- `additional_properties` (Map of String) Additional properties.
- `connection_manager` (Attributes) Configuration for connection manager connection. (see [below for nested schema](#nestedatt--clickhouse--connection_manager))
- `on_premise` (Attributes) Configuration for on-premise connection. (see [below for nested schema](#nestedatt--clickhouse--on_premise))

<a id="nestedatt--clickhouse--connection_manager"></a>
### Nested Schema for `clickhouse.connection_manager`

Required:

- `connection_id` (String) Connection ID.
- `database` (String) Database.

Optional:

- `connection_properties` (Map of String) Additional connection properties.


<a id="nestedatt--clickhouse--on_premise"></a>
### Nested Schema for `clickhouse.on_premise`

Required:

- `connection_url` (String) Connection to the clickhouse.
- `password` (String) Password of the clickhouse user.
- `user_name` (String) Name of the clickhouse user.



<a id="nestedatt--delta_lake"></a>
### Nested Schema for `delta_lake`

Required:

- `file_system` (Attributes) File system configuration. (see [below for nested schema](#nestedatt--delta_lake--file_system))
- `metastore` (Attributes) Metastore configuration. (see [below for nested schema](#nestedatt--delta_lake--metastore))

Optional:

- `additional_properties` (Map of String) Additional properties.

<a id="nestedatt--delta_lake--file_system"></a>
### Nested Schema for `delta_lake.file_system`

Optional:

- `external_s3` (Attributes) Describes External S3 compatible file system. (see [below for nested schema](#nestedatt--delta_lake--file_system--external_s3))
- `s3` (Attributes) Describes YandexCloud native S3 file system. (see [below for nested schema](#nestedatt--delta_lake--file_system--s3))

<a id="nestedatt--delta_lake--file_system--external_s3"></a>
### Nested Schema for `delta_lake.file_system.external_s3`

Required:

- `aws_access_key` (String, Sensitive) AWS access key ID for S3 authentication.
- `aws_endpoint` (String) AWS S3 compatible endpoint URL.
- `aws_region` (String) AWS region for S3 storage.
- `aws_secret_key` (String, Sensitive) AWS secret access key for S3 authentication.


<a id="nestedatt--delta_lake--file_system--s3"></a>
### Nested Schema for `delta_lake.file_system.s3`



<a id="nestedatt--delta_lake--metastore"></a>
### Nested Schema for `delta_lake.metastore`

Required:

- `uri` (String) The resource description.



<a id="nestedatt--hive"></a>
### Nested Schema for `hive`

Required:

- `file_system` (Attributes) File system configuration. (see [below for nested schema](#nestedatt--hive--file_system))
- `metastore` (Attributes) Metastore configuration. (see [below for nested schema](#nestedatt--hive--metastore))

Optional:

- `additional_properties` (Map of String) Additional properties.

<a id="nestedatt--hive--file_system"></a>
### Nested Schema for `hive.file_system`

Optional:

- `external_s3` (Attributes) Describes External S3 compatible file system. (see [below for nested schema](#nestedatt--hive--file_system--external_s3))
- `s3` (Attributes) Describes YandexCloud native S3 file system. (see [below for nested schema](#nestedatt--hive--file_system--s3))

<a id="nestedatt--hive--file_system--external_s3"></a>
### Nested Schema for `hive.file_system.external_s3`

Required:

- `aws_access_key` (String, Sensitive) AWS access key ID for S3 authentication.
- `aws_endpoint` (String) AWS S3 compatible endpoint URL.
- `aws_region` (String) AWS region for S3 storage.
- `aws_secret_key` (String, Sensitive) AWS secret access key for S3 authentication.


<a id="nestedatt--hive--file_system--s3"></a>
### Nested Schema for `hive.file_system.s3`



<a id="nestedatt--hive--metastore"></a>
### Nested Schema for `hive.metastore`

Required:

- `uri` (String) The resource description.



<a id="nestedatt--iceberg"></a>
### Nested Schema for `iceberg`

Required:

- `file_system` (Attributes) File system configuration. (see [below for nested schema](#nestedatt--iceberg--file_system))
- `metastore` (Attributes) Metastore configuration. (see [below for nested schema](#nestedatt--iceberg--metastore))

Optional:

- `additional_properties` (Map of String) Additional properties.

<a id="nestedatt--iceberg--file_system"></a>
### Nested Schema for `iceberg.file_system`

Optional:

- `external_s3` (Attributes) Describes External S3 compatible file system. (see [below for nested schema](#nestedatt--iceberg--file_system--external_s3))
- `s3` (Attributes) Describes YandexCloud native S3 file system. (see [below for nested schema](#nestedatt--iceberg--file_system--s3))

<a id="nestedatt--iceberg--file_system--external_s3"></a>
### Nested Schema for `iceberg.file_system.external_s3`

Required:

- `aws_access_key` (String, Sensitive) AWS access key ID for S3 authentication.
- `aws_endpoint` (String) AWS S3 compatible endpoint URL.
- `aws_region` (String) AWS region for S3 storage.
- `aws_secret_key` (String, Sensitive) AWS secret access key for S3 authentication.


<a id="nestedatt--iceberg--file_system--s3"></a>
### Nested Schema for `iceberg.file_system.s3`



<a id="nestedatt--iceberg--metastore"></a>
### Nested Schema for `iceberg.metastore`

Required:

- `uri` (String) The resource description.



<a id="nestedatt--oracle"></a>
### Nested Schema for `oracle`

Optional:

- `additional_properties` (Map of String) Additional properties.
- `on_premise` (Attributes) Configuration for on-premise connection. (see [below for nested schema](#nestedatt--oracle--on_premise))

<a id="nestedatt--oracle--on_premise"></a>
### Nested Schema for `oracle.on_premise`

Required:

- `connection_url` (String) Connection to the clickhouse.
- `password` (String) Password of the clickhouse user.
- `user_name` (String) Name of the clickhouse user.



<a id="nestedatt--postgresql"></a>
### Nested Schema for `postgresql`

Optional:

- `additional_properties` (Map of String) Additional properties.
- `connection_manager` (Attributes) Configuration for connection manager connection. (see [below for nested schema](#nestedatt--postgresql--connection_manager))
- `on_premise` (Attributes) Configuration for on-premise connection. (see [below for nested schema](#nestedatt--postgresql--on_premise))

<a id="nestedatt--postgresql--connection_manager"></a>
### Nested Schema for `postgresql.connection_manager`

Required:

- `connection_id` (String) Connection ID.
- `database` (String) Database.

Optional:

- `connection_properties` (Map of String) Additional connection properties.


<a id="nestedatt--postgresql--on_premise"></a>
### Nested Schema for `postgresql.on_premise`

Required:

- `connection_url` (String) Connection to the clickhouse.
- `password` (String) Password of the clickhouse user.
- `user_name` (String) Name of the clickhouse user.



<a id="nestedatt--sqlserver"></a>
### Nested Schema for `sqlserver`

Optional:

- `additional_properties` (Map of String) Additional properties.
- `on_premise` (Attributes) Configuration for on-premise connection. (see [below for nested schema](#nestedatt--sqlserver--on_premise))

<a id="nestedatt--sqlserver--on_premise"></a>
### Nested Schema for `sqlserver.on_premise`

Required:

- `connection_url` (String) Connection to the clickhouse.
- `password` (String) Password of the clickhouse user.
- `user_name` (String) Name of the clickhouse user.



<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
- `delete` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
- `update` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).


<a id="nestedatt--tpcds"></a>
### Nested Schema for `tpcds`

Optional:

- `additional_properties` (Map of String) Additional properties.


<a id="nestedatt--tpch"></a>
### Nested Schema for `tpch`

Optional:

- `additional_properties` (Map of String) Additional properties.

## Import

The resource can be imported by using their `resource ID`. For getting the resource ID you can use Yandex Cloud [Web Console](https://console.yandex.cloud) or [YC CLI](https://yandex.cloud/docs/cli/quickstart).

```bash
# terraform import yandex_trino_catalog.<resource Name> <cluster_id>:<resource_id>
terraform import yandex_trino_catalog.my_catalog cluster_id:catalog_id
```

