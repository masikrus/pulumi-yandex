// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetMdbClickhouseClusterShardResourcesInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Volume of the storage available to a ClickHouse host, in gigabytes.
        /// </summary>
        [Input("diskSize", required: true)]
        public Input<int> DiskSize { get; set; } = null!;

        /// <summary>
        /// Type of the storage of ClickHouse hosts. For more information see [the official documentation](https://yandex.cloud/docs/managed-clickhouse/concepts/storage).
        /// </summary>
        [Input("diskTypeId", required: true)]
        public Input<string> DiskTypeId { get; set; } = null!;

        /// <summary>
        /// The ID of the preset for computational resources available to a ClickHouse host (CPU, memory etc.). For more information, see [the official documentation](https://yandex.cloud/docs/managed-clickhouse/concepts).
        /// </summary>
        [Input("resourcePresetId", required: true)]
        public Input<string> ResourcePresetId { get; set; } = null!;

        public GetMdbClickhouseClusterShardResourcesInputArgs()
        {
        }
        public static new GetMdbClickhouseClusterShardResourcesInputArgs Empty => new GetMdbClickhouseClusterShardResourcesInputArgs();
    }
}
