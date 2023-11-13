package users

import (
    "context"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ida1f2c756ef54dd41dc613dcf8bf5a0caf413c1958bf4ea869be1bfb06e03bbc "github.com/octokit/go-sdk/users/item/hovercard"
)

// ItemHovercardRequestBuilder builds and executes requests for operations under \users\{username}\hovercard
type ItemHovercardRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemHovercardRequestBuilderGetQueryParameters provides hovercard information when authenticated through basic auth or OAuth with the `repo` scope. You can find out more about someone in relation to their pull requests, issues, repositories, and organizations.The `subject_type` and `subject_id` parameters provide context for the person's hovercard, which returns more information than without the parameters. For example, if you wanted to find out more about `octocat` who owns the `Spoon-Knife` repository via cURL, it would look like this:```shell curl -u username:token  https://api.github.com/users/octocat/hovercard?subject_type=repository&subject_id=1300192```
type ItemHovercardRequestBuilderGetQueryParameters struct {
    // Uses the ID for the `subject_type` you specified. **Required** when using `subject_type`.
    Subject_id *string `uriparametername:"subject_id"`
    // Identifies which additional information you'd like to receive about the person's hovercard. Can be `organization`, `repository`, `issue`, `pull_request`. **Required** when using `subject_id`.
    // Deprecated: This property is deprecated, use subject_typeAsGetSubject_typeQueryParameterType instead
    Subject_type *string `uriparametername:"subject_type"`
    // Identifies which additional information you'd like to receive about the person's hovercard. Can be `organization`, `repository`, `issue`, `pull_request`. **Required** when using `subject_id`.
    Subject_typeAsGetSubject_typeQueryParameterType *ida1f2c756ef54dd41dc613dcf8bf5a0caf413c1958bf4ea869be1bfb06e03bbc.GetSubject_typeQueryParameterType `uriparametername:"subject_type"`
}
// ItemHovercardRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemHovercardRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemHovercardRequestBuilderGetQueryParameters
}
// NewItemHovercardRequestBuilderInternal instantiates a new HovercardRequestBuilder and sets the default values.
func NewItemHovercardRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemHovercardRequestBuilder) {
    m := &ItemHovercardRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{username}/hovercard{?subject_type*,subject_id*}", pathParameters),
    }
    return m
}
// NewItemHovercardRequestBuilder instantiates a new HovercardRequestBuilder and sets the default values.
func NewItemHovercardRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemHovercardRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemHovercardRequestBuilderInternal(urlParams, requestAdapter)
}
// Get provides hovercard information when authenticated through basic auth or OAuth with the `repo` scope. You can find out more about someone in relation to their pull requests, issues, repositories, and organizations.The `subject_type` and `subject_id` parameters provide context for the person's hovercard, which returns more information than without the parameters. For example, if you wanted to find out more about `octocat` who owns the `Spoon-Knife` repository via cURL, it would look like this:```shell curl -u username:token  https://api.github.com/users/octocat/hovercard?subject_type=repository&subject_id=1300192```
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/users/users#get-contextual-information-for-a-user
func (m *ItemHovercardRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemHovercardRequestBuilderGetRequestConfiguration)(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Hovercardable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateBasicErrorFromDiscriminatorValue,
        "422": i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateHovercardFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.Hovercardable), nil
}
// ToGetRequestInformation provides hovercard information when authenticated through basic auth or OAuth with the `repo` scope. You can find out more about someone in relation to their pull requests, issues, repositories, and organizations.The `subject_type` and `subject_id` parameters provide context for the person's hovercard, which returns more information than without the parameters. For example, if you wanted to find out more about `octocat` who owns the `Spoon-Knife` repository via cURL, it would look like this:```shell curl -u username:token  https://api.github.com/users/octocat/hovercard?subject_type=repository&subject_id=1300192```
func (m *ItemHovercardRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemHovercardRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemHovercardRequestBuilder) WithUrl(rawUrl string)(*ItemHovercardRequestBuilder) {
    return NewItemHovercardRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
