// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MdbMongodbClusterClusterConfigMongocfgGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A set of network settings (see the [net](https://www.mongodb.com/docs/manual/reference/configuration-options/#net-options) option).
        /// </summary>
        [Input("net")]
        public Input<Inputs.MdbMongodbClusterClusterConfigMongocfgNetGetArgs>? Net { get; set; }

        /// <summary>
        /// A set of profiling settings (see the [operationProfiling](https://www.mongodb.com/docs/manual/reference/configuration-options/#operationprofiling-options) option).
        /// </summary>
        [Input("operationProfiling")]
        public Input<Inputs.MdbMongodbClusterClusterConfigMongocfgOperationProfilingGetArgs>? OperationProfiling { get; set; }

        /// <summary>
        /// A set of storage settings (see the [storage](https://www.mongodb.com/docs/manual/reference/configuration-options/#storage-options) option).
        /// </summary>
        [Input("storage")]
        public Input<Inputs.MdbMongodbClusterClusterConfigMongocfgStorageGetArgs>? Storage { get; set; }

        public MdbMongodbClusterClusterConfigMongocfgGetArgs()
        {
        }
        public static new MdbMongodbClusterClusterConfigMongocfgGetArgs Empty => new MdbMongodbClusterClusterConfigMongocfgGetArgs();
    }
}
