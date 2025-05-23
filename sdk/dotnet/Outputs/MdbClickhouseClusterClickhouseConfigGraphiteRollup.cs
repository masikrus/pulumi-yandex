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
    public sealed class MdbClickhouseClusterClickhouseConfigGraphiteRollup
    {
        /// <summary>
        /// Graphite rollup configuration name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The name of the column storing the metric name (Graphite sensor). Default value: Path.
        /// </summary>
        public readonly string? PathColumnName;
        /// <summary>
        /// Set of thinning rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.MdbClickhouseClusterClickhouseConfigGraphiteRollupPattern> Patterns;
        /// <summary>
        /// The name of the column storing the time of measuring the metric. Default value: Time.
        /// </summary>
        public readonly string? TimeColumnName;
        /// <summary>
        /// The name of the column storing the value of the metric at the time set in `time_column_name`. Default value: Value.
        /// </summary>
        public readonly string? ValueColumnName;
        /// <summary>
        /// The name of the column storing the version of the metric. Default value: Timestamp.
        /// </summary>
        public readonly string? VersionColumnName;

        [OutputConstructor]
        private MdbClickhouseClusterClickhouseConfigGraphiteRollup(
            string name,

            string? pathColumnName,

            ImmutableArray<Outputs.MdbClickhouseClusterClickhouseConfigGraphiteRollupPattern> patterns,

            string? timeColumnName,

            string? valueColumnName,

            string? versionColumnName)
        {
            Name = name;
            PathColumnName = pathColumnName;
            Patterns = patterns;
            TimeColumnName = timeColumnName;
            ValueColumnName = valueColumnName;
            VersionColumnName = versionColumnName;
        }
    }
}
