// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetMdbRedisCluster
    {
        public static Task<GetMdbRedisClusterResult> InvokeAsync(GetMdbRedisClusterArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMdbRedisClusterResult>("yandex:index/getMdbRedisCluster:getMdbRedisCluster", args ?? new GetMdbRedisClusterArgs(), options.WithDefaults());

        public static Output<GetMdbRedisClusterResult> Invoke(GetMdbRedisClusterInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMdbRedisClusterResult>("yandex:index/getMdbRedisCluster:getMdbRedisCluster", args ?? new GetMdbRedisClusterInvokeArgs(), options.WithDefaults());

        public static Output<GetMdbRedisClusterResult> Invoke(GetMdbRedisClusterInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMdbRedisClusterResult>("yandex:index/getMdbRedisCluster:getMdbRedisCluster", args ?? new GetMdbRedisClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMdbRedisClusterArgs : global::Pulumi.InvokeArgs
    {
        [Input("clusterId")]
        public string? ClusterId { get; set; }

        [Input("deletionProtection")]
        public bool? DeletionProtection { get; set; }

        [Input("folderId")]
        public string? FolderId { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        public GetMdbRedisClusterArgs()
        {
        }
        public static new GetMdbRedisClusterArgs Empty => new GetMdbRedisClusterArgs();
    }

    public sealed class GetMdbRedisClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetMdbRedisClusterInvokeArgs()
        {
        }
        public static new GetMdbRedisClusterInvokeArgs Empty => new GetMdbRedisClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetMdbRedisClusterResult
    {
        public readonly bool AnnounceHostnames;
        public readonly bool AuthSentinel;
        public readonly string ClusterId;
        public readonly ImmutableArray<Outputs.GetMdbRedisClusterConfigResult> Configs;
        public readonly string CreatedAt;
        public readonly bool DeletionProtection;
        public readonly string Description;
        public readonly ImmutableArray<Outputs.GetMdbRedisClusterDiskSizeAutoscalingResult> DiskSizeAutoscalings;
        public readonly string Environment;
        public readonly string FolderId;
        public readonly string Health;
        public readonly ImmutableArray<Outputs.GetMdbRedisClusterHostResult> Hosts;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly ImmutableArray<Outputs.GetMdbRedisClusterMaintenanceWindowResult> MaintenanceWindows;
        public readonly string Name;
        public readonly string NetworkId;
        public readonly string PersistenceMode;
        public readonly ImmutableArray<Outputs.GetMdbRedisClusterResourceResult> Resources;
        public readonly ImmutableArray<string> SecurityGroupIds;
        public readonly bool Sharded;
        public readonly string Status;
        public readonly bool TlsEnabled;

        [OutputConstructor]
        private GetMdbRedisClusterResult(
            bool announceHostnames,

            bool authSentinel,

            string clusterId,

            ImmutableArray<Outputs.GetMdbRedisClusterConfigResult> configs,

            string createdAt,

            bool deletionProtection,

            string description,

            ImmutableArray<Outputs.GetMdbRedisClusterDiskSizeAutoscalingResult> diskSizeAutoscalings,

            string environment,

            string folderId,

            string health,

            ImmutableArray<Outputs.GetMdbRedisClusterHostResult> hosts,

            string id,

            ImmutableDictionary<string, string> labels,

            ImmutableArray<Outputs.GetMdbRedisClusterMaintenanceWindowResult> maintenanceWindows,

            string name,

            string networkId,

            string persistenceMode,

            ImmutableArray<Outputs.GetMdbRedisClusterResourceResult> resources,

            ImmutableArray<string> securityGroupIds,

            bool sharded,

            string status,

            bool tlsEnabled)
        {
            AnnounceHostnames = announceHostnames;
            AuthSentinel = authSentinel;
            ClusterId = clusterId;
            Configs = configs;
            CreatedAt = createdAt;
            DeletionProtection = deletionProtection;
            Description = description;
            DiskSizeAutoscalings = diskSizeAutoscalings;
            Environment = environment;
            FolderId = folderId;
            Health = health;
            Hosts = hosts;
            Id = id;
            Labels = labels;
            MaintenanceWindows = maintenanceWindows;
            Name = name;
            NetworkId = networkId;
            PersistenceMode = persistenceMode;
            Resources = resources;
            SecurityGroupIds = securityGroupIds;
            Sharded = sharded;
            Status = status;
            TlsEnabled = tlsEnabled;
        }
    }
}
