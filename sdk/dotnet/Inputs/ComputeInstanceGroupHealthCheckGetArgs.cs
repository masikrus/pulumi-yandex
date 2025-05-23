// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class ComputeInstanceGroupHealthCheckGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of successful health checks before the managed instance is declared healthy.
        /// </summary>
        [Input("healthyThreshold")]
        public Input<int>? HealthyThreshold { get; set; }

        /// <summary>
        /// HTTP check options.
        /// </summary>
        [Input("httpOptions")]
        public Input<Inputs.ComputeInstanceGroupHealthCheckHttpOptionsGetArgs>? HttpOptions { get; set; }

        /// <summary>
        /// The interval to wait between health checks in seconds.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// TCP check options.
        /// </summary>
        [Input("tcpOptions")]
        public Input<Inputs.ComputeInstanceGroupHealthCheckTcpOptionsGetArgs>? TcpOptions { get; set; }

        /// <summary>
        /// The length of time to wait for a response before the health check times out in seconds.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The number of failed health checks before the managed instance is declared unhealthy.
        /// </summary>
        [Input("unhealthyThreshold")]
        public Input<int>? UnhealthyThreshold { get; set; }

        public ComputeInstanceGroupHealthCheckGetArgs()
        {
        }
        public static new ComputeInstanceGroupHealthCheckGetArgs Empty => new ComputeInstanceGroupHealthCheckGetArgs();
    }
}
