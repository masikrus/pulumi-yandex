// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class AuditTrailsTrailFilterEventFilterPathFilterGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("anyFilter")]
        public Input<Inputs.AuditTrailsTrailFilterEventFilterPathFilterAnyFilterGetArgs>? AnyFilter { get; set; }

        [Input("someFilter")]
        public Input<Inputs.AuditTrailsTrailFilterEventFilterPathFilterSomeFilterGetArgs>? SomeFilter { get; set; }

        public AuditTrailsTrailFilterEventFilterPathFilterGetArgs()
        {
        }
        public static new AuditTrailsTrailFilterEventFilterPathFilterGetArgs Empty => new AuditTrailsTrailFilterEventFilterPathFilterGetArgs();
    }
}
