// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MdbClickhouseClusterClickhouseConfigGraphiteRollupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Graphite rollup configuration name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the column storing the metric name (Graphite sensor). Default value: Path.
        /// </summary>
        [Input("pathColumnName")]
        public Input<string>? PathColumnName { get; set; }

        [Input("patterns")]
        private InputList<Inputs.MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternArgs>? _patterns;

        /// <summary>
        /// Set of thinning rules.
        /// </summary>
        public InputList<Inputs.MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternArgs> Patterns
        {
            get => _patterns ?? (_patterns = new InputList<Inputs.MdbClickhouseClusterClickhouseConfigGraphiteRollupPatternArgs>());
            set => _patterns = value;
        }

        /// <summary>
        /// The name of the column storing the time of measuring the metric. Default value: Time.
        /// </summary>
        [Input("timeColumnName")]
        public Input<string>? TimeColumnName { get; set; }

        /// <summary>
        /// The name of the column storing the value of the metric at the time set in `time_column_name`. Default value: Value.
        /// </summary>
        [Input("valueColumnName")]
        public Input<string>? ValueColumnName { get; set; }

        /// <summary>
        /// The name of the column storing the version of the metric. Default value: Timestamp.
        /// </summary>
        [Input("versionColumnName")]
        public Input<string>? VersionColumnName { get; set; }

        public MdbClickhouseClusterClickhouseConfigGraphiteRollupArgs()
        {
        }
        public static new MdbClickhouseClusterClickhouseConfigGraphiteRollupArgs Empty => new MdbClickhouseClusterClickhouseConfigGraphiteRollupArgs();
    }
}
