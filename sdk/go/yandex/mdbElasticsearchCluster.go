// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package yandex

import (
	"context"
	"reflect"

	"errors"
	"github.com/masikrus/pulumi-yandex/sdk/go/yandex/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MdbElasticsearchCluster struct {
	pulumi.CustomResourceState

	// Configuration of the Elasticsearch cluster.
	Config MdbElasticsearchClusterConfigOutput `pulumi:"config"`
	// The creation timestamp of the resource.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The `true` value means that resource is protected from accidental deletion.
	DeletionProtection pulumi.BoolOutput `pulumi:"deletionProtection"`
	// The resource description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Deployment environment of the Elasticsearch cluster. Can be either `PRESTABLE` or `PRODUCTION`.
	Environment pulumi.StringOutput `pulumi:"environment"`
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// Aggregated health of the cluster. Can be either `ALIVE`, `DEGRADED`, `DEAD` or `HEALTH_UNKNOWN`. For more information
	// see `health` field of JSON representation in [the official
	// documentation](https://yandex.cloud/docs/managed-elasticsearch/api-ref/Cluster/).
	Health pulumi.StringOutput `pulumi:"health"`
	// A host of the Elasticsearch cluster.
	Hosts MdbElasticsearchClusterHostArrayOutput `pulumi:"hosts"`
	// A set of key/value label pairs which assigned to resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Elasticsearch cluster maintenance window.
	MaintenanceWindow MdbElasticsearchClusterMaintenanceWindowOutput `pulumi:"maintenanceWindow"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The `VPC Network ID` of subnets which resource attached to.
	NetworkId        pulumi.StringOutput      `pulumi:"networkId"`
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// [Service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) which linked to the resource.
	ServiceAccountId pulumi.StringPtrOutput `pulumi:"serviceAccountId"`
	// Status of the cluster. Can be either `CREATING`, `STARTING`, `RUNNING`, `UPDATING`, `STOPPING`, `STOPPED`, `ERROR` or
	// `STATUS_UNKNOWN`. For more information see `status` field of JSON representation in [the official
	// documentation](https://yandex.cloud/docs/managed-elasticsearch/api-ref/Cluster/).
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewMdbElasticsearchCluster registers a new resource with the given unique name, arguments, and options.
func NewMdbElasticsearchCluster(ctx *pulumi.Context,
	name string, args *MdbElasticsearchClusterArgs, opts ...pulumi.ResourceOption) (*MdbElasticsearchCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Config == nil {
		return nil, errors.New("invalid value for required argument 'Config'")
	}
	if args.Environment == nil {
		return nil, errors.New("invalid value for required argument 'Environment'")
	}
	if args.NetworkId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MdbElasticsearchCluster
	err := ctx.RegisterResource("yandex:index/mdbElasticsearchCluster:MdbElasticsearchCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMdbElasticsearchCluster gets an existing MdbElasticsearchCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMdbElasticsearchCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MdbElasticsearchClusterState, opts ...pulumi.ResourceOption) (*MdbElasticsearchCluster, error) {
	var resource MdbElasticsearchCluster
	err := ctx.ReadResource("yandex:index/mdbElasticsearchCluster:MdbElasticsearchCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MdbElasticsearchCluster resources.
type mdbElasticsearchClusterState struct {
	// Configuration of the Elasticsearch cluster.
	Config *MdbElasticsearchClusterConfig `pulumi:"config"`
	// The creation timestamp of the resource.
	CreatedAt *string `pulumi:"createdAt"`
	// The `true` value means that resource is protected from accidental deletion.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// The resource description.
	Description *string `pulumi:"description"`
	// Deployment environment of the Elasticsearch cluster. Can be either `PRESTABLE` or `PRODUCTION`.
	Environment *string `pulumi:"environment"`
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId *string `pulumi:"folderId"`
	// Aggregated health of the cluster. Can be either `ALIVE`, `DEGRADED`, `DEAD` or `HEALTH_UNKNOWN`. For more information
	// see `health` field of JSON representation in [the official
	// documentation](https://yandex.cloud/docs/managed-elasticsearch/api-ref/Cluster/).
	Health *string `pulumi:"health"`
	// A host of the Elasticsearch cluster.
	Hosts []MdbElasticsearchClusterHost `pulumi:"hosts"`
	// A set of key/value label pairs which assigned to resource.
	Labels map[string]string `pulumi:"labels"`
	// Elasticsearch cluster maintenance window.
	MaintenanceWindow *MdbElasticsearchClusterMaintenanceWindow `pulumi:"maintenanceWindow"`
	// The resource name.
	Name *string `pulumi:"name"`
	// The `VPC Network ID` of subnets which resource attached to.
	NetworkId        *string  `pulumi:"networkId"`
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// [Service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) which linked to the resource.
	ServiceAccountId *string `pulumi:"serviceAccountId"`
	// Status of the cluster. Can be either `CREATING`, `STARTING`, `RUNNING`, `UPDATING`, `STOPPING`, `STOPPED`, `ERROR` or
	// `STATUS_UNKNOWN`. For more information see `status` field of JSON representation in [the official
	// documentation](https://yandex.cloud/docs/managed-elasticsearch/api-ref/Cluster/).
	Status *string `pulumi:"status"`
}

type MdbElasticsearchClusterState struct {
	// Configuration of the Elasticsearch cluster.
	Config MdbElasticsearchClusterConfigPtrInput
	// The creation timestamp of the resource.
	CreatedAt pulumi.StringPtrInput
	// The `true` value means that resource is protected from accidental deletion.
	DeletionProtection pulumi.BoolPtrInput
	// The resource description.
	Description pulumi.StringPtrInput
	// Deployment environment of the Elasticsearch cluster. Can be either `PRESTABLE` or `PRODUCTION`.
	Environment pulumi.StringPtrInput
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId pulumi.StringPtrInput
	// Aggregated health of the cluster. Can be either `ALIVE`, `DEGRADED`, `DEAD` or `HEALTH_UNKNOWN`. For more information
	// see `health` field of JSON representation in [the official
	// documentation](https://yandex.cloud/docs/managed-elasticsearch/api-ref/Cluster/).
	Health pulumi.StringPtrInput
	// A host of the Elasticsearch cluster.
	Hosts MdbElasticsearchClusterHostArrayInput
	// A set of key/value label pairs which assigned to resource.
	Labels pulumi.StringMapInput
	// Elasticsearch cluster maintenance window.
	MaintenanceWindow MdbElasticsearchClusterMaintenanceWindowPtrInput
	// The resource name.
	Name pulumi.StringPtrInput
	// The `VPC Network ID` of subnets which resource attached to.
	NetworkId        pulumi.StringPtrInput
	SecurityGroupIds pulumi.StringArrayInput
	// [Service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) which linked to the resource.
	ServiceAccountId pulumi.StringPtrInput
	// Status of the cluster. Can be either `CREATING`, `STARTING`, `RUNNING`, `UPDATING`, `STOPPING`, `STOPPED`, `ERROR` or
	// `STATUS_UNKNOWN`. For more information see `status` field of JSON representation in [the official
	// documentation](https://yandex.cloud/docs/managed-elasticsearch/api-ref/Cluster/).
	Status pulumi.StringPtrInput
}

func (MdbElasticsearchClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*mdbElasticsearchClusterState)(nil)).Elem()
}

type mdbElasticsearchClusterArgs struct {
	// Configuration of the Elasticsearch cluster.
	Config MdbElasticsearchClusterConfig `pulumi:"config"`
	// The `true` value means that resource is protected from accidental deletion.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// The resource description.
	Description *string `pulumi:"description"`
	// Deployment environment of the Elasticsearch cluster. Can be either `PRESTABLE` or `PRODUCTION`.
	Environment string `pulumi:"environment"`
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId *string `pulumi:"folderId"`
	// A host of the Elasticsearch cluster.
	Hosts []MdbElasticsearchClusterHost `pulumi:"hosts"`
	// A set of key/value label pairs which assigned to resource.
	Labels map[string]string `pulumi:"labels"`
	// Elasticsearch cluster maintenance window.
	MaintenanceWindow *MdbElasticsearchClusterMaintenanceWindow `pulumi:"maintenanceWindow"`
	// The resource name.
	Name *string `pulumi:"name"`
	// The `VPC Network ID` of subnets which resource attached to.
	NetworkId        string   `pulumi:"networkId"`
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// [Service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) which linked to the resource.
	ServiceAccountId *string `pulumi:"serviceAccountId"`
}

// The set of arguments for constructing a MdbElasticsearchCluster resource.
type MdbElasticsearchClusterArgs struct {
	// Configuration of the Elasticsearch cluster.
	Config MdbElasticsearchClusterConfigInput
	// The `true` value means that resource is protected from accidental deletion.
	DeletionProtection pulumi.BoolPtrInput
	// The resource description.
	Description pulumi.StringPtrInput
	// Deployment environment of the Elasticsearch cluster. Can be either `PRESTABLE` or `PRODUCTION`.
	Environment pulumi.StringInput
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId pulumi.StringPtrInput
	// A host of the Elasticsearch cluster.
	Hosts MdbElasticsearchClusterHostArrayInput
	// A set of key/value label pairs which assigned to resource.
	Labels pulumi.StringMapInput
	// Elasticsearch cluster maintenance window.
	MaintenanceWindow MdbElasticsearchClusterMaintenanceWindowPtrInput
	// The resource name.
	Name pulumi.StringPtrInput
	// The `VPC Network ID` of subnets which resource attached to.
	NetworkId        pulumi.StringInput
	SecurityGroupIds pulumi.StringArrayInput
	// [Service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) which linked to the resource.
	ServiceAccountId pulumi.StringPtrInput
}

func (MdbElasticsearchClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mdbElasticsearchClusterArgs)(nil)).Elem()
}

type MdbElasticsearchClusterInput interface {
	pulumi.Input

	ToMdbElasticsearchClusterOutput() MdbElasticsearchClusterOutput
	ToMdbElasticsearchClusterOutputWithContext(ctx context.Context) MdbElasticsearchClusterOutput
}

func (*MdbElasticsearchCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**MdbElasticsearchCluster)(nil)).Elem()
}

func (i *MdbElasticsearchCluster) ToMdbElasticsearchClusterOutput() MdbElasticsearchClusterOutput {
	return i.ToMdbElasticsearchClusterOutputWithContext(context.Background())
}

func (i *MdbElasticsearchCluster) ToMdbElasticsearchClusterOutputWithContext(ctx context.Context) MdbElasticsearchClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbElasticsearchClusterOutput)
}

// MdbElasticsearchClusterArrayInput is an input type that accepts MdbElasticsearchClusterArray and MdbElasticsearchClusterArrayOutput values.
// You can construct a concrete instance of `MdbElasticsearchClusterArrayInput` via:
//
//	MdbElasticsearchClusterArray{ MdbElasticsearchClusterArgs{...} }
type MdbElasticsearchClusterArrayInput interface {
	pulumi.Input

	ToMdbElasticsearchClusterArrayOutput() MdbElasticsearchClusterArrayOutput
	ToMdbElasticsearchClusterArrayOutputWithContext(context.Context) MdbElasticsearchClusterArrayOutput
}

type MdbElasticsearchClusterArray []MdbElasticsearchClusterInput

func (MdbElasticsearchClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MdbElasticsearchCluster)(nil)).Elem()
}

func (i MdbElasticsearchClusterArray) ToMdbElasticsearchClusterArrayOutput() MdbElasticsearchClusterArrayOutput {
	return i.ToMdbElasticsearchClusterArrayOutputWithContext(context.Background())
}

func (i MdbElasticsearchClusterArray) ToMdbElasticsearchClusterArrayOutputWithContext(ctx context.Context) MdbElasticsearchClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbElasticsearchClusterArrayOutput)
}

// MdbElasticsearchClusterMapInput is an input type that accepts MdbElasticsearchClusterMap and MdbElasticsearchClusterMapOutput values.
// You can construct a concrete instance of `MdbElasticsearchClusterMapInput` via:
//
//	MdbElasticsearchClusterMap{ "key": MdbElasticsearchClusterArgs{...} }
type MdbElasticsearchClusterMapInput interface {
	pulumi.Input

	ToMdbElasticsearchClusterMapOutput() MdbElasticsearchClusterMapOutput
	ToMdbElasticsearchClusterMapOutputWithContext(context.Context) MdbElasticsearchClusterMapOutput
}

type MdbElasticsearchClusterMap map[string]MdbElasticsearchClusterInput

func (MdbElasticsearchClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MdbElasticsearchCluster)(nil)).Elem()
}

func (i MdbElasticsearchClusterMap) ToMdbElasticsearchClusterMapOutput() MdbElasticsearchClusterMapOutput {
	return i.ToMdbElasticsearchClusterMapOutputWithContext(context.Background())
}

func (i MdbElasticsearchClusterMap) ToMdbElasticsearchClusterMapOutputWithContext(ctx context.Context) MdbElasticsearchClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbElasticsearchClusterMapOutput)
}

type MdbElasticsearchClusterOutput struct{ *pulumi.OutputState }

func (MdbElasticsearchClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MdbElasticsearchCluster)(nil)).Elem()
}

func (o MdbElasticsearchClusterOutput) ToMdbElasticsearchClusterOutput() MdbElasticsearchClusterOutput {
	return o
}

func (o MdbElasticsearchClusterOutput) ToMdbElasticsearchClusterOutputWithContext(ctx context.Context) MdbElasticsearchClusterOutput {
	return o
}

// Configuration of the Elasticsearch cluster.
func (o MdbElasticsearchClusterOutput) Config() MdbElasticsearchClusterConfigOutput {
	return o.ApplyT(func(v *MdbElasticsearchCluster) MdbElasticsearchClusterConfigOutput { return v.Config }).(MdbElasticsearchClusterConfigOutput)
}

// The creation timestamp of the resource.
func (o MdbElasticsearchClusterOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbElasticsearchCluster) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The `true` value means that resource is protected from accidental deletion.
func (o MdbElasticsearchClusterOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v *MdbElasticsearchCluster) pulumi.BoolOutput { return v.DeletionProtection }).(pulumi.BoolOutput)
}

// The resource description.
func (o MdbElasticsearchClusterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MdbElasticsearchCluster) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Deployment environment of the Elasticsearch cluster. Can be either `PRESTABLE` or `PRODUCTION`.
func (o MdbElasticsearchClusterOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbElasticsearchCluster) pulumi.StringOutput { return v.Environment }).(pulumi.StringOutput)
}

// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
func (o MdbElasticsearchClusterOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbElasticsearchCluster) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

// Aggregated health of the cluster. Can be either `ALIVE`, `DEGRADED`, `DEAD` or `HEALTH_UNKNOWN`. For more information
// see `health` field of JSON representation in [the official
// documentation](https://yandex.cloud/docs/managed-elasticsearch/api-ref/Cluster/).
func (o MdbElasticsearchClusterOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbElasticsearchCluster) pulumi.StringOutput { return v.Health }).(pulumi.StringOutput)
}

// A host of the Elasticsearch cluster.
func (o MdbElasticsearchClusterOutput) Hosts() MdbElasticsearchClusterHostArrayOutput {
	return o.ApplyT(func(v *MdbElasticsearchCluster) MdbElasticsearchClusterHostArrayOutput { return v.Hosts }).(MdbElasticsearchClusterHostArrayOutput)
}

// A set of key/value label pairs which assigned to resource.
func (o MdbElasticsearchClusterOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MdbElasticsearchCluster) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Elasticsearch cluster maintenance window.
func (o MdbElasticsearchClusterOutput) MaintenanceWindow() MdbElasticsearchClusterMaintenanceWindowOutput {
	return o.ApplyT(func(v *MdbElasticsearchCluster) MdbElasticsearchClusterMaintenanceWindowOutput {
		return v.MaintenanceWindow
	}).(MdbElasticsearchClusterMaintenanceWindowOutput)
}

// The resource name.
func (o MdbElasticsearchClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbElasticsearchCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The `VPC Network ID` of subnets which resource attached to.
func (o MdbElasticsearchClusterOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbElasticsearchCluster) pulumi.StringOutput { return v.NetworkId }).(pulumi.StringOutput)
}

func (o MdbElasticsearchClusterOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MdbElasticsearchCluster) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// [Service account](https://yandex.cloud/docs/iam/concepts/users/service-accounts) which linked to the resource.
func (o MdbElasticsearchClusterOutput) ServiceAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MdbElasticsearchCluster) pulumi.StringPtrOutput { return v.ServiceAccountId }).(pulumi.StringPtrOutput)
}

// Status of the cluster. Can be either `CREATING`, `STARTING`, `RUNNING`, `UPDATING`, `STOPPING`, `STOPPED`, `ERROR` or
// `STATUS_UNKNOWN`. For more information see `status` field of JSON representation in [the official
// documentation](https://yandex.cloud/docs/managed-elasticsearch/api-ref/Cluster/).
func (o MdbElasticsearchClusterOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbElasticsearchCluster) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type MdbElasticsearchClusterArrayOutput struct{ *pulumi.OutputState }

func (MdbElasticsearchClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MdbElasticsearchCluster)(nil)).Elem()
}

func (o MdbElasticsearchClusterArrayOutput) ToMdbElasticsearchClusterArrayOutput() MdbElasticsearchClusterArrayOutput {
	return o
}

func (o MdbElasticsearchClusterArrayOutput) ToMdbElasticsearchClusterArrayOutputWithContext(ctx context.Context) MdbElasticsearchClusterArrayOutput {
	return o
}

func (o MdbElasticsearchClusterArrayOutput) Index(i pulumi.IntInput) MdbElasticsearchClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MdbElasticsearchCluster {
		return vs[0].([]*MdbElasticsearchCluster)[vs[1].(int)]
	}).(MdbElasticsearchClusterOutput)
}

type MdbElasticsearchClusterMapOutput struct{ *pulumi.OutputState }

func (MdbElasticsearchClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MdbElasticsearchCluster)(nil)).Elem()
}

func (o MdbElasticsearchClusterMapOutput) ToMdbElasticsearchClusterMapOutput() MdbElasticsearchClusterMapOutput {
	return o
}

func (o MdbElasticsearchClusterMapOutput) ToMdbElasticsearchClusterMapOutputWithContext(ctx context.Context) MdbElasticsearchClusterMapOutput {
	return o
}

func (o MdbElasticsearchClusterMapOutput) MapIndex(k pulumi.StringInput) MdbElasticsearchClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MdbElasticsearchCluster {
		return vs[0].(map[string]*MdbElasticsearchCluster)[vs[1].(string)]
	}).(MdbElasticsearchClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MdbElasticsearchClusterInput)(nil)).Elem(), &MdbElasticsearchCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*MdbElasticsearchClusterArrayInput)(nil)).Elem(), MdbElasticsearchClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MdbElasticsearchClusterMapInput)(nil)).Elem(), MdbElasticsearchClusterMap{})
	pulumi.RegisterOutputType(MdbElasticsearchClusterOutput{})
	pulumi.RegisterOutputType(MdbElasticsearchClusterArrayOutput{})
	pulumi.RegisterOutputType(MdbElasticsearchClusterMapOutput{})
}
