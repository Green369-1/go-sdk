package orgs

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilder builds and executes requests for operations under \orgs\{org}\personal-access-token-requests\{pat_request_id}\repositories
type ItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilderGetQueryParameters lists the repositories a fine-grained personal access token request is requesting access to. Only GitHub Apps can call this API,using the `organization_personal_access_token_requests: read` permission.**Note**: Fine-grained PATs are in public beta. Related APIs, events, and functionality are subject to change.
type ItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilderGetQueryParameters
}
// NewItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilderInternal instantiates a new RepositoriesRequestBuilder and sets the default values.
func NewItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilder) {
    m := &ItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/personal-access-token-requests/{pat_request_id}/repositories{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilder instantiates a new RepositoriesRequestBuilder and sets the default values.
func NewItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the repositories a fine-grained personal access token request is requesting access to. Only GitHub Apps can call this API,using the `organization_personal_access_token_requests: read` permission.**Note**: Fine-grained PATs are in public beta. Related APIs, events, and functionality are subject to change.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/orgs/personal-access-tokens#list-repositories-requested-to-be-accessed-by-a-fine-grained-personal-access-token
func (m *ItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilderGetRequestConfiguration)([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.MinimalRepositoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "500": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateMinimalRepositoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.MinimalRepositoryable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.MinimalRepositoryable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists the repositories a fine-grained personal access token request is requesting access to. Only GitHub Apps can call this API,using the `organization_personal_access_token_requests: read` permission.**Note**: Fine-grained PATs are in public beta. Related APIs, events, and functionality are subject to change.
func (m *ItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
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
func (m *ItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilder) WithUrl(rawUrl string)(*ItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilder) {
    return NewItemPersonalAccessTokenRequestsItemRepositoriesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
