// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MdbKafkaClusterConfigRestApiGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables REST API on cluster. The default is `false`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public MdbKafkaClusterConfigRestApiGetArgs()
        {
        }
        public static new MdbKafkaClusterConfigRestApiGetArgs Empty => new MdbKafkaClusterConfigRestApiGetArgs();
    }
}
