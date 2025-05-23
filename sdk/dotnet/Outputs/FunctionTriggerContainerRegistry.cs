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
    public sealed class FunctionTriggerContainerRegistry
    {
        /// <summary>
        /// Batch Duration in seconds for Yandex Cloud Functions Trigger.
        /// </summary>
        public readonly string BatchCutoff;
        /// <summary>
        /// Batch Size for Yandex Cloud Functions Trigger.
        /// </summary>
        public readonly string? BatchSize;
        /// <summary>
        /// Boolean flag for setting `create image` event for Yandex Cloud Functions Trigger.
        /// </summary>
        public readonly bool? CreateImage;
        /// <summary>
        /// Boolean flag for setting `create image tag` event for Yandex Cloud Functions Trigger.
        /// </summary>
        public readonly bool? CreateImageTag;
        /// <summary>
        /// Boolean flag for setting `delete image` event for Yandex Cloud Functions Trigger.
        /// </summary>
        public readonly bool? DeleteImage;
        /// <summary>
        /// Boolean flag for setting `delete image tag` event for Yandex Cloud Functions Trigger.
        /// </summary>
        public readonly bool? DeleteImageTag;
        /// <summary>
        /// Image name filter setting for Yandex Cloud Functions Trigger.
        /// </summary>
        public readonly string? ImageName;
        /// <summary>
        /// Container Registry ID for Yandex Cloud Functions Trigger.
        /// </summary>
        public readonly string RegistryId;
        /// <summary>
        /// Image tag filter setting for Yandex Cloud Functions Trigger.
        /// </summary>
        public readonly string? Tag;

        [OutputConstructor]
        private FunctionTriggerContainerRegistry(
            string batchCutoff,

            string? batchSize,

            bool? createImage,

            bool? createImageTag,

            bool? deleteImage,

            bool? deleteImageTag,

            string? imageName,

            string registryId,

            string? tag)
        {
            BatchCutoff = batchCutoff;
            BatchSize = batchSize;
            CreateImage = createImage;
            CreateImageTag = createImageTag;
            DeleteImage = deleteImage;
            DeleteImageTag = deleteImageTag;
            ImageName = imageName;
            RegistryId = registryId;
            Tag = tag;
        }
    }
}
