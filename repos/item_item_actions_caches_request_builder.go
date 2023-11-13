package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ice0f71b4c3da01ae12e785c80778ea7158cd0f6291b8c08c147f7b41f8d021bc "github.com/octokit/go-sdk/repos/item/item/actions/caches"
)

// ItemItemActionsCachesRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\actions\caches
type ItemItemActionsCachesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemActionsCachesRequestBuilderDeleteQueryParameters deletes one or more GitHub Actions caches for a repository, using a complete cache key. By default, all caches that match the provided key are deleted, but you can optionally provide a Git ref to restrict deletions to caches that match both the provided key and the Git ref.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `actions:write` permission to use this endpoint.
type ItemItemActionsCachesRequestBuilderDeleteQueryParameters struct {
    // A key for identifying the cache.
    Key *string `uriparametername:"key"`
    // The full Git reference for narrowing down the cache. The `ref` for a branch should be formatted as `refs/heads/<branch name>`. To reference a pull request use `refs/pull/<number>/merge`.
    Ref *string `uriparametername:"ref"`
}
// ItemItemActionsCachesRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsCachesRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemActionsCachesRequestBuilderDeleteQueryParameters
}
// ItemItemActionsCachesRequestBuilderGetQueryParameters lists the GitHub Actions caches for a repository.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `actions:read` permission to use this endpoint.
type ItemItemActionsCachesRequestBuilderGetQueryParameters struct {
    // The direction to sort the results by.
    // Deprecated: This property is deprecated, use directionAsGetDirectionQueryParameterType instead
    Direction *string `uriparametername:"direction"`
    // The direction to sort the results by.
    DirectionAsGetDirectionQueryParameterType *ice0f71b4c3da01ae12e785c80778ea7158cd0f6291b8c08c147f7b41f8d021bc.GetDirectionQueryParameterType `uriparametername:"direction"`
    // An explicit key or prefix for identifying the cache
    Key *string `uriparametername:"key"`
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
    // The full Git reference for narrowing down the cache. The `ref` for a branch should be formatted as `refs/heads/<branch name>`. To reference a pull request use `refs/pull/<number>/merge`.
    Ref *string `uriparametername:"ref"`
    // The property to sort the results by. `created_at` means when the cache was created. `last_accessed_at` means when the cache was last accessed. `size_in_bytes` is the size of the cache in bytes.
    // Deprecated: This property is deprecated, use sortAsGetSortQueryParameterType instead
    Sort *string `uriparametername:"sort"`
    // The property to sort the results by. `created_at` means when the cache was created. `last_accessed_at` means when the cache was last accessed. `size_in_bytes` is the size of the cache in bytes.
    SortAsGetSortQueryParameterType *ice0f71b4c3da01ae12e785c80778ea7158cd0f6291b8c08c147f7b41f8d021bc.GetSortQueryParameterType `uriparametername:"sort"`
}
// ItemItemActionsCachesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemActionsCachesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemActionsCachesRequestBuilderGetQueryParameters
}
// ByCache_id gets an item from the go-sdk.repos.item.item.actions.caches.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemItemActionsCachesRequestBuilder) ByCache_id(cache_id string)(*ItemItemActionsCachesWithCache_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if cache_id != "" {
        urlTplParams["cache_id"] = cache_id
    }
    return NewItemItemActionsCachesWithCache_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByCache_idInteger gets an item from the go-sdk.repos.item.item.actions.caches.item collection
func (m *ItemItemActionsCachesRequestBuilder) ByCache_idInteger(cache_id int32)(*ItemItemActionsCachesWithCache_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["cache_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(cache_id), 10)
    return NewItemItemActionsCachesWithCache_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemActionsCachesRequestBuilderInternal instantiates a new CachesRequestBuilder and sets the default values.
func NewItemItemActionsCachesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsCachesRequestBuilder) {
    m := &ItemItemActionsCachesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/actions/caches{?per_page*,page*,ref*,key*,sort*,direction*}", pathParameters),
    }
    return m
}
// NewItemItemActionsCachesRequestBuilder instantiates a new CachesRequestBuilder and sets the default values.
func NewItemItemActionsCachesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemActionsCachesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemActionsCachesRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes one or more GitHub Actions caches for a repository, using a complete cache key. By default, all caches that match the provided key are deleted, but you can optionally provide a Git ref to restrict deletions to caches that match both the provided key and the Git ref.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `actions:write` permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/cache#delete-github-actions-caches-for-a-repository-using-a-cache-key
func (m *ItemItemActionsCachesRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemActionsCachesRequestBuilderDeleteRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ActionsCacheListable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateActionsCacheListFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ActionsCacheListable), nil
}
// Get lists the GitHub Actions caches for a repository.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `actions:read` permission to use this endpoint.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/actions/cache#list-github-actions-caches-for-a-repository
func (m *ItemItemActionsCachesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemActionsCachesRequestBuilderGetRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ActionsCacheListable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateActionsCacheListFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.ActionsCacheListable), nil
}
// ToDeleteRequestInformation deletes one or more GitHub Actions caches for a repository, using a complete cache key. By default, all caches that match the provided key are deleted, but you can optionally provide a Git ref to restrict deletions to caches that match both the provided key and the Git ref.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `actions:write` permission to use this endpoint.
func (m *ItemItemActionsCachesRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsCachesRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// ToGetRequestInformation lists the GitHub Actions caches for a repository.You must authenticate using an access token with the `repo` scope to use this endpoint.GitHub Apps must have the `actions:read` permission to use this endpoint.
func (m *ItemItemActionsCachesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemActionsCachesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemActionsCachesRequestBuilder) WithUrl(rawUrl string)(*ItemItemActionsCachesRequestBuilder) {
    return NewItemItemActionsCachesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
