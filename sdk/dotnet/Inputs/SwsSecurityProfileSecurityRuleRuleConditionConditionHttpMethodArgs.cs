// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class SwsSecurityProfileSecurityRuleRuleConditionConditionHttpMethodArgs : global::Pulumi.ResourceArgs
    {
        [Input("httpMethods")]
        private InputList<Inputs.SwsSecurityProfileSecurityRuleRuleConditionConditionHttpMethodHttpMethodArgs>? _httpMethods;
        public InputList<Inputs.SwsSecurityProfileSecurityRuleRuleConditionConditionHttpMethodHttpMethodArgs> HttpMethods
        {
            get => _httpMethods ?? (_httpMethods = new InputList<Inputs.SwsSecurityProfileSecurityRuleRuleConditionConditionHttpMethodHttpMethodArgs>());
            set => _httpMethods = value;
        }

        public SwsSecurityProfileSecurityRuleRuleConditionConditionHttpMethodArgs()
        {
        }
        public static new SwsSecurityProfileSecurityRuleRuleConditionConditionHttpMethodArgs Empty => new SwsSecurityProfileSecurityRuleRuleConditionConditionHttpMethodArgs();
    }
}
