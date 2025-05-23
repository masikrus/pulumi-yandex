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
    public sealed class MdbKafkaConnectorConnectorConfigMirrormaker
    {
        /// <summary>
        /// Replication factor for topics created in target cluster.
        /// </summary>
        public readonly int ReplicationFactor;
        /// <summary>
        /// Settings for source cluster.
        /// </summary>
        public readonly Outputs.MdbKafkaConnectorConnectorConfigMirrormakerSourceCluster SourceCluster;
        /// <summary>
        /// Settings for target cluster.
        /// </summary>
        public readonly Outputs.MdbKafkaConnectorConnectorConfigMirrormakerTargetCluster TargetCluster;
        /// <summary>
        /// The pattern for topic names to be replicated.
        /// </summary>
        public readonly string Topics;

        [OutputConstructor]
        private MdbKafkaConnectorConnectorConfigMirrormaker(
            int replicationFactor,

            Outputs.MdbKafkaConnectorConnectorConfigMirrormakerSourceCluster sourceCluster,

            Outputs.MdbKafkaConnectorConnectorConfigMirrormakerTargetCluster targetCluster,

            string topics)
        {
            ReplicationFactor = replicationFactor;
            SourceCluster = sourceCluster;
            TargetCluster = targetCluster;
            Topics = topics;
        }
    }
}
