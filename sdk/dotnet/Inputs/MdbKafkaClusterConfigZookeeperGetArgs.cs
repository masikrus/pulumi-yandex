// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MdbKafkaClusterConfigZookeeperGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resources allocated to hosts of the ZooKeeper subcluster.
        /// </summary>
        [Input("resources")]
        public Input<Inputs.MdbKafkaClusterConfigZookeeperResourcesGetArgs>? Resources { get; set; }

        public MdbKafkaClusterConfigZookeeperGetArgs()
        {
        }
        public static new MdbKafkaClusterConfigZookeeperGetArgs Empty => new MdbKafkaClusterConfigZookeeperGetArgs();
    }
}
