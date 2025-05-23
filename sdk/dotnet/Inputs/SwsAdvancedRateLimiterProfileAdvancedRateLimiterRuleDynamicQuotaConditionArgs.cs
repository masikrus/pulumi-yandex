// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleDynamicQuotaConditionArgs : global::Pulumi.ResourceArgs
    {
        [Input("authority")]
        public Input<Inputs.SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleDynamicQuotaConditionAuthorityArgs>? Authority { get; set; }

        [Input("headers")]
        private InputList<Inputs.SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleDynamicQuotaConditionHeaderArgs>? _headers;
        public InputList<Inputs.SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleDynamicQuotaConditionHeaderArgs> Headers
        {
            get => _headers ?? (_headers = new InputList<Inputs.SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleDynamicQuotaConditionHeaderArgs>());
            set => _headers = value;
        }

        [Input("httpMethod")]
        public Input<Inputs.SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleDynamicQuotaConditionHttpMethodArgs>? HttpMethod { get; set; }

        [Input("requestUri")]
        public Input<Inputs.SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleDynamicQuotaConditionRequestUriArgs>? RequestUri { get; set; }

        [Input("sourceIp")]
        public Input<Inputs.SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleDynamicQuotaConditionSourceIpArgs>? SourceIp { get; set; }

        public SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleDynamicQuotaConditionArgs()
        {
        }
        public static new SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleDynamicQuotaConditionArgs Empty => new SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleDynamicQuotaConditionArgs();
    }
}
