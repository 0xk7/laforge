// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/gen0cide/laforge/ent"
)

type AgentStatusBatch struct {
	AgentStatuses []*ent.AgentStatus `json:"agentStatuses"`
	PageInfo      *LaForgePageInfo   `json:"pageInfo"`
}

type LaForgePageInfo struct {
	Total      int `json:"total"`
	NextOffset int `json:"nextOffset"`
}

type StatusBatch struct {
	Statuses []*ent.Status    `json:"statuses"`
	PageInfo *LaForgePageInfo `json:"pageInfo"`
}

type ConfigMap struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type TagMap struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type VarsMap struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type AgentCommand string

const (
	AgentCommandDefault        AgentCommand = "DEFAULT"
	AgentCommandDelete         AgentCommand = "DELETE"
	AgentCommandReboot         AgentCommand = "REBOOT"
	AgentCommandExtract        AgentCommand = "EXTRACT"
	AgentCommandDownload       AgentCommand = "DOWNLOAD"
	AgentCommandCreateuser     AgentCommand = "CREATEUSER"
	AgentCommandCreateuserpass AgentCommand = "CREATEUSERPASS"
	AgentCommandAddtogroup     AgentCommand = "ADDTOGROUP"
	AgentCommandExecute        AgentCommand = "EXECUTE"
	AgentCommandValidate       AgentCommand = "VALIDATE"
	AgentCommandChangeperms    AgentCommand = "CHANGEPERMS"
	AgentCommandAppendfile     AgentCommand = "APPENDFILE"
)

var AllAgentCommand = []AgentCommand{
	AgentCommandDefault,
	AgentCommandDelete,
	AgentCommandReboot,
	AgentCommandExtract,
	AgentCommandDownload,
	AgentCommandCreateuser,
	AgentCommandCreateuserpass,
	AgentCommandAddtogroup,
	AgentCommandExecute,
	AgentCommandValidate,
	AgentCommandChangeperms,
	AgentCommandAppendfile,
}

func (e AgentCommand) IsValid() bool {
	switch e {
	case AgentCommandDefault, AgentCommandDelete, AgentCommandReboot, AgentCommandExtract, AgentCommandDownload, AgentCommandCreateuser, AgentCommandCreateuserpass, AgentCommandAddtogroup, AgentCommandExecute, AgentCommandValidate, AgentCommandChangeperms, AgentCommandAppendfile:
		return true
	}
	return false
}

func (e AgentCommand) String() string {
	return string(e)
}

func (e *AgentCommand) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AgentCommand(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AgentCommand", str)
	}
	return nil
}

func (e AgentCommand) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type AgentTaskState string

const (
	AgentTaskStateAwaiting   AgentTaskState = "AWAITING"
	AgentTaskStateInprogress AgentTaskState = "INPROGRESS"
	AgentTaskStateFailed     AgentTaskState = "FAILED"
	AgentTaskStateComplete   AgentTaskState = "COMPLETE"
)

var AllAgentTaskState = []AgentTaskState{
	AgentTaskStateAwaiting,
	AgentTaskStateInprogress,
	AgentTaskStateFailed,
	AgentTaskStateComplete,
}

func (e AgentTaskState) IsValid() bool {
	switch e {
	case AgentTaskStateAwaiting, AgentTaskStateInprogress, AgentTaskStateFailed, AgentTaskStateComplete:
		return true
	}
	return false
}

func (e AgentTaskState) String() string {
	return string(e)
}

func (e *AgentTaskState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AgentTaskState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AgentTaskState", str)
	}
	return nil
}

func (e AgentTaskState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type BuildCommitState string

const (
	BuildCommitStatePlanning   BuildCommitState = "PLANNING"
	BuildCommitStateInprogress BuildCommitState = "INPROGRESS"
	BuildCommitStateApplied    BuildCommitState = "APPLIED"
	BuildCommitStateCancelled  BuildCommitState = "CANCELLED"
	BuildCommitStateApproved   BuildCommitState = "APPROVED"
)

var AllBuildCommitState = []BuildCommitState{
	BuildCommitStatePlanning,
	BuildCommitStateInprogress,
	BuildCommitStateApplied,
	BuildCommitStateCancelled,
	BuildCommitStateApproved,
}

func (e BuildCommitState) IsValid() bool {
	switch e {
	case BuildCommitStatePlanning, BuildCommitStateInprogress, BuildCommitStateApplied, BuildCommitStateCancelled, BuildCommitStateApproved:
		return true
	}
	return false
}

func (e BuildCommitState) String() string {
	return string(e)
}

func (e *BuildCommitState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BuildCommitState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid BuildCommitState", str)
	}
	return nil
}

func (e BuildCommitState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type BuildCommitType string

const (
	BuildCommitTypeRoot    BuildCommitType = "ROOT"
	BuildCommitTypeRebuild BuildCommitType = "REBUILD"
	BuildCommitTypeDelete  BuildCommitType = "DELETE"
)

var AllBuildCommitType = []BuildCommitType{
	BuildCommitTypeRoot,
	BuildCommitTypeRebuild,
	BuildCommitTypeDelete,
}

func (e BuildCommitType) IsValid() bool {
	switch e {
	case BuildCommitTypeRoot, BuildCommitTypeRebuild, BuildCommitTypeDelete:
		return true
	}
	return false
}

func (e BuildCommitType) String() string {
	return string(e)
}

func (e *BuildCommitType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BuildCommitType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid BuildCommitType", str)
	}
	return nil
}

func (e BuildCommitType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FindingDifficulty string

const (
	FindingDifficultyZeroDifficulty     FindingDifficulty = "ZeroDifficulty"
	FindingDifficultyNoviceDifficulty   FindingDifficulty = "NoviceDifficulty"
	FindingDifficultyAdvancedDifficulty FindingDifficulty = "AdvancedDifficulty"
	FindingDifficultyExpertDifficulty   FindingDifficulty = "ExpertDifficulty"
	FindingDifficultyNullDifficulty     FindingDifficulty = "NullDifficulty"
)

var AllFindingDifficulty = []FindingDifficulty{
	FindingDifficultyZeroDifficulty,
	FindingDifficultyNoviceDifficulty,
	FindingDifficultyAdvancedDifficulty,
	FindingDifficultyExpertDifficulty,
	FindingDifficultyNullDifficulty,
}

func (e FindingDifficulty) IsValid() bool {
	switch e {
	case FindingDifficultyZeroDifficulty, FindingDifficultyNoviceDifficulty, FindingDifficultyAdvancedDifficulty, FindingDifficultyExpertDifficulty, FindingDifficultyNullDifficulty:
		return true
	}
	return false
}

func (e FindingDifficulty) String() string {
	return string(e)
}

func (e *FindingDifficulty) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FindingDifficulty(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FindingDifficulty", str)
	}
	return nil
}

func (e FindingDifficulty) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FindingSeverity string

const (
	FindingSeverityZeroSeverity     FindingSeverity = "ZeroSeverity"
	FindingSeverityLowSeverity      FindingSeverity = "LowSeverity"
	FindingSeverityMediumSeverity   FindingSeverity = "MediumSeverity"
	FindingSeverityHighSeverity     FindingSeverity = "HighSeverity"
	FindingSeverityCriticalSeverity FindingSeverity = "CriticalSeverity"
	FindingSeverityNullSeverity     FindingSeverity = "NullSeverity"
)

var AllFindingSeverity = []FindingSeverity{
	FindingSeverityZeroSeverity,
	FindingSeverityLowSeverity,
	FindingSeverityMediumSeverity,
	FindingSeverityHighSeverity,
	FindingSeverityCriticalSeverity,
	FindingSeverityNullSeverity,
}

func (e FindingSeverity) IsValid() bool {
	switch e {
	case FindingSeverityZeroSeverity, FindingSeverityLowSeverity, FindingSeverityMediumSeverity, FindingSeverityHighSeverity, FindingSeverityCriticalSeverity, FindingSeverityNullSeverity:
		return true
	}
	return false
}

func (e FindingSeverity) String() string {
	return string(e)
}

func (e *FindingSeverity) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FindingSeverity(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FindingSeverity", str)
	}
	return nil
}

func (e FindingSeverity) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PlanType string

const (
	PlanTypeStartBuild       PlanType = "start_build"
	PlanTypeStartTeam        PlanType = "start_team"
	PlanTypeProvisionNetwork PlanType = "provision_network"
	PlanTypeProvisionHost    PlanType = "provision_host"
	PlanTypeExecuteStep      PlanType = "execute_step"
	PlanTypeUndefined        PlanType = "undefined"
)

var AllPlanType = []PlanType{
	PlanTypeStartBuild,
	PlanTypeStartTeam,
	PlanTypeProvisionNetwork,
	PlanTypeProvisionHost,
	PlanTypeExecuteStep,
	PlanTypeUndefined,
}

func (e PlanType) IsValid() bool {
	switch e {
	case PlanTypeStartBuild, PlanTypeStartTeam, PlanTypeProvisionNetwork, PlanTypeProvisionHost, PlanTypeExecuteStep, PlanTypeUndefined:
		return true
	}
	return false
}

func (e PlanType) String() string {
	return string(e)
}

func (e *PlanType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PlanType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PlanType", str)
	}
	return nil
}

func (e PlanType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProviderType string

const (
	ProviderTypeLocal     ProviderType = "LOCAL"
	ProviderTypeGithub    ProviderType = "GITHUB"
	ProviderTypeOpenid    ProviderType = "OPENID"
	ProviderTypeUndefined ProviderType = "UNDEFINED"
)

var AllProviderType = []ProviderType{
	ProviderTypeLocal,
	ProviderTypeGithub,
	ProviderTypeOpenid,
	ProviderTypeUndefined,
}

func (e ProviderType) IsValid() bool {
	switch e {
	case ProviderTypeLocal, ProviderTypeGithub, ProviderTypeOpenid, ProviderTypeUndefined:
		return true
	}
	return false
}

func (e ProviderType) String() string {
	return string(e)
}

func (e *ProviderType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProviderType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProviderType", str)
	}
	return nil
}

func (e ProviderType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProvisionStatus string

const (
	ProvisionStatusPlanning         ProvisionStatus = "PLANNING"
	ProvisionStatusAwaiting         ProvisionStatus = "AWAITING"
	ProvisionStatusParentawaiting   ProvisionStatus = "PARENTAWAITING"
	ProvisionStatusInprogress       ProvisionStatus = "INPROGRESS"
	ProvisionStatusFailed           ProvisionStatus = "FAILED"
	ProvisionStatusComplete         ProvisionStatus = "COMPLETE"
	ProvisionStatusTainted          ProvisionStatus = "TAINTED"
	ProvisionStatusUndefined        ProvisionStatus = "UNDEFINED"
	ProvisionStatusTodelete         ProvisionStatus = "TODELETE"
	ProvisionStatusDeleteinprogress ProvisionStatus = "DELETEINPROGRESS"
	ProvisionStatusDeleted          ProvisionStatus = "DELETED"
	ProvisionStatusTorebuild        ProvisionStatus = "TOREBUILD"
	ProvisionStatusCancelled        ProvisionStatus = "CANCELLED"
)

var AllProvisionStatus = []ProvisionStatus{
	ProvisionStatusPlanning,
	ProvisionStatusAwaiting,
	ProvisionStatusParentawaiting,
	ProvisionStatusInprogress,
	ProvisionStatusFailed,
	ProvisionStatusComplete,
	ProvisionStatusTainted,
	ProvisionStatusUndefined,
	ProvisionStatusTodelete,
	ProvisionStatusDeleteinprogress,
	ProvisionStatusDeleted,
	ProvisionStatusTorebuild,
	ProvisionStatusCancelled,
}

func (e ProvisionStatus) IsValid() bool {
	switch e {
	case ProvisionStatusPlanning, ProvisionStatusAwaiting, ProvisionStatusParentawaiting, ProvisionStatusInprogress, ProvisionStatusFailed, ProvisionStatusComplete, ProvisionStatusTainted, ProvisionStatusUndefined, ProvisionStatusTodelete, ProvisionStatusDeleteinprogress, ProvisionStatusDeleted, ProvisionStatusTorebuild, ProvisionStatusCancelled:
		return true
	}
	return false
}

func (e ProvisionStatus) String() string {
	return string(e)
}

func (e *ProvisionStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProvisionStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProvisionStatus", str)
	}
	return nil
}

func (e ProvisionStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProvisionStatusFor string

const (
	ProvisionStatusForBuild              ProvisionStatusFor = "Build"
	ProvisionStatusForTeam               ProvisionStatusFor = "Team"
	ProvisionStatusForPlan               ProvisionStatusFor = "Plan"
	ProvisionStatusForProvisionedNetwork ProvisionStatusFor = "ProvisionedNetwork"
	ProvisionStatusForProvisionedHost    ProvisionStatusFor = "ProvisionedHost"
	ProvisionStatusForProvisioningStep   ProvisionStatusFor = "ProvisioningStep"
	ProvisionStatusForUndefined          ProvisionStatusFor = "Undefined"
)

var AllProvisionStatusFor = []ProvisionStatusFor{
	ProvisionStatusForBuild,
	ProvisionStatusForTeam,
	ProvisionStatusForPlan,
	ProvisionStatusForProvisionedNetwork,
	ProvisionStatusForProvisionedHost,
	ProvisionStatusForProvisioningStep,
	ProvisionStatusForUndefined,
}

func (e ProvisionStatusFor) IsValid() bool {
	switch e {
	case ProvisionStatusForBuild, ProvisionStatusForTeam, ProvisionStatusForPlan, ProvisionStatusForProvisionedNetwork, ProvisionStatusForProvisionedHost, ProvisionStatusForProvisioningStep, ProvisionStatusForUndefined:
		return true
	}
	return false
}

func (e ProvisionStatusFor) String() string {
	return string(e)
}

func (e *ProvisionStatusFor) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProvisionStatusFor(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProvisionStatusFor", str)
	}
	return nil
}

func (e ProvisionStatusFor) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProvisioningStepType string

const (
	ProvisioningStepTypeAnsible      ProvisioningStepType = "Ansible"
	ProvisioningStepTypeScript       ProvisioningStepType = "Script"
	ProvisioningStepTypeCommand      ProvisioningStepType = "Command"
	ProvisioningStepTypeDNSRecord    ProvisioningStepType = "DNSRecord"
	ProvisioningStepTypeFileDelete   ProvisioningStepType = "FileDelete"
	ProvisioningStepTypeFileDownload ProvisioningStepType = "FileDownload"
	ProvisioningStepTypeFileExtract  ProvisioningStepType = "FileExtract"
	ProvisioningStepTypeUndefined    ProvisioningStepType = "Undefined"
)

var AllProvisioningStepType = []ProvisioningStepType{
	ProvisioningStepTypeAnsible,
	ProvisioningStepTypeScript,
	ProvisioningStepTypeCommand,
	ProvisioningStepTypeDNSRecord,
	ProvisioningStepTypeFileDelete,
	ProvisioningStepTypeFileDownload,
	ProvisioningStepTypeFileExtract,
	ProvisioningStepTypeUndefined,
}

func (e ProvisioningStepType) IsValid() bool {
	switch e {
	case ProvisioningStepTypeAnsible, ProvisioningStepTypeScript, ProvisioningStepTypeCommand, ProvisioningStepTypeDNSRecord, ProvisioningStepTypeFileDelete, ProvisioningStepTypeFileDownload, ProvisioningStepTypeFileExtract, ProvisioningStepTypeUndefined:
		return true
	}
	return false
}

func (e ProvisioningStepType) String() string {
	return string(e)
}

func (e *ProvisioningStepType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProvisioningStepType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProvisioningStepType", str)
	}
	return nil
}

func (e ProvisioningStepType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RoleLevel string

const (
	RoleLevelAdmin     RoleLevel = "ADMIN"
	RoleLevelUser      RoleLevel = "USER"
	RoleLevelUndefined RoleLevel = "UNDEFINED"
)

var AllRoleLevel = []RoleLevel{
	RoleLevelAdmin,
	RoleLevelUser,
	RoleLevelUndefined,
}

func (e RoleLevel) IsValid() bool {
	switch e {
	case RoleLevelAdmin, RoleLevelUser, RoleLevelUndefined:
		return true
	}
	return false
}

func (e RoleLevel) String() string {
	return string(e)
}

func (e *RoleLevel) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RoleLevel(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RoleLevel", str)
	}
	return nil
}

func (e RoleLevel) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ServerTaskType string

const (
	ServerTaskTypeLoadenv      ServerTaskType = "LOADENV"
	ServerTaskTypeCreatebuild  ServerTaskType = "CREATEBUILD"
	ServerTaskTypeRenderfiles  ServerTaskType = "RENDERFILES"
	ServerTaskTypeDeletebuild  ServerTaskType = "DELETEBUILD"
	ServerTaskTypeRebuild      ServerTaskType = "REBUILD"
	ServerTaskTypeExecutebuild ServerTaskType = "EXECUTEBUILD"
)

var AllServerTaskType = []ServerTaskType{
	ServerTaskTypeLoadenv,
	ServerTaskTypeCreatebuild,
	ServerTaskTypeRenderfiles,
	ServerTaskTypeDeletebuild,
	ServerTaskTypeRebuild,
	ServerTaskTypeExecutebuild,
}

func (e ServerTaskType) IsValid() bool {
	switch e {
	case ServerTaskTypeLoadenv, ServerTaskTypeCreatebuild, ServerTaskTypeRenderfiles, ServerTaskTypeDeletebuild, ServerTaskTypeRebuild, ServerTaskTypeExecutebuild:
		return true
	}
	return false
}

func (e ServerTaskType) String() string {
	return string(e)
}

func (e *ServerTaskType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ServerTaskType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ServerTaskType", str)
	}
	return nil
}

func (e ServerTaskType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
