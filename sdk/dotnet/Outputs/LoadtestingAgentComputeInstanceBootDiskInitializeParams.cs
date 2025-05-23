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
    public sealed class LoadtestingAgentComputeInstanceBootDiskInitializeParams
    {
        /// <summary>
        /// Block size of the disk, specified in bytes.
        /// </summary>
        public readonly int? BlockSize;
        /// <summary>
        /// A description of the boot disk.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// A name of the boot disk.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The size of the disk in GB. Defaults to 15 GB.
        /// </summary>
        public readonly int? Size;
        /// <summary>
        /// The disk type.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private LoadtestingAgentComputeInstanceBootDiskInitializeParams(
            int? blockSize,

            string? description,

            string? name,

            int? size,

            string? type)
        {
            BlockSize = blockSize;
            Description = description;
            Name = name;
            Size = size;
            Type = type;
        }
    }
}
