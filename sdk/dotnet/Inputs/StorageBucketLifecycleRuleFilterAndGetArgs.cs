// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class StorageBucketLifecycleRuleFilterAndGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Minimum object size to which the rule applies.
        /// </summary>
        [Input("objectSizeGreaterThan")]
        public Input<int>? ObjectSizeGreaterThan { get; set; }

        /// <summary>
        /// Maximum object size to which the rule applies.
        /// </summary>
        [Input("objectSizeLessThan")]
        public Input<int>? ObjectSizeLessThan { get; set; }

        /// <summary>
        /// Object key prefix identifying one or more objects to which the rule applies.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The `tags` object for setting tags (or labels) for bucket. See [Tags](https://yandex.cloud/docs/storage/concepts/tags) for more information.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public StorageBucketLifecycleRuleFilterAndGetArgs()
        {
        }
        public static new StorageBucketLifecycleRuleFilterAndGetArgs Empty => new StorageBucketLifecycleRuleFilterAndGetArgs();
    }
}
