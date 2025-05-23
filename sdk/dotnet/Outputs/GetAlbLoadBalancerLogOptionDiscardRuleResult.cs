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
    public sealed class GetAlbLoadBalancerLogOptionDiscardRuleResult
    {
        /// <summary>
        /// The percent of logs which will be discarded.
        /// </summary>
        public readonly int DiscardPercent;
        /// <summary>
        /// list of grpc codes by name
        /// </summary>
        public readonly ImmutableArray<string> GrpcCodes;
        /// <summary>
        /// List of http code intervals *1XX*-*5XX* or *ALL*
        /// </summary>
        public readonly ImmutableArray<string> HttpCodeIntervals;
        /// <summary>
        /// List of http codes *100*-*599*.
        /// </summary>
        public readonly ImmutableArray<int> HttpCodes;

        [OutputConstructor]
        private GetAlbLoadBalancerLogOptionDiscardRuleResult(
            int discardPercent,

            ImmutableArray<string> grpcCodes,

            ImmutableArray<string> httpCodeIntervals,

            ImmutableArray<int> httpCodes)
        {
            DiscardPercent = discardPercent;
            GrpcCodes = grpcCodes;
            HttpCodeIntervals = httpCodeIntervals;
            HttpCodes = httpCodes;
        }
    }
}
