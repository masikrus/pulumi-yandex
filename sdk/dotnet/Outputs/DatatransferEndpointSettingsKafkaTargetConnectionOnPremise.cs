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
    public sealed class DatatransferEndpointSettingsKafkaTargetConnectionOnPremise
    {
        /// <summary>
        /// List of Kafka broker URLs.
        /// </summary>
        public readonly ImmutableArray<string> BrokerUrls;
        /// <summary>
        /// Identifier of the Yandex Cloud VPC subnetwork to user for accessing the database. If omitted, the server has to be accessible via Internet.
        /// </summary>
        public readonly string? SubnetId;
        /// <summary>
        /// TLS settings for the server connection. Empty implies plaintext connection.
        /// </summary>
        public readonly Outputs.DatatransferEndpointSettingsKafkaTargetConnectionOnPremiseTlsMode? TlsMode;

        [OutputConstructor]
        private DatatransferEndpointSettingsKafkaTargetConnectionOnPremise(
            ImmutableArray<string> brokerUrls,

            string? subnetId,

            Outputs.DatatransferEndpointSettingsKafkaTargetConnectionOnPremiseTlsMode? tlsMode)
        {
            BrokerUrls = brokerUrls;
            SubnetId = subnetId;
            TlsMode = tlsMode;
        }
    }
}
