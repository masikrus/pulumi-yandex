// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetCmCertificate
    {
        public static Task<GetCmCertificateResult> InvokeAsync(GetCmCertificateArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCmCertificateResult>("yandex:index/getCmCertificate:getCmCertificate", args ?? new GetCmCertificateArgs(), options.WithDefaults());

        public static Output<GetCmCertificateResult> Invoke(GetCmCertificateInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCmCertificateResult>("yandex:index/getCmCertificate:getCmCertificate", args ?? new GetCmCertificateInvokeArgs(), options.WithDefaults());

        public static Output<GetCmCertificateResult> Invoke(GetCmCertificateInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCmCertificateResult>("yandex:index/getCmCertificate:getCmCertificate", args ?? new GetCmCertificateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCmCertificateArgs : global::Pulumi.InvokeArgs
    {
        [Input("certificateId")]
        public string? CertificateId { get; set; }

        [Input("description")]
        public string? Description { get; set; }

        [Input("folderId")]
        public string? FolderId { get; set; }

        [Input("labels")]
        private Dictionary<string, string>? _labels;
        public Dictionary<string, string> Labels
        {
            get => _labels ?? (_labels = new Dictionary<string, string>());
            set => _labels = value;
        }

        [Input("name")]
        public string? Name { get; set; }

        [Input("waitValidation")]
        public bool? WaitValidation { get; set; }

        public GetCmCertificateArgs()
        {
        }
        public static new GetCmCertificateArgs Empty => new GetCmCertificateArgs();
    }

    public sealed class GetCmCertificateInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("certificateId")]
        public Input<string>? CertificateId { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("waitValidation")]
        public Input<bool>? WaitValidation { get; set; }

        public GetCmCertificateInvokeArgs()
        {
        }
        public static new GetCmCertificateInvokeArgs Empty => new GetCmCertificateInvokeArgs();
    }


    [OutputType]
    public sealed class GetCmCertificateResult
    {
        public readonly string CertificateId;
        public readonly ImmutableArray<Outputs.GetCmCertificateChallengeResult> Challenges;
        public readonly string CreatedAt;
        public readonly bool DeletionProtection;
        public readonly string Description;
        public readonly ImmutableArray<string> Domains;
        public readonly string FolderId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string IssuedAt;
        public readonly string Issuer;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly string Name;
        public readonly string NotAfter;
        public readonly string NotBefore;
        public readonly string Serial;
        public readonly string Status;
        public readonly string Subject;
        public readonly string Type;
        public readonly string UpdatedAt;
        public readonly bool? WaitValidation;

        [OutputConstructor]
        private GetCmCertificateResult(
            string certificateId,

            ImmutableArray<Outputs.GetCmCertificateChallengeResult> challenges,

            string createdAt,

            bool deletionProtection,

            string description,

            ImmutableArray<string> domains,

            string folderId,

            string id,

            string issuedAt,

            string issuer,

            ImmutableDictionary<string, string> labels,

            string name,

            string notAfter,

            string notBefore,

            string serial,

            string status,

            string subject,

            string type,

            string updatedAt,

            bool? waitValidation)
        {
            CertificateId = certificateId;
            Challenges = challenges;
            CreatedAt = createdAt;
            DeletionProtection = deletionProtection;
            Description = description;
            Domains = domains;
            FolderId = folderId;
            Id = id;
            IssuedAt = issuedAt;
            Issuer = issuer;
            Labels = labels;
            Name = name;
            NotAfter = notAfter;
            NotBefore = notBefore;
            Serial = serial;
            Status = status;
            Subject = subject;
            Type = type;
            UpdatedAt = updatedAt;
            WaitValidation = waitValidation;
        }
    }
}
