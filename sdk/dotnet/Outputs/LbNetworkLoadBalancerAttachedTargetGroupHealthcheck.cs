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
    public sealed class LbNetworkLoadBalancerAttachedTargetGroupHealthcheck
    {
        /// <summary>
        /// Number of successful health checks required in order to set the `HEALTHY` status for the target.
        /// </summary>
        public readonly int? HealthyThreshold;
        /// <summary>
        /// Options for HTTP health check.
        /// </summary>
        public readonly Outputs.LbNetworkLoadBalancerAttachedTargetGroupHealthcheckHttpOptions? HttpOptions;
        /// <summary>
        /// The interval between health checks. The default is 2 seconds.
        /// </summary>
        public readonly int? Interval;
        /// <summary>
        /// Name of the health check. The name must be unique for each target group that attached to a single load balancer.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Options for TCP health check.
        /// </summary>
        public readonly Outputs.LbNetworkLoadBalancerAttachedTargetGroupHealthcheckTcpOptions? TcpOptions;
        /// <summary>
        /// Timeout for a target to return a response for the health check. The default is 1 second.
        /// </summary>
        public readonly int? Timeout;
        /// <summary>
        /// Number of failed health checks before changing the status to `UNHEALTHY`. The default is 2.
        /// </summary>
        public readonly int? UnhealthyThreshold;

        [OutputConstructor]
        private LbNetworkLoadBalancerAttachedTargetGroupHealthcheck(
            int? healthyThreshold,

            Outputs.LbNetworkLoadBalancerAttachedTargetGroupHealthcheckHttpOptions? httpOptions,

            int? interval,

            string name,

            Outputs.LbNetworkLoadBalancerAttachedTargetGroupHealthcheckTcpOptions? tcpOptions,

            int? timeout,

            int? unhealthyThreshold)
        {
            HealthyThreshold = healthyThreshold;
            HttpOptions = httpOptions;
            Interval = interval;
            Name = name;
            TcpOptions = tcpOptions;
            Timeout = timeout;
            UnhealthyThreshold = unhealthyThreshold;
        }
    }
}
