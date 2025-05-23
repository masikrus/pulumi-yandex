// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class FunctionTriggerObjectStorageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Batch Duration in seconds for Yandex Cloud Functions Trigger.
        /// </summary>
        [Input("batchCutoff", required: true)]
        public Input<string> BatchCutoff { get; set; } = null!;

        /// <summary>
        /// Batch Size for Yandex Cloud Functions Trigger.
        /// </summary>
        [Input("batchSize")]
        public Input<string>? BatchSize { get; set; }

        /// <summary>
        /// Object Storage Bucket ID for Yandex Cloud Functions Trigger.
        /// </summary>
        [Input("bucketId", required: true)]
        public Input<string> BucketId { get; set; } = null!;

        /// <summary>
        /// Boolean flag for setting `create` event for Yandex Cloud Functions Trigger.
        /// </summary>
        [Input("create")]
        public Input<bool>? Create { get; set; }

        /// <summary>
        /// Boolean flag for setting `delete` event for Yandex Cloud Functions Trigger.
        /// </summary>
        [Input("delete")]
        public Input<bool>? Delete { get; set; }

        /// <summary>
        /// Prefix for Object Storage for Yandex Cloud Functions Trigger.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// Suffix for Object Storage for Yandex Cloud Functions Trigger.
        /// </summary>
        [Input("suffix")]
        public Input<string>? Suffix { get; set; }

        /// <summary>
        /// Boolean flag for setting `update` event for Yandex Cloud Functions Trigger.
        /// </summary>
        [Input("update")]
        public Input<bool>? Update { get; set; }

        public FunctionTriggerObjectStorageArgs()
        {
        }
        public static new FunctionTriggerObjectStorageArgs Empty => new FunctionTriggerObjectStorageArgs();
    }
}
