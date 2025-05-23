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
    public sealed class ServerlessContainerMountEphemeralDisk
    {
        /// <summary>
        /// Block size of the ephemeral disk in KB.
        /// </summary>
        public readonly int? BlockSizeKb;
        /// <summary>
        /// Size of the ephemeral disk in GB.
        /// </summary>
        public readonly int SizeGb;

        [OutputConstructor]
        private ServerlessContainerMountEphemeralDisk(
            int? blockSizeKb,

            int sizeGb)
        {
            BlockSizeKb = blockSizeKb;
            SizeGb = sizeGb;
        }
    }
}
