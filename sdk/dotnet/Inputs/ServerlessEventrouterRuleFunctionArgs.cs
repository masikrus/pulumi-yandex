// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class ServerlessEventrouterRuleFunctionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Batch settings
        /// </summary>
        [Input("batchSettings")]
        public Input<Inputs.ServerlessEventrouterRuleFunctionBatchSettingsArgs>? BatchSettings { get; set; }

        /// <summary>
        /// Function ID
        /// </summary>
        [Input("functionId", required: true)]
        public Input<string> FunctionId { get; set; } = null!;

        /// <summary>
        /// Function tag
        /// </summary>
        [Input("functionTag")]
        public Input<string>? FunctionTag { get; set; }

        /// <summary>
        /// Service account which has call permission on the function
        /// </summary>
        [Input("serviceAccountId")]
        public Input<string>? ServiceAccountId { get; set; }

        public ServerlessEventrouterRuleFunctionArgs()
        {
        }
        public static new ServerlessEventrouterRuleFunctionArgs Empty => new ServerlessEventrouterRuleFunctionArgs();
    }
}
