// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetMdbGreenplumCluster
    {
        public static Task<GetMdbGreenplumClusterResult> InvokeAsync(GetMdbGreenplumClusterArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMdbGreenplumClusterResult>("yandex:index/getMdbGreenplumCluster:getMdbGreenplumCluster", args ?? new GetMdbGreenplumClusterArgs(), options.WithDefaults());

        public static Output<GetMdbGreenplumClusterResult> Invoke(GetMdbGreenplumClusterInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMdbGreenplumClusterResult>("yandex:index/getMdbGreenplumCluster:getMdbGreenplumCluster", args ?? new GetMdbGreenplumClusterInvokeArgs(), options.WithDefaults());

        public static Output<GetMdbGreenplumClusterResult> Invoke(GetMdbGreenplumClusterInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMdbGreenplumClusterResult>("yandex:index/getMdbGreenplumCluster:getMdbGreenplumCluster", args ?? new GetMdbGreenplumClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMdbGreenplumClusterArgs : global::Pulumi.InvokeArgs
    {
        [Input("backgroundActivities")]
        private List<Inputs.GetMdbGreenplumClusterBackgroundActivityArgs>? _backgroundActivities;
        public List<Inputs.GetMdbGreenplumClusterBackgroundActivityArgs> BackgroundActivities
        {
            get => _backgroundActivities ?? (_backgroundActivities = new List<Inputs.GetMdbGreenplumClusterBackgroundActivityArgs>());
            set => _backgroundActivities = value;
        }

        [Input("clusterId")]
        public string? ClusterId { get; set; }

        [Input("folderId")]
        public string? FolderId { get; set; }

        [Input("greenplumConfig")]
        private Dictionary<string, string>? _greenplumConfig;
        public Dictionary<string, string> GreenplumConfig
        {
            get => _greenplumConfig ?? (_greenplumConfig = new Dictionary<string, string>());
            set => _greenplumConfig = value;
        }

        [Input("masterHostGroupIds")]
        private List<string>? _masterHostGroupIds;
        public List<string> MasterHostGroupIds
        {
            get => _masterHostGroupIds ?? (_masterHostGroupIds = new List<string>());
            set => _masterHostGroupIds = value;
        }

        [Input("name")]
        public string? Name { get; set; }

        [Input("poolerConfig")]
        public Inputs.GetMdbGreenplumClusterPoolerConfigArgs? PoolerConfig { get; set; }

        [Input("pxfConfigs")]
        private List<Inputs.GetMdbGreenplumClusterPxfConfigArgs>? _pxfConfigs;
        public List<Inputs.GetMdbGreenplumClusterPxfConfigArgs> PxfConfigs
        {
            get => _pxfConfigs ?? (_pxfConfigs = new List<Inputs.GetMdbGreenplumClusterPxfConfigArgs>());
            set => _pxfConfigs = value;
        }

        [Input("segmentHostGroupIds")]
        private List<string>? _segmentHostGroupIds;
        public List<string> SegmentHostGroupIds
        {
            get => _segmentHostGroupIds ?? (_segmentHostGroupIds = new List<string>());
            set => _segmentHostGroupIds = value;
        }

        public GetMdbGreenplumClusterArgs()
        {
        }
        public static new GetMdbGreenplumClusterArgs Empty => new GetMdbGreenplumClusterArgs();
    }

    public sealed class GetMdbGreenplumClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("backgroundActivities")]
        private InputList<Inputs.GetMdbGreenplumClusterBackgroundActivityInputArgs>? _backgroundActivities;
        public InputList<Inputs.GetMdbGreenplumClusterBackgroundActivityInputArgs> BackgroundActivities
        {
            get => _backgroundActivities ?? (_backgroundActivities = new InputList<Inputs.GetMdbGreenplumClusterBackgroundActivityInputArgs>());
            set => _backgroundActivities = value;
        }

        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("greenplumConfig")]
        private InputMap<string>? _greenplumConfig;
        public InputMap<string> GreenplumConfig
        {
            get => _greenplumConfig ?? (_greenplumConfig = new InputMap<string>());
            set => _greenplumConfig = value;
        }

        [Input("masterHostGroupIds")]
        private InputList<string>? _masterHostGroupIds;
        public InputList<string> MasterHostGroupIds
        {
            get => _masterHostGroupIds ?? (_masterHostGroupIds = new InputList<string>());
            set => _masterHostGroupIds = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("poolerConfig")]
        public Input<Inputs.GetMdbGreenplumClusterPoolerConfigInputArgs>? PoolerConfig { get; set; }

        [Input("pxfConfigs")]
        private InputList<Inputs.GetMdbGreenplumClusterPxfConfigInputArgs>? _pxfConfigs;
        public InputList<Inputs.GetMdbGreenplumClusterPxfConfigInputArgs> PxfConfigs
        {
            get => _pxfConfigs ?? (_pxfConfigs = new InputList<Inputs.GetMdbGreenplumClusterPxfConfigInputArgs>());
            set => _pxfConfigs = value;
        }

        [Input("segmentHostGroupIds")]
        private InputList<string>? _segmentHostGroupIds;
        public InputList<string> SegmentHostGroupIds
        {
            get => _segmentHostGroupIds ?? (_segmentHostGroupIds = new InputList<string>());
            set => _segmentHostGroupIds = value;
        }

        public GetMdbGreenplumClusterInvokeArgs()
        {
        }
        public static new GetMdbGreenplumClusterInvokeArgs Empty => new GetMdbGreenplumClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetMdbGreenplumClusterResult
    {
        public readonly ImmutableArray<Outputs.GetMdbGreenplumClusterAccessResult> Accesses;
        public readonly bool AssignPublicIp;
        public readonly ImmutableArray<Outputs.GetMdbGreenplumClusterBackgroundActivityResult> BackgroundActivities;
        public readonly ImmutableArray<Outputs.GetMdbGreenplumClusterBackupWindowStartResult> BackupWindowStarts;
        public readonly ImmutableArray<Outputs.GetMdbGreenplumClusterCloudStorageResult> CloudStorages;
        public readonly string ClusterId;
        public readonly string CreatedAt;
        public readonly bool DeletionProtection;
        public readonly string Description;
        public readonly string Environment;
        public readonly string FolderId;
        public readonly ImmutableDictionary<string, string> GreenplumConfig;
        public readonly string Health;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly ImmutableArray<Outputs.GetMdbGreenplumClusterLoggingResult> Loggings;
        public readonly ImmutableArray<Outputs.GetMdbGreenplumClusterMaintenanceWindowResult> MaintenanceWindows;
        public readonly int MasterHostCount;
        public readonly ImmutableArray<string> MasterHostGroupIds;
        public readonly ImmutableArray<Outputs.GetMdbGreenplumClusterMasterHostResult> MasterHosts;
        public readonly ImmutableArray<Outputs.GetMdbGreenplumClusterMasterSubclusterResult> MasterSubclusters;
        public readonly string Name;
        public readonly string NetworkId;
        public readonly Outputs.GetMdbGreenplumClusterPoolerConfigResult? PoolerConfig;
        public readonly ImmutableArray<Outputs.GetMdbGreenplumClusterPxfConfigResult> PxfConfigs;
        public readonly ImmutableArray<string> SecurityGroupIds;
        public readonly int SegmentHostCount;
        public readonly ImmutableArray<string> SegmentHostGroupIds;
        public readonly ImmutableArray<Outputs.GetMdbGreenplumClusterSegmentHostResult> SegmentHosts;
        public readonly int SegmentInHost;
        public readonly ImmutableArray<Outputs.GetMdbGreenplumClusterSegmentSubclusterResult> SegmentSubclusters;
        public readonly string ServiceAccountId;
        public readonly string Status;
        public readonly string SubnetId;
        public readonly string UserName;
        public readonly string Version;
        public readonly string Zone;

        [OutputConstructor]
        private GetMdbGreenplumClusterResult(
            ImmutableArray<Outputs.GetMdbGreenplumClusterAccessResult> accesses,

            bool assignPublicIp,

            ImmutableArray<Outputs.GetMdbGreenplumClusterBackgroundActivityResult> backgroundActivities,

            ImmutableArray<Outputs.GetMdbGreenplumClusterBackupWindowStartResult> backupWindowStarts,

            ImmutableArray<Outputs.GetMdbGreenplumClusterCloudStorageResult> cloudStorages,

            string clusterId,

            string createdAt,

            bool deletionProtection,

            string description,

            string environment,

            string folderId,

            ImmutableDictionary<string, string> greenplumConfig,

            string health,

            string id,

            ImmutableDictionary<string, string> labels,

            ImmutableArray<Outputs.GetMdbGreenplumClusterLoggingResult> loggings,

            ImmutableArray<Outputs.GetMdbGreenplumClusterMaintenanceWindowResult> maintenanceWindows,

            int masterHostCount,

            ImmutableArray<string> masterHostGroupIds,

            ImmutableArray<Outputs.GetMdbGreenplumClusterMasterHostResult> masterHosts,

            ImmutableArray<Outputs.GetMdbGreenplumClusterMasterSubclusterResult> masterSubclusters,

            string name,

            string networkId,

            Outputs.GetMdbGreenplumClusterPoolerConfigResult? poolerConfig,

            ImmutableArray<Outputs.GetMdbGreenplumClusterPxfConfigResult> pxfConfigs,

            ImmutableArray<string> securityGroupIds,

            int segmentHostCount,

            ImmutableArray<string> segmentHostGroupIds,

            ImmutableArray<Outputs.GetMdbGreenplumClusterSegmentHostResult> segmentHosts,

            int segmentInHost,

            ImmutableArray<Outputs.GetMdbGreenplumClusterSegmentSubclusterResult> segmentSubclusters,

            string serviceAccountId,

            string status,

            string subnetId,

            string userName,

            string version,

            string zone)
        {
            Accesses = accesses;
            AssignPublicIp = assignPublicIp;
            BackgroundActivities = backgroundActivities;
            BackupWindowStarts = backupWindowStarts;
            CloudStorages = cloudStorages;
            ClusterId = clusterId;
            CreatedAt = createdAt;
            DeletionProtection = deletionProtection;
            Description = description;
            Environment = environment;
            FolderId = folderId;
            GreenplumConfig = greenplumConfig;
            Health = health;
            Id = id;
            Labels = labels;
            Loggings = loggings;
            MaintenanceWindows = maintenanceWindows;
            MasterHostCount = masterHostCount;
            MasterHostGroupIds = masterHostGroupIds;
            MasterHosts = masterHosts;
            MasterSubclusters = masterSubclusters;
            Name = name;
            NetworkId = networkId;
            PoolerConfig = poolerConfig;
            PxfConfigs = pxfConfigs;
            SecurityGroupIds = securityGroupIds;
            SegmentHostCount = segmentHostCount;
            SegmentHostGroupIds = segmentHostGroupIds;
            SegmentHosts = segmentHosts;
            SegmentInHost = segmentInHost;
            SegmentSubclusters = segmentSubclusters;
            ServiceAccountId = serviceAccountId;
            Status = status;
            SubnetId = subnetId;
            UserName = userName;
            Version = version;
            Zone = zone;
        }
    }
}
