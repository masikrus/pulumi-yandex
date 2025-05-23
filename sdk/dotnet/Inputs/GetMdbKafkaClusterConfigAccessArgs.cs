// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetMdbKafkaClusterConfigAccessInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow access for [DataTransfer](https://yandex.cloud/services/data-transfer).
        /// </summary>
        [Input("dataTransfer")]
        public Input<bool>? DataTransfer { get; set; }

        public GetMdbKafkaClusterConfigAccessInputArgs()
        {
        }
        public static new GetMdbKafkaClusterConfigAccessInputArgs Empty => new GetMdbKafkaClusterConfigAccessInputArgs();
    }
}
