// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class KubernetesClusterMasterZonalGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the subnet. If no ID is specified, and there only one subnet in specified zone, an address in this subnet will be allocated.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// ID of the availability zone.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public KubernetesClusterMasterZonalGetArgs()
        {
        }
        public static new KubernetesClusterMasterZonalGetArgs Empty => new KubernetesClusterMasterZonalGetArgs();
    }
}
