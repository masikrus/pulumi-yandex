// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class LbNetworkLoadBalancerListenerGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// External IP address specification.
        /// </summary>
        [Input("externalAddressSpec")]
        public Input<Inputs.LbNetworkLoadBalancerListenerExternalAddressSpecGetArgs>? ExternalAddressSpec { get; set; }

        /// <summary>
        /// Internal IP address specification.
        /// </summary>
        [Input("internalAddressSpec")]
        public Input<Inputs.LbNetworkLoadBalancerListenerInternalAddressSpecGetArgs>? InternalAddressSpec { get; set; }

        /// <summary>
        /// Name of the listener. The name must be unique for each listener on a single load balancer.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Port for incoming traffic.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// Protocol for incoming traffic. TCP or UDP and the default is TCP.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Port of a target. The default is the same as listener's port.
        /// </summary>
        [Input("targetPort")]
        public Input<int>? TargetPort { get; set; }

        public LbNetworkLoadBalancerListenerGetArgs()
        {
        }
        public static new LbNetworkLoadBalancerListenerGetArgs Empty => new LbNetworkLoadBalancerListenerGetArgs();
    }
}
