// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    [YandexResourceType("yandex:index/iamWorkloadIdentityOidcFederation:IamWorkloadIdentityOidcFederation")]
    public partial class IamWorkloadIdentityOidcFederation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of trusted values for aud claim.
        /// </summary>
        [Output("audiences")]
        public Output<ImmutableArray<string>> Audiences { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Description of the OIDC workload identity federation.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Disabled flag.
        /// </summary>
        [Output("disabled")]
        public Output<bool?> Disabled { get; private set; } = null!;

        /// <summary>
        /// Enabled flag.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// Id of the OIDC workload identity federation.
        /// </summary>
        [Output("federationId")]
        public Output<string> FederationId { get; private set; } = null!;

        /// <summary>
        /// Id of the folder that the OIDC workload identity federation belongs to.
        /// </summary>
        [Output("folderId")]
        public Output<string> FolderId { get; private set; } = null!;

        /// <summary>
        /// Issuer identifier of the external IdP server to be used for authentication.
        /// </summary>
        [Output("issuer")]
        public Output<string> Issuer { get; private set; } = null!;

        /// <summary>
        /// URL reference to trusted keys in format of JSON Web Key Set.
        /// </summary>
        [Output("jwksUrl")]
        public Output<string> JwksUrl { get; private set; } = null!;

        /// <summary>
        /// Resource labels as key-value pairs.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Name of the OIDC workload identity federation. The name is unique within the folder.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a IamWorkloadIdentityOidcFederation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IamWorkloadIdentityOidcFederation(string name, IamWorkloadIdentityOidcFederationArgs args, CustomResourceOptions? options = null)
            : base("yandex:index/iamWorkloadIdentityOidcFederation:IamWorkloadIdentityOidcFederation", name, args ?? new IamWorkloadIdentityOidcFederationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IamWorkloadIdentityOidcFederation(string name, Input<string> id, IamWorkloadIdentityOidcFederationState? state = null, CustomResourceOptions? options = null)
            : base("yandex:index/iamWorkloadIdentityOidcFederation:IamWorkloadIdentityOidcFederation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IamWorkloadIdentityOidcFederation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IamWorkloadIdentityOidcFederation Get(string name, Input<string> id, IamWorkloadIdentityOidcFederationState? state = null, CustomResourceOptions? options = null)
        {
            return new IamWorkloadIdentityOidcFederation(name, id, state, options);
        }
    }

    public sealed class IamWorkloadIdentityOidcFederationArgs : global::Pulumi.ResourceArgs
    {
        [Input("audiences")]
        private InputList<string>? _audiences;

        /// <summary>
        /// List of trusted values for aud claim.
        /// </summary>
        public InputList<string> Audiences
        {
            get => _audiences ?? (_audiences = new InputList<string>());
            set => _audiences = value;
        }

        /// <summary>
        /// Description of the OIDC workload identity federation.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Disabled flag.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Id of the OIDC workload identity federation.
        /// </summary>
        [Input("federationId")]
        public Input<string>? FederationId { get; set; }

        /// <summary>
        /// Id of the folder that the OIDC workload identity federation belongs to.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        /// <summary>
        /// Issuer identifier of the external IdP server to be used for authentication.
        /// </summary>
        [Input("issuer", required: true)]
        public Input<string> Issuer { get; set; } = null!;

        /// <summary>
        /// URL reference to trusted keys in format of JSON Web Key Set.
        /// </summary>
        [Input("jwksUrl", required: true)]
        public Input<string> JwksUrl { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Resource labels as key-value pairs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Name of the OIDC workload identity federation. The name is unique within the folder.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public IamWorkloadIdentityOidcFederationArgs()
        {
        }
        public static new IamWorkloadIdentityOidcFederationArgs Empty => new IamWorkloadIdentityOidcFederationArgs();
    }

    public sealed class IamWorkloadIdentityOidcFederationState : global::Pulumi.ResourceArgs
    {
        [Input("audiences")]
        private InputList<string>? _audiences;

        /// <summary>
        /// List of trusted values for aud claim.
        /// </summary>
        public InputList<string> Audiences
        {
            get => _audiences ?? (_audiences = new InputList<string>());
            set => _audiences = value;
        }

        /// <summary>
        /// Creation timestamp.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Description of the OIDC workload identity federation.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Disabled flag.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Enabled flag.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Id of the OIDC workload identity federation.
        /// </summary>
        [Input("federationId")]
        public Input<string>? FederationId { get; set; }

        /// <summary>
        /// Id of the folder that the OIDC workload identity federation belongs to.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        /// <summary>
        /// Issuer identifier of the external IdP server to be used for authentication.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// URL reference to trusted keys in format of JSON Web Key Set.
        /// </summary>
        [Input("jwksUrl")]
        public Input<string>? JwksUrl { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Resource labels as key-value pairs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Name of the OIDC workload identity federation. The name is unique within the folder.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public IamWorkloadIdentityOidcFederationState()
        {
        }
        public static new IamWorkloadIdentityOidcFederationState Empty => new IamWorkloadIdentityOidcFederationState();
    }
}
