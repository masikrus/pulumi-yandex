// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetAlbBackendGroupHttpBackendHealthcheckStreamHealthcheckArgs : global::Pulumi.InvokeArgs
    {
        [Input("receive", required: true)]
        public string Receive { get; set; } = null!;

        [Input("send", required: true)]
        public string Send { get; set; } = null!;

        public GetAlbBackendGroupHttpBackendHealthcheckStreamHealthcheckArgs()
        {
        }
        public static new GetAlbBackendGroupHttpBackendHealthcheckStreamHealthcheckArgs Empty => new GetAlbBackendGroupHttpBackendHealthcheckStreamHealthcheckArgs();
    }
}
