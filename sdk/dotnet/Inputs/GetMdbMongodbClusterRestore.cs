// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetMdbMongodbClusterRestoreArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Backup ID. The cluster will be created from the specified backup. [How to get a list of PostgreSQL backups](https://yandex.cloud/docs/managed-mongodb/operations/cluster-backups).
        /// </summary>
        [Input("backupId")]
        public string? BackupId { get; set; }

        /// <summary>
        /// Timestamp of the moment to which the MongoDB cluster should be restored. (Format: `2006-01-02T15:04:05` - UTC). When not set, current time is used.
        /// </summary>
        [Input("time")]
        public string? Time { get; set; }

        public GetMdbMongodbClusterRestoreArgs()
        {
        }
        public static new GetMdbMongodbClusterRestoreArgs Empty => new GetMdbMongodbClusterRestoreArgs();
    }
}
