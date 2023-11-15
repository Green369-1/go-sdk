package users

import (
    "context"
    i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e "github.com/octokit/go-sdk/github/octokit/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemKeysRequestBuilder builds and executes requests for operations under \users\{username}\keys
type ItemKeysRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemKeysRequestBuilderGetQueryParameters lists the _verified_ public SSH keys for a user. This is accessible by anyone.
type ItemKeysRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemKeysRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemKeysRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemKeysRequestBuilderGetQueryParameters
}
// NewItemKeysRequestBuilderInternal instantiates a new KeysRequestBuilder and sets the default values.
func NewItemKeysRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemKeysRequestBuilder) {
    m := &ItemKeysRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{username}/keys{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemKeysRequestBuilder instantiates a new KeysRequestBuilder and sets the default values.
func NewItemKeysRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemKeysRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemKeysRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the _verified_ public SSH keys for a user. This is accessible by anyone.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/users/keys#list-public-keys-for-a-user
func (m *ItemKeysRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemKeysRequestBuilderGetRequestConfiguration)([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.KeySimpleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.CreateKeySimpleFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.KeySimpleable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i8bb20811a612dd15efa26f086111481a68f72cd9ac5da7a939a417131078d77e.KeySimpleable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists the _verified_ public SSH keys for a user. This is accessible by anyone.
func (m *ItemKeysRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemKeysRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemKeysRequestBuilder) WithUrl(rawUrl string)(*ItemKeysRequestBuilder) {
    return NewItemKeysRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}