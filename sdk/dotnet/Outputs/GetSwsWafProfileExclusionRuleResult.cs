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
    public sealed class GetSwsWafProfileExclusionRuleResult
    {
        public readonly ImmutableArray<Outputs.GetSwsWafProfileExclusionRuleConditionResult> Conditions;
        public readonly string Description;
        public readonly ImmutableArray<Outputs.GetSwsWafProfileExclusionRuleExcludeRuleResult> ExcludeRules;
        public readonly bool LogExcluded;
        public readonly string Name;

        [OutputConstructor]
        private GetSwsWafProfileExclusionRuleResult(
            ImmutableArray<Outputs.GetSwsWafProfileExclusionRuleConditionResult> conditions,

            string description,

            ImmutableArray<Outputs.GetSwsWafProfileExclusionRuleExcludeRuleResult> excludeRules,

            bool logExcluded,

            string name)
        {
            Conditions = conditions;
            Description = description;
            ExcludeRules = excludeRules;
            LogExcluded = logExcluded;
            Name = name;
        }
    }
}
