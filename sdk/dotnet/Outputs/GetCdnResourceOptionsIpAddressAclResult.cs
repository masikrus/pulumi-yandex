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
    public sealed class GetCdnResourceOptionsIpAddressAclResult
    {
        /// <summary>
        /// The list of specified IP addresses to be allowed or denied depending on acl policy type.
        /// </summary>
        public readonly ImmutableArray<string> ExceptedValues;
        /// <summary>
        /// The policy type for ACL. One of `allow` or `deny` values.
        /// </summary>
        public readonly string PolicyType;

        [OutputConstructor]
        private GetCdnResourceOptionsIpAddressAclResult(
            ImmutableArray<string> exceptedValues,

            string policyType)
        {
            ExceptedValues = exceptedValues;
            PolicyType = policyType;
        }
    }
}
