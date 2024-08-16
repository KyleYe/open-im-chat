package chat

import (
	"context"

	"github.com/KyleYe/open-im-chat/pkg/common/mctx"
	"github.com/KyleYe/open-im-chat/pkg/protocol/chat"
)

func (o *chatSvr) GetTokenForVideoMeeting(ctx context.Context, req *chat.GetTokenForVideoMeetingReq) (*chat.GetTokenForVideoMeetingResp, error) {
	if _, _, err := mctx.Check(ctx); err != nil {
		return nil, err
	}
	token, err := o.Livekit.GetLiveKitToken(req.Room, req.Identity)
	if err != nil {
		return nil, err
	}
	return &chat.GetTokenForVideoMeetingResp{
		ServerUrl: o.Livekit.GetLiveKitURL(),
		Token:     token,
	}, err
}
