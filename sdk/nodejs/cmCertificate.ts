// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class CmCertificate extends pulumi.CustomResource {
    /**
     * Get an existing CmCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CmCertificateState, opts?: pulumi.CustomResourceOptions): CmCertificate {
        return new CmCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/cmCertificate:CmCertificate';

    /**
     * Returns true if the given object is an instance of CmCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CmCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CmCertificate.__pulumiType;
    }

    /**
     * Array of challenges.
     */
    public /*out*/ readonly challenges!: pulumi.Output<outputs.CmCertificateChallenge[]>;
    /**
     * The creation timestamp of the resource.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The `true` value means that resource is protected from accidental deletion.
     */
    public readonly deletionProtection!: pulumi.Output<boolean | undefined>;
    /**
     * The resource description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Domains for this certificate. Should be specified for managed certificates.
     */
    public readonly domains!: pulumi.Output<string[] | undefined>;
    /**
     * The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * Certificate issue timestamp.
     */
    public /*out*/ readonly issuedAt!: pulumi.Output<string>;
    /**
     * Certificate Issuer.
     */
    public /*out*/ readonly issuer!: pulumi.Output<string>;
    /**
     * A set of key/value label pairs which assigned to resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Managed specification. > Resource creation awaits getting challenges from issue provider.
     */
    public readonly managed!: pulumi.Output<outputs.CmCertificateManaged | undefined>;
    /**
     * The resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Certificate end valid period.
     */
    public /*out*/ readonly notAfter!: pulumi.Output<string>;
    /**
     * Certificate start valid period.
     */
    public /*out*/ readonly notBefore!: pulumi.Output<string>;
    /**
     * Self-managed specification. > Only one type `privateKey` or `privateKeyLockboxSecret` should be specified.
     */
    public readonly selfManaged!: pulumi.Output<outputs.CmCertificateSelfManaged | undefined>;
    /**
     * Certificate Serial Number.
     */
    public /*out*/ readonly serial!: pulumi.Output<string>;
    /**
     * Certificate status: `VALIDATING`, `INVALID`, `ISSUED`, `REVOKED`, `RENEWING` or `RENEWAL_FAILED`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Certificate Subject.
     */
    public /*out*/ readonly subject!: pulumi.Output<string>;
    /**
     * Certificate type: `MANAGED` or `IMPORTED`.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Certificate update timestamp.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a CmCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: CmCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CmCertificateArgs | CmCertificateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CmCertificateState | undefined;
            resourceInputs["challenges"] = state ? state.challenges : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["domains"] = state ? state.domains : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["issuedAt"] = state ? state.issuedAt : undefined;
            resourceInputs["issuer"] = state ? state.issuer : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["managed"] = state ? state.managed : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["notAfter"] = state ? state.notAfter : undefined;
            resourceInputs["notBefore"] = state ? state.notBefore : undefined;
            resourceInputs["selfManaged"] = state ? state.selfManaged : undefined;
            resourceInputs["serial"] = state ? state.serial : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subject"] = state ? state.subject : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as CmCertificateArgs | undefined;
            resourceInputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["domains"] = args ? args.domains : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["managed"] = args ? args.managed : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["selfManaged"] = args ? args.selfManaged : undefined;
            resourceInputs["challenges"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["issuedAt"] = undefined /*out*/;
            resourceInputs["issuer"] = undefined /*out*/;
            resourceInputs["notAfter"] = undefined /*out*/;
            resourceInputs["notBefore"] = undefined /*out*/;
            resourceInputs["serial"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["subject"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CmCertificate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CmCertificate resources.
 */
export interface CmCertificateState {
    /**
     * Array of challenges.
     */
    challenges?: pulumi.Input<pulumi.Input<inputs.CmCertificateChallenge>[]>;
    /**
     * The creation timestamp of the resource.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The `true` value means that resource is protected from accidental deletion.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * The resource description.
     */
    description?: pulumi.Input<string>;
    /**
     * Domains for this certificate. Should be specified for managed certificates.
     */
    domains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Certificate issue timestamp.
     */
    issuedAt?: pulumi.Input<string>;
    /**
     * Certificate Issuer.
     */
    issuer?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs which assigned to resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Managed specification. > Resource creation awaits getting challenges from issue provider.
     */
    managed?: pulumi.Input<inputs.CmCertificateManaged>;
    /**
     * The resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * Certificate end valid period.
     */
    notAfter?: pulumi.Input<string>;
    /**
     * Certificate start valid period.
     */
    notBefore?: pulumi.Input<string>;
    /**
     * Self-managed specification. > Only one type `privateKey` or `privateKeyLockboxSecret` should be specified.
     */
    selfManaged?: pulumi.Input<inputs.CmCertificateSelfManaged>;
    /**
     * Certificate Serial Number.
     */
    serial?: pulumi.Input<string>;
    /**
     * Certificate status: `VALIDATING`, `INVALID`, `ISSUED`, `REVOKED`, `RENEWING` or `RENEWAL_FAILED`.
     */
    status?: pulumi.Input<string>;
    /**
     * Certificate Subject.
     */
    subject?: pulumi.Input<string>;
    /**
     * Certificate type: `MANAGED` or `IMPORTED`.
     */
    type?: pulumi.Input<string>;
    /**
     * Certificate update timestamp.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CmCertificate resource.
 */
export interface CmCertificateArgs {
    /**
     * The `true` value means that resource is protected from accidental deletion.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * The resource description.
     */
    description?: pulumi.Input<string>;
    /**
     * Domains for this certificate. Should be specified for managed certificates.
     */
    domains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs which assigned to resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Managed specification. > Resource creation awaits getting challenges from issue provider.
     */
    managed?: pulumi.Input<inputs.CmCertificateManaged>;
    /**
     * The resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * Self-managed specification. > Only one type `privateKey` or `privateKeyLockboxSecret` should be specified.
     */
    selfManaged?: pulumi.Input<inputs.CmCertificateSelfManaged>;
}
