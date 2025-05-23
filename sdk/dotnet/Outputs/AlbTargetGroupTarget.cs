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
    public sealed class AlbTargetGroupTarget
    {
        /// <summary>
        /// IP address of the target.
        /// </summary>
        public readonly string IpAddress;
        public readonly bool? PrivateIpv4Address;
        /// <summary>
        /// ID of the subnet that targets are connected to. All targets in the target group must be connected to the same subnet within a single availability zone.
        /// </summary>
        public readonly string? SubnetId;

        [OutputConstructor]
        private AlbTargetGroupTarget(
            string ipAddress,

            bool? privateIpv4Address,

            string? subnetId)
        {
            IpAddress = ipAddress;
            PrivateIpv4Address = privateIpv4Address;
            SubnetId = subnetId;
        }
    }
}
