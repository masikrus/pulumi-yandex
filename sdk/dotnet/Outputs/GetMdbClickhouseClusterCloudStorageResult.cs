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
    public sealed class GetMdbClickhouseClusterCloudStorageResult
    {
        /// <summary>
        /// Enables temporary storage in the cluster repository of data requested from the object repository.
        /// </summary>
        public readonly bool DataCacheEnabled;
        /// <summary>
        /// Defines the maximum amount of memory (in bytes) allocated in the cluster storage for temporary storage of data requested from the object storage.
        /// </summary>
        public readonly int DataCacheMaxSize;
        /// <summary>
        /// Whether to use Yandex Object Storage for storing ClickHouse data. Can be either `true` or `false`.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Sets the minimum free space ratio in the cluster storage. If the free space is lower than this value, the data is transferred to Yandex Object Storage. Acceptable values are 0 to 1, inclusive.
        /// </summary>
        public readonly double MoveFactor;
        /// <summary>
        /// Disables merging of data parts in `Yandex Object Storage`.
        /// </summary>
        public readonly bool PreferNotToMerge;

        [OutputConstructor]
        private GetMdbClickhouseClusterCloudStorageResult(
            bool dataCacheEnabled,

            int dataCacheMaxSize,

            bool? enabled,

            double moveFactor,

            bool preferNotToMerge)
        {
            DataCacheEnabled = dataCacheEnabled;
            DataCacheMaxSize = dataCacheMaxSize;
            Enabled = enabled;
            MoveFactor = moveFactor;
            PreferNotToMerge = preferNotToMerge;
        }
    }
}
