// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class SwsWafProfileCoreRuleSetRuleSetGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public SwsWafProfileCoreRuleSetRuleSetGetArgs()
        {
        }
        public static new SwsWafProfileCoreRuleSetRuleSetGetArgs Empty => new SwsWafProfileCoreRuleSetRuleSetGetArgs();
    }
}
