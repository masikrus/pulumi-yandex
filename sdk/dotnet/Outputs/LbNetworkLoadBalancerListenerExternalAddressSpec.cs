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
    public sealed class LbNetworkLoadBalancerListenerExternalAddressSpec
    {
        /// <summary>
        /// External IP address for a listener. IP address will be allocated if it wasn't been set.
        /// </summary>
        public readonly string? Address;
        /// <summary>
        /// IP version of the external addresses that the load balancer works with. Must be one of `ipv4` or `ipv6`. The default is `ipv4`.
        /// </summary>
        public readonly string? IpVersion;

        [OutputConstructor]
        private LbNetworkLoadBalancerListenerExternalAddressSpec(
            string? address,

            string? ipVersion)
        {
            Address = address;
            IpVersion = ipVersion;
        }
    }
}
