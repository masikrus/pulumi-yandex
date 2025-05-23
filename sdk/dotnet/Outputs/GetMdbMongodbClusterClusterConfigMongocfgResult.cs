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
    public sealed class GetMdbMongodbClusterClusterConfigMongocfgResult
    {
        /// <summary>
        /// A set of network settings (see the [net](https://www.mongodb.com/docs/manual/reference/configuration-options/#net-options) option).
        /// </summary>
        public readonly Outputs.GetMdbMongodbClusterClusterConfigMongocfgNetResult? Net;
        /// <summary>
        /// A set of profiling settings (see the [operationProfiling](https://www.mongodb.com/docs/manual/reference/configuration-options/#operationprofiling-options) option).
        /// </summary>
        public readonly Outputs.GetMdbMongodbClusterClusterConfigMongocfgOperationProfilingResult? OperationProfiling;
        /// <summary>
        /// A set of storage settings (see the [storage](https://www.mongodb.com/docs/manual/reference/configuration-options/#storage-options) option).
        /// </summary>
        public readonly Outputs.GetMdbMongodbClusterClusterConfigMongocfgStorageResult? Storage;

        [OutputConstructor]
        private GetMdbMongodbClusterClusterConfigMongocfgResult(
            Outputs.GetMdbMongodbClusterClusterConfigMongocfgNetResult? net,

            Outputs.GetMdbMongodbClusterClusterConfigMongocfgOperationProfilingResult? operationProfiling,

            Outputs.GetMdbMongodbClusterClusterConfigMongocfgStorageResult? storage)
        {
            Net = net;
            OperationProfiling = operationProfiling;
            Storage = storage;
        }
    }
}
