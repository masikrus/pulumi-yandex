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
    public sealed class GetServerlessEventrouterRuleFunctionBatchSettingResult
    {
        /// <summary>
        /// Maximum batch size: rule will send a batch if its lifetime exceeds this value
        /// </summary>
        public readonly string Cutoff;
        /// <summary>
        /// Maximum batch size: rule will send a batch if total size of events exceeds this value
        /// </summary>
        public readonly int MaxBytes;
        /// <summary>
        /// Maximum batch size: rule will send a batch if number of events exceeds this value
        /// </summary>
        public readonly int MaxCount;

        [OutputConstructor]
        private GetServerlessEventrouterRuleFunctionBatchSettingResult(
            string cutoff,

            int maxBytes,

            int maxCount)
        {
            Cutoff = cutoff;
            MaxBytes = maxBytes;
            MaxCount = maxCount;
        }
    }
}
