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
    public sealed class AlbBackendGroupSessionAffinity
    {
        /// <summary>
        /// Requests received from the same IP are combined into a session. Stream backend groups only support session affinity by client IP address.
        /// </summary>
        public readonly Outputs.AlbBackendGroupSessionAffinityConnection? Connection;
        /// <summary>
        /// Requests with the same cookie value and the specified file name are combined into a session. Allowed only for `HTTP` and `gRPC` backend groups.
        /// </summary>
        public readonly Outputs.AlbBackendGroupSessionAffinityCookie? Cookie;
        /// <summary>
        /// Requests with the same value of the specified HTTP header, such as with user authentication data, are combined into a session. Allowed only for `HTTP` and `gRPC` backend groups.
        /// </summary>
        public readonly Outputs.AlbBackendGroupSessionAffinityHeader? Header;

        [OutputConstructor]
        private AlbBackendGroupSessionAffinity(
            Outputs.AlbBackendGroupSessionAffinityConnection? connection,

            Outputs.AlbBackendGroupSessionAffinityCookie? cookie,

            Outputs.AlbBackendGroupSessionAffinityHeader? header)
        {
            Connection = connection;
            Cookie = cookie;
            Header = header;
        }
    }
}
