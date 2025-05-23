// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetMdbClickhouseClusterCloudStorageInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables temporary storage in the cluster repository of data requested from the object repository.
        /// </summary>
        [Input("dataCacheEnabled", required: true)]
        public Input<bool> DataCacheEnabled { get; set; } = null!;

        /// <summary>
        /// Defines the maximum amount of memory (in bytes) allocated in the cluster storage for temporary storage of data requested from the object storage.
        /// </summary>
        [Input("dataCacheMaxSize", required: true)]
        public Input<int> DataCacheMaxSize { get; set; } = null!;

        /// <summary>
        /// Whether to use Yandex Object Storage for storing ClickHouse data. Can be either `true` or `false`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Sets the minimum free space ratio in the cluster storage. If the free space is lower than this value, the data is transferred to Yandex Object Storage. Acceptable values are 0 to 1, inclusive.
        /// </summary>
        [Input("moveFactor", required: true)]
        public Input<double> MoveFactor { get; set; } = null!;

        /// <summary>
        /// Disables merging of data parts in `Yandex Object Storage`.
        /// </summary>
        [Input("preferNotToMerge", required: true)]
        public Input<bool> PreferNotToMerge { get; set; } = null!;

        public GetMdbClickhouseClusterCloudStorageInputArgs()
        {
        }
        public static new GetMdbClickhouseClusterCloudStorageInputArgs Empty => new GetMdbClickhouseClusterCloudStorageInputArgs();
    }
}
