// Auto-generated by avdl-compiler v1.3.20 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/chat1/notify.avdl

package chat1

import (
	"errors"
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type ChatActivityType int

const (
	ChatActivityType_RESERVED                      ChatActivityType = 0
	ChatActivityType_INCOMING_MESSAGE              ChatActivityType = 1
	ChatActivityType_READ_MESSAGE                  ChatActivityType = 2
	ChatActivityType_NEW_CONVERSATION              ChatActivityType = 3
	ChatActivityType_SET_STATUS                    ChatActivityType = 4
	ChatActivityType_FAILED_MESSAGE                ChatActivityType = 5
	ChatActivityType_MEMBERS_UPDATE                ChatActivityType = 6
	ChatActivityType_SET_APP_NOTIFICATION_SETTINGS ChatActivityType = 7
	ChatActivityType_TEAMTYPE                      ChatActivityType = 8
)

func (o ChatActivityType) DeepCopy() ChatActivityType { return o }

var ChatActivityTypeMap = map[string]ChatActivityType{
	"RESERVED":                      0,
	"INCOMING_MESSAGE":              1,
	"READ_MESSAGE":                  2,
	"NEW_CONVERSATION":              3,
	"SET_STATUS":                    4,
	"FAILED_MESSAGE":                5,
	"MEMBERS_UPDATE":                6,
	"SET_APP_NOTIFICATION_SETTINGS": 7,
	"TEAMTYPE":                      8,
}

var ChatActivityTypeRevMap = map[ChatActivityType]string{
	0: "RESERVED",
	1: "INCOMING_MESSAGE",
	2: "READ_MESSAGE",
	3: "NEW_CONVERSATION",
	4: "SET_STATUS",
	5: "FAILED_MESSAGE",
	6: "MEMBERS_UPDATE",
	7: "SET_APP_NOTIFICATION_SETTINGS",
	8: "TEAMTYPE",
}

func (e ChatActivityType) String() string {
	if v, ok := ChatActivityTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type IncomingMessage struct {
	Message                    UIMessage      `codec:"message" json:"message"`
	ConvID                     ConversationID `codec:"convID" json:"convID"`
	DisplayDesktopNotification bool           `codec:"displayDesktopNotification" json:"displayDesktopNotification"`
	Conv                       *InboxUIItem   `codec:"conv,omitempty" json:"conv,omitempty"`
	Pagination                 *UIPagination  `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

func (o IncomingMessage) DeepCopy() IncomingMessage {
	return IncomingMessage{
		Message: o.Message.DeepCopy(),
		ConvID:  o.ConvID.DeepCopy(),
		DisplayDesktopNotification: o.DisplayDesktopNotification,
		Conv: (func(x *InboxUIItem) *InboxUIItem {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Conv),
		Pagination: (func(x *UIPagination) *UIPagination {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Pagination),
	}
}

type ReadMessageInfo struct {
	ConvID ConversationID `codec:"convID" json:"convID"`
	MsgID  MessageID      `codec:"msgID" json:"msgID"`
	Conv   *InboxUIItem   `codec:"conv,omitempty" json:"conv,omitempty"`
}

func (o ReadMessageInfo) DeepCopy() ReadMessageInfo {
	return ReadMessageInfo{
		ConvID: o.ConvID.DeepCopy(),
		MsgID:  o.MsgID.DeepCopy(),
		Conv: (func(x *InboxUIItem) *InboxUIItem {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Conv),
	}
}

type NewConversationInfo struct {
	Conv InboxUIItem `codec:"conv" json:"conv"`
}

func (o NewConversationInfo) DeepCopy() NewConversationInfo {
	return NewConversationInfo{
		Conv: o.Conv.DeepCopy(),
	}
}

type SetStatusInfo struct {
	ConvID ConversationID     `codec:"convID" json:"convID"`
	Status ConversationStatus `codec:"status" json:"status"`
	Conv   *InboxUIItem       `codec:"conv,omitempty" json:"conv,omitempty"`
}

func (o SetStatusInfo) DeepCopy() SetStatusInfo {
	return SetStatusInfo{
		ConvID: o.ConvID.DeepCopy(),
		Status: o.Status.DeepCopy(),
		Conv: (func(x *InboxUIItem) *InboxUIItem {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Conv),
	}
}

type SetAppNotificationSettingsInfo struct {
	ConvID   ConversationID               `codec:"convID" json:"convID"`
	Settings ConversationNotificationInfo `codec:"settings" json:"settings"`
}

func (o SetAppNotificationSettingsInfo) DeepCopy() SetAppNotificationSettingsInfo {
	return SetAppNotificationSettingsInfo{
		ConvID:   o.ConvID.DeepCopy(),
		Settings: o.Settings.DeepCopy(),
	}
}

type FailedMessageInfo struct {
	OutboxRecords []OutboxRecord `codec:"outboxRecords" json:"outboxRecords"`
}

func (o FailedMessageInfo) DeepCopy() FailedMessageInfo {
	return FailedMessageInfo{
		OutboxRecords: (func(x []OutboxRecord) []OutboxRecord {
			if x == nil {
				return nil
			}
			var ret []OutboxRecord
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.OutboxRecords),
	}
}

type MembersUpdateInfo struct {
	ConvID ConversationID `codec:"convID" json:"convID"`
	Member string         `codec:"member" json:"member"`
	Joined bool           `codec:"joined" json:"joined"`
}

func (o MembersUpdateInfo) DeepCopy() MembersUpdateInfo {
	return MembersUpdateInfo{
		ConvID: o.ConvID.DeepCopy(),
		Member: o.Member,
		Joined: o.Joined,
	}
}

type TeamTypeInfo struct {
	ConvID   ConversationID `codec:"convID" json:"convID"`
	TeamType TeamType       `codec:"teamType" json:"teamType"`
	Conv     *InboxUIItem   `codec:"conv,omitempty" json:"conv,omitempty"`
}

func (o TeamTypeInfo) DeepCopy() TeamTypeInfo {
	return TeamTypeInfo{
		ConvID:   o.ConvID.DeepCopy(),
		TeamType: o.TeamType.DeepCopy(),
		Conv: (func(x *InboxUIItem) *InboxUIItem {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Conv),
	}
}

type ChatActivity struct {
	ActivityType__               ChatActivityType                `codec:"activityType" json:"activityType"`
	IncomingMessage__            *IncomingMessage                `codec:"incomingMessage,omitempty" json:"incomingMessage,omitempty"`
	ReadMessage__                *ReadMessageInfo                `codec:"readMessage,omitempty" json:"readMessage,omitempty"`
	NewConversation__            *NewConversationInfo            `codec:"newConversation,omitempty" json:"newConversation,omitempty"`
	SetStatus__                  *SetStatusInfo                  `codec:"setStatus,omitempty" json:"setStatus,omitempty"`
	FailedMessage__              *FailedMessageInfo              `codec:"failedMessage,omitempty" json:"failedMessage,omitempty"`
	MembersUpdate__              *MembersUpdateInfo              `codec:"membersUpdate,omitempty" json:"membersUpdate,omitempty"`
	SetAppNotificationSettings__ *SetAppNotificationSettingsInfo `codec:"setAppNotificationSettings,omitempty" json:"setAppNotificationSettings,omitempty"`
	Teamtype__                   *TeamTypeInfo                   `codec:"teamtype,omitempty" json:"teamtype,omitempty"`
}

func (o *ChatActivity) ActivityType() (ret ChatActivityType, err error) {
	switch o.ActivityType__ {
	case ChatActivityType_INCOMING_MESSAGE:
		if o.IncomingMessage__ == nil {
			err = errors.New("unexpected nil value for IncomingMessage__")
			return ret, err
		}
	case ChatActivityType_READ_MESSAGE:
		if o.ReadMessage__ == nil {
			err = errors.New("unexpected nil value for ReadMessage__")
			return ret, err
		}
	case ChatActivityType_NEW_CONVERSATION:
		if o.NewConversation__ == nil {
			err = errors.New("unexpected nil value for NewConversation__")
			return ret, err
		}
	case ChatActivityType_SET_STATUS:
		if o.SetStatus__ == nil {
			err = errors.New("unexpected nil value for SetStatus__")
			return ret, err
		}
	case ChatActivityType_FAILED_MESSAGE:
		if o.FailedMessage__ == nil {
			err = errors.New("unexpected nil value for FailedMessage__")
			return ret, err
		}
	case ChatActivityType_MEMBERS_UPDATE:
		if o.MembersUpdate__ == nil {
			err = errors.New("unexpected nil value for MembersUpdate__")
			return ret, err
		}
	case ChatActivityType_SET_APP_NOTIFICATION_SETTINGS:
		if o.SetAppNotificationSettings__ == nil {
			err = errors.New("unexpected nil value for SetAppNotificationSettings__")
			return ret, err
		}
	case ChatActivityType_TEAMTYPE:
		if o.Teamtype__ == nil {
			err = errors.New("unexpected nil value for Teamtype__")
			return ret, err
		}
	}
	return o.ActivityType__, nil
}

func (o ChatActivity) IncomingMessage() (res IncomingMessage) {
	if o.ActivityType__ != ChatActivityType_INCOMING_MESSAGE {
		panic("wrong case accessed")
	}
	if o.IncomingMessage__ == nil {
		return
	}
	return *o.IncomingMessage__
}

func (o ChatActivity) ReadMessage() (res ReadMessageInfo) {
	if o.ActivityType__ != ChatActivityType_READ_MESSAGE {
		panic("wrong case accessed")
	}
	if o.ReadMessage__ == nil {
		return
	}
	return *o.ReadMessage__
}

func (o ChatActivity) NewConversation() (res NewConversationInfo) {
	if o.ActivityType__ != ChatActivityType_NEW_CONVERSATION {
		panic("wrong case accessed")
	}
	if o.NewConversation__ == nil {
		return
	}
	return *o.NewConversation__
}

func (o ChatActivity) SetStatus() (res SetStatusInfo) {
	if o.ActivityType__ != ChatActivityType_SET_STATUS {
		panic("wrong case accessed")
	}
	if o.SetStatus__ == nil {
		return
	}
	return *o.SetStatus__
}

func (o ChatActivity) FailedMessage() (res FailedMessageInfo) {
	if o.ActivityType__ != ChatActivityType_FAILED_MESSAGE {
		panic("wrong case accessed")
	}
	if o.FailedMessage__ == nil {
		return
	}
	return *o.FailedMessage__
}

func (o ChatActivity) MembersUpdate() (res MembersUpdateInfo) {
	if o.ActivityType__ != ChatActivityType_MEMBERS_UPDATE {
		panic("wrong case accessed")
	}
	if o.MembersUpdate__ == nil {
		return
	}
	return *o.MembersUpdate__
}

func (o ChatActivity) SetAppNotificationSettings() (res SetAppNotificationSettingsInfo) {
	if o.ActivityType__ != ChatActivityType_SET_APP_NOTIFICATION_SETTINGS {
		panic("wrong case accessed")
	}
	if o.SetAppNotificationSettings__ == nil {
		return
	}
	return *o.SetAppNotificationSettings__
}

func (o ChatActivity) Teamtype() (res TeamTypeInfo) {
	if o.ActivityType__ != ChatActivityType_TEAMTYPE {
		panic("wrong case accessed")
	}
	if o.Teamtype__ == nil {
		return
	}
	return *o.Teamtype__
}

func NewChatActivityWithIncomingMessage(v IncomingMessage) ChatActivity {
	return ChatActivity{
		ActivityType__:    ChatActivityType_INCOMING_MESSAGE,
		IncomingMessage__: &v,
	}
}

func NewChatActivityWithReadMessage(v ReadMessageInfo) ChatActivity {
	return ChatActivity{
		ActivityType__: ChatActivityType_READ_MESSAGE,
		ReadMessage__:  &v,
	}
}

func NewChatActivityWithNewConversation(v NewConversationInfo) ChatActivity {
	return ChatActivity{
		ActivityType__:    ChatActivityType_NEW_CONVERSATION,
		NewConversation__: &v,
	}
}

func NewChatActivityWithSetStatus(v SetStatusInfo) ChatActivity {
	return ChatActivity{
		ActivityType__: ChatActivityType_SET_STATUS,
		SetStatus__:    &v,
	}
}

func NewChatActivityWithFailedMessage(v FailedMessageInfo) ChatActivity {
	return ChatActivity{
		ActivityType__:  ChatActivityType_FAILED_MESSAGE,
		FailedMessage__: &v,
	}
}

func NewChatActivityWithMembersUpdate(v MembersUpdateInfo) ChatActivity {
	return ChatActivity{
		ActivityType__:  ChatActivityType_MEMBERS_UPDATE,
		MembersUpdate__: &v,
	}
}

func NewChatActivityWithSetAppNotificationSettings(v SetAppNotificationSettingsInfo) ChatActivity {
	return ChatActivity{
		ActivityType__:               ChatActivityType_SET_APP_NOTIFICATION_SETTINGS,
		SetAppNotificationSettings__: &v,
	}
}

func NewChatActivityWithTeamtype(v TeamTypeInfo) ChatActivity {
	return ChatActivity{
		ActivityType__: ChatActivityType_TEAMTYPE,
		Teamtype__:     &v,
	}
}

func (o ChatActivity) DeepCopy() ChatActivity {
	return ChatActivity{
		ActivityType__: o.ActivityType__.DeepCopy(),
		IncomingMessage__: (func(x *IncomingMessage) *IncomingMessage {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.IncomingMessage__),
		ReadMessage__: (func(x *ReadMessageInfo) *ReadMessageInfo {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.ReadMessage__),
		NewConversation__: (func(x *NewConversationInfo) *NewConversationInfo {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.NewConversation__),
		SetStatus__: (func(x *SetStatusInfo) *SetStatusInfo {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.SetStatus__),
		FailedMessage__: (func(x *FailedMessageInfo) *FailedMessageInfo {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.FailedMessage__),
		MembersUpdate__: (func(x *MembersUpdateInfo) *MembersUpdateInfo {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.MembersUpdate__),
		SetAppNotificationSettings__: (func(x *SetAppNotificationSettingsInfo) *SetAppNotificationSettingsInfo {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.SetAppNotificationSettings__),
		Teamtype__: (func(x *TeamTypeInfo) *TeamTypeInfo {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Teamtype__),
	}
}

type TyperInfo struct {
	Uid        keybase1.UID      `codec:"uid" json:"uid"`
	Username   string            `codec:"username" json:"username"`
	DeviceID   keybase1.DeviceID `codec:"deviceID" json:"deviceID"`
	DeviceName string            `codec:"deviceName" json:"deviceName"`
	DeviceType string            `codec:"deviceType" json:"deviceType"`
}

func (o TyperInfo) DeepCopy() TyperInfo {
	return TyperInfo{
		Uid:        o.Uid.DeepCopy(),
		Username:   o.Username,
		DeviceID:   o.DeviceID.DeepCopy(),
		DeviceName: o.DeviceName,
		DeviceType: o.DeviceType,
	}
}

type ConvTypingUpdate struct {
	ConvID ConversationID `codec:"convID" json:"convID"`
	Typers []TyperInfo    `codec:"typers" json:"typers"`
}

func (o ConvTypingUpdate) DeepCopy() ConvTypingUpdate {
	return ConvTypingUpdate{
		ConvID: o.ConvID.DeepCopy(),
		Typers: (func(x []TyperInfo) []TyperInfo {
			if x == nil {
				return nil
			}
			var ret []TyperInfo
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Typers),
	}
}

type StaleUpdateType int

const (
	StaleUpdateType_CLEAR       StaleUpdateType = 0
	StaleUpdateType_NEWACTIVITY StaleUpdateType = 1
)

func (o StaleUpdateType) DeepCopy() StaleUpdateType { return o }

var StaleUpdateTypeMap = map[string]StaleUpdateType{
	"CLEAR":       0,
	"NEWACTIVITY": 1,
}

var StaleUpdateTypeRevMap = map[StaleUpdateType]string{
	0: "CLEAR",
	1: "NEWACTIVITY",
}

func (e StaleUpdateType) String() string {
	if v, ok := StaleUpdateTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type ConversationStaleUpdate struct {
	ConvID     ConversationID  `codec:"convID" json:"convID"`
	UpdateType StaleUpdateType `codec:"updateType" json:"updateType"`
}

func (o ConversationStaleUpdate) DeepCopy() ConversationStaleUpdate {
	return ConversationStaleUpdate{
		ConvID:     o.ConvID.DeepCopy(),
		UpdateType: o.UpdateType.DeepCopy(),
	}
}

type NewChatActivityArg struct {
	Uid      keybase1.UID `codec:"uid" json:"uid"`
	Activity ChatActivity `codec:"activity" json:"activity"`
}

func (o NewChatActivityArg) DeepCopy() NewChatActivityArg {
	return NewChatActivityArg{
		Uid:      o.Uid.DeepCopy(),
		Activity: o.Activity.DeepCopy(),
	}
}

type ChatIdentifyUpdateArg struct {
	Update keybase1.CanonicalTLFNameAndIDWithBreaks `codec:"update" json:"update"`
}

func (o ChatIdentifyUpdateArg) DeepCopy() ChatIdentifyUpdateArg {
	return ChatIdentifyUpdateArg{
		Update: o.Update.DeepCopy(),
	}
}

type ChatTLFFinalizeArg struct {
	Uid          keybase1.UID             `codec:"uid" json:"uid"`
	ConvID       ConversationID           `codec:"convID" json:"convID"`
	FinalizeInfo ConversationFinalizeInfo `codec:"finalizeInfo" json:"finalizeInfo"`
	Conv         *InboxUIItem             `codec:"conv,omitempty" json:"conv,omitempty"`
}

func (o ChatTLFFinalizeArg) DeepCopy() ChatTLFFinalizeArg {
	return ChatTLFFinalizeArg{
		Uid:          o.Uid.DeepCopy(),
		ConvID:       o.ConvID.DeepCopy(),
		FinalizeInfo: o.FinalizeInfo.DeepCopy(),
		Conv: (func(x *InboxUIItem) *InboxUIItem {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Conv),
	}
}

type ChatTLFResolveArg struct {
	Uid         keybase1.UID            `codec:"uid" json:"uid"`
	ConvID      ConversationID          `codec:"convID" json:"convID"`
	ResolveInfo ConversationResolveInfo `codec:"resolveInfo" json:"resolveInfo"`
}

func (o ChatTLFResolveArg) DeepCopy() ChatTLFResolveArg {
	return ChatTLFResolveArg{
		Uid:         o.Uid.DeepCopy(),
		ConvID:      o.ConvID.DeepCopy(),
		ResolveInfo: o.ResolveInfo.DeepCopy(),
	}
}

type ChatInboxStaleArg struct {
	Uid keybase1.UID `codec:"uid" json:"uid"`
}

func (o ChatInboxStaleArg) DeepCopy() ChatInboxStaleArg {
	return ChatInboxStaleArg{
		Uid: o.Uid.DeepCopy(),
	}
}

type ChatThreadsStaleArg struct {
	Uid     keybase1.UID              `codec:"uid" json:"uid"`
	Updates []ConversationStaleUpdate `codec:"updates" json:"updates"`
}

func (o ChatThreadsStaleArg) DeepCopy() ChatThreadsStaleArg {
	return ChatThreadsStaleArg{
		Uid: o.Uid.DeepCopy(),
		Updates: (func(x []ConversationStaleUpdate) []ConversationStaleUpdate {
			if x == nil {
				return nil
			}
			var ret []ConversationStaleUpdate
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Updates),
	}
}

type ChatTypingUpdateArg struct {
	TypingUpdates []ConvTypingUpdate `codec:"typingUpdates" json:"typingUpdates"`
}

func (o ChatTypingUpdateArg) DeepCopy() ChatTypingUpdateArg {
	return ChatTypingUpdateArg{
		TypingUpdates: (func(x []ConvTypingUpdate) []ConvTypingUpdate {
			if x == nil {
				return nil
			}
			var ret []ConvTypingUpdate
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.TypingUpdates),
	}
}

type ChatJoinedConversationArg struct {
	Uid  keybase1.UID `codec:"uid" json:"uid"`
	Conv InboxUIItem  `codec:"conv" json:"conv"`
}

func (o ChatJoinedConversationArg) DeepCopy() ChatJoinedConversationArg {
	return ChatJoinedConversationArg{
		Uid:  o.Uid.DeepCopy(),
		Conv: o.Conv.DeepCopy(),
	}
}

type ChatLeftConversationArg struct {
	Uid    keybase1.UID   `codec:"uid" json:"uid"`
	ConvID ConversationID `codec:"convID" json:"convID"`
}

func (o ChatLeftConversationArg) DeepCopy() ChatLeftConversationArg {
	return ChatLeftConversationArg{
		Uid:    o.Uid.DeepCopy(),
		ConvID: o.ConvID.DeepCopy(),
	}
}

type ChatInboxSyncStartedArg struct {
	Uid keybase1.UID `codec:"uid" json:"uid"`
}

func (o ChatInboxSyncStartedArg) DeepCopy() ChatInboxSyncStartedArg {
	return ChatInboxSyncStartedArg{
		Uid: o.Uid.DeepCopy(),
	}
}

type ChatInboxSyncedArg struct {
	Uid   keybase1.UID            `codec:"uid" json:"uid"`
	Convs []UnverifiedInboxUIItem `codec:"convs" json:"convs"`
}

func (o ChatInboxSyncedArg) DeepCopy() ChatInboxSyncedArg {
	return ChatInboxSyncedArg{
		Uid: o.Uid.DeepCopy(),
		Convs: (func(x []UnverifiedInboxUIItem) []UnverifiedInboxUIItem {
			if x == nil {
				return nil
			}
			var ret []UnverifiedInboxUIItem
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Convs),
	}
}

type NotifyChatInterface interface {
	NewChatActivity(context.Context, NewChatActivityArg) error
	ChatIdentifyUpdate(context.Context, keybase1.CanonicalTLFNameAndIDWithBreaks) error
	ChatTLFFinalize(context.Context, ChatTLFFinalizeArg) error
	ChatTLFResolve(context.Context, ChatTLFResolveArg) error
	ChatInboxStale(context.Context, keybase1.UID) error
	ChatThreadsStale(context.Context, ChatThreadsStaleArg) error
	ChatTypingUpdate(context.Context, []ConvTypingUpdate) error
	ChatJoinedConversation(context.Context, ChatJoinedConversationArg) error
	ChatLeftConversation(context.Context, ChatLeftConversationArg) error
	ChatInboxSyncStarted(context.Context, keybase1.UID) error
	ChatInboxSynced(context.Context, ChatInboxSyncedArg) error
}

func NotifyChatProtocol(i NotifyChatInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "chat.1.NotifyChat",
		Methods: map[string]rpc.ServeHandlerDescription{
			"NewChatActivity": {
				MakeArg: func() interface{} {
					ret := make([]NewChatActivityArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]NewChatActivityArg)
					if !ok {
						err = rpc.NewTypeError((*[]NewChatActivityArg)(nil), args)
						return
					}
					err = i.NewChatActivity(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"ChatIdentifyUpdate": {
				MakeArg: func() interface{} {
					ret := make([]ChatIdentifyUpdateArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatIdentifyUpdateArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatIdentifyUpdateArg)(nil), args)
						return
					}
					err = i.ChatIdentifyUpdate(ctx, (*typedArgs)[0].Update)
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"ChatTLFFinalize": {
				MakeArg: func() interface{} {
					ret := make([]ChatTLFFinalizeArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatTLFFinalizeArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatTLFFinalizeArg)(nil), args)
						return
					}
					err = i.ChatTLFFinalize(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"ChatTLFResolve": {
				MakeArg: func() interface{} {
					ret := make([]ChatTLFResolveArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatTLFResolveArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatTLFResolveArg)(nil), args)
						return
					}
					err = i.ChatTLFResolve(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"ChatInboxStale": {
				MakeArg: func() interface{} {
					ret := make([]ChatInboxStaleArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatInboxStaleArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatInboxStaleArg)(nil), args)
						return
					}
					err = i.ChatInboxStale(ctx, (*typedArgs)[0].Uid)
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"ChatThreadsStale": {
				MakeArg: func() interface{} {
					ret := make([]ChatThreadsStaleArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatThreadsStaleArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatThreadsStaleArg)(nil), args)
						return
					}
					err = i.ChatThreadsStale(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"ChatTypingUpdate": {
				MakeArg: func() interface{} {
					ret := make([]ChatTypingUpdateArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatTypingUpdateArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatTypingUpdateArg)(nil), args)
						return
					}
					err = i.ChatTypingUpdate(ctx, (*typedArgs)[0].TypingUpdates)
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"ChatJoinedConversation": {
				MakeArg: func() interface{} {
					ret := make([]ChatJoinedConversationArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatJoinedConversationArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatJoinedConversationArg)(nil), args)
						return
					}
					err = i.ChatJoinedConversation(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"ChatLeftConversation": {
				MakeArg: func() interface{} {
					ret := make([]ChatLeftConversationArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatLeftConversationArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatLeftConversationArg)(nil), args)
						return
					}
					err = i.ChatLeftConversation(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"ChatInboxSyncStarted": {
				MakeArg: func() interface{} {
					ret := make([]ChatInboxSyncStartedArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatInboxSyncStartedArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatInboxSyncStartedArg)(nil), args)
						return
					}
					err = i.ChatInboxSyncStarted(ctx, (*typedArgs)[0].Uid)
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"ChatInboxSynced": {
				MakeArg: func() interface{} {
					ret := make([]ChatInboxSyncedArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatInboxSyncedArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatInboxSyncedArg)(nil), args)
						return
					}
					err = i.ChatInboxSynced(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodNotify,
			},
		},
	}
}

type NotifyChatClient struct {
	Cli rpc.GenericClient
}

func (c NotifyChatClient) NewChatActivity(ctx context.Context, __arg NewChatActivityArg) (err error) {
	err = c.Cli.Notify(ctx, "chat.1.NotifyChat.NewChatActivity", []interface{}{__arg})
	return
}

func (c NotifyChatClient) ChatIdentifyUpdate(ctx context.Context, update keybase1.CanonicalTLFNameAndIDWithBreaks) (err error) {
	__arg := ChatIdentifyUpdateArg{Update: update}
	err = c.Cli.Notify(ctx, "chat.1.NotifyChat.ChatIdentifyUpdate", []interface{}{__arg})
	return
}

func (c NotifyChatClient) ChatTLFFinalize(ctx context.Context, __arg ChatTLFFinalizeArg) (err error) {
	err = c.Cli.Notify(ctx, "chat.1.NotifyChat.ChatTLFFinalize", []interface{}{__arg})
	return
}

func (c NotifyChatClient) ChatTLFResolve(ctx context.Context, __arg ChatTLFResolveArg) (err error) {
	err = c.Cli.Notify(ctx, "chat.1.NotifyChat.ChatTLFResolve", []interface{}{__arg})
	return
}

func (c NotifyChatClient) ChatInboxStale(ctx context.Context, uid keybase1.UID) (err error) {
	__arg := ChatInboxStaleArg{Uid: uid}
	err = c.Cli.Notify(ctx, "chat.1.NotifyChat.ChatInboxStale", []interface{}{__arg})
	return
}

func (c NotifyChatClient) ChatThreadsStale(ctx context.Context, __arg ChatThreadsStaleArg) (err error) {
	err = c.Cli.Notify(ctx, "chat.1.NotifyChat.ChatThreadsStale", []interface{}{__arg})
	return
}

func (c NotifyChatClient) ChatTypingUpdate(ctx context.Context, typingUpdates []ConvTypingUpdate) (err error) {
	__arg := ChatTypingUpdateArg{TypingUpdates: typingUpdates}
	err = c.Cli.Notify(ctx, "chat.1.NotifyChat.ChatTypingUpdate", []interface{}{__arg})
	return
}

func (c NotifyChatClient) ChatJoinedConversation(ctx context.Context, __arg ChatJoinedConversationArg) (err error) {
	err = c.Cli.Notify(ctx, "chat.1.NotifyChat.ChatJoinedConversation", []interface{}{__arg})
	return
}

func (c NotifyChatClient) ChatLeftConversation(ctx context.Context, __arg ChatLeftConversationArg) (err error) {
	err = c.Cli.Notify(ctx, "chat.1.NotifyChat.ChatLeftConversation", []interface{}{__arg})
	return
}

func (c NotifyChatClient) ChatInboxSyncStarted(ctx context.Context, uid keybase1.UID) (err error) {
	__arg := ChatInboxSyncStartedArg{Uid: uid}
	err = c.Cli.Notify(ctx, "chat.1.NotifyChat.ChatInboxSyncStarted", []interface{}{__arg})
	return
}

func (c NotifyChatClient) ChatInboxSynced(ctx context.Context, __arg ChatInboxSyncedArg) (err error) {
	err = c.Cli.Notify(ctx, "chat.1.NotifyChat.ChatInboxSynced", []interface{}{__arg})
	return
}
