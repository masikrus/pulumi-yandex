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
    public sealed class GetMdbGreenplumClusterBackgroundActivityQueryKillerIdleInTransactionResult
    {
        public readonly bool Enable;
        public readonly ImmutableArray<string> IgnoreUsers;
        public readonly int MaxAge;

        [OutputConstructor]
        private GetMdbGreenplumClusterBackgroundActivityQueryKillerIdleInTransactionResult(
            bool enable,

            ImmutableArray<string> ignoreUsers,

            int maxAge)
        {
            Enable = enable;
            IgnoreUsers = ignoreUsers;
            MaxAge = maxAge;
        }
    }
}
