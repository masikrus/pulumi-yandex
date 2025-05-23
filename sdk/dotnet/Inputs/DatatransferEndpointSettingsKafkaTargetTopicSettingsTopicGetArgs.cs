// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class DatatransferEndpointSettingsKafkaTargetTopicSettingsTopicGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Not to split events queue into separate per-table queues.
        /// </summary>
        [Input("saveTxOrder")]
        public Input<bool>? SaveTxOrder { get; set; }

        /// <summary>
        /// Full topic name.
        /// </summary>
        [Input("topicName")]
        public Input<string>? TopicName { get; set; }

        public DatatransferEndpointSettingsKafkaTargetTopicSettingsTopicGetArgs()
        {
        }
        public static new DatatransferEndpointSettingsKafkaTargetTopicSettingsTopicGetArgs Empty => new DatatransferEndpointSettingsKafkaTargetTopicSettingsTopicGetArgs();
    }
}
