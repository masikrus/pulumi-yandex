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
    public sealed class DatatransferEndpointSettingsMongoTarget
    {
        /// <summary>
        /// How to clean collections when activating the transfer. One of `DISABLED`, `DROP` or `TRUNCATE`.
        /// </summary>
        public readonly string? CleanupPolicy;
        /// <summary>
        /// Connection settings.
        /// </summary>
        public readonly Outputs.DatatransferEndpointSettingsMongoTargetConnection? Connection;
        /// <summary>
        /// If not empty, then all the data will be written to the database with the specified name; otherwise the database name is the same as in the source endpoint.
        /// </summary>
        public readonly string? Database;
        /// <summary>
        /// List of security groups that the transfer associated with this endpoint should use.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroups;
        /// <summary>
        /// Identifier of the Yandex Cloud VPC subnetwork to user for accessing the database. If omitted, the server has to be accessible via Internet.
        /// </summary>
        public readonly string? SubnetId;

        [OutputConstructor]
        private DatatransferEndpointSettingsMongoTarget(
            string? cleanupPolicy,

            Outputs.DatatransferEndpointSettingsMongoTargetConnection? connection,

            string? database,

            ImmutableArray<string> securityGroups,

            string? subnetId)
        {
            CleanupPolicy = cleanupPolicy;
            Connection = connection;
            Database = database;
            SecurityGroups = securityGroups;
            SubnetId = subnetId;
        }
    }
}
