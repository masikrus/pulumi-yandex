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
    public sealed class GetSwsWafProfileExclusionRuleConditionAuthorityResult
    {
        public readonly ImmutableArray<Outputs.GetSwsWafProfileExclusionRuleConditionAuthorityAuthorityResult> Authorities;

        [OutputConstructor]
        private GetSwsWafProfileExclusionRuleConditionAuthorityResult(ImmutableArray<Outputs.GetSwsWafProfileExclusionRuleConditionAuthorityAuthorityResult> authorities)
        {
            Authorities = authorities;
        }
    }
}
