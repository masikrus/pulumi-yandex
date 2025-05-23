// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class OrganizationmanagerUserSshKey extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationmanagerUserSshKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrganizationmanagerUserSshKeyState, opts?: pulumi.CustomResourceOptions): OrganizationmanagerUserSshKey {
        return new OrganizationmanagerUserSshKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/organizationmanagerUserSshKey:OrganizationmanagerUserSshKey';

    /**
     * Returns true if the given object is an instance of OrganizationmanagerUserSshKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationmanagerUserSshKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationmanagerUserSshKey.__pulumiType;
    }

    /**
     * The creation timestamp of the resource.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Data of the user ssh key.
     */
    public readonly data!: pulumi.Output<string>;
    /**
     * User ssh key will be no longer valid after expiration timestamp.
     */
    public readonly expiresAt!: pulumi.Output<string | undefined>;
    /**
     * SSH Key Fingerprint.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * The resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Organization that the user ssh key belongs to.
     */
    public readonly organizationId!: pulumi.Output<string>;
    /**
     * Subject that the user ssh key belongs to.
     */
    public readonly subjectId!: pulumi.Output<string>;

    /**
     * Create a OrganizationmanagerUserSshKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationmanagerUserSshKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrganizationmanagerUserSshKeyArgs | OrganizationmanagerUserSshKeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrganizationmanagerUserSshKeyState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["data"] = state ? state.data : undefined;
            resourceInputs["expiresAt"] = state ? state.expiresAt : undefined;
            resourceInputs["fingerprint"] = state ? state.fingerprint : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["subjectId"] = state ? state.subjectId : undefined;
        } else {
            const args = argsOrState as OrganizationmanagerUserSshKeyArgs | undefined;
            if ((!args || args.data === undefined) && !opts.urn) {
                throw new Error("Missing required property 'data'");
            }
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            if ((!args || args.subjectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subjectId'");
            }
            resourceInputs["data"] = args ? args.data : undefined;
            resourceInputs["expiresAt"] = args ? args.expiresAt : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
            resourceInputs["subjectId"] = args ? args.subjectId : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OrganizationmanagerUserSshKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrganizationmanagerUserSshKey resources.
 */
export interface OrganizationmanagerUserSshKeyState {
    /**
     * The creation timestamp of the resource.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Data of the user ssh key.
     */
    data?: pulumi.Input<string>;
    /**
     * User ssh key will be no longer valid after expiration timestamp.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * SSH Key Fingerprint.
     */
    fingerprint?: pulumi.Input<string>;
    /**
     * The resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * Organization that the user ssh key belongs to.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * Subject that the user ssh key belongs to.
     */
    subjectId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OrganizationmanagerUserSshKey resource.
 */
export interface OrganizationmanagerUserSshKeyArgs {
    /**
     * Data of the user ssh key.
     */
    data: pulumi.Input<string>;
    /**
     * User ssh key will be no longer valid after expiration timestamp.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * The resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * Organization that the user ssh key belongs to.
     */
    organizationId: pulumi.Input<string>;
    /**
     * Subject that the user ssh key belongs to.
     */
    subjectId: pulumi.Input<string>;
}
