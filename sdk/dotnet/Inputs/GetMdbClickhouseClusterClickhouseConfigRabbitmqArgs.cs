// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class GetMdbClickhouseClusterClickhouseConfigRabbitmqInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// RabbitMQ user password.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// RabbitMQ username.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        /// <summary>
        /// RabbitMQ vhost. Default: `\`.
        /// </summary>
        [Input("vhost", required: true)]
        public Input<string> Vhost { get; set; } = null!;

        public GetMdbClickhouseClusterClickhouseConfigRabbitmqInputArgs()
        {
        }
        public static new GetMdbClickhouseClusterClickhouseConfigRabbitmqInputArgs Empty => new GetMdbClickhouseClusterClickhouseConfigRabbitmqInputArgs();
    }
}
