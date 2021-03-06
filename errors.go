package qortexapi

import "errors"

var (
	ServerError                    = errors.New("Oops, something is wrong!")
	InvalidIdError                 = errors.New("Invalid Id error!")
	UnexpectedError                = errors.New("Unexpected error!")
	OrganizationNotFoundError      = errors.New("organization not found")
	OrganizationNotJoinedError     = errors.New("not join organization yet")
	OrganizationsNotFoundError     = errors.New("organizations not found")
	NonStandardOrganizationError   = errors.New("this organization is not standard")
	NonSandboxOrganizationError    = errors.New("this organization is not sandbox")
	GroupNotFoundError             = errors.New("group not found")
	GroupsNotFoundError            = errors.New("groups not found")
	MemberNotFoundError            = errors.New("member not found")
	MemberNotLogginError           = errors.New("member not loggin yet")
	ChangePasswordError            = errors.New("change password error")
	UserNotFoundError              = errors.New("user not found")
	UsersNotFoundError             = errors.New("users not found")
	PermissionDeniedError          = errors.New("Permiession Denied")
	SaveMemberError                = errors.New("can not save member")
	SaveUserError                  = errors.New("can not save user")
	SaveGroupError                 = errors.New("can not save group")
	SaveOrganizationError          = errors.New("can not save organization")
	DeleteGroupError               = errors.New("can not delete group")
	BackupEntryError               = errors.New("can not backup entry")
	EntryNotFoundError             = errors.New("entry not found")
	RemoveIndexError               = errors.New("can not remove index")
	SaveEntryError                 = errors.New("can not save entry")
	UserIsDeletedError             = errors.New("user is already deleted")
	UserIsStillAvailableError      = errors.New("user is still available")
	SelfActionError                = errors.New("can not do this for your self")
	FollowUserError                = errors.New("can not follow user")
	UnfollowUserError              = errors.New("can not unfollow user")
	AlreadyJoinedOrganizationError = errors.New("already joined the organization")
	JoinOrganizationError          = errors.New("join organization error")
	SystemMessagesNotFoundError    = errors.New("system messages not found")
	BroadcastTypeError             = errors.New("no such broadcast type")
	DatabaseNotExistError          = errors.New("no such database")
	UnknownEntryType               = errors.New("unknown entry type")
	UpdateLikeError                = errors.New("update like error")
	TokenInvalid                   = errors.New("Token is not available")
	AccountNotActivated            = errors.New("Account was not activated yet")
	AccountAlreadyActivated        = errors.New("Account was already activated")
	EntryBlankError                = errors.New("Entry can not be blank")
	GroupBlankError                = errors.New("Group can not be blank")
	UserBlankError                 = errors.New("User can not be blank")
	TakeActionError                = errors.New("Can not take the action")
	HighFrequencyError             = errors.New("High frequency requests")
	SaveTranscriptionError         = errors.New("Can not save transcription")
	ClaimTaskError                 = errors.New("Task was already cliamed by others")
	NeedPayError                   = errors.New("Qortex need to pay")
	InvalidLanguageCodeError       = errors.New("Language code is not supported")
	NotFoundError                  = errors.New("not found")
	InvitationAlreadyAcceptedError = errors.New("Invitation was accepted")
	InvitationAlreadyCanceledError = errors.New("Invitation was canceled")
	ChangeEmailTypeNotProvided     = errors.New("Should provide the type of changing email")
)
