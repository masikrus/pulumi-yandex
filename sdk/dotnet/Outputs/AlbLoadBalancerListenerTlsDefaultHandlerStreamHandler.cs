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
    public sealed class AlbLoadBalancerListenerTlsDefaultHandlerStreamHandler
    {
        /// <summary>
        /// Backend Group ID.
        /// </summary>
        public readonly string? BackendGroupId;
        /// <summary>
        /// The idle timeout is the interval after which the connection will be forcibly closed if no data has been transmitted or received on either the upstream or downstream connection. If not configured, the default idle timeout is 1 hour. Setting it to 0 disables the timeout.
        /// </summary>
        public readonly string? IdleTimeout;

        [OutputConstructor]
        private AlbLoadBalancerListenerTlsDefaultHandlerStreamHandler(
            string? backendGroupId,

            string? idleTimeout)
        {
            BackendGroupId = backendGroupId;
            IdleTimeout = idleTimeout;
        }
    }
}
