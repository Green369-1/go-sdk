package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2 "github.com/octokit/go-sdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemInvitationsRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\invitations
type ItemItemInvitationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemInvitationsRequestBuilderGetQueryParameters when authenticating as a user with admin rights to a repository, this endpoint will list all currently open repository invitations.
type ItemItemInvitationsRequestBuilderGetQueryParameters struct {
    // Page number of the results to fetch.
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100).
    Per_page *int32 `uriparametername:"per_page"`
}
// ItemItemInvitationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemInvitationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemInvitationsRequestBuilderGetQueryParameters
}
// ByInvitation_id gets an item from the go-sdk.repos.item.item.invitations.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemItemInvitationsRequestBuilder) ByInvitation_id(invitation_id string)(*ItemItemInvitationsWithInvitation_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if invitation_id != "" {
        urlTplParams["invitation_id"] = invitation_id
    }
    return NewItemItemInvitationsWithInvitation_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByInvitation_idInteger gets an item from the go-sdk.repos.item.item.invitations.item collection
func (m *ItemItemInvitationsRequestBuilder) ByInvitation_idInteger(invitation_id int32)(*ItemItemInvitationsWithInvitation_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["invitation_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(invitation_id), 10)
    return NewItemItemInvitationsWithInvitation_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemInvitationsRequestBuilderInternal instantiates a new InvitationsRequestBuilder and sets the default values.
func NewItemItemInvitationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemInvitationsRequestBuilder) {
    m := &ItemItemInvitationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/invitations{?per_page*,page*}", pathParameters),
    }
    return m
}
// NewItemItemInvitationsRequestBuilder instantiates a new InvitationsRequestBuilder and sets the default values.
func NewItemItemInvitationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemInvitationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemInvitationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get when authenticating as a user with admin rights to a repository, this endpoint will list all currently open repository invitations.
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/collaborators/invitations#list-repository-invitations
func (m *ItemItemInvitationsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemInvitationsRequestBuilderGetRequestConfiguration)([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryInvitationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.CreateRepositoryInvitationFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryInvitationable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i33bcfa37677dfd8b71eaf22ad3bfd140ba1702e7a8418558499c22418bf2fef2.RepositoryInvitationable)
        }
    }
    return val, nil
}
// ToGetRequestInformation when authenticating as a user with admin rights to a repository, this endpoint will list all currently open repository invitations.
func (m *ItemItemInvitationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemInvitationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemItemInvitationsRequestBuilder) WithUrl(rawUrl string)(*ItemItemInvitationsRequestBuilder) {
    return NewItemItemInvitationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
