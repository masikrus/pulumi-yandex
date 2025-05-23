// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetMdbMongodbClusterClusterConfigMongosInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A set of network settings (see the [net](https://www.mongodb.com/docs/manual/reference/configuration-options/#net-options) option).
        /// </summary>
        [Input("net")]
        public Input<Inputs.GetMdbMongodbClusterClusterConfigMongosNetInputArgs>? Net { get; set; }

        public GetMdbMongodbClusterClusterConfigMongosInputArgs()
        {
        }
        public static new GetMdbMongodbClusterClusterConfigMongosInputArgs Empty => new GetMdbMongodbClusterClusterConfigMongosInputArgs();
    }
}
