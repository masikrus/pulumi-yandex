// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class MdbClickhouseClusterClickhouseConfigRabbitmqArgs : global::Pulumi.ResourceArgs
    {
        [Input("password")]
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
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// RabbitMQ vhost. Default: `\`.
        /// </summary>
        [Input("vhost")]
        public Input<string>? Vhost { get; set; }

        public MdbClickhouseClusterClickhouseConfigRabbitmqArgs()
        {
        }
        public static new MdbClickhouseClusterClickhouseConfigRabbitmqArgs Empty => new MdbClickhouseClusterClickhouseConfigRabbitmqArgs();
    }
}
