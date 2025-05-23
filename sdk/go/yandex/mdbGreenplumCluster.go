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

type MdbGreenplumCluster struct {
	pulumi.CustomResourceState

	// Access policy to the Greenplum cluster.
	Access MdbGreenplumClusterAccessOutput `pulumi:"access"`
	// Sets whether the master hosts should get a public IP address on creation. Changing this parameter for an existing host
	// is not supported at the moment.
	AssignPublicIp pulumi.BoolOutput `pulumi:"assignPublicIp"`
	// Background activities settings.
	BackgroundActivities MdbGreenplumClusterBackgroundActivityArrayOutput `pulumi:"backgroundActivities"`
	// Time to start the daily backup, in the UTC timezone.
	BackupWindowStart MdbGreenplumClusterBackupWindowStartOutput `pulumi:"backupWindowStart"`
	// Cloud Storage settings of the Greenplum cluster.
	CloudStorage MdbGreenplumClusterCloudStorageOutput `pulumi:"cloudStorage"`
	// The creation timestamp of the resource.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The `true` value means that resource is protected from accidental deletion.
	DeletionProtection pulumi.BoolOutput `pulumi:"deletionProtection"`
	// The resource description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Deployment environment of the Greenplum cluster. (PRODUCTION, PRESTABLE)
	Environment pulumi.StringOutput `pulumi:"environment"`
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// Greenplum cluster config. Detail info in `Greenplum cluster settings` block.
	GreenplumConfig pulumi.StringMapOutput `pulumi:"greenplumConfig"`
	// Aggregated health of the cluster.
	Health pulumi.StringOutput `pulumi:"health"`
	// A set of key/value label pairs which assigned to resource.
	Labels  pulumi.StringMapOutput           `pulumi:"labels"`
	Logging MdbGreenplumClusterLoggingOutput `pulumi:"logging"`
	// Maintenance policy of the Greenplum cluster.
	MaintenanceWindow MdbGreenplumClusterMaintenanceWindowOutput `pulumi:"maintenanceWindow"`
	// Number of hosts in master subcluster (1 or 2).
	MasterHostCount pulumi.IntOutput `pulumi:"masterHostCount"`
	// A list of IDs of the host groups to place master subclusters' VMs of the cluster on.
	MasterHostGroupIds pulumi.StringArrayOutput `pulumi:"masterHostGroupIds"`
	// Info about hosts in master subcluster.
	MasterHosts MdbGreenplumClusterMasterHostArrayOutput `pulumi:"masterHosts"`
	// Settings for master subcluster.
	MasterSubcluster MdbGreenplumClusterMasterSubclusterOutput `pulumi:"masterSubcluster"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The `VPC Network ID` of subnets which resource attached to.
	NetworkId pulumi.StringOutput `pulumi:"networkId"`
	// Configuration of the connection pooler.
	PoolerConfig MdbGreenplumClusterPoolerConfigOutput `pulumi:"poolerConfig"`
	// Configuration of the PXF daemon.
	PxfConfig MdbGreenplumClusterPxfConfigOutput `pulumi:"pxfConfig"`
	// The list of security groups applied to resource or their components.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// Number of hosts in segment subcluster (from 1 to 32).
	SegmentHostCount pulumi.IntOutput `pulumi:"segmentHostCount"`
	// A list of IDs of the host groups to place segment subclusters' VMs of the cluster on.
	SegmentHostGroupIds pulumi.StringArrayOutput `pulumi:"segmentHostGroupIds"`
	// Info about hosts in segment subcluster.
	SegmentHosts MdbGreenplumClusterSegmentHostArrayOutput `pulumi:"segmentHosts"`
	// Number of segments on segment host (not more then 1 + RAM/8).
	SegmentInHost pulumi.IntOutput `pulumi:"segmentInHost"`
	// Settings for segment subcluster.
	SegmentSubcluster MdbGreenplumClusterSegmentSubclusterOutput `pulumi:"segmentSubcluster"`
	// ID of service account to use with Yandex Cloud resources (e.g. S3, Cloud Logging).
	ServiceAccountId pulumi.StringPtrOutput `pulumi:"serviceAccountId"`
	// Status of the cluster.
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID of the subnet, to which the hosts belongs. The subnet must be a part of the network to which the cluster belongs.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// Greenplum cluster admin user name.
	UserName pulumi.StringOutput `pulumi:"userName"`
	// Greenplum cluster admin password name.
	UserPassword pulumi.StringOutput `pulumi:"userPassword"`
	// Version of the Greenplum cluster. (`6.25`)
	Version pulumi.StringOutput `pulumi:"version"`
	// The [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) where resource is located. If it is not
	// provided, the default provider zone will be used.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewMdbGreenplumCluster registers a new resource with the given unique name, arguments, and options.
func NewMdbGreenplumCluster(ctx *pulumi.Context,
	name string, args *MdbGreenplumClusterArgs, opts ...pulumi.ResourceOption) (*MdbGreenplumCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssignPublicIp == nil {
		return nil, errors.New("invalid value for required argument 'AssignPublicIp'")
	}
	if args.Environment == nil {
		return nil, errors.New("invalid value for required argument 'Environment'")
	}
	if args.MasterHostCount == nil {
		return nil, errors.New("invalid value for required argument 'MasterHostCount'")
	}
	if args.MasterSubcluster == nil {
		return nil, errors.New("invalid value for required argument 'MasterSubcluster'")
	}
	if args.NetworkId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkId'")
	}
	if args.SegmentHostCount == nil {
		return nil, errors.New("invalid value for required argument 'SegmentHostCount'")
	}
	if args.SegmentInHost == nil {
		return nil, errors.New("invalid value for required argument 'SegmentInHost'")
	}
	if args.SegmentSubcluster == nil {
		return nil, errors.New("invalid value for required argument 'SegmentSubcluster'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	if args.UserPassword == nil {
		return nil, errors.New("invalid value for required argument 'UserPassword'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	if args.UserPassword != nil {
		args.UserPassword = pulumi.ToSecret(args.UserPassword).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"userPassword",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MdbGreenplumCluster
	err := ctx.RegisterResource("yandex:index/mdbGreenplumCluster:MdbGreenplumCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMdbGreenplumCluster gets an existing MdbGreenplumCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMdbGreenplumCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MdbGreenplumClusterState, opts ...pulumi.ResourceOption) (*MdbGreenplumCluster, error) {
	var resource MdbGreenplumCluster
	err := ctx.ReadResource("yandex:index/mdbGreenplumCluster:MdbGreenplumCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MdbGreenplumCluster resources.
type mdbGreenplumClusterState struct {
	// Access policy to the Greenplum cluster.
	Access *MdbGreenplumClusterAccess `pulumi:"access"`
	// Sets whether the master hosts should get a public IP address on creation. Changing this parameter for an existing host
	// is not supported at the moment.
	AssignPublicIp *bool `pulumi:"assignPublicIp"`
	// Background activities settings.
	BackgroundActivities []MdbGreenplumClusterBackgroundActivity `pulumi:"backgroundActivities"`
	// Time to start the daily backup, in the UTC timezone.
	BackupWindowStart *MdbGreenplumClusterBackupWindowStart `pulumi:"backupWindowStart"`
	// Cloud Storage settings of the Greenplum cluster.
	CloudStorage *MdbGreenplumClusterCloudStorage `pulumi:"cloudStorage"`
	// The creation timestamp of the resource.
	CreatedAt *string `pulumi:"createdAt"`
	// The `true` value means that resource is protected from accidental deletion.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// The resource description.
	Description *string `pulumi:"description"`
	// Deployment environment of the Greenplum cluster. (PRODUCTION, PRESTABLE)
	Environment *string `pulumi:"environment"`
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId *string `pulumi:"folderId"`
	// Greenplum cluster config. Detail info in `Greenplum cluster settings` block.
	GreenplumConfig map[string]string `pulumi:"greenplumConfig"`
	// Aggregated health of the cluster.
	Health *string `pulumi:"health"`
	// A set of key/value label pairs which assigned to resource.
	Labels  map[string]string           `pulumi:"labels"`
	Logging *MdbGreenplumClusterLogging `pulumi:"logging"`
	// Maintenance policy of the Greenplum cluster.
	MaintenanceWindow *MdbGreenplumClusterMaintenanceWindow `pulumi:"maintenanceWindow"`
	// Number of hosts in master subcluster (1 or 2).
	MasterHostCount *int `pulumi:"masterHostCount"`
	// A list of IDs of the host groups to place master subclusters' VMs of the cluster on.
	MasterHostGroupIds []string `pulumi:"masterHostGroupIds"`
	// Info about hosts in master subcluster.
	MasterHosts []MdbGreenplumClusterMasterHost `pulumi:"masterHosts"`
	// Settings for master subcluster.
	MasterSubcluster *MdbGreenplumClusterMasterSubcluster `pulumi:"masterSubcluster"`
	// The resource name.
	Name *string `pulumi:"name"`
	// The `VPC Network ID` of subnets which resource attached to.
	NetworkId *string `pulumi:"networkId"`
	// Configuration of the connection pooler.
	PoolerConfig *MdbGreenplumClusterPoolerConfig `pulumi:"poolerConfig"`
	// Configuration of the PXF daemon.
	PxfConfig *MdbGreenplumClusterPxfConfig `pulumi:"pxfConfig"`
	// The list of security groups applied to resource or their components.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Number of hosts in segment subcluster (from 1 to 32).
	SegmentHostCount *int `pulumi:"segmentHostCount"`
	// A list of IDs of the host groups to place segment subclusters' VMs of the cluster on.
	SegmentHostGroupIds []string `pulumi:"segmentHostGroupIds"`
	// Info about hosts in segment subcluster.
	SegmentHosts []MdbGreenplumClusterSegmentHost `pulumi:"segmentHosts"`
	// Number of segments on segment host (not more then 1 + RAM/8).
	SegmentInHost *int `pulumi:"segmentInHost"`
	// Settings for segment subcluster.
	SegmentSubcluster *MdbGreenplumClusterSegmentSubcluster `pulumi:"segmentSubcluster"`
	// ID of service account to use with Yandex Cloud resources (e.g. S3, Cloud Logging).
	ServiceAccountId *string `pulumi:"serviceAccountId"`
	// Status of the cluster.
	Status *string `pulumi:"status"`
	// The ID of the subnet, to which the hosts belongs. The subnet must be a part of the network to which the cluster belongs.
	SubnetId *string `pulumi:"subnetId"`
	// Greenplum cluster admin user name.
	UserName *string `pulumi:"userName"`
	// Greenplum cluster admin password name.
	UserPassword *string `pulumi:"userPassword"`
	// Version of the Greenplum cluster. (`6.25`)
	Version *string `pulumi:"version"`
	// The [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) where resource is located. If it is not
	// provided, the default provider zone will be used.
	Zone *string `pulumi:"zone"`
}

type MdbGreenplumClusterState struct {
	// Access policy to the Greenplum cluster.
	Access MdbGreenplumClusterAccessPtrInput
	// Sets whether the master hosts should get a public IP address on creation. Changing this parameter for an existing host
	// is not supported at the moment.
	AssignPublicIp pulumi.BoolPtrInput
	// Background activities settings.
	BackgroundActivities MdbGreenplumClusterBackgroundActivityArrayInput
	// Time to start the daily backup, in the UTC timezone.
	BackupWindowStart MdbGreenplumClusterBackupWindowStartPtrInput
	// Cloud Storage settings of the Greenplum cluster.
	CloudStorage MdbGreenplumClusterCloudStoragePtrInput
	// The creation timestamp of the resource.
	CreatedAt pulumi.StringPtrInput
	// The `true` value means that resource is protected from accidental deletion.
	DeletionProtection pulumi.BoolPtrInput
	// The resource description.
	Description pulumi.StringPtrInput
	// Deployment environment of the Greenplum cluster. (PRODUCTION, PRESTABLE)
	Environment pulumi.StringPtrInput
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId pulumi.StringPtrInput
	// Greenplum cluster config. Detail info in `Greenplum cluster settings` block.
	GreenplumConfig pulumi.StringMapInput
	// Aggregated health of the cluster.
	Health pulumi.StringPtrInput
	// A set of key/value label pairs which assigned to resource.
	Labels  pulumi.StringMapInput
	Logging MdbGreenplumClusterLoggingPtrInput
	// Maintenance policy of the Greenplum cluster.
	MaintenanceWindow MdbGreenplumClusterMaintenanceWindowPtrInput
	// Number of hosts in master subcluster (1 or 2).
	MasterHostCount pulumi.IntPtrInput
	// A list of IDs of the host groups to place master subclusters' VMs of the cluster on.
	MasterHostGroupIds pulumi.StringArrayInput
	// Info about hosts in master subcluster.
	MasterHosts MdbGreenplumClusterMasterHostArrayInput
	// Settings for master subcluster.
	MasterSubcluster MdbGreenplumClusterMasterSubclusterPtrInput
	// The resource name.
	Name pulumi.StringPtrInput
	// The `VPC Network ID` of subnets which resource attached to.
	NetworkId pulumi.StringPtrInput
	// Configuration of the connection pooler.
	PoolerConfig MdbGreenplumClusterPoolerConfigPtrInput
	// Configuration of the PXF daemon.
	PxfConfig MdbGreenplumClusterPxfConfigPtrInput
	// The list of security groups applied to resource or their components.
	SecurityGroupIds pulumi.StringArrayInput
	// Number of hosts in segment subcluster (from 1 to 32).
	SegmentHostCount pulumi.IntPtrInput
	// A list of IDs of the host groups to place segment subclusters' VMs of the cluster on.
	SegmentHostGroupIds pulumi.StringArrayInput
	// Info about hosts in segment subcluster.
	SegmentHosts MdbGreenplumClusterSegmentHostArrayInput
	// Number of segments on segment host (not more then 1 + RAM/8).
	SegmentInHost pulumi.IntPtrInput
	// Settings for segment subcluster.
	SegmentSubcluster MdbGreenplumClusterSegmentSubclusterPtrInput
	// ID of service account to use with Yandex Cloud resources (e.g. S3, Cloud Logging).
	ServiceAccountId pulumi.StringPtrInput
	// Status of the cluster.
	Status pulumi.StringPtrInput
	// The ID of the subnet, to which the hosts belongs. The subnet must be a part of the network to which the cluster belongs.
	SubnetId pulumi.StringPtrInput
	// Greenplum cluster admin user name.
	UserName pulumi.StringPtrInput
	// Greenplum cluster admin password name.
	UserPassword pulumi.StringPtrInput
	// Version of the Greenplum cluster. (`6.25`)
	Version pulumi.StringPtrInput
	// The [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) where resource is located. If it is not
	// provided, the default provider zone will be used.
	Zone pulumi.StringPtrInput
}

func (MdbGreenplumClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*mdbGreenplumClusterState)(nil)).Elem()
}

type mdbGreenplumClusterArgs struct {
	// Access policy to the Greenplum cluster.
	Access *MdbGreenplumClusterAccess `pulumi:"access"`
	// Sets whether the master hosts should get a public IP address on creation. Changing this parameter for an existing host
	// is not supported at the moment.
	AssignPublicIp bool `pulumi:"assignPublicIp"`
	// Background activities settings.
	BackgroundActivities []MdbGreenplumClusterBackgroundActivity `pulumi:"backgroundActivities"`
	// Time to start the daily backup, in the UTC timezone.
	BackupWindowStart *MdbGreenplumClusterBackupWindowStart `pulumi:"backupWindowStart"`
	// Cloud Storage settings of the Greenplum cluster.
	CloudStorage *MdbGreenplumClusterCloudStorage `pulumi:"cloudStorage"`
	// The `true` value means that resource is protected from accidental deletion.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// The resource description.
	Description *string `pulumi:"description"`
	// Deployment environment of the Greenplum cluster. (PRODUCTION, PRESTABLE)
	Environment string `pulumi:"environment"`
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId *string `pulumi:"folderId"`
	// Greenplum cluster config. Detail info in `Greenplum cluster settings` block.
	GreenplumConfig map[string]string `pulumi:"greenplumConfig"`
	// A set of key/value label pairs which assigned to resource.
	Labels  map[string]string           `pulumi:"labels"`
	Logging *MdbGreenplumClusterLogging `pulumi:"logging"`
	// Maintenance policy of the Greenplum cluster.
	MaintenanceWindow *MdbGreenplumClusterMaintenanceWindow `pulumi:"maintenanceWindow"`
	// Number of hosts in master subcluster (1 or 2).
	MasterHostCount int `pulumi:"masterHostCount"`
	// A list of IDs of the host groups to place master subclusters' VMs of the cluster on.
	MasterHostGroupIds []string `pulumi:"masterHostGroupIds"`
	// Settings for master subcluster.
	MasterSubcluster MdbGreenplumClusterMasterSubcluster `pulumi:"masterSubcluster"`
	// The resource name.
	Name *string `pulumi:"name"`
	// The `VPC Network ID` of subnets which resource attached to.
	NetworkId string `pulumi:"networkId"`
	// Configuration of the connection pooler.
	PoolerConfig *MdbGreenplumClusterPoolerConfig `pulumi:"poolerConfig"`
	// Configuration of the PXF daemon.
	PxfConfig *MdbGreenplumClusterPxfConfig `pulumi:"pxfConfig"`
	// The list of security groups applied to resource or their components.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Number of hosts in segment subcluster (from 1 to 32).
	SegmentHostCount int `pulumi:"segmentHostCount"`
	// A list of IDs of the host groups to place segment subclusters' VMs of the cluster on.
	SegmentHostGroupIds []string `pulumi:"segmentHostGroupIds"`
	// Number of segments on segment host (not more then 1 + RAM/8).
	SegmentInHost int `pulumi:"segmentInHost"`
	// Settings for segment subcluster.
	SegmentSubcluster MdbGreenplumClusterSegmentSubcluster `pulumi:"segmentSubcluster"`
	// ID of service account to use with Yandex Cloud resources (e.g. S3, Cloud Logging).
	ServiceAccountId *string `pulumi:"serviceAccountId"`
	// The ID of the subnet, to which the hosts belongs. The subnet must be a part of the network to which the cluster belongs.
	SubnetId string `pulumi:"subnetId"`
	// Greenplum cluster admin user name.
	UserName string `pulumi:"userName"`
	// Greenplum cluster admin password name.
	UserPassword string `pulumi:"userPassword"`
	// Version of the Greenplum cluster. (`6.25`)
	Version string `pulumi:"version"`
	// The [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) where resource is located. If it is not
	// provided, the default provider zone will be used.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a MdbGreenplumCluster resource.
type MdbGreenplumClusterArgs struct {
	// Access policy to the Greenplum cluster.
	Access MdbGreenplumClusterAccessPtrInput
	// Sets whether the master hosts should get a public IP address on creation. Changing this parameter for an existing host
	// is not supported at the moment.
	AssignPublicIp pulumi.BoolInput
	// Background activities settings.
	BackgroundActivities MdbGreenplumClusterBackgroundActivityArrayInput
	// Time to start the daily backup, in the UTC timezone.
	BackupWindowStart MdbGreenplumClusterBackupWindowStartPtrInput
	// Cloud Storage settings of the Greenplum cluster.
	CloudStorage MdbGreenplumClusterCloudStoragePtrInput
	// The `true` value means that resource is protected from accidental deletion.
	DeletionProtection pulumi.BoolPtrInput
	// The resource description.
	Description pulumi.StringPtrInput
	// Deployment environment of the Greenplum cluster. (PRODUCTION, PRESTABLE)
	Environment pulumi.StringInput
	// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
	FolderId pulumi.StringPtrInput
	// Greenplum cluster config. Detail info in `Greenplum cluster settings` block.
	GreenplumConfig pulumi.StringMapInput
	// A set of key/value label pairs which assigned to resource.
	Labels  pulumi.StringMapInput
	Logging MdbGreenplumClusterLoggingPtrInput
	// Maintenance policy of the Greenplum cluster.
	MaintenanceWindow MdbGreenplumClusterMaintenanceWindowPtrInput
	// Number of hosts in master subcluster (1 or 2).
	MasterHostCount pulumi.IntInput
	// A list of IDs of the host groups to place master subclusters' VMs of the cluster on.
	MasterHostGroupIds pulumi.StringArrayInput
	// Settings for master subcluster.
	MasterSubcluster MdbGreenplumClusterMasterSubclusterInput
	// The resource name.
	Name pulumi.StringPtrInput
	// The `VPC Network ID` of subnets which resource attached to.
	NetworkId pulumi.StringInput
	// Configuration of the connection pooler.
	PoolerConfig MdbGreenplumClusterPoolerConfigPtrInput
	// Configuration of the PXF daemon.
	PxfConfig MdbGreenplumClusterPxfConfigPtrInput
	// The list of security groups applied to resource or their components.
	SecurityGroupIds pulumi.StringArrayInput
	// Number of hosts in segment subcluster (from 1 to 32).
	SegmentHostCount pulumi.IntInput
	// A list of IDs of the host groups to place segment subclusters' VMs of the cluster on.
	SegmentHostGroupIds pulumi.StringArrayInput
	// Number of segments on segment host (not more then 1 + RAM/8).
	SegmentInHost pulumi.IntInput
	// Settings for segment subcluster.
	SegmentSubcluster MdbGreenplumClusterSegmentSubclusterInput
	// ID of service account to use with Yandex Cloud resources (e.g. S3, Cloud Logging).
	ServiceAccountId pulumi.StringPtrInput
	// The ID of the subnet, to which the hosts belongs. The subnet must be a part of the network to which the cluster belongs.
	SubnetId pulumi.StringInput
	// Greenplum cluster admin user name.
	UserName pulumi.StringInput
	// Greenplum cluster admin password name.
	UserPassword pulumi.StringInput
	// Version of the Greenplum cluster. (`6.25`)
	Version pulumi.StringInput
	// The [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) where resource is located. If it is not
	// provided, the default provider zone will be used.
	Zone pulumi.StringInput
}

func (MdbGreenplumClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mdbGreenplumClusterArgs)(nil)).Elem()
}

type MdbGreenplumClusterInput interface {
	pulumi.Input

	ToMdbGreenplumClusterOutput() MdbGreenplumClusterOutput
	ToMdbGreenplumClusterOutputWithContext(ctx context.Context) MdbGreenplumClusterOutput
}

func (*MdbGreenplumCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**MdbGreenplumCluster)(nil)).Elem()
}

func (i *MdbGreenplumCluster) ToMdbGreenplumClusterOutput() MdbGreenplumClusterOutput {
	return i.ToMdbGreenplumClusterOutputWithContext(context.Background())
}

func (i *MdbGreenplumCluster) ToMdbGreenplumClusterOutputWithContext(ctx context.Context) MdbGreenplumClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbGreenplumClusterOutput)
}

// MdbGreenplumClusterArrayInput is an input type that accepts MdbGreenplumClusterArray and MdbGreenplumClusterArrayOutput values.
// You can construct a concrete instance of `MdbGreenplumClusterArrayInput` via:
//
//	MdbGreenplumClusterArray{ MdbGreenplumClusterArgs{...} }
type MdbGreenplumClusterArrayInput interface {
	pulumi.Input

	ToMdbGreenplumClusterArrayOutput() MdbGreenplumClusterArrayOutput
	ToMdbGreenplumClusterArrayOutputWithContext(context.Context) MdbGreenplumClusterArrayOutput
}

type MdbGreenplumClusterArray []MdbGreenplumClusterInput

func (MdbGreenplumClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MdbGreenplumCluster)(nil)).Elem()
}

func (i MdbGreenplumClusterArray) ToMdbGreenplumClusterArrayOutput() MdbGreenplumClusterArrayOutput {
	return i.ToMdbGreenplumClusterArrayOutputWithContext(context.Background())
}

func (i MdbGreenplumClusterArray) ToMdbGreenplumClusterArrayOutputWithContext(ctx context.Context) MdbGreenplumClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbGreenplumClusterArrayOutput)
}

// MdbGreenplumClusterMapInput is an input type that accepts MdbGreenplumClusterMap and MdbGreenplumClusterMapOutput values.
// You can construct a concrete instance of `MdbGreenplumClusterMapInput` via:
//
//	MdbGreenplumClusterMap{ "key": MdbGreenplumClusterArgs{...} }
type MdbGreenplumClusterMapInput interface {
	pulumi.Input

	ToMdbGreenplumClusterMapOutput() MdbGreenplumClusterMapOutput
	ToMdbGreenplumClusterMapOutputWithContext(context.Context) MdbGreenplumClusterMapOutput
}

type MdbGreenplumClusterMap map[string]MdbGreenplumClusterInput

func (MdbGreenplumClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MdbGreenplumCluster)(nil)).Elem()
}

func (i MdbGreenplumClusterMap) ToMdbGreenplumClusterMapOutput() MdbGreenplumClusterMapOutput {
	return i.ToMdbGreenplumClusterMapOutputWithContext(context.Background())
}

func (i MdbGreenplumClusterMap) ToMdbGreenplumClusterMapOutputWithContext(ctx context.Context) MdbGreenplumClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MdbGreenplumClusterMapOutput)
}

type MdbGreenplumClusterOutput struct{ *pulumi.OutputState }

func (MdbGreenplumClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MdbGreenplumCluster)(nil)).Elem()
}

func (o MdbGreenplumClusterOutput) ToMdbGreenplumClusterOutput() MdbGreenplumClusterOutput {
	return o
}

func (o MdbGreenplumClusterOutput) ToMdbGreenplumClusterOutputWithContext(ctx context.Context) MdbGreenplumClusterOutput {
	return o
}

// Access policy to the Greenplum cluster.
func (o MdbGreenplumClusterOutput) Access() MdbGreenplumClusterAccessOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) MdbGreenplumClusterAccessOutput { return v.Access }).(MdbGreenplumClusterAccessOutput)
}

// Sets whether the master hosts should get a public IP address on creation. Changing this parameter for an existing host
// is not supported at the moment.
func (o MdbGreenplumClusterOutput) AssignPublicIp() pulumi.BoolOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) pulumi.BoolOutput { return v.AssignPublicIp }).(pulumi.BoolOutput)
}

// Background activities settings.
func (o MdbGreenplumClusterOutput) BackgroundActivities() MdbGreenplumClusterBackgroundActivityArrayOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) MdbGreenplumClusterBackgroundActivityArrayOutput {
		return v.BackgroundActivities
	}).(MdbGreenplumClusterBackgroundActivityArrayOutput)
}

// Time to start the daily backup, in the UTC timezone.
func (o MdbGreenplumClusterOutput) BackupWindowStart() MdbGreenplumClusterBackupWindowStartOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) MdbGreenplumClusterBackupWindowStartOutput { return v.BackupWindowStart }).(MdbGreenplumClusterBackupWindowStartOutput)
}

// Cloud Storage settings of the Greenplum cluster.
func (o MdbGreenplumClusterOutput) CloudStorage() MdbGreenplumClusterCloudStorageOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) MdbGreenplumClusterCloudStorageOutput { return v.CloudStorage }).(MdbGreenplumClusterCloudStorageOutput)
}

// The creation timestamp of the resource.
func (o MdbGreenplumClusterOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The `true` value means that resource is protected from accidental deletion.
func (o MdbGreenplumClusterOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) pulumi.BoolOutput { return v.DeletionProtection }).(pulumi.BoolOutput)
}

// The resource description.
func (o MdbGreenplumClusterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Deployment environment of the Greenplum cluster. (PRODUCTION, PRESTABLE)
func (o MdbGreenplumClusterOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) pulumi.StringOutput { return v.Environment }).(pulumi.StringOutput)
}

// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
func (o MdbGreenplumClusterOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

// Greenplum cluster config. Detail info in `Greenplum cluster settings` block.
func (o MdbGreenplumClusterOutput) GreenplumConfig() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) pulumi.StringMapOutput { return v.GreenplumConfig }).(pulumi.StringMapOutput)
}

// Aggregated health of the cluster.
func (o MdbGreenplumClusterOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) pulumi.StringOutput { return v.Health }).(pulumi.StringOutput)
}

// A set of key/value label pairs which assigned to resource.
func (o MdbGreenplumClusterOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o MdbGreenplumClusterOutput) Logging() MdbGreenplumClusterLoggingOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) MdbGreenplumClusterLoggingOutput { return v.Logging }).(MdbGreenplumClusterLoggingOutput)
}

// Maintenance policy of the Greenplum cluster.
func (o MdbGreenplumClusterOutput) MaintenanceWindow() MdbGreenplumClusterMaintenanceWindowOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) MdbGreenplumClusterMaintenanceWindowOutput { return v.MaintenanceWindow }).(MdbGreenplumClusterMaintenanceWindowOutput)
}

// Number of hosts in master subcluster (1 or 2).
func (o MdbGreenplumClusterOutput) MasterHostCount() pulumi.IntOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) pulumi.IntOutput { return v.MasterHostCount }).(pulumi.IntOutput)
}

// A list of IDs of the host groups to place master subclusters' VMs of the cluster on.
func (o MdbGreenplumClusterOutput) MasterHostGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) pulumi.StringArrayOutput { return v.MasterHostGroupIds }).(pulumi.StringArrayOutput)
}

// Info about hosts in master subcluster.
func (o MdbGreenplumClusterOutput) MasterHosts() MdbGreenplumClusterMasterHostArrayOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) MdbGreenplumClusterMasterHostArrayOutput { return v.MasterHosts }).(MdbGreenplumClusterMasterHostArrayOutput)
}

// Settings for master subcluster.
func (o MdbGreenplumClusterOutput) MasterSubcluster() MdbGreenplumClusterMasterSubclusterOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) MdbGreenplumClusterMasterSubclusterOutput { return v.MasterSubcluster }).(MdbGreenplumClusterMasterSubclusterOutput)
}

// The resource name.
func (o MdbGreenplumClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The `VPC Network ID` of subnets which resource attached to.
func (o MdbGreenplumClusterOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) pulumi.StringOutput { return v.NetworkId }).(pulumi.StringOutput)
}

// Configuration of the connection pooler.
func (o MdbGreenplumClusterOutput) PoolerConfig() MdbGreenplumClusterPoolerConfigOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) MdbGreenplumClusterPoolerConfigOutput { return v.PoolerConfig }).(MdbGreenplumClusterPoolerConfigOutput)
}

// Configuration of the PXF daemon.
func (o MdbGreenplumClusterOutput) PxfConfig() MdbGreenplumClusterPxfConfigOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) MdbGreenplumClusterPxfConfigOutput { return v.PxfConfig }).(MdbGreenplumClusterPxfConfigOutput)
}

// The list of security groups applied to resource or their components.
func (o MdbGreenplumClusterOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// Number of hosts in segment subcluster (from 1 to 32).
func (o MdbGreenplumClusterOutput) SegmentHostCount() pulumi.IntOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) pulumi.IntOutput { return v.SegmentHostCount }).(pulumi.IntOutput)
}

// A list of IDs of the host groups to place segment subclusters' VMs of the cluster on.
func (o MdbGreenplumClusterOutput) SegmentHostGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) pulumi.StringArrayOutput { return v.SegmentHostGroupIds }).(pulumi.StringArrayOutput)
}

// Info about hosts in segment subcluster.
func (o MdbGreenplumClusterOutput) SegmentHosts() MdbGreenplumClusterSegmentHostArrayOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) MdbGreenplumClusterSegmentHostArrayOutput { return v.SegmentHosts }).(MdbGreenplumClusterSegmentHostArrayOutput)
}

// Number of segments on segment host (not more then 1 + RAM/8).
func (o MdbGreenplumClusterOutput) SegmentInHost() pulumi.IntOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) pulumi.IntOutput { return v.SegmentInHost }).(pulumi.IntOutput)
}

// Settings for segment subcluster.
func (o MdbGreenplumClusterOutput) SegmentSubcluster() MdbGreenplumClusterSegmentSubclusterOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) MdbGreenplumClusterSegmentSubclusterOutput { return v.SegmentSubcluster }).(MdbGreenplumClusterSegmentSubclusterOutput)
}

// ID of service account to use with Yandex Cloud resources (e.g. S3, Cloud Logging).
func (o MdbGreenplumClusterOutput) ServiceAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) pulumi.StringPtrOutput { return v.ServiceAccountId }).(pulumi.StringPtrOutput)
}

// Status of the cluster.
func (o MdbGreenplumClusterOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The ID of the subnet, to which the hosts belongs. The subnet must be a part of the network to which the cluster belongs.
func (o MdbGreenplumClusterOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// Greenplum cluster admin user name.
func (o MdbGreenplumClusterOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

// Greenplum cluster admin password name.
func (o MdbGreenplumClusterOutput) UserPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) pulumi.StringOutput { return v.UserPassword }).(pulumi.StringOutput)
}

// Version of the Greenplum cluster. (`6.25`)
func (o MdbGreenplumClusterOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

// The [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) where resource is located. If it is not
// provided, the default provider zone will be used.
func (o MdbGreenplumClusterOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *MdbGreenplumCluster) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type MdbGreenplumClusterArrayOutput struct{ *pulumi.OutputState }

func (MdbGreenplumClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MdbGreenplumCluster)(nil)).Elem()
}

func (o MdbGreenplumClusterArrayOutput) ToMdbGreenplumClusterArrayOutput() MdbGreenplumClusterArrayOutput {
	return o
}

func (o MdbGreenplumClusterArrayOutput) ToMdbGreenplumClusterArrayOutputWithContext(ctx context.Context) MdbGreenplumClusterArrayOutput {
	return o
}

func (o MdbGreenplumClusterArrayOutput) Index(i pulumi.IntInput) MdbGreenplumClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MdbGreenplumCluster {
		return vs[0].([]*MdbGreenplumCluster)[vs[1].(int)]
	}).(MdbGreenplumClusterOutput)
}

type MdbGreenplumClusterMapOutput struct{ *pulumi.OutputState }

func (MdbGreenplumClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MdbGreenplumCluster)(nil)).Elem()
}

func (o MdbGreenplumClusterMapOutput) ToMdbGreenplumClusterMapOutput() MdbGreenplumClusterMapOutput {
	return o
}

func (o MdbGreenplumClusterMapOutput) ToMdbGreenplumClusterMapOutputWithContext(ctx context.Context) MdbGreenplumClusterMapOutput {
	return o
}

func (o MdbGreenplumClusterMapOutput) MapIndex(k pulumi.StringInput) MdbGreenplumClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MdbGreenplumCluster {
		return vs[0].(map[string]*MdbGreenplumCluster)[vs[1].(string)]
	}).(MdbGreenplumClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MdbGreenplumClusterInput)(nil)).Elem(), &MdbGreenplumCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*MdbGreenplumClusterArrayInput)(nil)).Elem(), MdbGreenplumClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MdbGreenplumClusterMapInput)(nil)).Elem(), MdbGreenplumClusterMap{})
	pulumi.RegisterOutputType(MdbGreenplumClusterOutput{})
	pulumi.RegisterOutputType(MdbGreenplumClusterArrayOutput{})
	pulumi.RegisterOutputType(MdbGreenplumClusterMapOutput{})
}
