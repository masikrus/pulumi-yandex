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
    public sealed class DatatransferTransferTransformationTransformerSharderTransformer
    {
        /// <summary>
        /// List of the columns to transfer to the target tables using lists of included and excluded columns.
        /// </summary>
        public readonly Outputs.DatatransferTransferTransformationTransformerSharderTransformerColumns? Columns;
        /// <summary>
        /// Number of shards.
        /// </summary>
        public readonly int? ShardsCount;
        /// <summary>
        /// Table filter.
        /// </summary>
        public readonly Outputs.DatatransferTransferTransformationTransformerSharderTransformerTables? Tables;

        [OutputConstructor]
        private DatatransferTransferTransformationTransformerSharderTransformer(
            Outputs.DatatransferTransferTransformationTransformerSharderTransformerColumns? columns,

            int? shardsCount,

            Outputs.DatatransferTransferTransformationTransformerSharderTransformerTables? tables)
        {
            Columns = columns;
            ShardsCount = shardsCount;
            Tables = tables;
        }
    }
}
