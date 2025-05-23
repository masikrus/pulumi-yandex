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
    public sealed class VpcPrivateEndpointEndpointAddress
    {
        /// <summary>
        /// Specifies IP address within `subnet_id`.
        /// </summary>
        public readonly string? Address;
        /// <summary>
        /// ID of the address.
        /// </summary>
        public readonly string? AddressId;
        /// <summary>
        /// Subnet of the IP address.
        /// </summary>
        public readonly string? SubnetId;

        [OutputConstructor]
        private VpcPrivateEndpointEndpointAddress(
            string? address,

            string? addressId,

            string? subnetId)
        {
            Address = address;
            AddressId = addressId;
            SubnetId = subnetId;
        }
    }
}
