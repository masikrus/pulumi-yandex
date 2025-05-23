// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex.Inputs
{

    public sealed class LockboxSecretPasswordPayloadSpecificationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// String of punctuation characters to exclude from the default. Requires `include_punctuation = true`. Default is empty.
        /// </summary>
        [Input("excludedPunctuation")]
        public Input<string>? ExcludedPunctuation { get; set; }

        /// <summary>
        /// Use digits in the generated password. Default is true.
        /// </summary>
        [Input("includeDigits")]
        public Input<bool>? IncludeDigits { get; set; }

        /// <summary>
        /// Use lowercase letters in the generated password. Default is true.
        /// </summary>
        [Input("includeLowercase")]
        public Input<bool>? IncludeLowercase { get; set; }

        /// <summary>
        /// Use punctuations (`!"#$%&amp;'()*+,-./:;&lt;=&gt;?@[\]^_`{|}~`) in the generated password. Default is true.
        /// </summary>
        [Input("includePunctuation")]
        public Input<bool>? IncludePunctuation { get; set; }

        /// <summary>
        /// Use capital letters in the generated password. Default is true.
        /// </summary>
        [Input("includeUppercase")]
        public Input<bool>? IncludeUppercase { get; set; }

        /// <summary>
        /// String of specific punctuation characters to use. Requires `include_punctuation = true`. Default is empty.
        /// </summary>
        [Input("includedPunctuation")]
        public Input<string>? IncludedPunctuation { get; set; }

        /// <summary>
        /// Length of generated password. Default is `36`.
        /// </summary>
        [Input("length")]
        public Input<int>? Length { get; set; }

        /// <summary>
        /// The key with which the generated password will be placed in the secret version.
        /// </summary>
        [Input("passwordKey", required: true)]
        public Input<string> PasswordKey { get; set; } = null!;

        public LockboxSecretPasswordPayloadSpecificationArgs()
        {
        }
        public static new LockboxSecretPasswordPayloadSpecificationArgs Empty => new LockboxSecretPasswordPayloadSpecificationArgs();
    }
}
