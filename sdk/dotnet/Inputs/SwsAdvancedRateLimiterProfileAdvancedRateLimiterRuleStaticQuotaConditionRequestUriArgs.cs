// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleStaticQuotaConditionRequestUriArgs : global::Pulumi.ResourceArgs
    {
        [Input("path")]
        public Input<Inputs.SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleStaticQuotaConditionRequestUriPathArgs>? Path { get; set; }

        [Input("queries")]
        private InputList<Inputs.SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleStaticQuotaConditionRequestUriQueryArgs>? _queries;
        public InputList<Inputs.SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleStaticQuotaConditionRequestUriQueryArgs> Queries
        {
            get => _queries ?? (_queries = new InputList<Inputs.SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleStaticQuotaConditionRequestUriQueryArgs>());
            set => _queries = value;
        }

        public SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleStaticQuotaConditionRequestUriArgs()
        {
        }
        public static new SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleStaticQuotaConditionRequestUriArgs Empty => new SwsAdvancedRateLimiterProfileAdvancedRateLimiterRuleStaticQuotaConditionRequestUriArgs();
    }
}
