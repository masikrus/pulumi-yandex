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
    public sealed class StorageBucketLifecycleRuleFilterAnd
    {
        /// <summary>
        /// Minimum object size to which the rule applies.
        /// </summary>
        public readonly int? ObjectSizeGreaterThan;
        /// <summary>
        /// Maximum object size to which the rule applies.
        /// </summary>
        public readonly int? ObjectSizeLessThan;
        /// <summary>
        /// Object key prefix identifying one or more objects to which the rule applies.
        /// </summary>
        public readonly string? Prefix;
        /// <summary>
        /// The `tags` object for setting tags (or labels) for bucket. See [Tags](https://yandex.cloud/docs/storage/concepts/tags) for more information.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;

        [OutputConstructor]
        private StorageBucketLifecycleRuleFilterAnd(
            int? objectSizeGreaterThan,

            int? objectSizeLessThan,

            string? prefix,

            ImmutableDictionary<string, string>? tags)
        {
            ObjectSizeGreaterThan = objectSizeGreaterThan;
            ObjectSizeLessThan = objectSizeLessThan;
            Prefix = prefix;
            Tags = tags;
        }
    }
}
