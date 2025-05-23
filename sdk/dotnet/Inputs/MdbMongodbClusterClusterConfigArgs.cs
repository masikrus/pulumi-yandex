// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MdbMongodbClusterClusterConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access policy to the MongoDB cluster.
        /// </summary>
        [Input("access")]
        public Input<Inputs.MdbMongodbClusterClusterConfigAccessArgs>? Access { get; set; }

        /// <summary>
        /// Retain period of automatically created backup in days.
        /// </summary>
        [Input("backupRetainPeriodDays")]
        public Input<int>? BackupRetainPeriodDays { get; set; }

        /// <summary>
        /// Time to start the daily backup, in the UTC timezone.
        /// </summary>
        [Input("backupWindowStart")]
        public Input<Inputs.MdbMongodbClusterClusterConfigBackupWindowStartArgs>? BackupWindowStart { get; set; }

        /// <summary>
        /// Feature compatibility version of MongoDB. If not provided version is taken. Can be either `6.0`, `5.0`, `4.4` and `4.2`.
        /// </summary>
        [Input("featureCompatibilityVersion")]
        public Input<string>? FeatureCompatibilityVersion { get; set; }

        /// <summary>
        /// Configuration of the mongocfg service.
        /// </summary>
        [Input("mongocfg")]
        public Input<Inputs.MdbMongodbClusterClusterConfigMongocfgArgs>? Mongocfg { get; set; }

        /// <summary>
        /// Configuration of the mongod service.
        /// </summary>
        [Input("mongod")]
        public Input<Inputs.MdbMongodbClusterClusterConfigMongodArgs>? Mongod { get; set; }

        /// <summary>
        /// Configuration of the mongos service.
        /// </summary>
        [Input("mongos")]
        public Input<Inputs.MdbMongodbClusterClusterConfigMongosArgs>? Mongos { get; set; }

        /// <summary>
        /// Performance diagnostics to the MongoDB cluster.
        /// </summary>
        [Input("performanceDiagnostics")]
        public Input<Inputs.MdbMongodbClusterClusterConfigPerformanceDiagnosticsArgs>? PerformanceDiagnostics { get; set; }

        /// <summary>
        /// Version of the MongoDB server software. Can be either `4.2`, `4.4`, `4.4-enterprise`, `5.0`, `5.0-enterprise`, `6.0` and `6.0-enterprise`.
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public MdbMongodbClusterClusterConfigArgs()
        {
        }
        public static new MdbMongodbClusterClusterConfigArgs Empty => new MdbMongodbClusterClusterConfigArgs();
    }
}
