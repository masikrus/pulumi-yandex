// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class AlbVirtualHostRateLimitArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Rate limit configuration applied to all incoming requests
        /// </summary>
        [Input("allRequests")]
        public Input<Inputs.AlbVirtualHostRateLimitAllRequestsArgs>? AllRequests { get; set; }

        /// <summary>
        /// Rate limit configuration applied separately for each set of requests grouped by client IP address
        /// </summary>
        [Input("requestsPerIp")]
        public Input<Inputs.AlbVirtualHostRateLimitRequestsPerIpArgs>? RequestsPerIp { get; set; }

        public AlbVirtualHostRateLimitArgs()
        {
        }
        public static new AlbVirtualHostRateLimitArgs Empty => new AlbVirtualHostRateLimitArgs();
    }
}
