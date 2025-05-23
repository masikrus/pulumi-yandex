// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    public static class GetLockboxSecretVersion
    {
        public static Task<GetLockboxSecretVersionResult> InvokeAsync(GetLockboxSecretVersionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLockboxSecretVersionResult>("yandex:index/getLockboxSecretVersion:getLockboxSecretVersion", args ?? new GetLockboxSecretVersionArgs(), options.WithDefaults());

        public static Output<GetLockboxSecretVersionResult> Invoke(GetLockboxSecretVersionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLockboxSecretVersionResult>("yandex:index/getLockboxSecretVersion:getLockboxSecretVersion", args ?? new GetLockboxSecretVersionInvokeArgs(), options.WithDefaults());

        public static Output<GetLockboxSecretVersionResult> Invoke(GetLockboxSecretVersionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetLockboxSecretVersionResult>("yandex:index/getLockboxSecretVersion:getLockboxSecretVersion", args ?? new GetLockboxSecretVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLockboxSecretVersionArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId", required: true)]
        public string SecretId { get; set; } = null!;

        [Input("versionId")]
        public string? VersionId { get; set; }

        public GetLockboxSecretVersionArgs()
        {
        }
        public static new GetLockboxSecretVersionArgs Empty => new GetLockboxSecretVersionArgs();
    }

    public sealed class GetLockboxSecretVersionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("secretId", required: true)]
        public Input<string> SecretId { get; set; } = null!;

        [Input("versionId")]
        public Input<string>? VersionId { get; set; }

        public GetLockboxSecretVersionInvokeArgs()
        {
        }
        public static new GetLockboxSecretVersionInvokeArgs Empty => new GetLockboxSecretVersionInvokeArgs();
    }


    [OutputType]
    public sealed class GetLockboxSecretVersionResult
    {
        public readonly ImmutableArray<Outputs.GetLockboxSecretVersionEntryResult> Entries;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string SecretId;
        public readonly string VersionId;

        [OutputConstructor]
        private GetLockboxSecretVersionResult(
            ImmutableArray<Outputs.GetLockboxSecretVersionEntryResult> entries,

            string id,

            string secretId,

            string versionId)
        {
            Entries = entries;
            Id = id;
            SecretId = secretId;
            VersionId = versionId;
        }
    }
}
