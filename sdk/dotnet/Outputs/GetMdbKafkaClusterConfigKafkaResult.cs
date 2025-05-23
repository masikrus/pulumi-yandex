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
    public sealed class GetMdbKafkaClusterConfigKafkaResult
    {
        /// <summary>
        /// User-defined settings for the Kafka cluster. For more information, see [the official documentation](https://yandex.cloud/docs/managed-kafka/operations/cluster-update) and [the Kafka documentation](https://kafka.apache.org/documentation/#configuration).
        /// </summary>
        public readonly Outputs.GetMdbKafkaClusterConfigKafkaKafkaConfigResult? KafkaConfig;
        /// <summary>
        /// Resources allocated to hosts of the Kafka subcluster.
        /// </summary>
        public readonly Outputs.GetMdbKafkaClusterConfigKafkaResourcesResult Resources;

        [OutputConstructor]
        private GetMdbKafkaClusterConfigKafkaResult(
            Outputs.GetMdbKafkaClusterConfigKafkaKafkaConfigResult? kafkaConfig,

            Outputs.GetMdbKafkaClusterConfigKafkaResourcesResult resources)
        {
            KafkaConfig = kafkaConfig;
            Resources = resources;
        }
    }
}
