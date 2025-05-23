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
    public sealed class ComputeInstanceGroupInstanceTemplateBootDiskInitializeParams
    {
        /// <summary>
        /// A description of the boot disk.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The disk image to initialize this disk from.
        /// </summary>
        public readonly string? ImageId;
        /// <summary>
        /// The size of the disk in GB.
        /// </summary>
        public readonly int? Size;
        /// <summary>
        /// The snapshot to initialize this disk from.
        /// </summary>
        public readonly string? SnapshotId;
        /// <summary>
        /// The disk type.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private ComputeInstanceGroupInstanceTemplateBootDiskInitializeParams(
            string? description,

            string? imageId,

            int? size,

            string? snapshotId,

            string? type)
        {
            Description = description;
            ImageId = imageId;
            Size = size;
            SnapshotId = snapshotId;
            Type = type;
        }
    }
}
