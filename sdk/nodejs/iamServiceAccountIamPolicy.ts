// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class IamServiceAccountIamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing IamServiceAccountIamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IamServiceAccountIamPolicyState, opts?: pulumi.CustomResourceOptions): IamServiceAccountIamPolicy {
        return new IamServiceAccountIamPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/iamServiceAccountIamPolicy:IamServiceAccountIamPolicy';

    /**
     * Returns true if the given object is an instance of IamServiceAccountIamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IamServiceAccountIamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IamServiceAccountIamPolicy.__pulumiType;
    }

    /**
     * Required only by `yandex.IamServiceAccountIamPolicy`. The policy data generated by a `yandex.getIamPolicy` data source.
     */
    public readonly policyData!: pulumi.Output<string>;
    /**
     * The service account ID to apply a binding to.
     */
    public readonly serviceAccountId!: pulumi.Output<string>;

    /**
     * Create a IamServiceAccountIamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IamServiceAccountIamPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IamServiceAccountIamPolicyArgs | IamServiceAccountIamPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IamServiceAccountIamPolicyState | undefined;
            resourceInputs["policyData"] = state ? state.policyData : undefined;
            resourceInputs["serviceAccountId"] = state ? state.serviceAccountId : undefined;
        } else {
            const args = argsOrState as IamServiceAccountIamPolicyArgs | undefined;
            if ((!args || args.policyData === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyData'");
            }
            if ((!args || args.serviceAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceAccountId'");
            }
            resourceInputs["policyData"] = args ? args.policyData : undefined;
            resourceInputs["serviceAccountId"] = args ? args.serviceAccountId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IamServiceAccountIamPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IamServiceAccountIamPolicy resources.
 */
export interface IamServiceAccountIamPolicyState {
    /**
     * Required only by `yandex.IamServiceAccountIamPolicy`. The policy data generated by a `yandex.getIamPolicy` data source.
     */
    policyData?: pulumi.Input<string>;
    /**
     * The service account ID to apply a binding to.
     */
    serviceAccountId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IamServiceAccountIamPolicy resource.
 */
export interface IamServiceAccountIamPolicyArgs {
    /**
     * Required only by `yandex.IamServiceAccountIamPolicy`. The policy data generated by a `yandex.getIamPolicy` data source.
     */
    policyData: pulumi.Input<string>;
    /**
     * The service account ID to apply a binding to.
     */
    serviceAccountId: pulumi.Input<string>;
}
