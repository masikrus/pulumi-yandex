// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class SwsSecurityProfileSecurityRuleSmartProtectionConditionRequestUriArgs : global::Pulumi.ResourceArgs
    {
        [Input("path")]
        public Input<Inputs.SwsSecurityProfileSecurityRuleSmartProtectionConditionRequestUriPathArgs>? Path { get; set; }

        [Input("queries")]
        private InputList<Inputs.SwsSecurityProfileSecurityRuleSmartProtectionConditionRequestUriQueryArgs>? _queries;
        public InputList<Inputs.SwsSecurityProfileSecurityRuleSmartProtectionConditionRequestUriQueryArgs> Queries
        {
            get => _queries ?? (_queries = new InputList<Inputs.SwsSecurityProfileSecurityRuleSmartProtectionConditionRequestUriQueryArgs>());
            set => _queries = value;
        }

        public SwsSecurityProfileSecurityRuleSmartProtectionConditionRequestUriArgs()
        {
        }
        public static new SwsSecurityProfileSecurityRuleSmartProtectionConditionRequestUriArgs Empty => new SwsSecurityProfileSecurityRuleSmartProtectionConditionRequestUriArgs();
    }
}
