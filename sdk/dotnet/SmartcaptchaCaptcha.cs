// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    [YandexResourceType("yandex:index/smartcaptchaCaptcha:SmartcaptchaCaptcha")]
    public partial class SmartcaptchaCaptcha : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of allowed host names, see [Domain validation](https://yandex.cloud/docs/smartcaptcha/concepts/domain-validation).
        /// </summary>
        [Output("allowedSites")]
        public Output<ImmutableArray<string>> AllowedSites { get; private set; } = null!;

        /// <summary>
        /// Additional task type of the captcha. Possible values: * `IMAGE_TEXT` - Text recognition: The user has to type a
        /// distorted text from the picture into a special field. * `SILHOUETTES` - Silhouettes: The user has to mark several icons
        /// from the picture in a particular order. * `KALEIDOSCOPE` - Kaleidoscope: The user has to build a picture from individual
        /// parts by shuffling them using a slider.
        /// </summary>
        [Output("challengeType")]
        public Output<string?> ChallengeType { get; private set; } = null!;

        /// <summary>
        /// Client key of the captcha, see [CAPTCHA keys](https://yandex.cloud/docs/smartcaptcha/concepts/keys).
        /// </summary>
        [Output("clientKey")]
        public Output<string> ClientKey { get; private set; } = null!;

        /// <summary>
        /// The `Cloud ID` which resource belongs to. If it is not provided, the default provider `cloud-id` is used.
        /// </summary>
        [Output("cloudId")]
        public Output<string> CloudId { get; private set; } = null!;

        /// <summary>
        /// Complexity of the captcha. Possible values: * `EASY` - High chance to pass pre-check and easy advanced challenge. *
        /// `MEDIUM` - Medium chance to pass pre-check and normal advanced challenge. * `HARD` - Little chance to pass pre-check and
        /// hard advanced challenge. * `FORCE_HARD` - Impossible to pass pre-check and hard advanced challenge.
        /// </summary>
        [Output("complexity")]
        public Output<string?> Complexity { get; private set; } = null!;

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
        /// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        /// </summary>
        [Output("folderId")]
        public Output<string> FolderId { get; private set; } = null!;

        /// <summary>
        /// The resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of variants to use in security_rules.
        /// </summary>
        [Output("overrideVariants")]
        public Output<ImmutableArray<Outputs.SmartcaptchaCaptchaOverrideVariant>> OverrideVariants { get; private set; } = null!;

        /// <summary>
        /// Basic check type of the captcha.Possible values: * `CHECKBOX` - User must click the 'I am not a robot' button. *
        /// `SLIDER` - User must move the slider from left to right.
        /// </summary>
        [Output("preCheckType")]
        public Output<string?> PreCheckType { get; private set; } = null!;

        /// <summary>
        /// List of security rules.
        /// </summary>
        [Output("securityRules")]
        public Output<ImmutableArray<Outputs.SmartcaptchaCaptchaSecurityRule>> SecurityRules { get; private set; } = null!;

        /// <summary>
        /// JSON with variables to define the captcha appearance. For more details see generated JSON in cloud console.
        /// </summary>
        [Output("styleJson")]
        public Output<string?> StyleJson { get; private set; } = null!;

        [Output("suspend")]
        public Output<bool> Suspend { get; private set; } = null!;

        /// <summary>
        /// Turn off host name check, see [Domain validation](https://yandex.cloud/docs/smartcaptcha/concepts/domain-validation).
        /// </summary>
        [Output("turnOffHostnameCheck")]
        public Output<bool?> TurnOffHostnameCheck { get; private set; } = null!;


        /// <summary>
        /// Create a SmartcaptchaCaptcha resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SmartcaptchaCaptcha(string name, SmartcaptchaCaptchaArgs? args = null, CustomResourceOptions? options = null)
            : base("yandex:index/smartcaptchaCaptcha:SmartcaptchaCaptcha", name, args ?? new SmartcaptchaCaptchaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SmartcaptchaCaptcha(string name, Input<string> id, SmartcaptchaCaptchaState? state = null, CustomResourceOptions? options = null)
            : base("yandex:index/smartcaptchaCaptcha:SmartcaptchaCaptcha", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SmartcaptchaCaptcha resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SmartcaptchaCaptcha Get(string name, Input<string> id, SmartcaptchaCaptchaState? state = null, CustomResourceOptions? options = null)
        {
            return new SmartcaptchaCaptcha(name, id, state, options);
        }
    }

    public sealed class SmartcaptchaCaptchaArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedSites")]
        private InputList<string>? _allowedSites;

        /// <summary>
        /// List of allowed host names, see [Domain validation](https://yandex.cloud/docs/smartcaptcha/concepts/domain-validation).
        /// </summary>
        public InputList<string> AllowedSites
        {
            get => _allowedSites ?? (_allowedSites = new InputList<string>());
            set => _allowedSites = value;
        }

        /// <summary>
        /// Additional task type of the captcha. Possible values: * `IMAGE_TEXT` - Text recognition: The user has to type a
        /// distorted text from the picture into a special field. * `SILHOUETTES` - Silhouettes: The user has to mark several icons
        /// from the picture in a particular order. * `KALEIDOSCOPE` - Kaleidoscope: The user has to build a picture from individual
        /// parts by shuffling them using a slider.
        /// </summary>
        [Input("challengeType")]
        public Input<string>? ChallengeType { get; set; }

        /// <summary>
        /// The `Cloud ID` which resource belongs to. If it is not provided, the default provider `cloud-id` is used.
        /// </summary>
        [Input("cloudId")]
        public Input<string>? CloudId { get; set; }

        /// <summary>
        /// Complexity of the captcha. Possible values: * `EASY` - High chance to pass pre-check and easy advanced challenge. *
        /// `MEDIUM` - Medium chance to pass pre-check and normal advanced challenge. * `HARD` - Little chance to pass pre-check and
        /// hard advanced challenge. * `FORCE_HARD` - Impossible to pass pre-check and hard advanced challenge.
        /// </summary>
        [Input("complexity")]
        public Input<string>? Complexity { get; set; }

        /// <summary>
        /// The `true` value means that resource is protected from accidental deletion.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        /// <summary>
        /// The resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("overrideVariants")]
        private InputList<Inputs.SmartcaptchaCaptchaOverrideVariantArgs>? _overrideVariants;

        /// <summary>
        /// List of variants to use in security_rules.
        /// </summary>
        public InputList<Inputs.SmartcaptchaCaptchaOverrideVariantArgs> OverrideVariants
        {
            get => _overrideVariants ?? (_overrideVariants = new InputList<Inputs.SmartcaptchaCaptchaOverrideVariantArgs>());
            set => _overrideVariants = value;
        }

        /// <summary>
        /// Basic check type of the captcha.Possible values: * `CHECKBOX` - User must click the 'I am not a robot' button. *
        /// `SLIDER` - User must move the slider from left to right.
        /// </summary>
        [Input("preCheckType")]
        public Input<string>? PreCheckType { get; set; }

        [Input("securityRules")]
        private InputList<Inputs.SmartcaptchaCaptchaSecurityRuleArgs>? _securityRules;

        /// <summary>
        /// List of security rules.
        /// </summary>
        public InputList<Inputs.SmartcaptchaCaptchaSecurityRuleArgs> SecurityRules
        {
            get => _securityRules ?? (_securityRules = new InputList<Inputs.SmartcaptchaCaptchaSecurityRuleArgs>());
            set => _securityRules = value;
        }

        /// <summary>
        /// JSON with variables to define the captcha appearance. For more details see generated JSON in cloud console.
        /// </summary>
        [Input("styleJson")]
        public Input<string>? StyleJson { get; set; }

        /// <summary>
        /// Turn off host name check, see [Domain validation](https://yandex.cloud/docs/smartcaptcha/concepts/domain-validation).
        /// </summary>
        [Input("turnOffHostnameCheck")]
        public Input<bool>? TurnOffHostnameCheck { get; set; }

        public SmartcaptchaCaptchaArgs()
        {
        }
        public static new SmartcaptchaCaptchaArgs Empty => new SmartcaptchaCaptchaArgs();
    }

    public sealed class SmartcaptchaCaptchaState : global::Pulumi.ResourceArgs
    {
        [Input("allowedSites")]
        private InputList<string>? _allowedSites;

        /// <summary>
        /// List of allowed host names, see [Domain validation](https://yandex.cloud/docs/smartcaptcha/concepts/domain-validation).
        /// </summary>
        public InputList<string> AllowedSites
        {
            get => _allowedSites ?? (_allowedSites = new InputList<string>());
            set => _allowedSites = value;
        }

        /// <summary>
        /// Additional task type of the captcha. Possible values: * `IMAGE_TEXT` - Text recognition: The user has to type a
        /// distorted text from the picture into a special field. * `SILHOUETTES` - Silhouettes: The user has to mark several icons
        /// from the picture in a particular order. * `KALEIDOSCOPE` - Kaleidoscope: The user has to build a picture from individual
        /// parts by shuffling them using a slider.
        /// </summary>
        [Input("challengeType")]
        public Input<string>? ChallengeType { get; set; }

        /// <summary>
        /// Client key of the captcha, see [CAPTCHA keys](https://yandex.cloud/docs/smartcaptcha/concepts/keys).
        /// </summary>
        [Input("clientKey")]
        public Input<string>? ClientKey { get; set; }

        /// <summary>
        /// The `Cloud ID` which resource belongs to. If it is not provided, the default provider `cloud-id` is used.
        /// </summary>
        [Input("cloudId")]
        public Input<string>? CloudId { get; set; }

        /// <summary>
        /// Complexity of the captcha. Possible values: * `EASY` - High chance to pass pre-check and easy advanced challenge. *
        /// `MEDIUM` - Medium chance to pass pre-check and normal advanced challenge. * `HARD` - Little chance to pass pre-check and
        /// hard advanced challenge. * `FORCE_HARD` - Impossible to pass pre-check and hard advanced challenge.
        /// </summary>
        [Input("complexity")]
        public Input<string>? Complexity { get; set; }

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
        /// The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
        /// </summary>
        [Input("folderId")]
        public Input<string>? FolderId { get; set; }

        /// <summary>
        /// The resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("overrideVariants")]
        private InputList<Inputs.SmartcaptchaCaptchaOverrideVariantGetArgs>? _overrideVariants;

        /// <summary>
        /// List of variants to use in security_rules.
        /// </summary>
        public InputList<Inputs.SmartcaptchaCaptchaOverrideVariantGetArgs> OverrideVariants
        {
            get => _overrideVariants ?? (_overrideVariants = new InputList<Inputs.SmartcaptchaCaptchaOverrideVariantGetArgs>());
            set => _overrideVariants = value;
        }

        /// <summary>
        /// Basic check type of the captcha.Possible values: * `CHECKBOX` - User must click the 'I am not a robot' button. *
        /// `SLIDER` - User must move the slider from left to right.
        /// </summary>
        [Input("preCheckType")]
        public Input<string>? PreCheckType { get; set; }

        [Input("securityRules")]
        private InputList<Inputs.SmartcaptchaCaptchaSecurityRuleGetArgs>? _securityRules;

        /// <summary>
        /// List of security rules.
        /// </summary>
        public InputList<Inputs.SmartcaptchaCaptchaSecurityRuleGetArgs> SecurityRules
        {
            get => _securityRules ?? (_securityRules = new InputList<Inputs.SmartcaptchaCaptchaSecurityRuleGetArgs>());
            set => _securityRules = value;
        }

        /// <summary>
        /// JSON with variables to define the captcha appearance. For more details see generated JSON in cloud console.
        /// </summary>
        [Input("styleJson")]
        public Input<string>? StyleJson { get; set; }

        [Input("suspend")]
        public Input<bool>? Suspend { get; set; }

        /// <summary>
        /// Turn off host name check, see [Domain validation](https://yandex.cloud/docs/smartcaptcha/concepts/domain-validation).
        /// </summary>
        [Input("turnOffHostnameCheck")]
        public Input<bool>? TurnOffHostnameCheck { get; set; }

        public SmartcaptchaCaptchaState()
        {
        }
        public static new SmartcaptchaCaptchaState Empty => new SmartcaptchaCaptchaState();
    }
}
