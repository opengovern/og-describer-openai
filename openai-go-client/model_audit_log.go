/*
OpenAI API

The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.

API version: 2.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AuditLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditLog{}

// AuditLog A log of a user action or configuration change within this organization.
type AuditLog struct {
	// The ID of this log.
	Id string `json:"id"`
	Type AuditLogEventType `json:"type"`
	// The Unix timestamp (in seconds) of the event.
	EffectiveAt int32 `json:"effective_at"`
	Project *AuditLogProject `json:"project,omitempty"`
	Actor AuditLogActor `json:"actor"`
	ApiKeyCreated *AuditLogApiKeyCreated `json:"api_key.created,omitempty"`
	ApiKeyUpdated *AuditLogApiKeyUpdated `json:"api_key.updated,omitempty"`
	ApiKeyDeleted *AuditLogApiKeyDeleted `json:"api_key.deleted,omitempty"`
	InviteSent *AuditLogInviteSent `json:"invite.sent,omitempty"`
	InviteAccepted *AuditLogInviteAccepted `json:"invite.accepted,omitempty"`
	InviteDeleted *AuditLogInviteAccepted `json:"invite.deleted,omitempty"`
	LoginFailed *AuditLogLoginFailed `json:"login.failed,omitempty"`
	LogoutFailed *AuditLogLoginFailed `json:"logout.failed,omitempty"`
	OrganizationUpdated *AuditLogOrganizationUpdated `json:"organization.updated,omitempty"`
	ProjectCreated *AuditLogProjectCreated `json:"project.created,omitempty"`
	ProjectUpdated *AuditLogProjectUpdated `json:"project.updated,omitempty"`
	ProjectArchived *AuditLogProjectArchived `json:"project.archived,omitempty"`
	RateLimitUpdated *AuditLogRateLimitUpdated `json:"rate_limit.updated,omitempty"`
	RateLimitDeleted *AuditLogRateLimitDeleted `json:"rate_limit.deleted,omitempty"`
	ServiceAccountCreated *AuditLogServiceAccountCreated `json:"service_account.created,omitempty"`
	ServiceAccountUpdated *AuditLogServiceAccountUpdated `json:"service_account.updated,omitempty"`
	ServiceAccountDeleted *AuditLogServiceAccountDeleted `json:"service_account.deleted,omitempty"`
	UserAdded *AuditLogUserAdded `json:"user.added,omitempty"`
	UserUpdated *AuditLogUserUpdated `json:"user.updated,omitempty"`
	UserDeleted *AuditLogUserDeleted `json:"user.deleted,omitempty"`
}

type _AuditLog AuditLog

// NewAuditLog instantiates a new AuditLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLog(id string, type_ AuditLogEventType, effectiveAt int32, actor AuditLogActor) *AuditLog {
	this := AuditLog{}
	this.Id = id
	this.Type = type_
	this.EffectiveAt = effectiveAt
	this.Actor = actor
	return &this
}

// NewAuditLogWithDefaults instantiates a new AuditLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditLogWithDefaults() *AuditLog {
	this := AuditLog{}
	return &this
}

// GetId returns the Id field value
func (o *AuditLog) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AuditLog) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AuditLog) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *AuditLog) GetType() AuditLogEventType {
	if o == nil {
		var ret AuditLogEventType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AuditLog) GetTypeOk() (*AuditLogEventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AuditLog) SetType(v AuditLogEventType) {
	o.Type = v
}

// GetEffectiveAt returns the EffectiveAt field value
func (o *AuditLog) GetEffectiveAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EffectiveAt
}

// GetEffectiveAtOk returns a tuple with the EffectiveAt field value
// and a boolean to check if the value has been set.
func (o *AuditLog) GetEffectiveAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectiveAt, true
}

// SetEffectiveAt sets field value
func (o *AuditLog) SetEffectiveAt(v int32) {
	o.EffectiveAt = v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *AuditLog) GetProject() AuditLogProject {
	if o == nil || IsNil(o.Project) {
		var ret AuditLogProject
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLog) GetProjectOk() (*AuditLogProject, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *AuditLog) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given AuditLogProject and assigns it to the Project field.
func (o *AuditLog) SetProject(v AuditLogProject) {
	o.Project = &v
}

// GetActor returns the Actor field value
func (o *AuditLog) GetActor() AuditLogActor {
	if o == nil {
		var ret AuditLogActor
		return ret
	}

	return o.Actor
}

// GetActorOk returns a tuple with the Actor field value
// and a boolean to check if the value has been set.
func (o *AuditLog) GetActorOk() (*AuditLogActor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actor, true
}

// SetActor sets field value
func (o *AuditLog) SetActor(v AuditLogActor) {
	o.Actor = v
}

// GetApiKeyCreated returns the ApiKeyCreated field value if set, zero value otherwise.
func (o *AuditLog) GetApiKeyCreated() AuditLogApiKeyCreated {
	if o == nil || IsNil(o.ApiKeyCreated) {
		var ret AuditLogApiKeyCreated
		return ret
	}
	return *o.ApiKeyCreated
}

// GetApiKeyCreatedOk returns a tuple with the ApiKeyCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLog) GetApiKeyCreatedOk() (*AuditLogApiKeyCreated, bool) {
	if o == nil || IsNil(o.ApiKeyCreated) {
		return nil, false
	}
	return o.ApiKeyCreated, true
}

// HasApiKeyCreated returns a boolean if a field has been set.
func (o *AuditLog) HasApiKeyCreated() bool {
	if o != nil && !IsNil(o.ApiKeyCreated) {
		return true
	}

	return false
}

// SetApiKeyCreated gets a reference to the given AuditLogApiKeyCreated and assigns it to the ApiKeyCreated field.
func (o *AuditLog) SetApiKeyCreated(v AuditLogApiKeyCreated) {
	o.ApiKeyCreated = &v
}

// GetApiKeyUpdated returns the ApiKeyUpdated field value if set, zero value otherwise.
func (o *AuditLog) GetApiKeyUpdated() AuditLogApiKeyUpdated {
	if o == nil || IsNil(o.ApiKeyUpdated) {
		var ret AuditLogApiKeyUpdated
		return ret
	}
	return *o.ApiKeyUpdated
}

// GetApiKeyUpdatedOk returns a tuple with the ApiKeyUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLog) GetApiKeyUpdatedOk() (*AuditLogApiKeyUpdated, bool) {
	if o == nil || IsNil(o.ApiKeyUpdated) {
		return nil, false
	}
	return o.ApiKeyUpdated, true
}

// HasApiKeyUpdated returns a boolean if a field has been set.
func (o *AuditLog) HasApiKeyUpdated() bool {
	if o != nil && !IsNil(o.ApiKeyUpdated) {
		return true
	}

	return false
}

// SetApiKeyUpdated gets a reference to the given AuditLogApiKeyUpdated and assigns it to the ApiKeyUpdated field.
func (o *AuditLog) SetApiKeyUpdated(v AuditLogApiKeyUpdated) {
	o.ApiKeyUpdated = &v
}

// GetApiKeyDeleted returns the ApiKeyDeleted field value if set, zero value otherwise.
func (o *AuditLog) GetApiKeyDeleted() AuditLogApiKeyDeleted {
	if o == nil || IsNil(o.ApiKeyDeleted) {
		var ret AuditLogApiKeyDeleted
		return ret
	}
	return *o.ApiKeyDeleted
}

// GetApiKeyDeletedOk returns a tuple with the ApiKeyDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLog) GetApiKeyDeletedOk() (*AuditLogApiKeyDeleted, bool) {
	if o == nil || IsNil(o.ApiKeyDeleted) {
		return nil, false
	}
	return o.ApiKeyDeleted, true
}

// HasApiKeyDeleted returns a boolean if a field has been set.
func (o *AuditLog) HasApiKeyDeleted() bool {
	if o != nil && !IsNil(o.ApiKeyDeleted) {
		return true
	}

	return false
}

// SetApiKeyDeleted gets a reference to the given AuditLogApiKeyDeleted and assigns it to the ApiKeyDeleted field.
func (o *AuditLog) SetApiKeyDeleted(v AuditLogApiKeyDeleted) {
	o.ApiKeyDeleted = &v
}

// GetInviteSent returns the InviteSent field value if set, zero value otherwise.
func (o *AuditLog) GetInviteSent() AuditLogInviteSent {
	if o == nil || IsNil(o.InviteSent) {
		var ret AuditLogInviteSent
		return ret
	}
	return *o.InviteSent
}

// GetInviteSentOk returns a tuple with the InviteSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLog) GetInviteSentOk() (*AuditLogInviteSent, bool) {
	if o == nil || IsNil(o.InviteSent) {
		return nil, false
	}
	return o.InviteSent, true
}

// HasInviteSent returns a boolean if a field has been set.
func (o *AuditLog) HasInviteSent() bool {
	if o != nil && !IsNil(o.InviteSent) {
		return true
	}

	return false
}

// SetInviteSent gets a reference to the given AuditLogInviteSent and assigns it to the InviteSent field.
func (o *AuditLog) SetInviteSent(v AuditLogInviteSent) {
	o.InviteSent = &v
}

// GetInviteAccepted returns the InviteAccepted field value if set, zero value otherwise.
func (o *AuditLog) GetInviteAccepted() AuditLogInviteAccepted {
	if o == nil || IsNil(o.InviteAccepted) {
		var ret AuditLogInviteAccepted
		return ret
	}
	return *o.InviteAccepted
}

// GetInviteAcceptedOk returns a tuple with the InviteAccepted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLog) GetInviteAcceptedOk() (*AuditLogInviteAccepted, bool) {
	if o == nil || IsNil(o.InviteAccepted) {
		return nil, false
	}
	return o.InviteAccepted, true
}

// HasInviteAccepted returns a boolean if a field has been set.
func (o *AuditLog) HasInviteAccepted() bool {
	if o != nil && !IsNil(o.InviteAccepted) {
		return true
	}

	return false
}

// SetInviteAccepted gets a reference to the given AuditLogInviteAccepted and assigns it to the InviteAccepted field.
func (o *AuditLog) SetInviteAccepted(v AuditLogInviteAccepted) {
	o.InviteAccepted = &v
}

// GetInviteDeleted returns the InviteDeleted field value if set, zero value otherwise.
func (o *AuditLog) GetInviteDeleted() AuditLogInviteAccepted {
	if o == nil || IsNil(o.InviteDeleted) {
		var ret AuditLogInviteAccepted
		return ret
	}
	return *o.InviteDeleted
}

// GetInviteDeletedOk returns a tuple with the InviteDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLog) GetInviteDeletedOk() (*AuditLogInviteAccepted, bool) {
	if o == nil || IsNil(o.InviteDeleted) {
		return nil, false
	}
	return o.InviteDeleted, true
}

// HasInviteDeleted returns a boolean if a field has been set.
func (o *AuditLog) HasInviteDeleted() bool {
	if o != nil && !IsNil(o.InviteDeleted) {
		return true
	}

	return false
}

// SetInviteDeleted gets a reference to the given AuditLogInviteAccepted and assigns it to the InviteDeleted field.
func (o *AuditLog) SetInviteDeleted(v AuditLogInviteAccepted) {
	o.InviteDeleted = &v
}

// GetLoginFailed returns the LoginFailed field value if set, zero value otherwise.
func (o *AuditLog) GetLoginFailed() AuditLogLoginFailed {
	if o == nil || IsNil(o.LoginFailed) {
		var ret AuditLogLoginFailed
		return ret
	}
	return *o.LoginFailed
}

// GetLoginFailedOk returns a tuple with the LoginFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLog) GetLoginFailedOk() (*AuditLogLoginFailed, bool) {
	if o == nil || IsNil(o.LoginFailed) {
		return nil, false
	}
	return o.LoginFailed, true
}

// HasLoginFailed returns a boolean if a field has been set.
func (o *AuditLog) HasLoginFailed() bool {
	if o != nil && !IsNil(o.LoginFailed) {
		return true
	}

	return false
}

// SetLoginFailed gets a reference to the given AuditLogLoginFailed and assigns it to the LoginFailed field.
func (o *AuditLog) SetLoginFailed(v AuditLogLoginFailed) {
	o.LoginFailed = &v
}

// GetLogoutFailed returns the LogoutFailed field value if set, zero value otherwise.
func (o *AuditLog) GetLogoutFailed() AuditLogLoginFailed {
	if o == nil || IsNil(o.LogoutFailed) {
		var ret AuditLogLoginFailed
		return ret
	}
	return *o.LogoutFailed
}

// GetLogoutFailedOk returns a tuple with the LogoutFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLog) GetLogoutFailedOk() (*AuditLogLoginFailed, bool) {
	if o == nil || IsNil(o.LogoutFailed) {
		return nil, false
	}
	return o.LogoutFailed, true
}

// HasLogoutFailed returns a boolean if a field has been set.
func (o *AuditLog) HasLogoutFailed() bool {
	if o != nil && !IsNil(o.LogoutFailed) {
		return true
	}

	return false
}

// SetLogoutFailed gets a reference to the given AuditLogLoginFailed and assigns it to the LogoutFailed field.
func (o *AuditLog) SetLogoutFailed(v AuditLogLoginFailed) {
	o.LogoutFailed = &v
}

// GetOrganizationUpdated returns the OrganizationUpdated field value if set, zero value otherwise.
func (o *AuditLog) GetOrganizationUpdated() AuditLogOrganizationUpdated {
	if o == nil || IsNil(o.OrganizationUpdated) {
		var ret AuditLogOrganizationUpdated
		return ret
	}
	return *o.OrganizationUpdated
}

// GetOrganizationUpdatedOk returns a tuple with the OrganizationUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLog) GetOrganizationUpdatedOk() (*AuditLogOrganizationUpdated, bool) {
	if o == nil || IsNil(o.OrganizationUpdated) {
		return nil, false
	}
	return o.OrganizationUpdated, true
}

// HasOrganizationUpdated returns a boolean if a field has been set.
func (o *AuditLog) HasOrganizationUpdated() bool {
	if o != nil && !IsNil(o.OrganizationUpdated) {
		return true
	}

	return false
}

// SetOrganizationUpdated gets a reference to the given AuditLogOrganizationUpdated and assigns it to the OrganizationUpdated field.
func (o *AuditLog) SetOrganizationUpdated(v AuditLogOrganizationUpdated) {
	o.OrganizationUpdated = &v
}

// GetProjectCreated returns the ProjectCreated field value if set, zero value otherwise.
func (o *AuditLog) GetProjectCreated() AuditLogProjectCreated {
	if o == nil || IsNil(o.ProjectCreated) {
		var ret AuditLogProjectCreated
		return ret
	}
	return *o.ProjectCreated
}

// GetProjectCreatedOk returns a tuple with the ProjectCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLog) GetProjectCreatedOk() (*AuditLogProjectCreated, bool) {
	if o == nil || IsNil(o.ProjectCreated) {
		return nil, false
	}
	return o.ProjectCreated, true
}

// HasProjectCreated returns a boolean if a field has been set.
func (o *AuditLog) HasProjectCreated() bool {
	if o != nil && !IsNil(o.ProjectCreated) {
		return true
	}

	return false
}

// SetProjectCreated gets a reference to the given AuditLogProjectCreated and assigns it to the ProjectCreated field.
func (o *AuditLog) SetProjectCreated(v AuditLogProjectCreated) {
	o.ProjectCreated = &v
}

// GetProjectUpdated returns the ProjectUpdated field value if set, zero value otherwise.
func (o *AuditLog) GetProjectUpdated() AuditLogProjectUpdated {
	if o == nil || IsNil(o.ProjectUpdated) {
		var ret AuditLogProjectUpdated
		return ret
	}
	return *o.ProjectUpdated
}

// GetProjectUpdatedOk returns a tuple with the ProjectUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLog) GetProjectUpdatedOk() (*AuditLogProjectUpdated, bool) {
	if o == nil || IsNil(o.ProjectUpdated) {
		return nil, false
	}
	return o.ProjectUpdated, true
}

// HasProjectUpdated returns a boolean if a field has been set.
func (o *AuditLog) HasProjectUpdated() bool {
	if o != nil && !IsNil(o.ProjectUpdated) {
		return true
	}

	return false
}

// SetProjectUpdated gets a reference to the given AuditLogProjectUpdated and assigns it to the ProjectUpdated field.
func (o *AuditLog) SetProjectUpdated(v AuditLogProjectUpdated) {
	o.ProjectUpdated = &v
}

// GetProjectArchived returns the ProjectArchived field value if set, zero value otherwise.
func (o *AuditLog) GetProjectArchived() AuditLogProjectArchived {
	if o == nil || IsNil(o.ProjectArchived) {
		var ret AuditLogProjectArchived
		return ret
	}
	return *o.ProjectArchived
}

// GetProjectArchivedOk returns a tuple with the ProjectArchived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLog) GetProjectArchivedOk() (*AuditLogProjectArchived, bool) {
	if o == nil || IsNil(o.ProjectArchived) {
		return nil, false
	}
	return o.ProjectArchived, true
}

// HasProjectArchived returns a boolean if a field has been set.
func (o *AuditLog) HasProjectArchived() bool {
	if o != nil && !IsNil(o.ProjectArchived) {
		return true
	}

	return false
}

// SetProjectArchived gets a reference to the given AuditLogProjectArchived and assigns it to the ProjectArchived field.
func (o *AuditLog) SetProjectArchived(v AuditLogProjectArchived) {
	o.ProjectArchived = &v
}

// GetRateLimitUpdated returns the RateLimitUpdated field value if set, zero value otherwise.
func (o *AuditLog) GetRateLimitUpdated() AuditLogRateLimitUpdated {
	if o == nil || IsNil(o.RateLimitUpdated) {
		var ret AuditLogRateLimitUpdated
		return ret
	}
	return *o.RateLimitUpdated
}

// GetRateLimitUpdatedOk returns a tuple with the RateLimitUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLog) GetRateLimitUpdatedOk() (*AuditLogRateLimitUpdated, bool) {
	if o == nil || IsNil(o.RateLimitUpdated) {
		return nil, false
	}
	return o.RateLimitUpdated, true
}

// HasRateLimitUpdated returns a boolean if a field has been set.
func (o *AuditLog) HasRateLimitUpdated() bool {
	if o != nil && !IsNil(o.RateLimitUpdated) {
		return true
	}

	return false
}

// SetRateLimitUpdated gets a reference to the given AuditLogRateLimitUpdated and assigns it to the RateLimitUpdated field.
func (o *AuditLog) SetRateLimitUpdated(v AuditLogRateLimitUpdated) {
	o.RateLimitUpdated = &v
}

// GetRateLimitDeleted returns the RateLimitDeleted field value if set, zero value otherwise.
func (o *AuditLog) GetRateLimitDeleted() AuditLogRateLimitDeleted {
	if o == nil || IsNil(o.RateLimitDeleted) {
		var ret AuditLogRateLimitDeleted
		return ret
	}
	return *o.RateLimitDeleted
}

// GetRateLimitDeletedOk returns a tuple with the RateLimitDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLog) GetRateLimitDeletedOk() (*AuditLogRateLimitDeleted, bool) {
	if o == nil || IsNil(o.RateLimitDeleted) {
		return nil, false
	}
	return o.RateLimitDeleted, true
}

// HasRateLimitDeleted returns a boolean if a field has been set.
func (o *AuditLog) HasRateLimitDeleted() bool {
	if o != nil && !IsNil(o.RateLimitDeleted) {
		return true
	}

	return false
}

// SetRateLimitDeleted gets a reference to the given AuditLogRateLimitDeleted and assigns it to the RateLimitDeleted field.
func (o *AuditLog) SetRateLimitDeleted(v AuditLogRateLimitDeleted) {
	o.RateLimitDeleted = &v
}

// GetServiceAccountCreated returns the ServiceAccountCreated field value if set, zero value otherwise.
func (o *AuditLog) GetServiceAccountCreated() AuditLogServiceAccountCreated {
	if o == nil || IsNil(o.ServiceAccountCreated) {
		var ret AuditLogServiceAccountCreated
		return ret
	}
	return *o.ServiceAccountCreated
}

// GetServiceAccountCreatedOk returns a tuple with the ServiceAccountCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLog) GetServiceAccountCreatedOk() (*AuditLogServiceAccountCreated, bool) {
	if o == nil || IsNil(o.ServiceAccountCreated) {
		return nil, false
	}
	return o.ServiceAccountCreated, true
}

// HasServiceAccountCreated returns a boolean if a field has been set.
func (o *AuditLog) HasServiceAccountCreated() bool {
	if o != nil && !IsNil(o.ServiceAccountCreated) {
		return true
	}

	return false
}

// SetServiceAccountCreated gets a reference to the given AuditLogServiceAccountCreated and assigns it to the ServiceAccountCreated field.
func (o *AuditLog) SetServiceAccountCreated(v AuditLogServiceAccountCreated) {
	o.ServiceAccountCreated = &v
}

// GetServiceAccountUpdated returns the ServiceAccountUpdated field value if set, zero value otherwise.
func (o *AuditLog) GetServiceAccountUpdated() AuditLogServiceAccountUpdated {
	if o == nil || IsNil(o.ServiceAccountUpdated) {
		var ret AuditLogServiceAccountUpdated
		return ret
	}
	return *o.ServiceAccountUpdated
}

// GetServiceAccountUpdatedOk returns a tuple with the ServiceAccountUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLog) GetServiceAccountUpdatedOk() (*AuditLogServiceAccountUpdated, bool) {
	if o == nil || IsNil(o.ServiceAccountUpdated) {
		return nil, false
	}
	return o.ServiceAccountUpdated, true
}

// HasServiceAccountUpdated returns a boolean if a field has been set.
func (o *AuditLog) HasServiceAccountUpdated() bool {
	if o != nil && !IsNil(o.ServiceAccountUpdated) {
		return true
	}

	return false
}

// SetServiceAccountUpdated gets a reference to the given AuditLogServiceAccountUpdated and assigns it to the ServiceAccountUpdated field.
func (o *AuditLog) SetServiceAccountUpdated(v AuditLogServiceAccountUpdated) {
	o.ServiceAccountUpdated = &v
}

// GetServiceAccountDeleted returns the ServiceAccountDeleted field value if set, zero value otherwise.
func (o *AuditLog) GetServiceAccountDeleted() AuditLogServiceAccountDeleted {
	if o == nil || IsNil(o.ServiceAccountDeleted) {
		var ret AuditLogServiceAccountDeleted
		return ret
	}
	return *o.ServiceAccountDeleted
}

// GetServiceAccountDeletedOk returns a tuple with the ServiceAccountDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLog) GetServiceAccountDeletedOk() (*AuditLogServiceAccountDeleted, bool) {
	if o == nil || IsNil(o.ServiceAccountDeleted) {
		return nil, false
	}
	return o.ServiceAccountDeleted, true
}

// HasServiceAccountDeleted returns a boolean if a field has been set.
func (o *AuditLog) HasServiceAccountDeleted() bool {
	if o != nil && !IsNil(o.ServiceAccountDeleted) {
		return true
	}

	return false
}

// SetServiceAccountDeleted gets a reference to the given AuditLogServiceAccountDeleted and assigns it to the ServiceAccountDeleted field.
func (o *AuditLog) SetServiceAccountDeleted(v AuditLogServiceAccountDeleted) {
	o.ServiceAccountDeleted = &v
}

// GetUserAdded returns the UserAdded field value if set, zero value otherwise.
func (o *AuditLog) GetUserAdded() AuditLogUserAdded {
	if o == nil || IsNil(o.UserAdded) {
		var ret AuditLogUserAdded
		return ret
	}
	return *o.UserAdded
}

// GetUserAddedOk returns a tuple with the UserAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLog) GetUserAddedOk() (*AuditLogUserAdded, bool) {
	if o == nil || IsNil(o.UserAdded) {
		return nil, false
	}
	return o.UserAdded, true
}

// HasUserAdded returns a boolean if a field has been set.
func (o *AuditLog) HasUserAdded() bool {
	if o != nil && !IsNil(o.UserAdded) {
		return true
	}

	return false
}

// SetUserAdded gets a reference to the given AuditLogUserAdded and assigns it to the UserAdded field.
func (o *AuditLog) SetUserAdded(v AuditLogUserAdded) {
	o.UserAdded = &v
}

// GetUserUpdated returns the UserUpdated field value if set, zero value otherwise.
func (o *AuditLog) GetUserUpdated() AuditLogUserUpdated {
	if o == nil || IsNil(o.UserUpdated) {
		var ret AuditLogUserUpdated
		return ret
	}
	return *o.UserUpdated
}

// GetUserUpdatedOk returns a tuple with the UserUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLog) GetUserUpdatedOk() (*AuditLogUserUpdated, bool) {
	if o == nil || IsNil(o.UserUpdated) {
		return nil, false
	}
	return o.UserUpdated, true
}

// HasUserUpdated returns a boolean if a field has been set.
func (o *AuditLog) HasUserUpdated() bool {
	if o != nil && !IsNil(o.UserUpdated) {
		return true
	}

	return false
}

// SetUserUpdated gets a reference to the given AuditLogUserUpdated and assigns it to the UserUpdated field.
func (o *AuditLog) SetUserUpdated(v AuditLogUserUpdated) {
	o.UserUpdated = &v
}

// GetUserDeleted returns the UserDeleted field value if set, zero value otherwise.
func (o *AuditLog) GetUserDeleted() AuditLogUserDeleted {
	if o == nil || IsNil(o.UserDeleted) {
		var ret AuditLogUserDeleted
		return ret
	}
	return *o.UserDeleted
}

// GetUserDeletedOk returns a tuple with the UserDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLog) GetUserDeletedOk() (*AuditLogUserDeleted, bool) {
	if o == nil || IsNil(o.UserDeleted) {
		return nil, false
	}
	return o.UserDeleted, true
}

// HasUserDeleted returns a boolean if a field has been set.
func (o *AuditLog) HasUserDeleted() bool {
	if o != nil && !IsNil(o.UserDeleted) {
		return true
	}

	return false
}

// SetUserDeleted gets a reference to the given AuditLogUserDeleted and assigns it to the UserDeleted field.
func (o *AuditLog) SetUserDeleted(v AuditLogUserDeleted) {
	o.UserDeleted = &v
}

func (o AuditLog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["effective_at"] = o.EffectiveAt
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	toSerialize["actor"] = o.Actor
	if !IsNil(o.ApiKeyCreated) {
		toSerialize["api_key.created"] = o.ApiKeyCreated
	}
	if !IsNil(o.ApiKeyUpdated) {
		toSerialize["api_key.updated"] = o.ApiKeyUpdated
	}
	if !IsNil(o.ApiKeyDeleted) {
		toSerialize["api_key.deleted"] = o.ApiKeyDeleted
	}
	if !IsNil(o.InviteSent) {
		toSerialize["invite.sent"] = o.InviteSent
	}
	if !IsNil(o.InviteAccepted) {
		toSerialize["invite.accepted"] = o.InviteAccepted
	}
	if !IsNil(o.InviteDeleted) {
		toSerialize["invite.deleted"] = o.InviteDeleted
	}
	if !IsNil(o.LoginFailed) {
		toSerialize["login.failed"] = o.LoginFailed
	}
	if !IsNil(o.LogoutFailed) {
		toSerialize["logout.failed"] = o.LogoutFailed
	}
	if !IsNil(o.OrganizationUpdated) {
		toSerialize["organization.updated"] = o.OrganizationUpdated
	}
	if !IsNil(o.ProjectCreated) {
		toSerialize["project.created"] = o.ProjectCreated
	}
	if !IsNil(o.ProjectUpdated) {
		toSerialize["project.updated"] = o.ProjectUpdated
	}
	if !IsNil(o.ProjectArchived) {
		toSerialize["project.archived"] = o.ProjectArchived
	}
	if !IsNil(o.RateLimitUpdated) {
		toSerialize["rate_limit.updated"] = o.RateLimitUpdated
	}
	if !IsNil(o.RateLimitDeleted) {
		toSerialize["rate_limit.deleted"] = o.RateLimitDeleted
	}
	if !IsNil(o.ServiceAccountCreated) {
		toSerialize["service_account.created"] = o.ServiceAccountCreated
	}
	if !IsNil(o.ServiceAccountUpdated) {
		toSerialize["service_account.updated"] = o.ServiceAccountUpdated
	}
	if !IsNil(o.ServiceAccountDeleted) {
		toSerialize["service_account.deleted"] = o.ServiceAccountDeleted
	}
	if !IsNil(o.UserAdded) {
		toSerialize["user.added"] = o.UserAdded
	}
	if !IsNil(o.UserUpdated) {
		toSerialize["user.updated"] = o.UserUpdated
	}
	if !IsNil(o.UserDeleted) {
		toSerialize["user.deleted"] = o.UserDeleted
	}
	return toSerialize, nil
}

func (o *AuditLog) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"effective_at",
		"actor",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAuditLog := _AuditLog{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuditLog)

	if err != nil {
		return err
	}

	*o = AuditLog(varAuditLog)

	return err
}

type NullableAuditLog struct {
	value *AuditLog
	isSet bool
}

func (v NullableAuditLog) Get() *AuditLog {
	return v.value
}

func (v *NullableAuditLog) Set(val *AuditLog) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLog) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLog(val *AuditLog) *NullableAuditLog {
	return &NullableAuditLog{value: val, isSet: true}
}

func (v NullableAuditLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


