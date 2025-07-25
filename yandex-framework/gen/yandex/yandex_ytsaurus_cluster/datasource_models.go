// Code generated by tfgen. DO NOT EDIT.

package yandex_ytsaurus_cluster

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	ytsaurus "github.com/yandex-cloud/go-genproto/yandex/cloud/ytsaurus/v1"
)

type yandexYtsaurusClusterDatasourceModel struct {
	ClusterId        types.String   `tfsdk:"cluster_id"`
	ID               types.String   `tfsdk:"id"`
	CreatedAt        types.String   `tfsdk:"created_at"`
	CreatedBy        types.String   `tfsdk:"created_by"`
	Description      types.String   `tfsdk:"description"`
	Endpoints        types.Object   `tfsdk:"endpoints"`
	FolderId         types.String   `tfsdk:"folder_id"`
	Health           types.String   `tfsdk:"health"`
	Labels           types.Map      `tfsdk:"labels"`
	Name             types.String   `tfsdk:"name"`
	SecurityGroupIds types.List     `tfsdk:"security_group_ids"`
	Spec             types.Object   `tfsdk:"spec"`
	Status           types.String   `tfsdk:"status"`
	SubnetId         types.String   `tfsdk:"subnet_id"`
	UpdatedAt        types.String   `tfsdk:"updated_at"`
	UpdatedBy        types.String   `tfsdk:"updated_by"`
	ZoneId           types.String   `tfsdk:"zone_id"`
	Timeouts         timeouts.Value `tfsdk:"timeouts"`
}

var yandexYtsaurusClusterDatasourceModelType = types.ObjectType{
	AttrTypes: map[string]attr.Type{
		"cluster_id":         types.StringType,
		"id":                 types.StringType,
		"created_at":         types.StringType,
		"created_by":         types.StringType,
		"description":        types.StringType,
		"endpoints":          yandexYtsaurusClusterEndpointsModelType,
		"folder_id":          types.StringType,
		"health":             types.StringType,
		"labels":             types.MapType{ElemType: types.StringType},
		"name":               types.StringType,
		"security_group_ids": types.ListType{ElemType: types.StringType},
		"spec":               yandexYtsaurusClusterSpecModelType,
		"status":             types.StringType,
		"subnet_id":          types.StringType,
		"updated_at":         types.StringType,
		"updated_by":         types.StringType,
		"zone_id":            types.StringType,
		"timeouts":           timeouts.AttributesAll(context.Background()).GetType(),
	},
}

func flattenYandexYtsaurusClusterDatasource(ctx context.Context,
	yandexYtsaurusClusterDatasource *ytsaurus.Cluster,
	state yandexYtsaurusClusterDatasourceModel,
	to timeouts.Value,
	diags *diag.Diagnostics) types.Object {
	if yandexYtsaurusClusterDatasource == nil {
		return types.ObjectNull(yandexYtsaurusClusterDatasourceModelType.AttrTypes)
	}
	value, diag := types.ObjectValueFrom(ctx, yandexYtsaurusClusterDatasourceModelType.AttrTypes, yandexYtsaurusClusterDatasourceModel{
		ClusterId:        types.StringValue(yandexYtsaurusClusterDatasource.GetId()),
		ID:               types.StringValue(yandexYtsaurusClusterDatasource.GetId()),
		CreatedAt:        types.StringValue(yandexYtsaurusClusterDatasource.GetCreatedAt().AsTime().Format(time.RFC3339)),
		CreatedBy:        types.StringValue(yandexYtsaurusClusterDatasource.GetCreatedBy()),
		Description:      types.StringValue(yandexYtsaurusClusterDatasource.GetDescription()),
		Endpoints:        flattenYandexYtsaurusClusterEndpoints(ctx, yandexYtsaurusClusterDatasource.GetEndpoints(), diags),
		FolderId:         types.StringValue(yandexYtsaurusClusterDatasource.GetFolderId()),
		Health:           types.StringValue(yandexYtsaurusClusterDatasource.GetHealth().String()),
		Labels:           flattenYandexYtsaurusClusterLabels(ctx, yandexYtsaurusClusterDatasource.GetLabels(), diags),
		Name:             types.StringValue(yandexYtsaurusClusterDatasource.GetName()),
		SecurityGroupIds: flattenYandexYtsaurusClusterSecurityGroupIds(ctx, yandexYtsaurusClusterDatasource.GetSecurityGroupIds(), diags),
		Spec:             flattenYandexYtsaurusClusterSpec(ctx, yandexYtsaurusClusterDatasource.GetSpec(), diags),
		Status:           types.StringValue(yandexYtsaurusClusterDatasource.GetStatus().String()),
		SubnetId:         types.StringValue(yandexYtsaurusClusterDatasource.GetSubnetId()),
		UpdatedAt:        types.StringValue(yandexYtsaurusClusterDatasource.GetUpdatedAt().AsTime().Format(time.RFC3339)),
		UpdatedBy:        types.StringValue(yandexYtsaurusClusterDatasource.GetUpdatedBy()),
		ZoneId:           types.StringValue(yandexYtsaurusClusterDatasource.GetZoneId()),
		Timeouts:         to,
	})
	diags.Append(diag...)
	return value
}
