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
    public sealed class KubernetesNodeGroupInstanceTemplateNetworkInterface
    {
        /// <summary>
        /// Allocate an IPv4 address for the interface. The default value is `true`.
        /// </summary>
        public readonly bool? Ipv4;
        /// <summary>
        /// List of configurations for creating ipv4 DNS records.
        /// </summary>
        public readonly ImmutableArray<Outputs.KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecord> Ipv4DnsRecords;
        /// <summary>
        /// If true, allocate an IPv6 address for the interface. The address will be automatically assigned from the specified subnet.
        /// </summary>
        public readonly bool? Ipv6;
        /// <summary>
        /// List of configurations for creating ipv6 DNS records.
        /// </summary>
        public readonly ImmutableArray<Outputs.KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv6DnsRecord> Ipv6DnsRecords;
        /// <summary>
        /// A public address that can be used to access the internet over NAT.
        /// </summary>
        public readonly bool? Nat;
        /// <summary>
        /// Security group IDs for network interface.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// The IDs of the subnets.
        /// </summary>
        public readonly ImmutableArray<string> SubnetIds;

        [OutputConstructor]
        private KubernetesNodeGroupInstanceTemplateNetworkInterface(
            bool? ipv4,

            ImmutableArray<Outputs.KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv4DnsRecord> ipv4DnsRecords,

            bool? ipv6,

            ImmutableArray<Outputs.KubernetesNodeGroupInstanceTemplateNetworkInterfaceIpv6DnsRecord> ipv6DnsRecords,

            bool? nat,

            ImmutableArray<string> securityGroupIds,

            ImmutableArray<string> subnetIds)
        {
            Ipv4 = ipv4;
            Ipv4DnsRecords = ipv4DnsRecords;
            Ipv6 = ipv6;
            Ipv6DnsRecords = ipv6DnsRecords;
            Nat = nat;
            SecurityGroupIds = securityGroupIds;
            SubnetIds = subnetIds;
        }
    }
}
