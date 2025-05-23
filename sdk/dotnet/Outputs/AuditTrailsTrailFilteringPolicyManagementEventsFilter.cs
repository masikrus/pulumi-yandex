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
    public sealed class AuditTrailsTrailFilteringPolicyManagementEventsFilter
    {
        /// <summary>
        /// Structure describing that events will be gathered from the specified resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.AuditTrailsTrailFilteringPolicyManagementEventsFilterResourceScope> ResourceScopes;

        [OutputConstructor]
        private AuditTrailsTrailFilteringPolicyManagementEventsFilter(ImmutableArray<Outputs.AuditTrailsTrailFilteringPolicyManagementEventsFilterResourceScope> resourceScopes)
        {
            ResourceScopes = resourceScopes;
        }
    }
}
