// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class AlbHttpRouter extends pulumi.CustomResource {
    /**
     * Get an existing AlbHttpRouter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlbHttpRouterState, opts?: pulumi.CustomResourceOptions): AlbHttpRouter {
        return new AlbHttpRouter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/albHttpRouter:AlbHttpRouter';

    /**
     * Returns true if the given object is an instance of AlbHttpRouter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AlbHttpRouter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AlbHttpRouter.__pulumiType;
    }

    /**
     * The creation timestamp of the resource.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The resource description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * A set of key/value label pairs which assigned to resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Route options for the virtual host.
     */
    public readonly routeOptions!: pulumi.Output<outputs.AlbHttpRouterRouteOptions | undefined>;

    /**
     * Create a AlbHttpRouter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AlbHttpRouterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlbHttpRouterArgs | AlbHttpRouterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AlbHttpRouterState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["routeOptions"] = state ? state.routeOptions : undefined;
        } else {
            const args = argsOrState as AlbHttpRouterArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["routeOptions"] = args ? args.routeOptions : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AlbHttpRouter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AlbHttpRouter resources.
 */
export interface AlbHttpRouterState {
    /**
     * The creation timestamp of the resource.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The resource description.
     */
    description?: pulumi.Input<string>;
    /**
     * The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs which assigned to resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * Route options for the virtual host.
     */
    routeOptions?: pulumi.Input<inputs.AlbHttpRouterRouteOptions>;
}

/**
 * The set of arguments for constructing a AlbHttpRouter resource.
 */
export interface AlbHttpRouterArgs {
    /**
     * The resource description.
     */
    description?: pulumi.Input<string>;
    /**
     * The folder identifier that resource belongs to. If it is not provided, the default provider `folder-id` is used.
     */
    folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs which assigned to resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource name.
     */
    name?: pulumi.Input<string>;
    /**
     * Route options for the virtual host.
     */
    routeOptions?: pulumi.Input<inputs.AlbHttpRouterRouteOptions>;
}
