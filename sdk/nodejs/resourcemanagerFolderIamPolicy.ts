// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class ResourcemanagerFolderIamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ResourcemanagerFolderIamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResourcemanagerFolderIamPolicyState, opts?: pulumi.CustomResourceOptions): ResourcemanagerFolderIamPolicy {
        return new ResourcemanagerFolderIamPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/resourcemanagerFolderIamPolicy:ResourcemanagerFolderIamPolicy';

    /**
     * Returns true if the given object is an instance of ResourcemanagerFolderIamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResourcemanagerFolderIamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResourcemanagerFolderIamPolicy.__pulumiType;
    }

    /**
     * The ID of the folder to attach a policy to.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * Required only by `yandex.IamServiceAccountIamPolicy`. The policy data generated by a `yandex.getIamPolicy` data source.
     */
    public readonly policyData!: pulumi.Output<string>;

    /**
     * Create a ResourcemanagerFolderIamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResourcemanagerFolderIamPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResourcemanagerFolderIamPolicyArgs | ResourcemanagerFolderIamPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ResourcemanagerFolderIamPolicyState | undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["policyData"] = state ? state.policyData : undefined;
        } else {
            const args = argsOrState as ResourcemanagerFolderIamPolicyArgs | undefined;
            if ((!args || args.folderId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'folderId'");
            }
            if ((!args || args.policyData === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyData'");
            }
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["policyData"] = args ? args.policyData : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ResourcemanagerFolderIamPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ResourcemanagerFolderIamPolicy resources.
 */
export interface ResourcemanagerFolderIamPolicyState {
    /**
     * The ID of the folder to attach a policy to.
     */
    folderId?: pulumi.Input<string>;
    /**
     * Required only by `yandex.IamServiceAccountIamPolicy`. The policy data generated by a `yandex.getIamPolicy` data source.
     */
    policyData?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ResourcemanagerFolderIamPolicy resource.
 */
export interface ResourcemanagerFolderIamPolicyArgs {
    /**
     * The ID of the folder to attach a policy to.
     */
    folderId: pulumi.Input<string>;
    /**
     * Required only by `yandex.IamServiceAccountIamPolicy`. The policy data generated by a `yandex.getIamPolicy` data source.
     */
    policyData: pulumi.Input<string>;
}
