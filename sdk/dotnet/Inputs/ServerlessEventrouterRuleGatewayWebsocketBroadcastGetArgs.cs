// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class ServerlessEventrouterRuleGatewayWebsocketBroadcastGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Batch settings
        /// </summary>
        [Input("batchSettings")]
        public Input<Inputs.ServerlessEventrouterRuleGatewayWebsocketBroadcastBatchSettingsGetArgs>? BatchSettings { get; set; }

        /// <summary>
        /// Gateway ID
        /// </summary>
        [Input("gatewayId", required: true)]
        public Input<string> GatewayId { get; set; } = null!;

        /// <summary>
        /// Path
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// Service account which has permission for writing to websockets
        /// </summary>
        [Input("serviceAccountId", required: true)]
        public Input<string> ServiceAccountId { get; set; } = null!;

        public ServerlessEventrouterRuleGatewayWebsocketBroadcastGetArgs()
        {
        }
        public static new ServerlessEventrouterRuleGatewayWebsocketBroadcastGetArgs Empty => new ServerlessEventrouterRuleGatewayWebsocketBroadcastGetArgs();
    }
}
