package user

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CodespacesItemExportsWithExport_ItemRequestBuilder builds and executes requests for operations under \user\codespaces\{codespace_name}\exports\{export_id}
type CodespacesItemExportsWithExport_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CodespacesItemExportsWithExport_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CodespacesItemExportsWithExport_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCodespacesItemExportsWithExport_ItemRequestBuilderInternal instantiates a new WithExport_ItemRequestBuilder and sets the default values.
func NewCodespacesItemExportsWithExport_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesItemExportsWithExport_ItemRequestBuilder) {
    m := &CodespacesItemExportsWithExport_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/codespaces/{codespace_name}/exports/{export_id}", pathParameters),
    }
    return m
}
// NewCodespacesItemExportsWithExport_ItemRequestBuilder instantiates a new WithExport_ItemRequestBuilder and sets the default values.
func NewCodespacesItemExportsWithExport_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CodespacesItemExportsWithExport_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCodespacesItemExportsWithExport_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets information about an export of a codespace.You must authenticate using a personal access token with the `codespace` scope to use this endpoint.GitHub Apps must have read access to the `codespaces_lifecycle_admin` repository permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/codespaces/codespaces#get-details-about-a-codespace-export
func (m *CodespacesItemExportsWithExport_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CodespacesItemExportsWithExport_ItemRequestBuilderGetRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CodespaceExportDetailsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateCodespaceExportDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CodespaceExportDetailsable), nil
}
// ToGetRequestInformation gets information about an export of a codespace.You must authenticate using a personal access token with the `codespace` scope to use this endpoint.GitHub Apps must have read access to the `codespaces_lifecycle_admin` repository permission to use this endpoint.
func (m *CodespacesItemExportsWithExport_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CodespacesItemExportsWithExport_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *CodespacesItemExportsWithExport_ItemRequestBuilder) WithUrl(rawUrl string)(*CodespacesItemExportsWithExport_ItemRequestBuilder) {
    return NewCodespacesItemExportsWithExport_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
