// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class DatatransferEndpointSettingsClickhouseTargetShardingGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Shard data by the hash value of the specified column.
        /// </summary>
        [Input("columnValueHash")]
        public Input<Inputs.DatatransferEndpointSettingsClickhouseTargetShardingColumnValueHashGetArgs>? ColumnValueHash { get; set; }

        /// <summary>
        /// A custom shard mapping by the value of the specified column.
        /// </summary>
        [Input("customMapping")]
        public Input<Inputs.DatatransferEndpointSettingsClickhouseTargetShardingCustomMappingGetArgs>? CustomMapping { get; set; }

        /// <summary>
        /// Distribute incoming rows between ClickHouse shards in a round-robin manner. Specify as an empty block to enable.
        /// </summary>
        [Input("roundRobin")]
        public Input<Inputs.DatatransferEndpointSettingsClickhouseTargetShardingRoundRobinGetArgs>? RoundRobin { get; set; }

        /// <summary>
        /// Shard data by ID of the transfer.
        /// </summary>
        [Input("transferId")]
        public Input<Inputs.DatatransferEndpointSettingsClickhouseTargetShardingTransferIdGetArgs>? TransferId { get; set; }

        public DatatransferEndpointSettingsClickhouseTargetShardingGetArgs()
        {
        }
        public static new DatatransferEndpointSettingsClickhouseTargetShardingGetArgs Empty => new DatatransferEndpointSettingsClickhouseTargetShardingGetArgs();
    }
}
