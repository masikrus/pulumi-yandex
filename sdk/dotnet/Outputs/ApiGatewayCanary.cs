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
    public sealed class ApiGatewayCanary
    {
        /// <summary>
        /// A list of values for variables in gateway specification of canary release.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Variables;
        /// <summary>
        /// Percentage of requests, which will be processed by canary release.
        /// </summary>
        public readonly int? Weight;

        [OutputConstructor]
        private ApiGatewayCanary(
            ImmutableDictionary<string, string>? variables,

            int? weight)
        {
            Variables = variables;
            Weight = weight;
        }
    }
}
