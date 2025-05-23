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
    public sealed class DatatransferEndpointSettingsClickhouseSource
    {
        public readonly string? ClickhouseClusterName;
        /// <summary>
        /// Connection settings.
        /// </summary>
        public readonly Outputs.DatatransferEndpointSettingsClickhouseSourceConnection? Connection;
        /// <summary>
        /// The list of tables that should not be transferred.
        /// </summary>
        public readonly ImmutableArray<string> ExcludeTables;
        /// <summary>
        /// The list of tables that should be transferred. Leave empty if all tables should be transferred.
        /// </summary>
        public readonly ImmutableArray<string> IncludeTables;
        /// <summary>
        /// List of security groups that the transfer associated with this endpoint should use.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroups;
        /// <summary>
        /// Identifier of the Yandex Cloud VPC subnetwork to user for accessing the database. If omitted, the server has to be accessible via Internet.
        /// </summary>
        public readonly string? SubnetId;

        [OutputConstructor]
        private DatatransferEndpointSettingsClickhouseSource(
            string? clickhouseClusterName,

            Outputs.DatatransferEndpointSettingsClickhouseSourceConnection? connection,

            ImmutableArray<string> excludeTables,

            ImmutableArray<string> includeTables,

            ImmutableArray<string> securityGroups,

            string? subnetId)
        {
            ClickhouseClusterName = clickhouseClusterName;
            Connection = connection;
            ExcludeTables = excludeTables;
            IncludeTables = includeTables;
            SecurityGroups = securityGroups;
            SubnetId = subnetId;
        }
    }
}
