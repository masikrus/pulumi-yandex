// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class SwsSecurityProfileSecurityRuleWafConditionArgs : global::Pulumi.ResourceArgs
    {
        [Input("authority")]
        public Input<Inputs.SwsSecurityProfileSecurityRuleWafConditionAuthorityArgs>? Authority { get; set; }

        [Input("headers")]
        private InputList<Inputs.SwsSecurityProfileSecurityRuleWafConditionHeaderArgs>? _headers;
        public InputList<Inputs.SwsSecurityProfileSecurityRuleWafConditionHeaderArgs> Headers
        {
            get => _headers ?? (_headers = new InputList<Inputs.SwsSecurityProfileSecurityRuleWafConditionHeaderArgs>());
            set => _headers = value;
        }

        [Input("httpMethod")]
        public Input<Inputs.SwsSecurityProfileSecurityRuleWafConditionHttpMethodArgs>? HttpMethod { get; set; }

        [Input("requestUri")]
        public Input<Inputs.SwsSecurityProfileSecurityRuleWafConditionRequestUriArgs>? RequestUri { get; set; }

        [Input("sourceIp")]
        public Input<Inputs.SwsSecurityProfileSecurityRuleWafConditionSourceIpArgs>? SourceIp { get; set; }

        public SwsSecurityProfileSecurityRuleWafConditionArgs()
        {
        }
        public static new SwsSecurityProfileSecurityRuleWafConditionArgs Empty => new SwsSecurityProfileSecurityRuleWafConditionArgs();
    }
}
