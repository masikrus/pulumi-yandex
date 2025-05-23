// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Database name.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// Identifier of the Managed ClickHouse cluster.
        /// </summary>
        [Input("mdbClusterId")]
        public Input<string>? MdbClusterId { get; set; }

        /// <summary>
        /// Connection settings of the on-premise ClickHouse server.
        /// </summary>
        [Input("onPremise")]
        public Input<Inputs.DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsOnPremiseArgs>? OnPremise { get; set; }

        /// <summary>
        /// Password for the database access.
        /// </summary>
        [Input("password")]
        public Input<Inputs.DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsPasswordArgs>? Password { get; set; }

        /// <summary>
        /// User for database access.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsArgs()
        {
        }
        public static new DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsArgs Empty => new DatatransferEndpointSettingsClickhouseTargetConnectionConnectionOptionsArgs();
    }
}
