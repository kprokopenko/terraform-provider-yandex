package yandex

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/ydb/v1"
	"github.com/yandex-cloud/terraform-provider-yandex/common"
	"google.golang.org/genproto/protobuf/field_mask"
)

const yandexYDBDedicatedDefaultTimeout = 30 * time.Minute

func resourceYandexYDBDatabaseDedicated() *schema.Resource {
	return &schema.Resource{
		Description: "Yandex Database (dedicated) resource. For more information, see [the official documentation](https://yandex.cloud/docs/ydb/concepts/serverless_and_dedicated).",

		Create: resourceYandexYDBDatabaseDedicatedCreate,
		Read:   resourceYandexYDBDatabaseDedicatedRead,
		Update: resourceYandexYDBDatabaseDedicatedUpdate,
		Delete: performYandexYDBDatabaseDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(yandexYDBDedicatedDefaultTimeout),
		},

		SchemaVersion: 0,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Description:  common.ResourceDescriptions["name"],
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			"network_id": {
				Type:         schema.TypeString,
				Description:  common.ResourceDescriptions["network_id"],
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			"subnet_ids": {
				Type:        schema.TypeSet,
				Description: common.ResourceDescriptions["subnet_ids"],
				Required:    true,
				MinItems:    1,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
			},

			"security_group_ids": {
				Type:        schema.TypeSet,
				Description: common.ResourceDescriptions["security_group_ids"],
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
			},

			"resource_preset_id": {
				Type:         schema.TypeString,
				Description:  "The Yandex Database cluster preset. Available presets can be obtained via `yc ydb resource-preset list` command.",
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			"scale_policy": {
				Type:        schema.TypeList,
				Description: "Scaling policy for the Yandex Database cluster.\n\n~> Currently, only `fixed_scale` is supported.\n",
				MaxItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fixed_scale": {
							Type:        schema.TypeList,
							Description: "Fixed scaling policy for the Yandex Database cluster.",
							Required:    true,
							MaxItems:    1,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"size": {
										Type:         schema.TypeInt,
										Description:  "Number of instances for the Yandex Database cluster.",
										Required:     true,
										ValidateFunc: validation.IntAtLeast(1),
									},
								},
							},
						},
					},
				},
			},

			"storage_config": {
				Type:        schema.TypeList,
				Description: "A list of storage configuration options for the Yandex Database cluster.",
				Required:    true,
				MinItems:    1,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"storage_type_id": {
							Type:         schema.TypeString,
							Description:  "Storage type ID for the Yandex Database cluster. Available presets can be obtained via `yc ydb storage-type list` command.",
							Required:     true,
							ValidateFunc: validation.NoZeroValues,
						},
						"group_count": {
							Type:         schema.TypeInt,
							Description:  "Amount of storage groups of selected type for the Yandex Database cluster.",
							Required:     true,
							ValidateFunc: validation.IntAtLeast(1),
						},
					},
				},
			},

			"location": {
				Type:        schema.TypeList,
				Description: "Location for the Yandex Database cluster.",
				MaxItems:    1,
				Optional:    true,
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region": {
							Type:        schema.TypeList,
							Description: "Region for the Yandex Database cluster.",
							Optional:    true,
							MaxItems:    1,
							ForceNew:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:         schema.TypeString,
										Description:  "Region ID for the Yandex Database cluster.",
										Required:     true,
										ForceNew:     true,
										ValidateFunc: validation.NoZeroValues,
									},
								},
							},
						},
					},
				},
			},

			"location_id": {
				Type:        schema.TypeString,
				Description: "Location ID for the Yandex Database cluster.",
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},

			"assign_public_ips": {
				Type:        schema.TypeBool,
				Description: "Whether public IP addresses should be assigned to the Yandex Database cluster.",
				Optional:    true,
				Default:     false,
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

			"ydb_full_endpoint": {
				Type:        schema.TypeString,
				Description: "Full endpoint of the Yandex Database cluster.",
				Computed:    true,
			},

			"ydb_api_endpoint": {
				Type:        schema.TypeString,
				Description: "API endpoint of the Yandex Database cluster. Useful for SDK configuration.",
				Computed:    true,
			},

			"database_path": {
				Type:        schema.TypeString,
				Description: "Full database path of the Yandex Database cluster. Useful for SDK configuration.",
				Computed:    true,
			},

			"tls_enabled": {
				Type:        schema.TypeBool,
				Description: "Whether TLS is enabled for the Yandex Database cluster. Useful for SDK configuration.",
				Computed:    true,
			},

			"status": {
				Type:        schema.TypeString,
				Description: "Status of the Yandex Database cluster.",
				Computed:    true,
			},

			"created_at": {
				Type:        schema.TypeString,
				Description: common.ResourceDescriptions["created_at"],
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
		},
	}
}

func resourceYandexYDBDatabaseDedicatedCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	folderID, err := getFolderID(d, config)
	if err != nil {
		return fmt.Errorf("Error getting folder ID while creating database: %s", err)
	}

	labels, err := expandLabels(d.Get("labels"))
	if err != nil {
		return fmt.Errorf("Error expanding labels while creating database: %s", err)
	}

	subnetIDs := convertStringSet(d.Get("subnet_ids").(*schema.Set))
	if len(subnetIDs) == 0 {
		return fmt.Errorf("Error expanding subnet IDs while creating database: %s", err)
	}
	for _, subnetID := range subnetIDs {
		if len(subnetID) == 0 {
			return fmt.Errorf("Error checking subnet IDs while creating database: %s", err)
		}
	}

	securityGroupIDs := convertStringSet(d.Get("security_group_ids").(*schema.Set))
	for _, securityGroupID := range securityGroupIDs {
		if len(securityGroupID) == 0 {
			return fmt.Errorf("Error checking security_group IDs while creating database: %s", err)
		}
	}

	storageConfig, err := expandYDBStorageConfigSpec(d)
	if err != nil {
		return fmt.Errorf("Error expanding storage configuration while creating database: %s", err)
	}

	scalePolicy, err := expandYDBScalePolicySpec(d)
	if err != nil {
		return fmt.Errorf("Error expanding scale policy while creating database: %s", err)
	}

	dbType, err := expandYDBLocationSpec(d)
	if err != nil {
		return fmt.Errorf("Error expanding database type while creating database: %s", err)
	}
	req := ydb.CreateDatabaseRequest{
		FolderId:           folderID,
		Name:               d.Get("name").(string),
		DatabaseType:       dbType,
		Description:        d.Get("description").(string),
		ResourcePresetId:   d.Get("resource_preset_id").(string),
		StorageConfig:      storageConfig,
		ScalePolicy:        scalePolicy,
		NetworkId:          d.Get("network_id").(string),
		SubnetIds:          subnetIDs,
		SecurityGroupIds:   securityGroupIDs,
		AssignPublicIps:    d.Get("assign_public_ips").(bool),
		LocationId:         d.Get("location_id").(string),
		Labels:             labels,
		DeletionProtection: d.Get("deletion_protection").(bool),
	}

	if err := performYandexYDBDatabaseCreate(d, config, &req); err != nil {
		return err
	}

	return resourceYandexYDBDatabaseDedicatedRead(d, meta)
}

func performYandexYDBDatabaseCreate(d *schema.ResourceData, config *Config, req *ydb.CreateDatabaseRequest) error {
	ctx, cancel := context.WithTimeout(config.Context(), d.Timeout(schema.TimeoutCreate))
	defer cancel()

	op, err := config.sdk.WrapOperation(config.sdk.YDB().Database().Create(ctx, req))
	if err != nil {
		return fmt.Errorf("Error while requesting API to create database: %s", err)
	}

	protoMetadata, err := op.Metadata()
	if err != nil {
		return fmt.Errorf("Error while get database create operation metadata: %s", err)
	}

	md, ok := protoMetadata.(*ydb.CreateDatabaseMetadata)
	if !ok {
		return fmt.Errorf("could not get database ID from create operation metadata")
	}

	d.SetId(md.DatabaseId)

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Error while waiting operation to create database: %s", err)
	}

	if _, err := op.Response(); err != nil {
		return fmt.Errorf("Database creation failed: %s", err)
	}

	if slp := d.Get("sleep_after").(int); slp > 0 {
		log.Printf("[INFO] Waiting additional duration: %d", slp)
		time.Sleep(time.Duration(slp) * time.Second)
	}

	return nil
}

func resourceYandexYDBDatabaseDedicatedUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	req := ydb.UpdateDatabaseRequest{
		DatabaseId: d.Id(),
		UpdateMask: &field_mask.FieldMask{},
	}

	if d.HasChange("assign_public_ips") {
		req.AssignPublicIps = d.Get("assign_public_ips").(bool)
		req.UpdateMask.Paths = append(req.UpdateMask.Paths, "assign_public_ips")
	}

	if d.HasChange("resource_preset_id") {
		req.ResourcePresetId = d.Get("resource_preset_id").(string)
		req.UpdateMask.Paths = append(req.UpdateMask.Paths, "resource_preset_id")
	}

	if d.HasChange("storage_config") {
		storageConfig, err := expandYDBStorageConfigSpec(d)
		if err != nil {
			return fmt.Errorf("Error expanding storage configuration while updating database: %s", err)
		}
		req.StorageConfig = storageConfig
		req.UpdateMask.Paths = append(req.UpdateMask.Paths, "storage_config")
	}

	if d.HasChange("scale_policy") {
		scalePolicy, err := expandYDBScalePolicySpec(d)
		if err != nil {
			return fmt.Errorf("Error expanding scale policy while updating database: %s", err)
		}
		req.ScalePolicy = scalePolicy
		req.UpdateMask.Paths = append(req.UpdateMask.Paths, "scale_policy")
	}

	if d.HasChange("network_id") {
		networkId, err := changeYDBnetworkIdSpec(d)
		if err != nil {
			return fmt.Errorf("Error changing network_id while updating database: %s", err)
		}
		req.NetworkId = networkId
		req.UpdateMask.Paths = append(req.UpdateMask.Paths, "network_id")
	}

	if d.HasChange("subnet_ids") {
		subnetIds, err := changeYDBsubnetIdsSpec(d)
		if err != nil {
			return fmt.Errorf("Error changing subnet_ids while updating database: %s", err)
		}
		req.SubnetIds = subnetIds
		req.UpdateMask.Paths = append(req.UpdateMask.Paths, "subnet_ids")
	}

	if d.HasChange("security_group_ids") {
		securityGroupIds, err := changeYDBsecurityGroupIdsSpec(d)
		if err != nil {
			return fmt.Errorf("Error changing security_group_ids while updating database: %s", err)
		}
		req.SecurityGroupIds = securityGroupIds
		req.UpdateMask.Paths = append(req.UpdateMask.Paths, "security_group_ids")
	}

	if err := performYandexYDBDatabaseUpdate(d, config, &req); err != nil {
		return err
	}

	return resourceYandexYDBDatabaseDedicatedRead(d, meta)
}

func performYandexYDBDatabaseUpdate(d *schema.ResourceData, config *Config, req *ydb.UpdateDatabaseRequest) error {
	d.Partial(true)
	// common parameters
	if d.HasChange("labels") {
		labelsProp, err := expandLabels(d.Get("labels"))
		if err != nil {
			return err
		}

		req.Labels = labelsProp
		req.UpdateMask.Paths = append(req.UpdateMask.Paths, "labels")
	}

	if d.HasChange("name") {
		req.Name = d.Get("name").(string)
		req.UpdateMask.Paths = append(req.UpdateMask.Paths, "name")
	}

	if d.HasChange("description") {
		req.Description = d.Get("description").(string)
		req.UpdateMask.Paths = append(req.UpdateMask.Paths, "description")
	}

	if d.HasChange("deletion_protection") {
		req.DeletionProtection = d.Get("deletion_protection").(bool)
		req.UpdateMask.Paths = append(req.UpdateMask.Paths, "deletion_protection")
	}

	ctx, cancel := context.WithTimeout(config.Context(), d.Timeout(schema.TimeoutUpdate))
	defer cancel()

	op, err := config.sdk.WrapOperation(config.sdk.YDB().Database().Update(ctx, req))
	if err != nil {
		return fmt.Errorf("Error while requesting API to update database: %s", err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Error updating database %q: %s", d.Id(), err)
	}

	d.Partial(false)

	return nil
}

func performYandexYDBDatabaseDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	ctx, cancel := config.ContextWithTimeout(d.Timeout(schema.TimeoutDelete))
	defer cancel()

	op, err := config.sdk.YDB().Database().Delete(ctx, &ydb.DeleteDatabaseRequest{DatabaseId: d.Id()})
	err = waitOperation(ctx, config, op, err)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("YDB Database %q", d.Id()))
	}

	return nil
}

func resourceYandexYDBDatabaseDedicatedRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	database, err := performYandexYDBDatabaseRead(d, config)
	if err != nil {
		return err
	}

	return flattenYandexYDBDatabaseDedicated(d, database)
}

func performYandexYDBDatabaseRead(d *schema.ResourceData, config *Config) (*ydb.Database, error) {
	ctx, cancel := config.ContextWithTimeout(d.Timeout(schema.TimeoutRead))
	defer cancel()

	database, err := config.sdk.YDB().Database().Get(ctx, &ydb.GetDatabaseRequest{
		DatabaseId: d.Id(),
	})
	if err != nil {
		return nil, handleNotFoundError(err, d, fmt.Sprintf("YDB Database %q", d.Get("name").(string)))
	}

	return database, nil
}

func flattenYandexYDBDatabaseDedicated(d *schema.ResourceData, database *ydb.Database) error {
	if database == nil {
		// NOTE(shmel1k@): database existed before but was removed outside of terraform.
		d.SetId("")
		return nil
	}

	switch database.DatabaseType.(type) {
	case *ydb.Database_RegionalDatabase,
		*ydb.Database_ZonalDatabase,
		*ydb.Database_DedicatedDatabase: // we actually expect it
	case *ydb.Database_ServerlessDatabase:
		return fmt.Errorf("expect dedicated database, got serverless")
	default:
		return fmt.Errorf("unknown database type")
	}

	location, err := flattenYDBLocation(database)
	if err != nil {
		return err
	}
	d.Set("location", location)

	d.Set("assign_public_ips", database.AssignPublicIps)
	d.Set("resource_preset_id", database.ResourcePresetId)

	d.Set("network_id", database.NetworkId)
	if err := d.Set("subnet_ids", database.SubnetIds); err != nil {
		return err
	}
	if err := d.Set("security_group_ids", database.SecurityGroupIds); err != nil {
		return err
	}

	storageConfig, err := flattenYDBStorageConfig(database.StorageConfig)
	if err != nil {
		return err
	}

	if err := d.Set("storage_config", storageConfig); err != nil {
		return err
	}

	scalePolicy, err := flattenYDBScalePolicy(database)
	if err != nil {
		return err
	}
	if err := d.Set("scale_policy", scalePolicy); err != nil {
		return err
	}

	return flattenYandexYDBDatabase(d, database)
}

func flattenYandexYDBDatabase(d *schema.ResourceData, database *ydb.Database) error {
	baseEP, dbPath, useTLS, err := parseYandexYDBDatabaseEndpoint(database.Endpoint)
	if err != nil {
		return err
	}

	d.Set("name", database.Name)
	d.Set("folder_id", database.FolderId)
	d.Set("description", database.Description)
	d.Set("created_at", getTimestamp(database.CreatedAt))
	if err := d.Set("labels", database.Labels); err != nil {
		return err
	}
	d.Set("location_id", database.LocationId)
	d.Set("ydb_full_endpoint", database.Endpoint)
	d.Set("ydb_api_endpoint", baseEP)
	d.Set("database_path", dbPath)
	d.Set("tls_enabled", useTLS)
	d.Set("deletion_protection", database.DeletionProtection)

	return d.Set("status", database.Status.String())
}

func parseYandexYDBDatabaseEndpoint(endpoint string) (baseEP, databasePath string, useTLS bool, err error) {
	dbSplit := strings.Split(endpoint, "/?database=")
	if len(dbSplit) != 2 {
		return "", "", false, fmt.Errorf("cannot parse endpoint %q", endpoint)
	}
	parts := strings.SplitN(dbSplit[0], "/", 3)
	if len(parts) < 3 {
		return "", "", false, fmt.Errorf("cannot parse endpoint schema %q", dbSplit[0])
	}

	const (
		protocolGRPCS = "grpcs:"
		protocolGRPC  = "grpc:"
	)

	switch protocol := parts[0]; protocol {
	case protocolGRPCS:
		useTLS = true
	case protocolGRPC:
		useTLS = false
	default:
		return "", "", false, fmt.Errorf("unknown protocol %q", protocol)
	}
	return parts[2], dbSplit[1], useTLS, nil
}
