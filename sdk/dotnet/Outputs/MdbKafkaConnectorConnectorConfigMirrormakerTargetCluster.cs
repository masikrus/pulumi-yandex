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
    public sealed class MdbKafkaConnectorConnectorConfigMirrormakerTargetCluster
    {
        /// <summary>
        /// Name of the cluster. Used also as a topic prefix.
        /// </summary>
        public readonly string? Alias;
        /// <summary>
        /// Connection settings for external cluster.
        /// </summary>
        public readonly ImmutableArray<Outputs.MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalCluster> ExternalClusters;
        /// <summary>
        /// Using this section in the cluster definition (source or target) means it's this cluster.
        /// </summary>
        public readonly ImmutableArray<Outputs.MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterThisCluster> ThisClusters;

        [OutputConstructor]
        private MdbKafkaConnectorConnectorConfigMirrormakerTargetCluster(
            string? alias,

            ImmutableArray<Outputs.MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterExternalCluster> externalClusters,

            ImmutableArray<Outputs.MdbKafkaConnectorConnectorConfigMirrormakerTargetClusterThisCluster> thisClusters)
        {
            Alias = alias;
            ExternalClusters = externalClusters;
            ThisClusters = thisClusters;
        }
    }
}
