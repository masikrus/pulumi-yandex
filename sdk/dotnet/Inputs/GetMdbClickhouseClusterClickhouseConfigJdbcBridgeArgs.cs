// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetMdbClickhouseClusterClickhouseConfigJdbcBridgeInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Host of jdbc bridge.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Port of jdbc bridge. Default value: 9019.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        public GetMdbClickhouseClusterClickhouseConfigJdbcBridgeInputArgs()
        {
        }
        public static new GetMdbClickhouseClusterClickhouseConfigJdbcBridgeInputArgs Empty => new GetMdbClickhouseClusterClickhouseConfigJdbcBridgeInputArgs();
    }
}
