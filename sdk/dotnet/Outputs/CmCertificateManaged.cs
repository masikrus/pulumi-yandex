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
    public sealed class CmCertificateManaged
    {
        /// <summary>
        /// Expected number of challenge count needed to validate certificate. Resource creation will fail if the specified value does not match the actual number of challenges received from issue provider. This argument is helpful for safe automatic resource creation for passing challenges for multi-domain certificates.
        /// </summary>
        public readonly int? ChallengeCount;
        /// <summary>
        /// Domain owner-check method. Possible values:
        /// * `DNS_CNAME` - you will need to create a CNAME dns record with the specified value. Recommended for fully automated certificate renewal.
        /// * `DNS_TXT` - you will need to create a TXT dns record with specified value.
        /// * `HTTP` - you will need to place specified value into specified url.
        /// </summary>
        public readonly string ChallengeType;

        [OutputConstructor]
        private CmCertificateManaged(
            int? challengeCount,

            string challengeType)
        {
            ChallengeCount = challengeCount;
            ChallengeType = challengeType;
        }
    }
}
