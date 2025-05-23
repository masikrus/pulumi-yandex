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
    public sealed class SwsWafProfileExclusionRuleConditionSourceIp
    {
        public readonly Outputs.SwsWafProfileExclusionRuleConditionSourceIpGeoIpMatch? GeoIpMatch;
        public readonly Outputs.SwsWafProfileExclusionRuleConditionSourceIpGeoIpNotMatch? GeoIpNotMatch;
        public readonly Outputs.SwsWafProfileExclusionRuleConditionSourceIpIpRangesMatch? IpRangesMatch;
        public readonly Outputs.SwsWafProfileExclusionRuleConditionSourceIpIpRangesNotMatch? IpRangesNotMatch;

        [OutputConstructor]
        private SwsWafProfileExclusionRuleConditionSourceIp(
            Outputs.SwsWafProfileExclusionRuleConditionSourceIpGeoIpMatch? geoIpMatch,

            Outputs.SwsWafProfileExclusionRuleConditionSourceIpGeoIpNotMatch? geoIpNotMatch,

            Outputs.SwsWafProfileExclusionRuleConditionSourceIpIpRangesMatch? ipRangesMatch,

            Outputs.SwsWafProfileExclusionRuleConditionSourceIpIpRangesNotMatch? ipRangesNotMatch)
        {
            GeoIpMatch = geoIpMatch;
            GeoIpNotMatch = geoIpNotMatch;
            IpRangesMatch = ipRangesMatch;
            IpRangesNotMatch = ipRangesNotMatch;
        }
    }
}
