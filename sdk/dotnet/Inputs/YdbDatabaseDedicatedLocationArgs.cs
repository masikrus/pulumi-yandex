// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class YdbDatabaseDedicatedLocationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Region for the Yandex Database cluster.
        /// </summary>
        [Input("region")]
        public Input<Inputs.YdbDatabaseDedicatedLocationRegionArgs>? Region { get; set; }

        public YdbDatabaseDedicatedLocationArgs()
        {
        }
        public static new YdbDatabaseDedicatedLocationArgs Empty => new YdbDatabaseDedicatedLocationArgs();
    }
}
