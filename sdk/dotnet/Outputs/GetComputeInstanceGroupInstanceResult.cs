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
    public sealed class GetComputeInstanceGroupInstanceResult
    {
        public readonly string Fqdn;
        public readonly string InstanceId;
        public readonly string InstanceTag;
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetComputeInstanceGroupInstanceNetworkInterfaceResult> NetworkInterfaces;
        public readonly string Status;
        public readonly string StatusChangedAt;
        public readonly string StatusMessage;
        public readonly string ZoneId;

        [OutputConstructor]
        private GetComputeInstanceGroupInstanceResult(
            string fqdn,

            string instanceId,

            string instanceTag,

            string name,

            ImmutableArray<Outputs.GetComputeInstanceGroupInstanceNetworkInterfaceResult> networkInterfaces,

            string status,

            string statusChangedAt,

            string statusMessage,

            string zoneId)
        {
            Fqdn = fqdn;
            InstanceId = instanceId;
            InstanceTag = instanceTag;
            Name = name;
            NetworkInterfaces = networkInterfaces;
            Status = status;
            StatusChangedAt = statusChangedAt;
            StatusMessage = statusMessage;
            ZoneId = zoneId;
        }
    }
}
