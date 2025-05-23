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
    public sealed class MdbMysqlClusterHost
    {
        /// <summary>
        /// Sets whether the host should get a public IP address. It can be changed on the fly only when `name` is set.
        /// </summary>
        public readonly bool? AssignPublicIp;
        /// <summary>
        /// Host backup priority. Value is between 0 and 100, default is 0.
        /// </summary>
        public readonly int? BackupPriority;
        /// <summary>
        /// The fully qualified domain name of the host.
        /// </summary>
        public readonly string? Fqdn;
        /// <summary>
        /// Host state name. It should be set for all hosts or unset for all hosts. This field can be used by another host, to select which host will be its replication source. Please refer to `replication_source_name` parameter.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Host master promotion priority. Value is between 0 and 100, default is 0.
        /// </summary>
        public readonly int? Priority;
        /// <summary>
        /// Host replication source (fqdn), when replication_source is empty then host is in HA group.
        /// </summary>
        public readonly string? ReplicationSource;
        /// <summary>
        /// Host replication source name points to host's `name` from which this host should replicate. When not set then host in HA group. It works only when `name` is set.
        /// </summary>
        public readonly string? ReplicationSourceName;
        /// <summary>
        /// The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.
        /// </summary>
        public readonly string? SubnetId;
        /// <summary>
        /// The [availability zone](https://yandex.cloud/docs/overview/concepts/geo-scope) where resource is located. If it is not provided, the default provider zone will be used.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private MdbMysqlClusterHost(
            bool? assignPublicIp,

            int? backupPriority,

            string? fqdn,

            string? name,

            int? priority,

            string? replicationSource,

            string? replicationSourceName,

            string? subnetId,

            string zone)
        {
            AssignPublicIp = assignPublicIp;
            BackupPriority = backupPriority;
            Fqdn = fqdn;
            Name = name;
            Priority = priority;
            ReplicationSource = replicationSource;
            ReplicationSourceName = replicationSourceName;
            SubnetId = subnetId;
            Zone = zone;
        }
    }
}
