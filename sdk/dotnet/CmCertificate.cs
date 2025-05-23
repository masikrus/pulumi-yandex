// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    [YandexResourceType("yandex:index/cmCertificate:CmCertificate")]
    public partial class CmCertificate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Array of challenges.
        /// </summary>
        [Output("challenges")]
        public Output<ImmutableArray<Outputs.CmCertificateChallenge>> Challenges { get; private set; } = null!;

        /// <summary>
        /// The creation timestamp of the resource.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The `true` value means that resource is protected from accidental deletion.
        /// </summary>
        [Output("deletionProtection")]
        public Output<bool?> DeletionProtection { get; private set; } = null!;

        /// <summary>
        /// The resource description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Domains for this certificate. Should be specified for managed certificates.
        /// </summary>
        [Output("domains")]
        public Output<ImmutableArray<string>> Domains { get; private set; } = null!;

        /// <summary>
        /// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        /// </summary>
        [Output("folderId")]
        public Output<string> FolderId { get; private set; } = null!;

        /// <summary>
        /// Certificate issue timestamp.
        /// </summary>
        [Output("issuedAt")]
        public Output<string> IssuedAt { get; private set; } = null!;

        /// <summary>
        /// Certificate Issuer.
        /// </summary>
        [Output("issuer")]
        public Output<string> Issuer { get; private set; } = null!;

        /// <summary>
        /// A set of key/value label pairs which assigned to resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Managed specification. &gt; Resource creation awaits getting challenges from issue provider.
        /// </summary>
        [Output("managed")]
        public Output<Outputs.CmCertificateManaged?> Managed { get; private set; } = null!;

        /// <summary>
        /// The resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Certificate end valid period.
        /// </summary>
        [Output("notAfter")]
        public Output<string> NotAfter { get; private set; } = null!;

        /// <summary>
        /// Certificate start valid period.
        /// </summary>
        [Output("notBefore")]
        public Output<string> NotBefore { get; private set; } = null!;

        /// <summary>
        /// Self-managed specification. &gt; Only one type `private_key` or `private_key_lockbox_secret` should be specified.
        /// </summary>
        [Output("selfManaged")]
        public Output<Outputs.CmCertificateSelfManaged?> SelfManaged { get; private set; } = null!;

        /// <summary>
        /// Certificate Serial Number.
        /// </summary>
        [Output("serial")]
        public Output<string> Serial { get; private set; } = null!;

        /// <summary>
        /// Certificate status: `VALIDATING`, `INVALID`, `ISSUED`, `REVOKED`, `RENEWING` or `RENEWAL_FAILED`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Certificate Subject.
        /// </summary>
        [Output("subject")]
        public Output<string> Subject { get; private set; } = null!;

        /// <summary>
        /// Certificate type: `MANAGED` or `IMPORTED`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Certificate update timestamp.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a CmCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CmCertificate(string name, CmCertificateArgs? args = null, CustomResourceOptions? options = null)
            : base("yandex:index/cmCertificate:CmCertificate", name, args ?? new CmCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CmCertificate(string name, Input<string> id, CmCertificateState? state = null, CustomResourceOptions? options = null)
            : base("yandex:index/cmCertificate:CmCertificate", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CmCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CmCertificate Get(string name, Input<string> id, CmCertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new CmCertificate(name, id, state, options);
        }
    }

    public sealed class CmCertificateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The `true` value means that resource is protected from accidental deletion.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// The resource description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("domains")]
        private InputList<string>? _domains;

        /// <summary>
        /// Domains for this certificate. Should be specified for managed certificates.
        /// </summary>
        public InputList<string> Domains
        {
            get => _domains ?? (_domains = new InputList<string>());
            set => _domains = value;
        }

        /// <summary>
        /// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs which assigned to resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Managed specification. &gt; Resource creation awaits getting challenges from issue provider.
        /// </summary>
        [Input("managed")]
        public Input<Inputs.CmCertificateManagedArgs>? Managed { get; set; }

        /// <summary>
        /// The resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Self-managed specification. &gt; Only one type `private_key` or `private_key_lockbox_secret` should be specified.
        /// </summary>
        [Input("selfManaged")]
        public Input<Inputs.CmCertificateSelfManagedArgs>? SelfManaged { get; set; }

        public CmCertificateArgs()
        {
        }
        public static new CmCertificateArgs Empty => new CmCertificateArgs();
    }

    public sealed class CmCertificateState : global::Pulumi.ResourceArgs
    {
        [Input("challenges")]
        private InputList<Inputs.CmCertificateChallengeGetArgs>? _challenges;

        /// <summary>
        /// Array of challenges.
        /// </summary>
        public InputList<Inputs.CmCertificateChallengeGetArgs> Challenges
        {
            get => _challenges ?? (_challenges = new InputList<Inputs.CmCertificateChallengeGetArgs>());
            set => _challenges = value;
        }

        /// <summary>
        /// The creation timestamp of the resource.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The `true` value means that resource is protected from accidental deletion.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// The resource description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("domains")]
        private InputList<string>? _domains;

        /// <summary>
        /// Domains for this certificate. Should be specified for managed certificates.
        /// </summary>
        public InputList<string> Domains
        {
            get => _domains ?? (_domains = new InputList<string>());
            set => _domains = value;
        }

        /// <summary>
        /// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        /// <summary>
        /// Certificate issue timestamp.
        /// </summary>
        [Input("issuedAt")]
        public Input<string>? IssuedAt { get; set; }

        /// <summary>
        /// Certificate Issuer.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs which assigned to resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Managed specification. &gt; Resource creation awaits getting challenges from issue provider.
        /// </summary>
        [Input("managed")]
        public Input<Inputs.CmCertificateManagedGetArgs>? Managed { get; set; }

        /// <summary>
        /// The resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Certificate end valid period.
        /// </summary>
        [Input("notAfter")]
        public Input<string>? NotAfter { get; set; }

        /// <summary>
        /// Certificate start valid period.
        /// </summary>
        [Input("notBefore")]
        public Input<string>? NotBefore { get; set; }

        /// <summary>
        /// Self-managed specification. &gt; Only one type `private_key` or `private_key_lockbox_secret` should be specified.
        /// </summary>
        [Input("selfManaged")]
        public Input<Inputs.CmCertificateSelfManagedGetArgs>? SelfManaged { get; set; }

        /// <summary>
        /// Certificate Serial Number.
        /// </summary>
        [Input("serial")]
        public Input<string>? Serial { get; set; }

        /// <summary>
        /// Certificate status: `VALIDATING`, `INVALID`, `ISSUED`, `REVOKED`, `RENEWING` or `RENEWAL_FAILED`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Certificate Subject.
        /// </summary>
        [Input("subject")]
        public Input<string>? Subject { get; set; }

        /// <summary>
        /// Certificate type: `MANAGED` or `IMPORTED`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Certificate update timestamp.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public CmCertificateState()
        {
        }
        public static new CmCertificateState Empty => new CmCertificateState();
    }
}
