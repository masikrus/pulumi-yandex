// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    [YandexResourceType("yandex:index/mdbGreenplumCluster:MdbGreenplumCluster")]
    public partial class MdbGreenplumCluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Access policy to the Greenplum cluster.
        /// </summary>
        [Output("access")]
        public Output<Outputs.MdbGreenplumClusterAccess> Access { get; private set; } = null!;

        /// <summary>
        /// Sets whether the master hosts should get a public IP address on creation. Changing this parameter for an existing host
        /// is not supported at the moment.
        /// </summary>
        [Output("assignPublicIp")]
        public Output<bool> AssignPublicIp { get; private set; } = null!;

        /// <summary>
        /// Background activities settings.
        /// </summary>
        [Output("backgroundActivities")]
        public Output<ImmutableArray<Outputs.MdbGreenplumClusterBackgroundActivity>> BackgroundActivities { get; private set; } = null!;

        /// <summary>
        /// Time to start the daily backup, in the UTC timezone.
        /// </summary>
        [Output("backupWindowStart")]
        public Output<Outputs.MdbGreenplumClusterBackupWindowStart> BackupWindowStart { get; private set; } = null!;

        /// <summary>
        /// Cloud Storage settings of the Greenplum cluster.
        /// </summary>
        [Output("cloudStorage")]
        public Output<Outputs.MdbGreenplumClusterCloudStorage> CloudStorage { get; private set; } = null!;

        /// <summary>
        /// The creation timestamp of the resource.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The `true` value means that resource is protected from accidental deletion.
        /// </summary>
        [Output("deletionProtection")]
        public Output<bool> DeletionProtection { get; private set; } = null!;

        /// <summary>
        /// The resource description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Deployment environment of the Greenplum cluster. (PRODUCTION, PRESTABLE)
        /// </summary>
        [Output("environment")]
        public Output<string> Environment { get; private set; } = null!;

        /// <summary>
        /// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        /// </summary>
        [Output("folderId")]
        public Output<string> FolderId { get; private set; } = null!;

        /// <summary>
        /// Greenplum cluster config. Detail info in `Greenplum cluster settings` block.
        /// </summary>
        [Output("greenplumConfig")]
        public Output<ImmutableDictionary<string, string>> GreenplumConfig { get; private set; } = null!;

        /// <summary>
        /// Aggregated health of the cluster.
        /// </summary>
        [Output("health")]
        public Output<string> Health { get; private set; } = null!;

        /// <summary>
        /// A set of key/value label pairs which assigned to resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        [Output("logging")]
        public Output<Outputs.MdbGreenplumClusterLogging> Logging { get; private set; } = null!;

        /// <summary>
        /// Maintenance policy of the Greenplum cluster.
        /// </summary>
        [Output("maintenanceWindow")]
        public Output<Outputs.MdbGreenplumClusterMaintenanceWindow> MaintenanceWindow { get; private set; } = null!;

        /// <summary>
        /// Number of hosts in master subcluster (1 or 2).
        /// </summary>
        [Output("masterHostCount")]
        public Output<int> MasterHostCount { get; private set; } = null!;

        /// <summary>
        /// A list of IDs of the host groups to place master subclusters' VMs of the cluster on.
        /// </summary>
        [Output("masterHostGroupIds")]
        public Output<ImmutableArray<string>> MasterHostGroupIds { get; private set; } = null!;

        /// <summary>
        /// Info about hosts in master subcluster.
        /// </summary>
        [Output("masterHosts")]
        public Output<ImmutableArray<Outputs.MdbGreenplumClusterMasterHost>> MasterHosts { get; private set; } = null!;

        /// <summary>
        /// Settings for master subcluster.
        /// </summary>
        [Output("masterSubcluster")]
        public Output<Outputs.MdbGreenplumClusterMasterSubcluster> MasterSubcluster { get; private set; } = null!;

        /// <summary>
        /// The resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The `VPC Network ID` of subnets which resource attached to.
        /// </summary>
        [Output("networkId")]
        public Output<string> NetworkId { get; private set; } = null!;

        /// <summary>
        /// Configuration of the connection pooler.
        /// </summary>
        [Output("poolerConfig")]
        public Output<Outputs.MdbGreenplumClusterPoolerConfig> PoolerConfig { get; private set; } = null!;

        /// <summary>
        /// Configuration of the PXF daemon.
        /// </summary>
        [Output("pxfConfig")]
        public Output<Outputs.MdbGreenplumClusterPxfConfig> PxfConfig { get; private set; } = null!;

        /// <summary>
        /// The list of security groups applied to resource or their components.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// Number of hosts in segment subcluster (from 1 to 32).
        /// </summary>
        [Output("segmentHostCount")]
        public Output<int> SegmentHostCount { get; private set; } = null!;

        /// <summary>
        /// A list of IDs of the host groups to place segment subclusters' VMs of the cluster on.
        /// </summary>
        [Output("segmentHostGroupIds")]
        public Output<ImmutableArray<string>> SegmentHostGroupIds { get; private set; } = null!;

        /// <summary>
        /// Info about hosts in segment subcluster.
        /// </summary>
        [Output("segmentHosts")]
        public Output<ImmutableArray<Outputs.MdbGreenplumClusterSegmentHost>> SegmentHosts { get; private set; } = null!;

        /// <summary>
        /// Number of segments on segment host (not more then 1 + RAM/8).
        /// </summary>
        [Output("segmentInHost")]
        public Output<int> SegmentInHost { get; private set; } = null!;

        /// <summary>
        /// Settings for segment subcluster.
        /// </summary>
        [Output("segmentSubcluster")]
        public Output<Outputs.MdbGreenplumClusterSegmentSubcluster> SegmentSubcluster { get; private set; } = null!;

        /// <summary>
        /// ID of service account to use with Yandex Cloud resources (e.g. S3, Cloud Logging).
        /// </summary>
        [Output("serviceAccountId")]
        public Output<string?> ServiceAccountId { get; private set; } = null!;

        /// <summary>
        /// Status of the cluster.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The ID of the subnet, to which the hosts belongs. The subnet must be a part of the network to which the cluster belongs.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// Greenplum cluster admin user name.
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;

        /// <summary>
        /// Greenplum cluster admin password name.
        /// </summary>
        [Output("userPassword")]
        public Output<string> UserPassword { get; private set; } = null!;

        /// <summary>
        /// Version of the Greenplum cluster. (`6.25`)
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// The [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) where resource is located. If it is not
        /// provided, the default provider zone will be used.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a MdbGreenplumCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MdbGreenplumCluster(string name, MdbGreenplumClusterArgs args, CustomResourceOptions? options = null)
            : base("yandex:index/mdbGreenplumCluster:MdbGreenplumCluster", name, args ?? new MdbGreenplumClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MdbGreenplumCluster(string name, Input<string> id, MdbGreenplumClusterState? state = null, CustomResourceOptions? options = null)
            : base("yandex:index/mdbGreenplumCluster:MdbGreenplumCluster", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "userPassword",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MdbGreenplumCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MdbGreenplumCluster Get(string name, Input<string> id, MdbGreenplumClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new MdbGreenplumCluster(name, id, state, options);
        }
    }

    public sealed class MdbGreenplumClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access policy to the Greenplum cluster.
        /// </summary>
        [Input("access")]
        public Input<Inputs.MdbGreenplumClusterAccessArgs>? Access { get; set; }

        /// <summary>
        /// Sets whether the master hosts should get a public IP address on creation. Changing this parameter for an existing host
        /// is not supported at the moment.
        /// </summary>
        [Input("assignPublicIp", required: true)]
        public Input<bool> AssignPublicIp { get; set; } = null!;

        [Input("backgroundActivities")]
        private InputList<Inputs.MdbGreenplumClusterBackgroundActivityArgs>? _backgroundActivities;

        /// <summary>
        /// Background activities settings.
        /// </summary>
        public InputList<Inputs.MdbGreenplumClusterBackgroundActivityArgs> BackgroundActivities
        {
            get => _backgroundActivities ?? (_backgroundActivities = new InputList<Inputs.MdbGreenplumClusterBackgroundActivityArgs>());
            set => _backgroundActivities = value;
        }

        /// <summary>
        /// Time to start the daily backup, in the UTC timezone.
        /// </summary>
        [Input("backupWindowStart")]
        public Input<Inputs.MdbGreenplumClusterBackupWindowStartArgs>? BackupWindowStart { get; set; }

        /// <summary>
        /// Cloud Storage settings of the Greenplum cluster.
        /// </summary>
        [Input("cloudStorage")]
        public Input<Inputs.MdbGreenplumClusterCloudStorageArgs>? CloudStorage { get; set; }

        /// <summary>
        /// The `true` value means that resource is protected from accidental deletion.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// The resource description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Deployment environment of the Greenplum cluster. (PRODUCTION, PRESTABLE)
        /// </summary>
        [Input("environment", required: true)]
        public Input<string> Environment { get; set; } = null!;

        /// <summary>
        /// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("greenplumConfig")]
        private InputMap<string>? _greenplumConfig;

        /// <summary>
        /// Greenplum cluster config. Detail info in `Greenplum cluster settings` block.
        /// </summary>
        public InputMap<string> GreenplumConfig
        {
            get => _greenplumConfig ?? (_greenplumConfig = new InputMap<string>());
            set => _greenplumConfig = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs which assigned to resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("logging")]
        public Input<Inputs.MdbGreenplumClusterLoggingArgs>? Logging { get; set; }

        /// <summary>
        /// Maintenance policy of the Greenplum cluster.
        /// </summary>
        [Input("maintenanceWindow")]
        public Input<Inputs.MdbGreenplumClusterMaintenanceWindowArgs>? MaintenanceWindow { get; set; }

        /// <summary>
        /// Number of hosts in master subcluster (1 or 2).
        /// </summary>
        [Input("masterHostCount", required: true)]
        public Input<int> MasterHostCount { get; set; } = null!;

        [Input("masterHostGroupIds")]
        private InputList<string>? _masterHostGroupIds;

        /// <summary>
        /// A list of IDs of the host groups to place master subclusters' VMs of the cluster on.
        /// </summary>
        public InputList<string> MasterHostGroupIds
        {
            get => _masterHostGroupIds ?? (_masterHostGroupIds = new InputList<string>());
            set => _masterHostGroupIds = value;
        }

        /// <summary>
        /// Settings for master subcluster.
        /// </summary>
        [Input("masterSubcluster", required: true)]
        public Input<Inputs.MdbGreenplumClusterMasterSubclusterArgs> MasterSubcluster { get; set; } = null!;

        /// <summary>
        /// The resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The `VPC Network ID` of subnets which resource attached to.
        /// </summary>
        [Input("networkId", required: true)]
        public Input<string> NetworkId { get; set; } = null!;

        /// <summary>
        /// Configuration of the connection pooler.
        /// </summary>
        [Input("poolerConfig")]
        public Input<Inputs.MdbGreenplumClusterPoolerConfigArgs>? PoolerConfig { get; set; }

        /// <summary>
        /// Configuration of the PXF daemon.
        /// </summary>
        [Input("pxfConfig")]
        public Input<Inputs.MdbGreenplumClusterPxfConfigArgs>? PxfConfig { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The list of security groups applied to resource or their components.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Number of hosts in segment subcluster (from 1 to 32).
        /// </summary>
        [Input("segmentHostCount", required: true)]
        public Input<int> SegmentHostCount { get; set; } = null!;

        [Input("segmentHostGroupIds")]
        private InputList<string>? _segmentHostGroupIds;

        /// <summary>
        /// A list of IDs of the host groups to place segment subclusters' VMs of the cluster on.
        /// </summary>
        public InputList<string> SegmentHostGroupIds
        {
            get => _segmentHostGroupIds ?? (_segmentHostGroupIds = new InputList<string>());
            set => _segmentHostGroupIds = value;
        }

        /// <summary>
        /// Number of segments on segment host (not more then 1 + RAM/8).
        /// </summary>
        [Input("segmentInHost", required: true)]
        public Input<int> SegmentInHost { get; set; } = null!;

        /// <summary>
        /// Settings for segment subcluster.
        /// </summary>
        [Input("segmentSubcluster", required: true)]
        public Input<Inputs.MdbGreenplumClusterSegmentSubclusterArgs> SegmentSubcluster { get; set; } = null!;

        /// <summary>
        /// ID of service account to use with Yandex Cloud resources (e.g. S3, Cloud Logging).
        /// </summary>
        [Input("serviceAccountId")]
        public Input<string>? ServiceAccountId { get; set; }

        /// <summary>
        /// The ID of the subnet, to which the hosts belongs. The subnet must be a part of the network to which the cluster belongs.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        /// <summary>
        /// Greenplum cluster admin user name.
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        [Input("userPassword", required: true)]
        private Input<string>? _userPassword;

        /// <summary>
        /// Greenplum cluster admin password name.
        /// </summary>
        public Input<string>? UserPassword
        {
            get => _userPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _userPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Version of the Greenplum cluster. (`6.25`)
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        /// <summary>
        /// The [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) where resource is located. If it is not
        /// provided, the default provider zone will be used.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public MdbGreenplumClusterArgs()
        {
        }
        public static new MdbGreenplumClusterArgs Empty => new MdbGreenplumClusterArgs();
    }

    public sealed class MdbGreenplumClusterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access policy to the Greenplum cluster.
        /// </summary>
        [Input("access")]
        public Input<Inputs.MdbGreenplumClusterAccessGetArgs>? Access { get; set; }

        /// <summary>
        /// Sets whether the master hosts should get a public IP address on creation. Changing this parameter for an existing host
        /// is not supported at the moment.
        /// </summary>
        [Input("assignPublicIp")]
        public Input<bool>? AssignPublicIp { get; set; }

        [Input("backgroundActivities")]
        private InputList<Inputs.MdbGreenplumClusterBackgroundActivityGetArgs>? _backgroundActivities;

        /// <summary>
        /// Background activities settings.
        /// </summary>
        public InputList<Inputs.MdbGreenplumClusterBackgroundActivityGetArgs> BackgroundActivities
        {
            get => _backgroundActivities ?? (_backgroundActivities = new InputList<Inputs.MdbGreenplumClusterBackgroundActivityGetArgs>());
            set => _backgroundActivities = value;
        }

        /// <summary>
        /// Time to start the daily backup, in the UTC timezone.
        /// </summary>
        [Input("backupWindowStart")]
        public Input<Inputs.MdbGreenplumClusterBackupWindowStartGetArgs>? BackupWindowStart { get; set; }

        /// <summary>
        /// Cloud Storage settings of the Greenplum cluster.
        /// </summary>
        [Input("cloudStorage")]
        public Input<Inputs.MdbGreenplumClusterCloudStorageGetArgs>? CloudStorage { get; set; }

        /// <summary>
        /// The creation timestamp of the resource.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The `true` value means that resource is protected from accidental deletion.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// The resource description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Deployment environment of the Greenplum cluster. (PRODUCTION, PRESTABLE)
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("greenplumConfig")]
        private InputMap<string>? _greenplumConfig;

        /// <summary>
        /// Greenplum cluster config. Detail info in `Greenplum cluster settings` block.
        /// </summary>
        public InputMap<string> GreenplumConfig
        {
            get => _greenplumConfig ?? (_greenplumConfig = new InputMap<string>());
            set => _greenplumConfig = value;
        }

        /// <summary>
        /// Aggregated health of the cluster.
        /// </summary>
        [Input("health")]
        public Input<string>? Health { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs which assigned to resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("logging")]
        public Input<Inputs.MdbGreenplumClusterLoggingGetArgs>? Logging { get; set; }

        /// <summary>
        /// Maintenance policy of the Greenplum cluster.
        /// </summary>
        [Input("maintenanceWindow")]
        public Input<Inputs.MdbGreenplumClusterMaintenanceWindowGetArgs>? MaintenanceWindow { get; set; }

        /// <summary>
        /// Number of hosts in master subcluster (1 or 2).
        /// </summary>
        [Input("masterHostCount")]
        public Input<int>? MasterHostCount { get; set; }

        [Input("masterHostGroupIds")]
        private InputList<string>? _masterHostGroupIds;

        /// <summary>
        /// A list of IDs of the host groups to place master subclusters' VMs of the cluster on.
        /// </summary>
        public InputList<string> MasterHostGroupIds
        {
            get => _masterHostGroupIds ?? (_masterHostGroupIds = new InputList<string>());
            set => _masterHostGroupIds = value;
        }

        [Input("masterHosts")]
        private InputList<Inputs.MdbGreenplumClusterMasterHostGetArgs>? _masterHosts;

        /// <summary>
        /// Info about hosts in master subcluster.
        /// </summary>
        public InputList<Inputs.MdbGreenplumClusterMasterHostGetArgs> MasterHosts
        {
            get => _masterHosts ?? (_masterHosts = new InputList<Inputs.MdbGreenplumClusterMasterHostGetArgs>());
            set => _masterHosts = value;
        }

        /// <summary>
        /// Settings for master subcluster.
        /// </summary>
        [Input("masterSubcluster")]
        public Input<Inputs.MdbGreenplumClusterMasterSubclusterGetArgs>? MasterSubcluster { get; set; }

        /// <summary>
        /// The resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The `VPC Network ID` of subnets which resource attached to.
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        /// <summary>
        /// Configuration of the connection pooler.
        /// </summary>
        [Input("poolerConfig")]
        public Input<Inputs.MdbGreenplumClusterPoolerConfigGetArgs>? PoolerConfig { get; set; }

        /// <summary>
        /// Configuration of the PXF daemon.
        /// </summary>
        [Input("pxfConfig")]
        public Input<Inputs.MdbGreenplumClusterPxfConfigGetArgs>? PxfConfig { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The list of security groups applied to resource or their components.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Number of hosts in segment subcluster (from 1 to 32).
        /// </summary>
        [Input("segmentHostCount")]
        public Input<int>? SegmentHostCount { get; set; }

        [Input("segmentHostGroupIds")]
        private InputList<string>? _segmentHostGroupIds;

        /// <summary>
        /// A list of IDs of the host groups to place segment subclusters' VMs of the cluster on.
        /// </summary>
        public InputList<string> SegmentHostGroupIds
        {
            get => _segmentHostGroupIds ?? (_segmentHostGroupIds = new InputList<string>());
            set => _segmentHostGroupIds = value;
        }

        [Input("segmentHosts")]
        private InputList<Inputs.MdbGreenplumClusterSegmentHostGetArgs>? _segmentHosts;

        /// <summary>
        /// Info about hosts in segment subcluster.
        /// </summary>
        public InputList<Inputs.MdbGreenplumClusterSegmentHostGetArgs> SegmentHosts
        {
            get => _segmentHosts ?? (_segmentHosts = new InputList<Inputs.MdbGreenplumClusterSegmentHostGetArgs>());
            set => _segmentHosts = value;
        }

        /// <summary>
        /// Number of segments on segment host (not more then 1 + RAM/8).
        /// </summary>
        [Input("segmentInHost")]
        public Input<int>? SegmentInHost { get; set; }

        /// <summary>
        /// Settings for segment subcluster.
        /// </summary>
        [Input("segmentSubcluster")]
        public Input<Inputs.MdbGreenplumClusterSegmentSubclusterGetArgs>? SegmentSubcluster { get; set; }

        /// <summary>
        /// ID of service account to use with Yandex Cloud resources (e.g. S3, Cloud Logging).
        /// </summary>
        [Input("serviceAccountId")]
        public Input<string>? ServiceAccountId { get; set; }

        /// <summary>
        /// Status of the cluster.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The ID of the subnet, to which the hosts belongs. The subnet must be a part of the network to which the cluster belongs.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// Greenplum cluster admin user name.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        [Input("userPassword")]
        private Input<string>? _userPassword;

        /// <summary>
        /// Greenplum cluster admin password name.
        /// </summary>
        public Input<string>? UserPassword
        {
            get => _userPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _userPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Version of the Greenplum cluster. (`6.25`)
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) where resource is located. If it is not
        /// provided, the default provider zone will be used.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public MdbGreenplumClusterState()
        {
        }
        public static new MdbGreenplumClusterState Empty => new MdbGreenplumClusterState();
    }
}
