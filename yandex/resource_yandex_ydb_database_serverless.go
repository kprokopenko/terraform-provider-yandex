package yandex

import (
	"fmt"
	"time"

	"github.com/c2h5oh/datasize"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/ydb/v1"
	"github.com/yandex-cloud/terraform-provider-yandex/common"
	"google.golang.org/genproto/protobuf/field_mask"
)

const yandexYDBServerlessDefaultTimeout = 10 * time.Minute

func resourceYandexYDBDatabaseServerless() *schema.Resource {
	return &schema.Resource{
		Description: "Yandex Database (serverless) resource. For more information, see [the official documentation](https://yandex.cloud/docs/ydb/concepts/serverless_and_dedicated).",

		Create: resourceYandexYDBDatabaseServerlessCreate,
		Read:   resourceYandexYDBDatabaseServerlessRead,
		Update: resourceYandexYDBDatabaseServerlessUpdate,
		Delete: performYandexYDBDatabaseDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(yandexYDBServerlessDefaultTimeout),
		},

		SchemaVersion: 0,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Description:  common.ResourceDescriptions["name"],
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			"location_id": {
				Type:        schema.TypeString,
				Description: "Location ID for the Yandex Database serverless cluster.",
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},

			"folder_id": {
				Type:         schema.TypeString,
				Description:  common.ResourceDescriptions["folder_id"],
				Computed:     true,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			"description": {
				Type:        schema.TypeString,
				Description: common.ResourceDescriptions["description"],
				Optional:    true,
			},

			"labels": {
				Type:        schema.TypeMap,
				Description: common.ResourceDescriptions["labels"],
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
			},

			"document_api_endpoint": {
				Type:        schema.TypeString,
				Description: "Document API endpoint of the Yandex Database serverless cluster.",
				Computed:    true,
			},

			"ydb_full_endpoint": {
				Type:        schema.TypeString,
				Description: "Full endpoint of the Yandex Database serverless cluster.",
				Computed:    true,
			},

			"ydb_api_endpoint": {
				Type:        schema.TypeString,
				Description: "API endpoint of the Yandex Database serverless cluster. Useful for SDK configuration.",
				Computed:    true,
			},

			"database_path": {
				Type:        schema.TypeString,
				Description: "Full database path of the Yandex Database serverless cluster. Useful for SDK configuration.",
				Computed:    true,
			},

			"tls_enabled": {
				Type:        schema.TypeBool,
				Description: "Whether TLS is enabled for the Yandex Database serverless cluster. Useful for SDK configuration.",
				Computed:    true,
			},

			"created_at": {
				Type:        schema.TypeString,
				Description: common.ResourceDescriptions["created_at"],
				Computed:    true,
			},

			"status": {
				Type:        schema.TypeString,
				Description: "Status of the Yandex Database serverless cluster.",
				Computed:    true,
			},

			"deletion_protection": {
				Type:        schema.TypeBool,
				Description: common.ResourceDescriptions["deletion_protection"],
				Optional:    true,
				Default:     false,
			},
			"sleep_after": {
				Type:        schema.TypeInt,
				Description: "",
				Optional:    true,
				Default:     0,
				DiffSuppressFunc: func(_, _, _ string, d *schema.ResourceData) bool {
					return d.Id() != "" // suppress for updates
				},
			},
			"serverless_database": {
				Type:        schema.TypeSet,
				Description: "",
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"throttling_rcu_limit": {
							Type:        schema.TypeInt,
							Description: "",
							Optional:    true,
							Computed:    true,
						},
						"storage_size_limit": {
							Type:        schema.TypeInt,
							Description: "",
							Optional:    true,
							Computed:    true,
						},
						"enable_throttling_rcu_limit": {
							Type:        schema.TypeBool,
							Description: "",
							Optional:    true,
							Computed:    true,
						},
						"provisioned_rcu_limit": {
							Type:        schema.TypeInt,
							Description: "",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func flattenYDBServerlessDatabaseSettings(d *schema.ResourceData, database *ydb.Database_ServerlessDatabase) {
	output := make([]interface{}, 0, 1)
	serverlessDatabase := make(map[string]interface{})
	serverlessDatabase["throttling_rcu_limit"] = database.ServerlessDatabase.ThrottlingRcuLimit
	serverlessDatabase["storage_size_limit"] = database.ServerlessDatabase.StorageSizeLimit / int64(datasize.GB)
	serverlessDatabase["enable_throttling_rcu_limit"] = database.ServerlessDatabase.EnableThrottlingRcuLimit
	serverlessDatabase["provisioned_rcu_limit"] = database.ServerlessDatabase.ProvisionedRcuLimit

	output = append(output, serverlessDatabase)
	_ = d.Set("serverless_database", output)
}

func expandYDBServerlessDatabaseSettings(d *schema.ResourceData) *ydb.ServerlessDatabase {
	v, ok := d.GetOk("serverless_database")
	if !ok {
		return nil
	}
	serverlessDatabase := &ydb.ServerlessDatabase{}
	ttlSet := v.(*schema.Set)
	for _, l := range ttlSet.List() {
		m := l.(map[string]interface{})
		serverlessDatabase.ThrottlingRcuLimit = int64(m["throttling_rcu_limit"].(int))
		serverlessDatabase.StorageSizeLimit = int64(datasize.GB) * int64(m["storage_size_limit"].(int))
		serverlessDatabase.EnableThrottlingRcuLimit = m["enable_throttling_rcu_limit"].(bool)
		serverlessDatabase.ProvisionedRcuLimit = int64(m["provisioned_rcu_limit"].(int))
	}
	return serverlessDatabase
}

func resourceYandexYDBDatabaseServerlessCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	folderID, err := getFolderID(d, config)
	if err != nil {
		return fmt.Errorf("Error getting folder ID while creating database: %s", err)
	}
	labels, err := expandLabels(d.Get("labels"))
	if err != nil {
		return fmt.Errorf("Error expanding labels while creating database: %s", err)
	}
	req := ydb.CreateDatabaseRequest{
		FolderId:    folderID,
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		DatabaseType: &ydb.CreateDatabaseRequest_ServerlessDatabase{
			ServerlessDatabase: expandYDBServerlessDatabaseSettings(d),
		},
		LocationId:         d.Get("location_id").(string),
		Labels:             labels,
		DeletionProtection: d.Get("deletion_protection").(bool),
	}

	if err := performYandexYDBDatabaseCreate(d, config, &req); err != nil {
		return err
	}

	return resourceYandexYDBDatabaseServerlessRead(d, meta)
}

func resourceYandexYDBDatabaseServerlessUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	req := ydb.UpdateDatabaseRequest{
		DatabaseId: d.Id(),
		UpdateMask: &field_mask.FieldMask{},
	}
	if d.HasChange("serverless_database") {
		req.DatabaseType = &ydb.UpdateDatabaseRequest_ServerlessDatabase{
			ServerlessDatabase: expandYDBServerlessDatabaseSettings(d),
		}
		req.UpdateMask.Paths = append(req.UpdateMask.Paths, "serverless_database")
	}

	if err := performYandexYDBDatabaseUpdate(d, config, &req); err != nil {
		return err
	}

	return resourceYandexYDBDatabaseServerlessRead(d, meta)
}

func resourceYandexYDBDatabaseServerlessRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	database, err := performYandexYDBDatabaseRead(d, config)
	if err != nil {
		return err
	}

	return flattenYandexYDBDatabaseServerless(d, database)
}

func flattenYandexYDBDatabaseServerless(d *schema.ResourceData, database *ydb.Database) error {
	if database == nil {
		// NOTE(shmel1k@): database existed before but was removed outside of terraform.
		d.SetId("")
		return nil
	}

	switch t := database.DatabaseType.(type) {
	case *ydb.Database_ServerlessDatabase: // we actually expect it
		flattenYDBServerlessDatabaseSettings(d, t)
	case *ydb.Database_DedicatedDatabase:
		return fmt.Errorf("expect serverless database, got dedicated")
	case *ydb.Database_RegionalDatabase:
		return fmt.Errorf("expect serverless database, got regional")
	case *ydb.Database_ZonalDatabase:
		return fmt.Errorf("expect serverless database, got zonal")
	default:
		return fmt.Errorf("unknown database type")
	}

	d.Set("document_api_endpoint", database.DocumentApiEndpoint)

	return flattenYandexYDBDatabase(d, database)
}
