// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"github.com/masikrus/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMdbMongodbCluster(ctx *pulumi.Context, args *LookupMdbMongodbClusterArgs, opts ...pulumi.InvokeOption) (*LookupMdbMongodbClusterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupMdbMongodbClusterResult
	err := ctx.Invoke("yandex:index/getMdbMongodbCluster:getMdbMongodbCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMdbMongodbCluster.
type LookupMdbMongodbClusterArgs struct {
	ClusterConfig *GetMdbMongodbClusterClusterConfig `pulumi:"clusterConfig"`
	ClusterId     *string                            `pulumi:"clusterId"`
	CreatedAt     *string                            `pulumi:"createdAt"`
	// Deprecated: to manage databases, please switch to using a separate resource type yandex_mdb_mongodb_database
	Databases                     []GetMdbMongodbClusterDatabase                     `pulumi:"databases"`
	DeletionProtection            *bool                                              `pulumi:"deletionProtection"`
	Description                   *string                                            `pulumi:"description"`
	DiskSizeAutoscalingMongocfg   *GetMdbMongodbClusterDiskSizeAutoscalingMongocfg   `pulumi:"diskSizeAutoscalingMongocfg"`
	DiskSizeAutoscalingMongod     *GetMdbMongodbClusterDiskSizeAutoscalingMongod     `pulumi:"diskSizeAutoscalingMongod"`
	DiskSizeAutoscalingMongoinfra *GetMdbMongodbClusterDiskSizeAutoscalingMongoinfra `pulumi:"diskSizeAutoscalingMongoinfra"`
	DiskSizeAutoscalingMongos     *GetMdbMongodbClusterDiskSizeAutoscalingMongos     `pulumi:"diskSizeAutoscalingMongos"`
	Environment                   *string                                            `pulumi:"environment"`
	FolderId                      *string                                            `pulumi:"folderId"`
	Health                        *string                                            `pulumi:"health"`
	Hosts                         []GetMdbMongodbClusterHost                         `pulumi:"hosts"`
	Labels                        map[string]string                                  `pulumi:"labels"`
	MaintenanceWindow             *GetMdbMongodbClusterMaintenanceWindow             `pulumi:"maintenanceWindow"`
	Name                          *string                                            `pulumi:"name"`
	NetworkId                     *string                                            `pulumi:"networkId"`
	// Deprecated: to manage `resources`s, please switch to using a separate resource type `resources_mongo*`
	Resources           *GetMdbMongodbClusterResources           `pulumi:"resources"`
	ResourcesMongocfg   *GetMdbMongodbClusterResourcesMongocfg   `pulumi:"resourcesMongocfg"`
	ResourcesMongod     *GetMdbMongodbClusterResourcesMongod     `pulumi:"resourcesMongod"`
	ResourcesMongoinfra *GetMdbMongodbClusterResourcesMongoinfra `pulumi:"resourcesMongoinfra"`
	ResourcesMongos     *GetMdbMongodbClusterResourcesMongos     `pulumi:"resourcesMongos"`
	Restore             *GetMdbMongodbClusterRestore             `pulumi:"restore"`
	SecurityGroupIds    []string                                 `pulumi:"securityGroupIds"`
	Sharded             *bool                                    `pulumi:"sharded"`
	Status              *string                                  `pulumi:"status"`
	// Deprecated: to manage users, please switch to using a separate resource type yandex_mdb_mongodb_user
	Users []GetMdbMongodbClusterUser `pulumi:"users"`
}

// A collection of values returned by getMdbMongodbCluster.
type LookupMdbMongodbClusterResult struct {
	ClusterConfig *GetMdbMongodbClusterClusterConfig `pulumi:"clusterConfig"`
	ClusterId     string                             `pulumi:"clusterId"`
	CreatedAt     string                             `pulumi:"createdAt"`
	// Deprecated: to manage databases, please switch to using a separate resource type yandex_mdb_mongodb_database
	Databases                     []GetMdbMongodbClusterDatabase                    `pulumi:"databases"`
	DeletionProtection            bool                                              `pulumi:"deletionProtection"`
	Description                   *string                                           `pulumi:"description"`
	DiskSizeAutoscalingMongocfg   GetMdbMongodbClusterDiskSizeAutoscalingMongocfg   `pulumi:"diskSizeAutoscalingMongocfg"`
	DiskSizeAutoscalingMongod     GetMdbMongodbClusterDiskSizeAutoscalingMongod     `pulumi:"diskSizeAutoscalingMongod"`
	DiskSizeAutoscalingMongoinfra GetMdbMongodbClusterDiskSizeAutoscalingMongoinfra `pulumi:"diskSizeAutoscalingMongoinfra"`
	DiskSizeAutoscalingMongos     GetMdbMongodbClusterDiskSizeAutoscalingMongos     `pulumi:"diskSizeAutoscalingMongos"`
	Environment                   *string                                           `pulumi:"environment"`
	FolderId                      string                                            `pulumi:"folderId"`
	Health                        string                                            `pulumi:"health"`
	Hosts                         []GetMdbMongodbClusterHost                        `pulumi:"hosts"`
	// The provider-assigned unique ID for this managed resource.
	Id                string                                `pulumi:"id"`
	Labels            map[string]string                     `pulumi:"labels"`
	MaintenanceWindow GetMdbMongodbClusterMaintenanceWindow `pulumi:"maintenanceWindow"`
	Name              *string                               `pulumi:"name"`
	NetworkId         *string                               `pulumi:"networkId"`
	// Deprecated: to manage `resources`s, please switch to using a separate resource type `resources_mongo*`
	Resources           *GetMdbMongodbClusterResources           `pulumi:"resources"`
	ResourcesMongocfg   *GetMdbMongodbClusterResourcesMongocfg   `pulumi:"resourcesMongocfg"`
	ResourcesMongod     *GetMdbMongodbClusterResourcesMongod     `pulumi:"resourcesMongod"`
	ResourcesMongoinfra *GetMdbMongodbClusterResourcesMongoinfra `pulumi:"resourcesMongoinfra"`
	ResourcesMongos     *GetMdbMongodbClusterResourcesMongos     `pulumi:"resourcesMongos"`
	Restore             *GetMdbMongodbClusterRestore             `pulumi:"restore"`
	SecurityGroupIds    []string                                 `pulumi:"securityGroupIds"`
	Sharded             bool                                     `pulumi:"sharded"`
	Status              string                                   `pulumi:"status"`
	// Deprecated: to manage users, please switch to using a separate resource type yandex_mdb_mongodb_user
	Users []GetMdbMongodbClusterUser `pulumi:"users"`
}

func LookupMdbMongodbClusterOutput(ctx *pulumi.Context, args LookupMdbMongodbClusterOutputArgs, opts ...pulumi.InvokeOption) LookupMdbMongodbClusterResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupMdbMongodbClusterResultOutput, error) {
			args := v.(LookupMdbMongodbClusterArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("yandex:index/getMdbMongodbCluster:getMdbMongodbCluster", args, LookupMdbMongodbClusterResultOutput{}, options).(LookupMdbMongodbClusterResultOutput), nil
		}).(LookupMdbMongodbClusterResultOutput)
}

// A collection of arguments for invoking getMdbMongodbCluster.
type LookupMdbMongodbClusterOutputArgs struct {
	ClusterConfig GetMdbMongodbClusterClusterConfigPtrInput `pulumi:"clusterConfig"`
	ClusterId     pulumi.StringPtrInput                     `pulumi:"clusterId"`
	CreatedAt     pulumi.StringPtrInput                     `pulumi:"createdAt"`
	// Deprecated: to manage databases, please switch to using a separate resource type yandex_mdb_mongodb_database
	Databases                     GetMdbMongodbClusterDatabaseArrayInput                    `pulumi:"databases"`
	DeletionProtection            pulumi.BoolPtrInput                                       `pulumi:"deletionProtection"`
	Description                   pulumi.StringPtrInput                                     `pulumi:"description"`
	DiskSizeAutoscalingMongocfg   GetMdbMongodbClusterDiskSizeAutoscalingMongocfgPtrInput   `pulumi:"diskSizeAutoscalingMongocfg"`
	DiskSizeAutoscalingMongod     GetMdbMongodbClusterDiskSizeAutoscalingMongodPtrInput     `pulumi:"diskSizeAutoscalingMongod"`
	DiskSizeAutoscalingMongoinfra GetMdbMongodbClusterDiskSizeAutoscalingMongoinfraPtrInput `pulumi:"diskSizeAutoscalingMongoinfra"`
	DiskSizeAutoscalingMongos     GetMdbMongodbClusterDiskSizeAutoscalingMongosPtrInput     `pulumi:"diskSizeAutoscalingMongos"`
	Environment                   pulumi.StringPtrInput                                     `pulumi:"environment"`
	FolderId                      pulumi.StringPtrInput                                     `pulumi:"folderId"`
	Health                        pulumi.StringPtrInput                                     `pulumi:"health"`
	Hosts                         GetMdbMongodbClusterHostArrayInput                        `pulumi:"hosts"`
	Labels                        pulumi.StringMapInput                                     `pulumi:"labels"`
	MaintenanceWindow             GetMdbMongodbClusterMaintenanceWindowPtrInput             `pulumi:"maintenanceWindow"`
	Name                          pulumi.StringPtrInput                                     `pulumi:"name"`
	NetworkId                     pulumi.StringPtrInput                                     `pulumi:"networkId"`
	// Deprecated: to manage `resources`s, please switch to using a separate resource type `resources_mongo*`
	Resources           GetMdbMongodbClusterResourcesPtrInput           `pulumi:"resources"`
	ResourcesMongocfg   GetMdbMongodbClusterResourcesMongocfgPtrInput   `pulumi:"resourcesMongocfg"`
	ResourcesMongod     GetMdbMongodbClusterResourcesMongodPtrInput     `pulumi:"resourcesMongod"`
	ResourcesMongoinfra GetMdbMongodbClusterResourcesMongoinfraPtrInput `pulumi:"resourcesMongoinfra"`
	ResourcesMongos     GetMdbMongodbClusterResourcesMongosPtrInput     `pulumi:"resourcesMongos"`
	Restore             GetMdbMongodbClusterRestorePtrInput             `pulumi:"restore"`
	SecurityGroupIds    pulumi.StringArrayInput                         `pulumi:"securityGroupIds"`
	Sharded             pulumi.BoolPtrInput                             `pulumi:"sharded"`
	Status              pulumi.StringPtrInput                           `pulumi:"status"`
	// Deprecated: to manage users, please switch to using a separate resource type yandex_mdb_mongodb_user
	Users GetMdbMongodbClusterUserArrayInput `pulumi:"users"`
}

func (LookupMdbMongodbClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMdbMongodbClusterArgs)(nil)).Elem()
}

// A collection of values returned by getMdbMongodbCluster.
type LookupMdbMongodbClusterResultOutput struct{ *pulumi.OutputState }

func (LookupMdbMongodbClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMdbMongodbClusterResult)(nil)).Elem()
}

func (o LookupMdbMongodbClusterResultOutput) ToLookupMdbMongodbClusterResultOutput() LookupMdbMongodbClusterResultOutput {
	return o
}

func (o LookupMdbMongodbClusterResultOutput) ToLookupMdbMongodbClusterResultOutputWithContext(ctx context.Context) LookupMdbMongodbClusterResultOutput {
	return o
}

func (o LookupMdbMongodbClusterResultOutput) ClusterConfig() GetMdbMongodbClusterClusterConfigPtrOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) *GetMdbMongodbClusterClusterConfig { return v.ClusterConfig }).(GetMdbMongodbClusterClusterConfigPtrOutput)
}

func (o LookupMdbMongodbClusterResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o LookupMdbMongodbClusterResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Deprecated: to manage databases, please switch to using a separate resource type yandex_mdb_mongodb_database
func (o LookupMdbMongodbClusterResultOutput) Databases() GetMdbMongodbClusterDatabaseArrayOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) []GetMdbMongodbClusterDatabase { return v.Databases }).(GetMdbMongodbClusterDatabaseArrayOutput)
}

func (o LookupMdbMongodbClusterResultOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) bool { return v.DeletionProtection }).(pulumi.BoolOutput)
}

func (o LookupMdbMongodbClusterResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupMdbMongodbClusterResultOutput) DiskSizeAutoscalingMongocfg() GetMdbMongodbClusterDiskSizeAutoscalingMongocfgOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) GetMdbMongodbClusterDiskSizeAutoscalingMongocfg {
		return v.DiskSizeAutoscalingMongocfg
	}).(GetMdbMongodbClusterDiskSizeAutoscalingMongocfgOutput)
}

func (o LookupMdbMongodbClusterResultOutput) DiskSizeAutoscalingMongod() GetMdbMongodbClusterDiskSizeAutoscalingMongodOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) GetMdbMongodbClusterDiskSizeAutoscalingMongod {
		return v.DiskSizeAutoscalingMongod
	}).(GetMdbMongodbClusterDiskSizeAutoscalingMongodOutput)
}

func (o LookupMdbMongodbClusterResultOutput) DiskSizeAutoscalingMongoinfra() GetMdbMongodbClusterDiskSizeAutoscalingMongoinfraOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) GetMdbMongodbClusterDiskSizeAutoscalingMongoinfra {
		return v.DiskSizeAutoscalingMongoinfra
	}).(GetMdbMongodbClusterDiskSizeAutoscalingMongoinfraOutput)
}

func (o LookupMdbMongodbClusterResultOutput) DiskSizeAutoscalingMongos() GetMdbMongodbClusterDiskSizeAutoscalingMongosOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) GetMdbMongodbClusterDiskSizeAutoscalingMongos {
		return v.DiskSizeAutoscalingMongos
	}).(GetMdbMongodbClusterDiskSizeAutoscalingMongosOutput)
}

func (o LookupMdbMongodbClusterResultOutput) Environment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) *string { return v.Environment }).(pulumi.StringPtrOutput)
}

func (o LookupMdbMongodbClusterResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) string { return v.FolderId }).(pulumi.StringOutput)
}

func (o LookupMdbMongodbClusterResultOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) string { return v.Health }).(pulumi.StringOutput)
}

func (o LookupMdbMongodbClusterResultOutput) Hosts() GetMdbMongodbClusterHostArrayOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) []GetMdbMongodbClusterHost { return v.Hosts }).(GetMdbMongodbClusterHostArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMdbMongodbClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMdbMongodbClusterResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupMdbMongodbClusterResultOutput) MaintenanceWindow() GetMdbMongodbClusterMaintenanceWindowOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) GetMdbMongodbClusterMaintenanceWindow {
		return v.MaintenanceWindow
	}).(GetMdbMongodbClusterMaintenanceWindowOutput)
}

func (o LookupMdbMongodbClusterResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupMdbMongodbClusterResultOutput) NetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) *string { return v.NetworkId }).(pulumi.StringPtrOutput)
}

// Deprecated: to manage `resources`s, please switch to using a separate resource type `resources_mongo*`
func (o LookupMdbMongodbClusterResultOutput) Resources() GetMdbMongodbClusterResourcesPtrOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) *GetMdbMongodbClusterResources { return v.Resources }).(GetMdbMongodbClusterResourcesPtrOutput)
}

func (o LookupMdbMongodbClusterResultOutput) ResourcesMongocfg() GetMdbMongodbClusterResourcesMongocfgPtrOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) *GetMdbMongodbClusterResourcesMongocfg {
		return v.ResourcesMongocfg
	}).(GetMdbMongodbClusterResourcesMongocfgPtrOutput)
}

func (o LookupMdbMongodbClusterResultOutput) ResourcesMongod() GetMdbMongodbClusterResourcesMongodPtrOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) *GetMdbMongodbClusterResourcesMongod { return v.ResourcesMongod }).(GetMdbMongodbClusterResourcesMongodPtrOutput)
}

func (o LookupMdbMongodbClusterResultOutput) ResourcesMongoinfra() GetMdbMongodbClusterResourcesMongoinfraPtrOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) *GetMdbMongodbClusterResourcesMongoinfra {
		return v.ResourcesMongoinfra
	}).(GetMdbMongodbClusterResourcesMongoinfraPtrOutput)
}

func (o LookupMdbMongodbClusterResultOutput) ResourcesMongos() GetMdbMongodbClusterResourcesMongosPtrOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) *GetMdbMongodbClusterResourcesMongos { return v.ResourcesMongos }).(GetMdbMongodbClusterResourcesMongosPtrOutput)
}

func (o LookupMdbMongodbClusterResultOutput) Restore() GetMdbMongodbClusterRestorePtrOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) *GetMdbMongodbClusterRestore { return v.Restore }).(GetMdbMongodbClusterRestorePtrOutput)
}

func (o LookupMdbMongodbClusterResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o LookupMdbMongodbClusterResultOutput) Sharded() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) bool { return v.Sharded }).(pulumi.BoolOutput)
}

func (o LookupMdbMongodbClusterResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) string { return v.Status }).(pulumi.StringOutput)
}

// Deprecated: to manage users, please switch to using a separate resource type yandex_mdb_mongodb_user
func (o LookupMdbMongodbClusterResultOutput) Users() GetMdbMongodbClusterUserArrayOutput {
	return o.ApplyT(func(v LookupMdbMongodbClusterResult) []GetMdbMongodbClusterUser { return v.Users }).(GetMdbMongodbClusterUserArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMdbMongodbClusterResultOutput{})
}
