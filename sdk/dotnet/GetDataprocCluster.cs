// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetDataprocCluster
    {
        public static Task<GetDataprocClusterResult> InvokeAsync(GetDataprocClusterArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDataprocClusterResult>("yandex:index/getDataprocCluster:getDataprocCluster", args ?? new GetDataprocClusterArgs(), options.WithDefaults());

        public static Output<GetDataprocClusterResult> Invoke(GetDataprocClusterInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDataprocClusterResult>("yandex:index/getDataprocCluster:getDataprocCluster", args ?? new GetDataprocClusterInvokeArgs(), options.WithDefaults());

        public static Output<GetDataprocClusterResult> Invoke(GetDataprocClusterInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDataprocClusterResult>("yandex:index/getDataprocCluster:getDataprocCluster", args ?? new GetDataprocClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDataprocClusterArgs : global::Pulumi.InvokeArgs
    {
        [Input("clusterId")]
        public string? ClusterId { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        public GetDataprocClusterArgs()
        {
        }
        public static new GetDataprocClusterArgs Empty => new GetDataprocClusterArgs();
    }

    public sealed class GetDataprocClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetDataprocClusterInvokeArgs()
        {
        }
        public static new GetDataprocClusterInvokeArgs Empty => new GetDataprocClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetDataprocClusterResult
    {
        public readonly string Bucket;
        public readonly ImmutableArray<Outputs.GetDataprocClusterClusterConfigResult> ClusterConfigs;
        public readonly string ClusterId;
        public readonly string CreatedAt;
        public readonly bool DeletionProtection;
        public readonly string Description;
        public readonly string Environment;
        public readonly string FolderId;
        public readonly ImmutableArray<string> HostGroupIds;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly string LogGroupId;
        public readonly string Name;
        public readonly ImmutableArray<string> SecurityGroupIds;
        public readonly string ServiceAccountId;
        public readonly bool UiProxy;
        public readonly string ZoneId;

        [OutputConstructor]
        private GetDataprocClusterResult(
            string bucket,

            ImmutableArray<Outputs.GetDataprocClusterClusterConfigResult> clusterConfigs,

            string clusterId,

            string createdAt,

            bool deletionProtection,

            string description,

            string environment,

            string folderId,

            ImmutableArray<string> hostGroupIds,

            string id,

            ImmutableDictionary<string, string> labels,

            string logGroupId,

            string name,

            ImmutableArray<string> securityGroupIds,

            string serviceAccountId,

            bool uiProxy,

            string zoneId)
        {
            Bucket = bucket;
            ClusterConfigs = clusterConfigs;
            ClusterId = clusterId;
            CreatedAt = createdAt;
            DeletionProtection = deletionProtection;
            Description = description;
            Environment = environment;
            FolderId = folderId;
            HostGroupIds = hostGroupIds;
            Id = id;
            Labels = labels;
            LogGroupId = logGroupId;
            Name = name;
            SecurityGroupIds = securityGroupIds;
            ServiceAccountId = serviceAccountId;
            UiProxy = uiProxy;
            ZoneId = zoneId;
        }
    }
}
