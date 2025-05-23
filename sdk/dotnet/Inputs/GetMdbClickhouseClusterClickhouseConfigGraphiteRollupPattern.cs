// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetMdbClickhouseClusterClickhouseConfigGraphiteRollupPatternArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Aggregation function name.
        /// </summary>
        [Input("function")]
        public string? Function { get; set; }

        /// <summary>
        /// Regular expression that the metric name must match.
        /// </summary>
        [Input("regexp", required: true)]
        public string Regexp { get; set; } = null!;

        [Input("retentions")]
        private List<Inputs.GetMdbClickhouseClusterClickhouseConfigGraphiteRollupPatternRetentionArgs>? _retentions;

        /// <summary>
        /// Retain parameters.
        /// </summary>
        public List<Inputs.GetMdbClickhouseClusterClickhouseConfigGraphiteRollupPatternRetentionArgs> Retentions
        {
            get => _retentions ?? (_retentions = new List<Inputs.GetMdbClickhouseClusterClickhouseConfigGraphiteRollupPatternRetentionArgs>());
            set => _retentions = value;
        }

        public GetMdbClickhouseClusterClickhouseConfigGraphiteRollupPatternArgs()
        {
        }
        public static new GetMdbClickhouseClusterClickhouseConfigGraphiteRollupPatternArgs Empty => new GetMdbClickhouseClusterClickhouseConfigGraphiteRollupPatternArgs();
    }
}
