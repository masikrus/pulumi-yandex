// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class OrganizationmanagerGroupMapping extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationmanagerGroupMapping resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrganizationmanagerGroupMappingState, opts?: pulumi.CustomResourceOptions): OrganizationmanagerGroupMapping {
        return new OrganizationmanagerGroupMapping(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/organizationmanagerGroupMapping:OrganizationmanagerGroupMapping';

    /**
     * Returns true if the given object is an instance of OrganizationmanagerGroupMapping.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationmanagerGroupMapping {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationmanagerGroupMapping.__pulumiType;
    }

    /**
     * Set "true" to enable organization manager group mapping.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * ID of the SAML Federation.
     */
    public readonly federationId!: pulumi.Output<string>;

    /**
     * Create a OrganizationmanagerGroupMapping resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationmanagerGroupMappingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrganizationmanagerGroupMappingArgs | OrganizationmanagerGroupMappingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrganizationmanagerGroupMappingState | undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["federationId"] = state ? state.federationId : undefined;
        } else {
            const args = argsOrState as OrganizationmanagerGroupMappingArgs | undefined;
            if ((!args || args.federationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'federationId'");
            }
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["federationId"] = args ? args.federationId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OrganizationmanagerGroupMapping.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrganizationmanagerGroupMapping resources.
 */
export interface OrganizationmanagerGroupMappingState {
    /**
     * Set "true" to enable organization manager group mapping.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * ID of the SAML Federation.
     */
    federationId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OrganizationmanagerGroupMapping resource.
 */
export interface OrganizationmanagerGroupMappingArgs {
    /**
     * Set "true" to enable organization manager group mapping.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * ID of the SAML Federation.
     */
    federationId: pulumi.Input<string>;
}
