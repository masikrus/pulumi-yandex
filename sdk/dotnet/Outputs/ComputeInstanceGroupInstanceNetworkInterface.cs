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
    public sealed class ComputeInstanceGroupInstanceNetworkInterface
    {
        /// <summary>
        /// The index of the network interface as generated by the server.
        /// </summary>
        public readonly int? Index;
        /// <summary>
        /// The private IP address to assign to the instance. If empty, the address is automatically assigned from the specified subnet.
        /// </summary>
        public readonly string? IpAddress;
        /// <summary>
        /// `True` if IPv4 address allocated for the network interface.
        /// </summary>
        public readonly bool? Ipv4;
        public readonly bool? Ipv6;
        public readonly string? Ipv6Address;
        /// <summary>
        /// The MAC address assigned to the network interface.
        /// </summary>
        public readonly string? MacAddress;
        /// <summary>
        /// The instance's public address for accessing the internet over NAT.
        /// </summary>
        public readonly bool? Nat;
        /// <summary>
        /// The public IP address of the instance.
        /// </summary>
        public readonly string? NatIpAddress;
        /// <summary>
        /// The IP version for the public address.
        /// </summary>
        public readonly string? NatIpVersion;
        /// <summary>
        /// The ID of the subnet to attach this interface to. The subnet must reside in the same zone where this instance was created.
        /// </summary>
        public readonly string? SubnetId;

        [OutputConstructor]
        private ComputeInstanceGroupInstanceNetworkInterface(
            int? index,

            string? ipAddress,

            bool? ipv4,

            bool? ipv6,

            string? ipv6Address,

            string? macAddress,

            bool? nat,

            string? natIpAddress,

            string? natIpVersion,

            string? subnetId)
        {
            Index = index;
            IpAddress = ipAddress;
            Ipv4 = ipv4;
            Ipv6 = ipv6;
            Ipv6Address = ipv6Address;
            MacAddress = macAddress;
            Nat = nat;
            NatIpAddress = natIpAddress;
            NatIpVersion = natIpVersion;
            SubnetId = subnetId;
        }
    }
}
