// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class SwsSecurityProfileSecurityRuleSmartProtectionConditionAuthorityArgs : global::Pulumi.ResourceArgs
    {
        [Input("authorities")]
        private InputList<Inputs.SwsSecurityProfileSecurityRuleSmartProtectionConditionAuthorityAuthorityArgs>? _authorities;
        public InputList<Inputs.SwsSecurityProfileSecurityRuleSmartProtectionConditionAuthorityAuthorityArgs> Authorities
        {
            get => _authorities ?? (_authorities = new InputList<Inputs.SwsSecurityProfileSecurityRuleSmartProtectionConditionAuthorityAuthorityArgs>());
            set => _authorities = value;
        }

        public SwsSecurityProfileSecurityRuleSmartProtectionConditionAuthorityArgs()
        {
        }
        public static new SwsSecurityProfileSecurityRuleSmartProtectionConditionAuthorityArgs Empty => new SwsSecurityProfileSecurityRuleSmartProtectionConditionAuthorityArgs();
    }
}
