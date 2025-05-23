// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetCmCertificateContent
    {
        public static Task<GetCmCertificateContentResult> InvokeAsync(GetCmCertificateContentArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCmCertificateContentResult>("yandex:index/getCmCertificateContent:getCmCertificateContent", args ?? new GetCmCertificateContentArgs(), options.WithDefaults());

        public static Output<GetCmCertificateContentResult> Invoke(GetCmCertificateContentInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCmCertificateContentResult>("yandex:index/getCmCertificateContent:getCmCertificateContent", args ?? new GetCmCertificateContentInvokeArgs(), options.WithDefaults());

        public static Output<GetCmCertificateContentResult> Invoke(GetCmCertificateContentInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCmCertificateContentResult>("yandex:index/getCmCertificateContent:getCmCertificateContent", args ?? new GetCmCertificateContentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCmCertificateContentArgs : global::Pulumi.InvokeArgs
    {
        [Input("certificateId")]
        public string? CertificateId { get; set; }

        [Input("folderId")]
        public string? FolderId { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("privateKeyFormat")]
        public string? PrivateKeyFormat { get; set; }

        [Input("waitValidation")]
        public bool? WaitValidation { get; set; }

        public GetCmCertificateContentArgs()
        {
        }
        public static new GetCmCertificateContentArgs Empty => new GetCmCertificateContentArgs();
    }

    public sealed class GetCmCertificateContentInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("certificateId")]
        public Input<string>? CertificateId { get; set; }

        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("privateKeyFormat")]
        public Input<string>? PrivateKeyFormat { get; set; }

        [Input("waitValidation")]
        public Input<bool>? WaitValidation { get; set; }

        public GetCmCertificateContentInvokeArgs()
        {
        }
        public static new GetCmCertificateContentInvokeArgs Empty => new GetCmCertificateContentInvokeArgs();
    }


    [OutputType]
    public sealed class GetCmCertificateContentResult
    {
        public readonly string? CertificateId;
        public readonly ImmutableArray<string> Certificates;
        public readonly string? FolderId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        public readonly string PrivateKey;
        public readonly string? PrivateKeyFormat;
        public readonly bool? WaitValidation;

        [OutputConstructor]
        private GetCmCertificateContentResult(
            string? certificateId,

            ImmutableArray<string> certificates,

            string? folderId,

            string id,

            string? name,

            string privateKey,

            string? privateKeyFormat,

            bool? waitValidation)
        {
            CertificateId = certificateId;
            Certificates = certificates;
            FolderId = folderId;
            Id = id;
            Name = name;
            PrivateKey = privateKey;
            PrivateKeyFormat = privateKeyFormat;
            WaitValidation = waitValidation;
        }
    }
}
