// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleStaticQuotaConditionGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("authority")]
        public Input<Inputs.SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleStaticQuotaConditionAuthorityGetArgs>? Authority { get; set; }

        [Input("headers")]
        private InputList<Inputs.SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleStaticQuotaConditionHeaderGetArgs>? _headers;
        public InputList<Inputs.SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleStaticQuotaConditionHeaderGetArgs> Headers
        {
            get => _headers ?? (_headers = new InputList<Inputs.SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleStaticQuotaConditionHeaderGetArgs>());
            set => _headers = value;
        }

        [Input("httpMethod")]
        public Input<Inputs.SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleStaticQuotaConditionHttpMethodGetArgs>? HttpMethod { get; set; }

        [Input("requestUri")]
        public Input<Inputs.SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleStaticQuotaConditionRequestUriGetArgs>? RequestUri { get; set; }

        [Input("sourceIp")]
        public Input<Inputs.SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleStaticQuotaConditionSourceIpGetArgs>? SourceIp { get; set; }

        public SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleStaticQuotaConditionGetArgs()
        {
        }
        public static new SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleStaticQuotaConditionGetArgs Empty => new SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleStaticQuotaConditionGetArgs();
    }
}
