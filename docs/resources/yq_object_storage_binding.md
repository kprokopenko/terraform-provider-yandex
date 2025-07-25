---
subcategory: "Yandex Query"
page_title: "Yandex: yandex_yq_object_storage_binding"
description: |-
  Manages Object Storage binding.
---

# yandex_yq_object_storage_binding (Resource)

Manages Object Storage binding in Yandex Query service. For more information, see [the official documentation](https://yandex.cloud/docs/query/concepts/glossary#Binding).

## Example usage

```terraform
//
// Create a new Object Storage binding.
//

resource "yandex_yq_object_storage_binding" "my_os_binding1" {
  name          = "tf-test-os-binding1"
  description   = "Binding has been created from Terraform"
  connection_id = yandex_yq_object_storage_connection.my_os_connection.id
  compression   = "gzip"
  format        = "json_each_row"

  path_pattern = "my_logs/"
  column {
    name = "ts"
    type = "Timestamp"
  }
  column {
    name = "message"
    type = "Utf8"
  }
}
```

```terraform
//
// Create a new Object Storage binding with Hive partitioning.
//

resource "yandex_yq_object_storage_binding" "my_os_binding2" {
  name          = "tf-test-os-binding2"
  description   = "Binding has been created from Terraform"
  connection_id = yandex_yq_object_storage_connection.my_os_connection.id
  format        = "csv_with_names"
  path_pattern  = "my_logs/"
  format_setting = {
    "file_pattern" = "*.csv"
  }
  column {
    name     = "year"
    type     = "Int32"
    not_null = true
  }
  column {
    name     = "month"
    type     = "Int32"
    not_null = true
  }
  column {
    name     = "day"
    type     = "Int32"
    not_null = true
  }

  partitioned_by = ["year", "month", "day"]
  column {
    name = "ts"
    type = "Timestamp"
  }
  column {
    name = "message"
    type = "Utf8"
  }
}
```

```terraform
//
// Create a new Object Storage binding with extended partitioning.
//

resource "yandex_yq_object_storage_binding" "my_os_binding3" {
  name          = "tf-test-os-binding3"
  description   = "Binding has been created from Terraform"
  connection_id = yandex_yq_object_storage_connection.my_os_connection.id
  compression   = "gzip"
  format        = "json_each_row"

  partitioned_by = [
    "date",
    "severity",
  ]
  path_pattern = "/cold"
  projection = {
    "projection.date.format"     = "/year=%Y/month=%m/day=%d"
    "projection.date.interval"   = "1"
    "projection.date.max"        = "NOW"
    "projection.date.min"        = "2022-12-01"
    "projection.date.type"       = "date"
    "projection.date.unit"       = "DAYS"
    "projection.enabled"         = "true"
    "projection.severity.type"   = "enum"
    "projection.severity.values" = "error,info,fatal"
    "storage.location.template"  = "/$${date}/$${severity}"
  }

  column {
    name     = "timestamp"
    not_null = false
    type     = "String"
  }
  column {
    name     = "message"
    not_null = false
    type     = "String"
  }
  column {
    name     = "date"
    not_null = true
    type     = "Date"
  }
  column {
    name     = "severity"
    not_null = true
    type     = "String"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `connection_id` (String) The connection identifier.
- `format` (String) The data format, e.g. csv_with_names, json_as_string, json_each_row, json_list, parquet, raw, tsv_with_names.
- `name` (String) The resource name.
- `path_pattern` (String) The path pattern within Object Storage's bucket.

### Optional

- `column` (Block List) (see [below for nested schema](#nestedblock--column))
- `compression` (String) The data compression algorithm, e.g. brotli, bzip2, gzip, lz4, xz, zstd.
- `description` (String) The resource description.
- `format_setting` (Map of String) Special format setting.
- `partitioned_by` (List of String) The list of partitioning column names.
- `projection` (Map of String) Projection rules.

### Read-Only

- `id` (String) The resource identifier.

<a id="nestedblock--column"></a>
### Nested Schema for `column`

Required:

- `name` (String) Column name.

Optional:

- `not_null` (Boolean) A column cannot have the NULL data type. Default: `false`.
- `type` (String) Column data type. YQL data types are used.

## Import

The resource can be imported by using their `resource ID`. For getting the resource ID you can use Yandex Cloud [Web Console](https://console.yandex.cloud).

```shell
# terraform import yandex_yq_object_storage_binding.<resource Name> <resource Id>
terraform import yandex_yq_object_storage_binding.my_os_binding ...
```
