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
    public sealed class ServerlessContainerMount
    {
        /// <summary>
        /// One of the available mount types. Disk available during the function execution time.
        /// </summary>
        public readonly Outputs.ServerlessContainerMountEphemeralDisk? EphemeralDisk;
        /// <summary>
        /// Mount’s accessibility mode. Valid values are `ro` and `rw`.
        /// </summary>
        public readonly string? Mode;
        /// <summary>
        /// Path inside the container to access the directory in which the target is mounted.
        /// </summary>
        public readonly string MountPointPath;
        /// <summary>
        /// Available mount types. Object storage as a mount.
        /// </summary>
        public readonly Outputs.ServerlessContainerMountObjectStorage? ObjectStorage;

        [OutputConstructor]
        private ServerlessContainerMount(
            Outputs.ServerlessContainerMountEphemeralDisk? ephemeralDisk,

            string? mode,

            string mountPointPath,

            Outputs.ServerlessContainerMountObjectStorage? objectStorage)
        {
            EphemeralDisk = ephemeralDisk;
            Mode = mode;
            MountPointPath = mountPointPath;
            ObjectStorage = objectStorage;
        }
    }
}
