// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/gen0cide/laforge/ent/adhocplan"
	"github.com/gen0cide/laforge/ent/agentstatus"
	"github.com/gen0cide/laforge/ent/agenttask"
	"github.com/gen0cide/laforge/ent/authuser"
	"github.com/gen0cide/laforge/ent/build"
	"github.com/gen0cide/laforge/ent/buildcommit"
	"github.com/gen0cide/laforge/ent/command"
	"github.com/gen0cide/laforge/ent/competition"
	"github.com/gen0cide/laforge/ent/disk"
	"github.com/gen0cide/laforge/ent/dns"
	"github.com/gen0cide/laforge/ent/dnsrecord"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/filedelete"
	"github.com/gen0cide/laforge/ent/filedownload"
	"github.com/gen0cide/laforge/ent/fileextract"
	"github.com/gen0cide/laforge/ent/finding"
	"github.com/gen0cide/laforge/ent/ginfilemiddleware"
	"github.com/gen0cide/laforge/ent/host"
	"github.com/gen0cide/laforge/ent/hostdependency"
	"github.com/gen0cide/laforge/ent/identity"
	"github.com/gen0cide/laforge/ent/includednetwork"
	"github.com/gen0cide/laforge/ent/network"
	"github.com/gen0cide/laforge/ent/plan"
	"github.com/gen0cide/laforge/ent/plandiff"
	"github.com/gen0cide/laforge/ent/provisionedhost"
	"github.com/gen0cide/laforge/ent/provisionednetwork"
	"github.com/gen0cide/laforge/ent/provisioningstep"
	"github.com/gen0cide/laforge/ent/repository"
	"github.com/gen0cide/laforge/ent/schema"
	"github.com/gen0cide/laforge/ent/script"
	"github.com/gen0cide/laforge/ent/servertask"
	"github.com/gen0cide/laforge/ent/status"
	"github.com/gen0cide/laforge/ent/tag"
	"github.com/gen0cide/laforge/ent/team"
	"github.com/gen0cide/laforge/ent/token"
	"github.com/gen0cide/laforge/ent/user"
	"github.com/gen0cide/laforge/ent/validation"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	adhocplanFields := schema.AdhocPlan{}.Fields()
	_ = adhocplanFields
	// adhocplanDescID is the schema descriptor for id field.
	adhocplanDescID := adhocplanFields[0].Descriptor()
	// adhocplan.DefaultID holds the default value on creation for the id field.
	adhocplan.DefaultID = adhocplanDescID.Default.(func() uuid.UUID)
	agentstatusFields := schema.AgentStatus{}.Fields()
	_ = agentstatusFields
	// agentstatusDescID is the schema descriptor for id field.
	agentstatusDescID := agentstatusFields[0].Descriptor()
	// agentstatus.DefaultID holds the default value on creation for the id field.
	agentstatus.DefaultID = agentstatusDescID.Default.(func() uuid.UUID)
	agenttaskFields := schema.AgentTask{}.Fields()
	_ = agenttaskFields
	// agenttaskDescOutput is the schema descriptor for output field.
	agenttaskDescOutput := agenttaskFields[4].Descriptor()
	// agenttask.DefaultOutput holds the default value on creation for the output field.
	agenttask.DefaultOutput = agenttaskDescOutput.Default.(string)
	// agenttaskDescErrorMessage is the schema descriptor for error_message field.
	agenttaskDescErrorMessage := agenttaskFields[6].Descriptor()
	// agenttask.DefaultErrorMessage holds the default value on creation for the error_message field.
	agenttask.DefaultErrorMessage = agenttaskDescErrorMessage.Default.(string)
	// agenttaskDescID is the schema descriptor for id field.
	agenttaskDescID := agenttaskFields[0].Descriptor()
	// agenttask.DefaultID holds the default value on creation for the id field.
	agenttask.DefaultID = agenttaskDescID.Default.(func() uuid.UUID)
	authuserFields := schema.AuthUser{}.Fields()
	_ = authuserFields
	// authuserDescFirstName is the schema descriptor for first_name field.
	authuserDescFirstName := authuserFields[3].Descriptor()
	// authuser.DefaultFirstName holds the default value on creation for the first_name field.
	authuser.DefaultFirstName = authuserDescFirstName.Default.(string)
	// authuserDescLastName is the schema descriptor for last_name field.
	authuserDescLastName := authuserFields[4].Descriptor()
	// authuser.DefaultLastName holds the default value on creation for the last_name field.
	authuser.DefaultLastName = authuserDescLastName.Default.(string)
	// authuserDescEmail is the schema descriptor for email field.
	authuserDescEmail := authuserFields[5].Descriptor()
	// authuser.DefaultEmail holds the default value on creation for the email field.
	authuser.DefaultEmail = authuserDescEmail.Default.(string)
	// authuserDescPhone is the schema descriptor for phone field.
	authuserDescPhone := authuserFields[6].Descriptor()
	// authuser.DefaultPhone holds the default value on creation for the phone field.
	authuser.DefaultPhone = authuserDescPhone.Default.(string)
	// authuserDescCompany is the schema descriptor for company field.
	authuserDescCompany := authuserFields[7].Descriptor()
	// authuser.DefaultCompany holds the default value on creation for the company field.
	authuser.DefaultCompany = authuserDescCompany.Default.(string)
	// authuserDescOccupation is the schema descriptor for occupation field.
	authuserDescOccupation := authuserFields[8].Descriptor()
	// authuser.DefaultOccupation holds the default value on creation for the occupation field.
	authuser.DefaultOccupation = authuserDescOccupation.Default.(string)
	// authuserDescPrivateKeyPath is the schema descriptor for private_key_path field.
	authuserDescPrivateKeyPath := authuserFields[9].Descriptor()
	// authuser.DefaultPrivateKeyPath holds the default value on creation for the private_key_path field.
	authuser.DefaultPrivateKeyPath = authuserDescPrivateKeyPath.Default.(string)
	// authuserDescID is the schema descriptor for id field.
	authuserDescID := authuserFields[0].Descriptor()
	// authuser.DefaultID holds the default value on creation for the id field.
	authuser.DefaultID = authuserDescID.Default.(func() uuid.UUID)
	buildFields := schema.Build{}.Fields()
	_ = buildFields
	// buildDescCompletedPlan is the schema descriptor for completed_plan field.
	buildDescCompletedPlan := buildFields[3].Descriptor()
	// build.DefaultCompletedPlan holds the default value on creation for the completed_plan field.
	build.DefaultCompletedPlan = buildDescCompletedPlan.Default.(bool)
	// buildDescID is the schema descriptor for id field.
	buildDescID := buildFields[0].Descriptor()
	// build.DefaultID holds the default value on creation for the id field.
	build.DefaultID = buildDescID.Default.(func() uuid.UUID)
	buildcommitFields := schema.BuildCommit{}.Fields()
	_ = buildcommitFields
	// buildcommitDescID is the schema descriptor for id field.
	buildcommitDescID := buildcommitFields[0].Descriptor()
	// buildcommit.DefaultID holds the default value on creation for the id field.
	buildcommit.DefaultID = buildcommitDescID.Default.(func() uuid.UUID)
	commandFields := schema.Command{}.Fields()
	_ = commandFields
	// commandDescCooldown is the schema descriptor for cooldown field.
	commandDescCooldown := commandFields[8].Descriptor()
	// command.CooldownValidator is a validator for the "cooldown" field. It is called by the builders before save.
	command.CooldownValidator = commandDescCooldown.Validators[0].(func(int) error)
	// commandDescTimeout is the schema descriptor for timeout field.
	commandDescTimeout := commandFields[9].Descriptor()
	// command.TimeoutValidator is a validator for the "timeout" field. It is called by the builders before save.
	command.TimeoutValidator = commandDescTimeout.Validators[0].(func(int) error)
	// commandDescID is the schema descriptor for id field.
	commandDescID := commandFields[0].Descriptor()
	// command.DefaultID holds the default value on creation for the id field.
	command.DefaultID = commandDescID.Default.(func() uuid.UUID)
	competitionFields := schema.Competition{}.Fields()
	_ = competitionFields
	// competitionDescID is the schema descriptor for id field.
	competitionDescID := competitionFields[0].Descriptor()
	// competition.DefaultID holds the default value on creation for the id field.
	competition.DefaultID = competitionDescID.Default.(func() uuid.UUID)
	dnsFields := schema.DNS{}.Fields()
	_ = dnsFields
	// dnsDescID is the schema descriptor for id field.
	dnsDescID := dnsFields[0].Descriptor()
	// dns.DefaultID holds the default value on creation for the id field.
	dns.DefaultID = dnsDescID.Default.(func() uuid.UUID)
	dnsrecordFields := schema.DNSRecord{}.Fields()
	_ = dnsrecordFields
	// dnsrecordDescID is the schema descriptor for id field.
	dnsrecordDescID := dnsrecordFields[0].Descriptor()
	// dnsrecord.DefaultID holds the default value on creation for the id field.
	dnsrecord.DefaultID = dnsrecordDescID.Default.(func() uuid.UUID)
	diskFields := schema.Disk{}.Fields()
	_ = diskFields
	// diskDescSize is the schema descriptor for size field.
	diskDescSize := diskFields[1].Descriptor()
	// disk.SizeValidator is a validator for the "size" field. It is called by the builders before save.
	disk.SizeValidator = diskDescSize.Validators[0].(func(int) error)
	// diskDescID is the schema descriptor for id field.
	diskDescID := diskFields[0].Descriptor()
	// disk.DefaultID holds the default value on creation for the id field.
	disk.DefaultID = diskDescID.Default.(func() uuid.UUID)
	environmentFields := schema.Environment{}.Fields()
	_ = environmentFields
	// environmentDescID is the schema descriptor for id field.
	environmentDescID := environmentFields[0].Descriptor()
	// environment.DefaultID holds the default value on creation for the id field.
	environment.DefaultID = environmentDescID.Default.(func() uuid.UUID)
	filedeleteFields := schema.FileDelete{}.Fields()
	_ = filedeleteFields
	// filedeleteDescID is the schema descriptor for id field.
	filedeleteDescID := filedeleteFields[0].Descriptor()
	// filedelete.DefaultID holds the default value on creation for the id field.
	filedelete.DefaultID = filedeleteDescID.Default.(func() uuid.UUID)
	filedownloadFields := schema.FileDownload{}.Fields()
	_ = filedownloadFields
	// filedownloadDescID is the schema descriptor for id field.
	filedownloadDescID := filedownloadFields[0].Descriptor()
	// filedownload.DefaultID holds the default value on creation for the id field.
	filedownload.DefaultID = filedownloadDescID.Default.(func() uuid.UUID)
	fileextractFields := schema.FileExtract{}.Fields()
	_ = fileextractFields
	// fileextractDescID is the schema descriptor for id field.
	fileextractDescID := fileextractFields[0].Descriptor()
	// fileextract.DefaultID holds the default value on creation for the id field.
	fileextract.DefaultID = fileextractDescID.Default.(func() uuid.UUID)
	findingFields := schema.Finding{}.Fields()
	_ = findingFields
	// findingDescID is the schema descriptor for id field.
	findingDescID := findingFields[0].Descriptor()
	// finding.DefaultID holds the default value on creation for the id field.
	finding.DefaultID = findingDescID.Default.(func() uuid.UUID)
	ginfilemiddlewareFields := schema.GinFileMiddleware{}.Fields()
	_ = ginfilemiddlewareFields
	// ginfilemiddlewareDescAccessed is the schema descriptor for accessed field.
	ginfilemiddlewareDescAccessed := ginfilemiddlewareFields[3].Descriptor()
	// ginfilemiddleware.DefaultAccessed holds the default value on creation for the accessed field.
	ginfilemiddleware.DefaultAccessed = ginfilemiddlewareDescAccessed.Default.(bool)
	// ginfilemiddlewareDescID is the schema descriptor for id field.
	ginfilemiddlewareDescID := ginfilemiddlewareFields[0].Descriptor()
	// ginfilemiddleware.DefaultID holds the default value on creation for the id field.
	ginfilemiddleware.DefaultID = ginfilemiddlewareDescID.Default.(func() uuid.UUID)
	hostFields := schema.Host{}.Fields()
	_ = hostFields
	// hostDescID is the schema descriptor for id field.
	hostDescID := hostFields[0].Descriptor()
	// host.DefaultID holds the default value on creation for the id field.
	host.DefaultID = hostDescID.Default.(func() uuid.UUID)
	hostdependencyFields := schema.HostDependency{}.Fields()
	_ = hostdependencyFields
	// hostdependencyDescID is the schema descriptor for id field.
	hostdependencyDescID := hostdependencyFields[0].Descriptor()
	// hostdependency.DefaultID holds the default value on creation for the id field.
	hostdependency.DefaultID = hostdependencyDescID.Default.(func() uuid.UUID)
	identityFields := schema.Identity{}.Fields()
	_ = identityFields
	// identityDescID is the schema descriptor for id field.
	identityDescID := identityFields[0].Descriptor()
	// identity.DefaultID holds the default value on creation for the id field.
	identity.DefaultID = identityDescID.Default.(func() uuid.UUID)
	includednetworkFields := schema.IncludedNetwork{}.Fields()
	_ = includednetworkFields
	// includednetworkDescID is the schema descriptor for id field.
	includednetworkDescID := includednetworkFields[0].Descriptor()
	// includednetwork.DefaultID holds the default value on creation for the id field.
	includednetwork.DefaultID = includednetworkDescID.Default.(func() uuid.UUID)
	networkFields := schema.Network{}.Fields()
	_ = networkFields
	// networkDescID is the schema descriptor for id field.
	networkDescID := networkFields[0].Descriptor()
	// network.DefaultID holds the default value on creation for the id field.
	network.DefaultID = networkDescID.Default.(func() uuid.UUID)
	planFields := schema.Plan{}.Fields()
	_ = planFields
	// planDescID is the schema descriptor for id field.
	planDescID := planFields[0].Descriptor()
	// plan.DefaultID holds the default value on creation for the id field.
	plan.DefaultID = planDescID.Default.(func() uuid.UUID)
	plandiffFields := schema.PlanDiff{}.Fields()
	_ = plandiffFields
	// plandiffDescID is the schema descriptor for id field.
	plandiffDescID := plandiffFields[0].Descriptor()
	// plandiff.DefaultID holds the default value on creation for the id field.
	plandiff.DefaultID = plandiffDescID.Default.(func() uuid.UUID)
	provisionedhostFields := schema.ProvisionedHost{}.Fields()
	_ = provisionedhostFields
	// provisionedhostDescID is the schema descriptor for id field.
	provisionedhostDescID := provisionedhostFields[0].Descriptor()
	// provisionedhost.DefaultID holds the default value on creation for the id field.
	provisionedhost.DefaultID = provisionedhostDescID.Default.(func() uuid.UUID)
	provisionednetworkFields := schema.ProvisionedNetwork{}.Fields()
	_ = provisionednetworkFields
	// provisionednetworkDescID is the schema descriptor for id field.
	provisionednetworkDescID := provisionednetworkFields[0].Descriptor()
	// provisionednetwork.DefaultID holds the default value on creation for the id field.
	provisionednetwork.DefaultID = provisionednetworkDescID.Default.(func() uuid.UUID)
	provisioningstepFields := schema.ProvisioningStep{}.Fields()
	_ = provisioningstepFields
	// provisioningstepDescID is the schema descriptor for id field.
	provisioningstepDescID := provisioningstepFields[0].Descriptor()
	// provisioningstep.DefaultID holds the default value on creation for the id field.
	provisioningstep.DefaultID = provisioningstepDescID.Default.(func() uuid.UUID)
	repositoryFields := schema.Repository{}.Fields()
	_ = repositoryFields
	// repositoryDescBranchName is the schema descriptor for branch_name field.
	repositoryDescBranchName := repositoryFields[2].Descriptor()
	// repository.DefaultBranchName holds the default value on creation for the branch_name field.
	repository.DefaultBranchName = repositoryDescBranchName.Default.(string)
	// repositoryDescFolderPath is the schema descriptor for folder_path field.
	repositoryDescFolderPath := repositoryFields[4].Descriptor()
	// repository.DefaultFolderPath holds the default value on creation for the folder_path field.
	repository.DefaultFolderPath = repositoryDescFolderPath.Default.(string)
	// repositoryDescCommitInfo is the schema descriptor for commit_info field.
	repositoryDescCommitInfo := repositoryFields[5].Descriptor()
	// repository.DefaultCommitInfo holds the default value on creation for the commit_info field.
	repository.DefaultCommitInfo = repositoryDescCommitInfo.Default.(string)
	// repositoryDescID is the schema descriptor for id field.
	repositoryDescID := repositoryFields[0].Descriptor()
	// repository.DefaultID holds the default value on creation for the id field.
	repository.DefaultID = repositoryDescID.Default.(func() uuid.UUID)
	scriptFields := schema.Script{}.Fields()
	_ = scriptFields
	// scriptDescID is the schema descriptor for id field.
	scriptDescID := scriptFields[0].Descriptor()
	// script.DefaultID holds the default value on creation for the id field.
	script.DefaultID = scriptDescID.Default.(func() uuid.UUID)
	servertaskFields := schema.ServerTask{}.Fields()
	_ = servertaskFields
	// servertaskDescID is the schema descriptor for id field.
	servertaskDescID := servertaskFields[0].Descriptor()
	// servertask.DefaultID holds the default value on creation for the id field.
	servertask.DefaultID = servertaskDescID.Default.(func() uuid.UUID)
	statusFields := schema.Status{}.Fields()
	_ = statusFields
	// statusDescFailed is the schema descriptor for failed field.
	statusDescFailed := statusFields[5].Descriptor()
	// status.DefaultFailed holds the default value on creation for the failed field.
	status.DefaultFailed = statusDescFailed.Default.(bool)
	// statusDescCompleted is the schema descriptor for completed field.
	statusDescCompleted := statusFields[6].Descriptor()
	// status.DefaultCompleted holds the default value on creation for the completed field.
	status.DefaultCompleted = statusDescCompleted.Default.(bool)
	// statusDescID is the schema descriptor for id field.
	statusDescID := statusFields[0].Descriptor()
	// status.DefaultID holds the default value on creation for the id field.
	status.DefaultID = statusDescID.Default.(func() uuid.UUID)
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescID is the schema descriptor for id field.
	tagDescID := tagFields[0].Descriptor()
	// tag.DefaultID holds the default value on creation for the id field.
	tag.DefaultID = tagDescID.Default.(func() uuid.UUID)
	teamFields := schema.Team{}.Fields()
	_ = teamFields
	// teamDescID is the schema descriptor for id field.
	teamDescID := teamFields[0].Descriptor()
	// team.DefaultID holds the default value on creation for the id field.
	team.DefaultID = teamDescID.Default.(func() uuid.UUID)
	tokenFields := schema.Token{}.Fields()
	_ = tokenFields
	// tokenDescID is the schema descriptor for id field.
	tokenDescID := tokenFields[0].Descriptor()
	// token.DefaultID holds the default value on creation for the id field.
	token.DefaultID = tokenDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
	validationFields := schema.Validation{}.Fields()
	_ = validationFields
	// validationDescValidationType is the schema descriptor for validation_type field.
	validationDescValidationType := validationFields[2].Descriptor()
	// validation.DefaultValidationType holds the default value on creation for the validation_type field.
	validation.DefaultValidationType = validationDescValidationType.Default.(string)
	// validationDescOutput is the schema descriptor for output field.
	validationDescOutput := validationFields[3].Descriptor()
	// validation.DefaultOutput holds the default value on creation for the output field.
	validation.DefaultOutput = validationDescOutput.Default.(string)
	// validationDescErrorMessage is the schema descriptor for error_message field.
	validationDescErrorMessage := validationFields[5].Descriptor()
	// validation.DefaultErrorMessage holds the default value on creation for the error_message field.
	validation.DefaultErrorMessage = validationDescErrorMessage.Default.(string)
	// validationDescID is the schema descriptor for id field.
	validationDescID := validationFields[0].Descriptor()
	// validation.DefaultID holds the default value on creation for the id field.
	validation.DefaultID = validationDescID.Default.(func() uuid.UUID)
}
