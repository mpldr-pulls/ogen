// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/encoding/json"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
)

func encodeEnterpriseAdminSetGithubActionsPermissionsEnterpriseRequest(req EnterpriseAdminSetGithubActionsPermissionsEnterpriseApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeEnterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterpriseRequest(req EnterpriseAdminSetSelectedOrganizationsEnabledGithubActionsEnterpriseApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeEnterpriseAdminSetAllowedActionsEnterpriseRequest(req SelectedActions) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeEnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseRequest(req EnterpriseAdminCreateSelfHostedRunnerGroupForEnterpriseApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeEnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseRequest(req *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeEnterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterpriseRequest(req EnterpriseAdminSetOrgAccessToSelfHostedRunnerGroupInEnterpriseApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeEnterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseRequest(req EnterpriseAdminSetSelfHostedRunnersInGroupForEnterpriseApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeGistsCreateCommentRequest(req GistsCreateCommentApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeGistsUpdateCommentRequest(req GistsUpdateCommentApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeMarkdownRenderRequest(req MarkdownRenderApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeMarkdownRenderRawRequest(req MarkdownRenderRawRequest) ([]byte, string, error) {
	switch req := req.(type) {
	case *MarkdownRenderRawTextPlainRequest:
		return nil, "", fmt.Errorf("text/plain encoder not implemented")
	case *MarkdownRenderRawTextXMarkdownRequest:
		return nil, "", fmt.Errorf("text/x-markdown encoder not implemented")
	default:
		return nil, "", fmt.Errorf("unexpected request type: %T", req)
	}
}

func encodeActivityMarkNotificationsAsReadRequest(req *ActivityMarkNotificationsAsReadApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeActivitySetThreadSubscriptionRequest(req *ActivitySetThreadSubscriptionApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeActionsSetGithubActionsPermissionsOrganizationRequest(req ActionsSetGithubActionsPermissionsOrganizationApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeActionsSetSelectedRepositoriesEnabledGithubActionsOrganizationRequest(req ActionsSetSelectedRepositoriesEnabledGithubActionsOrganizationApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeActionsSetAllowedActionsOrganizationRequest(req *SelectedActions) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeActionsCreateSelfHostedRunnerGroupForOrgRequest(req ActionsCreateSelfHostedRunnerGroupForOrgApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeActionsUpdateSelfHostedRunnerGroupForOrgRequest(req ActionsUpdateSelfHostedRunnerGroupForOrgApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeActionsSetRepoAccessToSelfHostedRunnerGroupInOrgRequest(req ActionsSetRepoAccessToSelfHostedRunnerGroupInOrgApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeActionsSetSelfHostedRunnersInGroupForOrgRequest(req ActionsSetSelfHostedRunnersInGroupForOrgApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeActionsCreateOrUpdateOrgSecretRequest(req ActionsCreateOrUpdateOrgSecretApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeActionsSetSelectedReposForOrgSecretRequest(req ActionsSetSelectedReposForOrgSecretApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeProjectsCreateForOrgRequest(req ProjectsCreateForOrgApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeTeamsUpdateInOrgRequest(req *TeamsUpdateInOrgApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeTeamsCreateDiscussionInOrgRequest(req TeamsCreateDiscussionInOrgApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeTeamsUpdateDiscussionInOrgRequest(req *TeamsUpdateDiscussionInOrgApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeTeamsCreateDiscussionCommentInOrgRequest(req TeamsCreateDiscussionCommentInOrgApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeTeamsUpdateDiscussionCommentInOrgRequest(req TeamsUpdateDiscussionCommentInOrgApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeReactionsCreateForTeamDiscussionCommentInOrgRequest(req ReactionsCreateForTeamDiscussionCommentInOrgApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeReactionsCreateForTeamDiscussionInOrgRequest(req ReactionsCreateForTeamDiscussionInOrgApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeTeamsAddOrUpdateMembershipForUserInOrgRequest(req *TeamsAddOrUpdateMembershipForUserInOrgApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeTeamsAddOrUpdateProjectPermissionsInOrgRequest(req *TeamsAddOrUpdateProjectPermissionsInOrgApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeTeamsAddOrUpdateRepoPermissionsInOrgRequest(req *TeamsAddOrUpdateRepoPermissionsInOrgApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeTeamsCreateOrUpdateIdpGroupConnectionsInOrgRequest(req TeamsCreateOrUpdateIdpGroupConnectionsInOrgApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeProjectsUpdateCardRequest(req *ProjectsUpdateCardApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeProjectsUpdateColumnRequest(req ProjectsUpdateColumnApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeProjectsMoveColumnRequest(req ProjectsMoveColumnApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeProjectsUpdateRequest(req *ProjectsUpdateApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeProjectsCreateColumnRequest(req ProjectsCreateColumnApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeActionsSetGithubActionsPermissionsRepositoryRequest(req ActionsSetGithubActionsPermissionsRepositoryApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeActionsSetAllowedActionsRepositoryRequest(req *SelectedActions) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeActionsCreateOrUpdateRepoSecretRequest(req ActionsCreateOrUpdateRepoSecretApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeReposUpdateBranchProtectionRequest(req ReposUpdateBranchProtectionApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeChecksCreateSuiteRequest(req ChecksCreateSuiteApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeChecksSetSuitesPreferencesRequest(req ChecksSetSuitesPreferencesApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeCodeScanningUpdateAlertRequest(req CodeScanningUpdateAlertApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeCodeScanningUploadSarifRequest(req CodeScanningUploadSarifApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeReposUpdateCommitCommentRequest(req ReposUpdateCommitCommentApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeMigrationsUpdateImportRequest(req *MigrationsUpdateImportApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeInteractionsSetRestrictionsForRepoRequest(req InteractionLimit) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeReposUpdateInvitationRequest(req *ReposUpdateInvitationApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeIssuesAddAssigneesRequest(req *IssuesAddAssigneesApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeIssuesRemoveAssigneesRequest(req *IssuesRemoveAssigneesApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeIssuesUpdateLabelRequest(req *IssuesUpdateLabelApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeReposMergeUpstreamRequest(req ReposMergeUpstreamApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeIssuesUpdateMilestoneRequest(req *IssuesUpdateMilestoneApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeActivityMarkRepoNotificationsAsReadRequest(req *ActivityMarkRepoNotificationsAsReadApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeProjectsCreateForRepoRequest(req ProjectsCreateForRepoApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodePullsUpdateReviewCommentRequest(req PullsUpdateReviewCommentApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodePullsCreateReplyForReviewCommentRequest(req PullsCreateReplyForReviewCommentApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodePullsCreateReviewRequest(req *PullsCreateReviewApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodePullsUpdateReviewRequest(req PullsUpdateReviewApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodePullsDismissReviewRequest(req PullsDismissReviewApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodePullsSubmitReviewRequest(req PullsSubmitReviewApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeReposUpdateReleaseAssetRequest(req *ReposUpdateReleaseAssetApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeReposUpdateReleaseRequest(req *ReposUpdateReleaseApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeReposUploadReleaseAssetRequest(req *string) ([]byte, string, error) {
	return nil, "", fmt.Errorf("*/* encoder not implemented")
}

func encodeSecretScanningUpdateAlertRequest(req SecretScanningUpdateAlertApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeReposCreateCommitStatusRequest(req ReposCreateCommitStatusApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeActivitySetRepoSubscriptionRequest(req *ActivitySetRepoSubscriptionApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeReposReplaceAllTopicsRequest(req ReposReplaceAllTopicsApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeReposTransferRequest(req ReposTransferApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeReposCreateUsingTemplateRequest(req ReposCreateUsingTemplateApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeActionsCreateOrUpdateEnvironmentSecretRequest(req ActionsCreateOrUpdateEnvironmentSecretApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeEnterpriseAdminProvisionAndInviteEnterpriseGroupRequest(req EnterpriseAdminProvisionAndInviteEnterpriseGroupApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeEnterpriseAdminSetInformationForProvisionedEnterpriseGroupRequest(req EnterpriseAdminSetInformationForProvisionedEnterpriseGroupApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeEnterpriseAdminProvisionAndInviteEnterpriseUserRequest(req EnterpriseAdminProvisionAndInviteEnterpriseUserApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeEnterpriseAdminSetInformationForProvisionedEnterpriseUserRequest(req EnterpriseAdminSetInformationForProvisionedEnterpriseUserApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeEnterpriseAdminUpdateAttributeForEnterpriseUserRequest(req EnterpriseAdminUpdateAttributeForEnterpriseUserApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeTeamsCreateDiscussionLegacyRequest(req TeamsCreateDiscussionLegacyApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeTeamsUpdateDiscussionLegacyRequest(req *TeamsUpdateDiscussionLegacyApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeTeamsCreateDiscussionCommentLegacyRequest(req TeamsCreateDiscussionCommentLegacyApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeTeamsUpdateDiscussionCommentLegacyRequest(req TeamsUpdateDiscussionCommentLegacyApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeReactionsCreateForTeamDiscussionCommentLegacyRequest(req ReactionsCreateForTeamDiscussionCommentLegacyApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeReactionsCreateForTeamDiscussionLegacyRequest(req ReactionsCreateForTeamDiscussionLegacyApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeTeamsAddOrUpdateMembershipForUserLegacyRequest(req *TeamsAddOrUpdateMembershipForUserLegacyApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}

func encodeProjectsCreateForAuthenticatedUserRequest(req ProjectsCreateForAuthenticatedUserApplicationJSONRequest) ([]byte, string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, "", err
	}

	return b, "application/json", nil
}