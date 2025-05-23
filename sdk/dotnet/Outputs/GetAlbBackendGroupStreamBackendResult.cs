// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Outputs
{

    [OutputType]
    public sealed class GetAlbBackendGroupStreamBackendResult
    {
        public readonly bool EnableProxyProtocol;
        public readonly Outputs.GetAlbBackendGroupStreamBackendHealthcheckResult Healthcheck;
        public readonly bool? KeepConnectionsOnHostHealthFailure;
        public readonly Outputs.GetAlbBackendGroupStreamBackendLoadBalancingConfigResult LoadBalancingConfig;
        public readonly string Name;
        public readonly int Port;
        public readonly ImmutableArray<string> TargetGroupIds;
        public readonly Outputs.GetAlbBackendGroupStreamBackendTlsResult Tls;
        public readonly int Weight;

        [OutputConstructor]
        private GetAlbBackendGroupStreamBackendResult(
            bool enableProxyProtocol,

            Outputs.GetAlbBackendGroupStreamBackendHealthcheckResult healthcheck,

            bool? keepConnectionsOnHostHealthFailure,

            Outputs.GetAlbBackendGroupStreamBackendLoadBalancingConfigResult loadBalancingConfig,

            string name,

            int port,

            ImmutableArray<string> targetGroupIds,

            Outputs.GetAlbBackendGroupStreamBackendTlsResult tls,

            int weight)
        {
            EnableProxyProtocol = enableProxyProtocol;
            Healthcheck = healthcheck;
            KeepConnectionsOnHostHealthFailure = keepConnectionsOnHostHealthFailure;
            LoadBalancingConfig = loadBalancingConfig;
            Name = name;
            Port = port;
            TargetGroupIds = targetGroupIds;
            Tls = tls;
            Weight = weight;
        }
    }
}
