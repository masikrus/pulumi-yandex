// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetMdbMongodbClusterClusterConfigPerformanceDiagnosticsInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable or disable performance diagnostics.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public GetMdbMongodbClusterClusterConfigPerformanceDiagnosticsInputArgs()
        {
        }
        public static new GetMdbMongodbClusterClusterConfigPerformanceDiagnosticsInputArgs Empty => new GetMdbMongodbClusterClusterConfigPerformanceDiagnosticsInputArgs();
    }
}
