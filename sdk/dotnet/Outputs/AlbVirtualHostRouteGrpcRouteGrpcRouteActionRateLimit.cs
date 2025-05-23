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
    public sealed class AlbVirtualHostRouteGrpcRouteGrpcRouteActionRateLimit
    {
        /// <summary>
        /// Rate limit configuration applied to all incoming requests
        /// </summary>
        public readonly Outputs.AlbVirtualHostRouteGrpcRouteGrpcRouteActionRateLimitAllRequests? AllRequests;
        /// <summary>
        /// Rate limit configuration applied separately for each set of requests grouped by client IP address
        /// </summary>
        public readonly Outputs.AlbVirtualHostRouteGrpcRouteGrpcRouteActionRateLimitRequestsPerIp? RequestsPerIp;

        [OutputConstructor]
        private AlbVirtualHostRouteGrpcRouteGrpcRouteActionRateLimit(
            Outputs.AlbVirtualHostRouteGrpcRouteGrpcRouteActionRateLimitAllRequests? allRequests,

            Outputs.AlbVirtualHostRouteGrpcRouteGrpcRouteActionRateLimitRequestsPerIp? requestsPerIp)
        {
            AllRequests = allRequests;
            RequestsPerIp = requestsPerIp;
        }
    }
}
