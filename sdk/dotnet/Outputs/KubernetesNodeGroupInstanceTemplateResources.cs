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
    public sealed class KubernetesNodeGroupInstanceTemplateResources
    {
        /// <summary>
        /// Baseline core performance as a percent.
        /// </summary>
        public readonly int? CoreFraction;
        /// <summary>
        /// Number of CPU cores allocated to the instance.
        /// </summary>
        public readonly int? Cores;
        /// <summary>
        /// Number of GPU cores allocated to the instance.
        /// </summary>
        public readonly int? Gpus;
        /// <summary>
        /// The memory size allocated to the instance.
        /// </summary>
        public readonly double? Memory;

        [OutputConstructor]
        private KubernetesNodeGroupInstanceTemplateResources(
            int? coreFraction,

            int? cores,

            int? gpus,

            double? memory)
        {
            CoreFraction = coreFraction;
            Cores = cores;
            Gpus = gpus;
            Memory = memory;
        }
    }
}
