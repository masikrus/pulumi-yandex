// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetFunctionMountEphemeralDiskInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("blockSizeKb", required: true)]
        public Input<int> BlockSizeKb { get; set; } = null!;

        [Input("sizeGb", required: true)]
        public Input<int> SizeGb { get; set; } = null!;

        public GetFunctionMountEphemeralDiskInputArgs()
        {
        }
        public static new GetFunctionMountEphemeralDiskInputArgs Empty => new GetFunctionMountEphemeralDiskInputArgs();
    }
}
