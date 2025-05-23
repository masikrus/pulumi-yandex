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
    public sealed class GetAlbBackendGroupSessionAffinityHeaderResult
    {
        /// <summary>
        /// The name of the request header that will be used
        /// </summary>
        public readonly string HeaderName;

        [OutputConstructor]
        private GetAlbBackendGroupSessionAffinityHeaderResult(string headerName)
        {
            HeaderName = headerName;
        }
    }
}
