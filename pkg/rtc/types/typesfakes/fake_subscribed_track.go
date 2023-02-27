// Code generated by counterfeiter. DO NOT EDIT.
package typesfakes

import (
	"sync"

	"github.com/livekit/livekit-server/pkg/rtc/types"
	"github.com/livekit/protocol/livekit"
	webrtc "github.com/pion/webrtc/v3"
)

type FakeSubscribedTrack struct {
	AddOnBindStub        func(func())
	addOnBindMutex       sync.RWMutex
	addOnBindArgsForCall []struct {
		arg1 func()
	}
	CloseStub        func(bool)
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
		arg1 bool
	}
	DownTrackStub        func() types.DownTrack
	downTrackMutex       sync.RWMutex
	downTrackArgsForCall []struct {
	}
	downTrackReturns struct {
		result1 types.DownTrack
	}
	downTrackReturnsOnCall map[int]struct {
		result1 types.DownTrack
	}
	IDStub        func() livekit.TrackID
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
	}
	iDReturns struct {
		result1 livekit.TrackID
	}
	iDReturnsOnCall map[int]struct {
		result1 livekit.TrackID
	}
	IsBoundStub        func() bool
	isBoundMutex       sync.RWMutex
	isBoundArgsForCall []struct {
	}
	isBoundReturns struct {
		result1 bool
	}
	isBoundReturnsOnCall map[int]struct {
		result1 bool
	}
	IsMutedStub        func() bool
	isMutedMutex       sync.RWMutex
	isMutedArgsForCall []struct {
	}
	isMutedReturns struct {
		result1 bool
	}
	isMutedReturnsOnCall map[int]struct {
		result1 bool
	}
	IsMuxedTrackStub        func() bool
	isMuxedTrackMutex       sync.RWMutex
	isMuxedTrackArgsForCall []struct {
	}
	isMuxedTrackReturns struct {
		result1 bool
	}
	isMuxedTrackReturnsOnCall map[int]struct {
		result1 bool
	}
	MediaTrackStub        func() types.MediaTrack
	mediaTrackMutex       sync.RWMutex
	mediaTrackArgsForCall []struct {
	}
	mediaTrackReturns struct {
		result1 types.MediaTrack
	}
	mediaTrackReturnsOnCall map[int]struct {
		result1 types.MediaTrack
	}
	NeedsNegotiationStub        func() bool
	needsNegotiationMutex       sync.RWMutex
	needsNegotiationArgsForCall []struct {
	}
	needsNegotiationReturns struct {
		result1 bool
	}
	needsNegotiationReturnsOnCall map[int]struct {
		result1 bool
	}
	OnCloseStub        func(func(willBeResumed bool))
	onCloseMutex       sync.RWMutex
	onCloseArgsForCall []struct {
		arg1 func(willBeResumed bool)
	}
	PublisherIDStub        func() livekit.ParticipantID
	publisherIDMutex       sync.RWMutex
	publisherIDArgsForCall []struct {
	}
	publisherIDReturns struct {
		result1 livekit.ParticipantID
	}
	publisherIDReturnsOnCall map[int]struct {
		result1 livekit.ParticipantID
	}
	PublisherIdentityStub        func() livekit.ParticipantIdentity
	publisherIdentityMutex       sync.RWMutex
	publisherIdentityArgsForCall []struct {
	}
	publisherIdentityReturns struct {
		result1 livekit.ParticipantIdentity
	}
	publisherIdentityReturnsOnCall map[int]struct {
		result1 livekit.ParticipantIdentity
	}
	PublisherVersionStub        func() uint32
	publisherVersionMutex       sync.RWMutex
	publisherVersionArgsForCall []struct {
	}
	publisherVersionReturns struct {
		result1 uint32
	}
	publisherVersionReturnsOnCall map[int]struct {
		result1 uint32
	}
	RTPSenderStub        func() *webrtc.RTPSender
	rTPSenderMutex       sync.RWMutex
	rTPSenderArgsForCall []struct {
	}
	rTPSenderReturns struct {
		result1 *webrtc.RTPSender
	}
	rTPSenderReturnsOnCall map[int]struct {
		result1 *webrtc.RTPSender
	}
	SetPublisherMutedStub        func(bool)
	setPublisherMutedMutex       sync.RWMutex
	setPublisherMutedArgsForCall []struct {
		arg1 bool
	}
	SubscriberStub        func() types.LocalParticipant
	subscriberMutex       sync.RWMutex
	subscriberArgsForCall []struct {
	}
	subscriberReturns struct {
		result1 types.LocalParticipant
	}
	subscriberReturnsOnCall map[int]struct {
		result1 types.LocalParticipant
	}
	SubscriberIDStub        func() livekit.ParticipantID
	subscriberIDMutex       sync.RWMutex
	subscriberIDArgsForCall []struct {
	}
	subscriberIDReturns struct {
		result1 livekit.ParticipantID
	}
	subscriberIDReturnsOnCall map[int]struct {
		result1 livekit.ParticipantID
	}
	SubscriberIdentityStub        func() livekit.ParticipantIdentity
	subscriberIdentityMutex       sync.RWMutex
	subscriberIdentityArgsForCall []struct {
	}
	subscriberIdentityReturns struct {
		result1 livekit.ParticipantIdentity
	}
	subscriberIdentityReturnsOnCall map[int]struct {
		result1 livekit.ParticipantIdentity
	}
	UpdateSubscriberSettingsStub        func(*livekit.UpdateTrackSettings)
	updateSubscriberSettingsMutex       sync.RWMutex
	updateSubscriberSettingsArgsForCall []struct {
		arg1 *livekit.UpdateTrackSettings
	}
	UpdateVideoLayerStub        func()
	updateVideoLayerMutex       sync.RWMutex
	updateVideoLayerArgsForCall []struct {
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSubscribedTrack) AddOnBind(arg1 func()) {
	fake.addOnBindMutex.Lock()
	fake.addOnBindArgsForCall = append(fake.addOnBindArgsForCall, struct {
		arg1 func()
	}{arg1})
	stub := fake.AddOnBindStub
	fake.recordInvocation("AddOnBind", []interface{}{arg1})
	fake.addOnBindMutex.Unlock()
	if stub != nil {
		fake.AddOnBindStub(arg1)
	}
}

func (fake *FakeSubscribedTrack) AddOnBindCallCount() int {
	fake.addOnBindMutex.RLock()
	defer fake.addOnBindMutex.RUnlock()
	return len(fake.addOnBindArgsForCall)
}

func (fake *FakeSubscribedTrack) AddOnBindCalls(stub func(func())) {
	fake.addOnBindMutex.Lock()
	defer fake.addOnBindMutex.Unlock()
	fake.AddOnBindStub = stub
}

func (fake *FakeSubscribedTrack) AddOnBindArgsForCall(i int) func() {
	fake.addOnBindMutex.RLock()
	defer fake.addOnBindMutex.RUnlock()
	argsForCall := fake.addOnBindArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSubscribedTrack) Close(arg1 bool) {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
		arg1 bool
	}{arg1})
	stub := fake.CloseStub
	fake.recordInvocation("Close", []interface{}{arg1})
	fake.closeMutex.Unlock()
	if stub != nil {
		fake.CloseStub(arg1)
	}
}

func (fake *FakeSubscribedTrack) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeSubscribedTrack) CloseCalls(stub func(bool)) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeSubscribedTrack) CloseArgsForCall(i int) bool {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	argsForCall := fake.closeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSubscribedTrack) DownTrack() types.DownTrack {
	fake.downTrackMutex.Lock()
	ret, specificReturn := fake.downTrackReturnsOnCall[len(fake.downTrackArgsForCall)]
	fake.downTrackArgsForCall = append(fake.downTrackArgsForCall, struct {
	}{})
	stub := fake.DownTrackStub
	fakeReturns := fake.downTrackReturns
	fake.recordInvocation("DownTrack", []interface{}{})
	fake.downTrackMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubscribedTrack) DownTrackCallCount() int {
	fake.downTrackMutex.RLock()
	defer fake.downTrackMutex.RUnlock()
	return len(fake.downTrackArgsForCall)
}

func (fake *FakeSubscribedTrack) DownTrackCalls(stub func() types.DownTrack) {
	fake.downTrackMutex.Lock()
	defer fake.downTrackMutex.Unlock()
	fake.DownTrackStub = stub
}

func (fake *FakeSubscribedTrack) DownTrackReturns(result1 types.DownTrack) {
	fake.downTrackMutex.Lock()
	defer fake.downTrackMutex.Unlock()
	fake.DownTrackStub = nil
	fake.downTrackReturns = struct {
		result1 types.DownTrack
	}{result1}
}

func (fake *FakeSubscribedTrack) DownTrackReturnsOnCall(i int, result1 types.DownTrack) {
	fake.downTrackMutex.Lock()
	defer fake.downTrackMutex.Unlock()
	fake.DownTrackStub = nil
	if fake.downTrackReturnsOnCall == nil {
		fake.downTrackReturnsOnCall = make(map[int]struct {
			result1 types.DownTrack
		})
	}
	fake.downTrackReturnsOnCall[i] = struct {
		result1 types.DownTrack
	}{result1}
}

func (fake *FakeSubscribedTrack) ID() livekit.TrackID {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct {
	}{})
	stub := fake.IDStub
	fakeReturns := fake.iDReturns
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubscribedTrack) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeSubscribedTrack) IDCalls(stub func() livekit.TrackID) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = stub
}

func (fake *FakeSubscribedTrack) IDReturns(result1 livekit.TrackID) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 livekit.TrackID
	}{result1}
}

func (fake *FakeSubscribedTrack) IDReturnsOnCall(i int, result1 livekit.TrackID) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 livekit.TrackID
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 livekit.TrackID
	}{result1}
}

func (fake *FakeSubscribedTrack) IsBound() bool {
	fake.isBoundMutex.Lock()
	ret, specificReturn := fake.isBoundReturnsOnCall[len(fake.isBoundArgsForCall)]
	fake.isBoundArgsForCall = append(fake.isBoundArgsForCall, struct {
	}{})
	stub := fake.IsBoundStub
	fakeReturns := fake.isBoundReturns
	fake.recordInvocation("IsBound", []interface{}{})
	fake.isBoundMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubscribedTrack) IsBoundCallCount() int {
	fake.isBoundMutex.RLock()
	defer fake.isBoundMutex.RUnlock()
	return len(fake.isBoundArgsForCall)
}

func (fake *FakeSubscribedTrack) IsBoundCalls(stub func() bool) {
	fake.isBoundMutex.Lock()
	defer fake.isBoundMutex.Unlock()
	fake.IsBoundStub = stub
}

func (fake *FakeSubscribedTrack) IsBoundReturns(result1 bool) {
	fake.isBoundMutex.Lock()
	defer fake.isBoundMutex.Unlock()
	fake.IsBoundStub = nil
	fake.isBoundReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeSubscribedTrack) IsBoundReturnsOnCall(i int, result1 bool) {
	fake.isBoundMutex.Lock()
	defer fake.isBoundMutex.Unlock()
	fake.IsBoundStub = nil
	if fake.isBoundReturnsOnCall == nil {
		fake.isBoundReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isBoundReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeSubscribedTrack) IsMuted() bool {
	fake.isMutedMutex.Lock()
	ret, specificReturn := fake.isMutedReturnsOnCall[len(fake.isMutedArgsForCall)]
	fake.isMutedArgsForCall = append(fake.isMutedArgsForCall, struct {
	}{})
	stub := fake.IsMutedStub
	fakeReturns := fake.isMutedReturns
	fake.recordInvocation("IsMuted", []interface{}{})
	fake.isMutedMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubscribedTrack) IsMutedCallCount() int {
	fake.isMutedMutex.RLock()
	defer fake.isMutedMutex.RUnlock()
	return len(fake.isMutedArgsForCall)
}

func (fake *FakeSubscribedTrack) IsMutedCalls(stub func() bool) {
	fake.isMutedMutex.Lock()
	defer fake.isMutedMutex.Unlock()
	fake.IsMutedStub = stub
}

func (fake *FakeSubscribedTrack) IsMutedReturns(result1 bool) {
	fake.isMutedMutex.Lock()
	defer fake.isMutedMutex.Unlock()
	fake.IsMutedStub = nil
	fake.isMutedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeSubscribedTrack) IsMutedReturnsOnCall(i int, result1 bool) {
	fake.isMutedMutex.Lock()
	defer fake.isMutedMutex.Unlock()
	fake.IsMutedStub = nil
	if fake.isMutedReturnsOnCall == nil {
		fake.isMutedReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isMutedReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeSubscribedTrack) IsMuxedTrack() bool {
	fake.isMuxedTrackMutex.Lock()
	ret, specificReturn := fake.isMuxedTrackReturnsOnCall[len(fake.isMuxedTrackArgsForCall)]
	fake.isMuxedTrackArgsForCall = append(fake.isMuxedTrackArgsForCall, struct {
	}{})
	stub := fake.IsMuxedTrackStub
	fakeReturns := fake.isMuxedTrackReturns
	fake.recordInvocation("IsMuxedTrack", []interface{}{})
	fake.isMuxedTrackMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubscribedTrack) IsMuxedTrackCallCount() int {
	fake.isMuxedTrackMutex.RLock()
	defer fake.isMuxedTrackMutex.RUnlock()
	return len(fake.isMuxedTrackArgsForCall)
}

func (fake *FakeSubscribedTrack) IsMuxedTrackCalls(stub func() bool) {
	fake.isMuxedTrackMutex.Lock()
	defer fake.isMuxedTrackMutex.Unlock()
	fake.IsMuxedTrackStub = stub
}

func (fake *FakeSubscribedTrack) IsMuxedTrackReturns(result1 bool) {
	fake.isMuxedTrackMutex.Lock()
	defer fake.isMuxedTrackMutex.Unlock()
	fake.IsMuxedTrackStub = nil
	fake.isMuxedTrackReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeSubscribedTrack) IsMuxedTrackReturnsOnCall(i int, result1 bool) {
	fake.isMuxedTrackMutex.Lock()
	defer fake.isMuxedTrackMutex.Unlock()
	fake.IsMuxedTrackStub = nil
	if fake.isMuxedTrackReturnsOnCall == nil {
		fake.isMuxedTrackReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isMuxedTrackReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeSubscribedTrack) MediaTrack() types.MediaTrack {
	fake.mediaTrackMutex.Lock()
	ret, specificReturn := fake.mediaTrackReturnsOnCall[len(fake.mediaTrackArgsForCall)]
	fake.mediaTrackArgsForCall = append(fake.mediaTrackArgsForCall, struct {
	}{})
	stub := fake.MediaTrackStub
	fakeReturns := fake.mediaTrackReturns
	fake.recordInvocation("MediaTrack", []interface{}{})
	fake.mediaTrackMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubscribedTrack) MediaTrackCallCount() int {
	fake.mediaTrackMutex.RLock()
	defer fake.mediaTrackMutex.RUnlock()
	return len(fake.mediaTrackArgsForCall)
}

func (fake *FakeSubscribedTrack) MediaTrackCalls(stub func() types.MediaTrack) {
	fake.mediaTrackMutex.Lock()
	defer fake.mediaTrackMutex.Unlock()
	fake.MediaTrackStub = stub
}

func (fake *FakeSubscribedTrack) MediaTrackReturns(result1 types.MediaTrack) {
	fake.mediaTrackMutex.Lock()
	defer fake.mediaTrackMutex.Unlock()
	fake.MediaTrackStub = nil
	fake.mediaTrackReturns = struct {
		result1 types.MediaTrack
	}{result1}
}

func (fake *FakeSubscribedTrack) MediaTrackReturnsOnCall(i int, result1 types.MediaTrack) {
	fake.mediaTrackMutex.Lock()
	defer fake.mediaTrackMutex.Unlock()
	fake.MediaTrackStub = nil
	if fake.mediaTrackReturnsOnCall == nil {
		fake.mediaTrackReturnsOnCall = make(map[int]struct {
			result1 types.MediaTrack
		})
	}
	fake.mediaTrackReturnsOnCall[i] = struct {
		result1 types.MediaTrack
	}{result1}
}

func (fake *FakeSubscribedTrack) NeedsNegotiation() bool {
	fake.needsNegotiationMutex.Lock()
	ret, specificReturn := fake.needsNegotiationReturnsOnCall[len(fake.needsNegotiationArgsForCall)]
	fake.needsNegotiationArgsForCall = append(fake.needsNegotiationArgsForCall, struct {
	}{})
	stub := fake.NeedsNegotiationStub
	fakeReturns := fake.needsNegotiationReturns
	fake.recordInvocation("NeedsNegotiation", []interface{}{})
	fake.needsNegotiationMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubscribedTrack) NeedsNegotiationCallCount() int {
	fake.needsNegotiationMutex.RLock()
	defer fake.needsNegotiationMutex.RUnlock()
	return len(fake.needsNegotiationArgsForCall)
}

func (fake *FakeSubscribedTrack) NeedsNegotiationCalls(stub func() bool) {
	fake.needsNegotiationMutex.Lock()
	defer fake.needsNegotiationMutex.Unlock()
	fake.NeedsNegotiationStub = stub
}

func (fake *FakeSubscribedTrack) NeedsNegotiationReturns(result1 bool) {
	fake.needsNegotiationMutex.Lock()
	defer fake.needsNegotiationMutex.Unlock()
	fake.NeedsNegotiationStub = nil
	fake.needsNegotiationReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeSubscribedTrack) NeedsNegotiationReturnsOnCall(i int, result1 bool) {
	fake.needsNegotiationMutex.Lock()
	defer fake.needsNegotiationMutex.Unlock()
	fake.NeedsNegotiationStub = nil
	if fake.needsNegotiationReturnsOnCall == nil {
		fake.needsNegotiationReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.needsNegotiationReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeSubscribedTrack) OnClose(arg1 func(willBeResumed bool)) {
	fake.onCloseMutex.Lock()
	fake.onCloseArgsForCall = append(fake.onCloseArgsForCall, struct {
		arg1 func(willBeResumed bool)
	}{arg1})
	stub := fake.OnCloseStub
	fake.recordInvocation("OnClose", []interface{}{arg1})
	fake.onCloseMutex.Unlock()
	if stub != nil {
		fake.OnCloseStub(arg1)
	}
}

func (fake *FakeSubscribedTrack) OnCloseCallCount() int {
	fake.onCloseMutex.RLock()
	defer fake.onCloseMutex.RUnlock()
	return len(fake.onCloseArgsForCall)
}

func (fake *FakeSubscribedTrack) OnCloseCalls(stub func(func(willBeResumed bool))) {
	fake.onCloseMutex.Lock()
	defer fake.onCloseMutex.Unlock()
	fake.OnCloseStub = stub
}

func (fake *FakeSubscribedTrack) OnCloseArgsForCall(i int) func(willBeResumed bool) {
	fake.onCloseMutex.RLock()
	defer fake.onCloseMutex.RUnlock()
	argsForCall := fake.onCloseArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSubscribedTrack) PublisherID() livekit.ParticipantID {
	fake.publisherIDMutex.Lock()
	ret, specificReturn := fake.publisherIDReturnsOnCall[len(fake.publisherIDArgsForCall)]
	fake.publisherIDArgsForCall = append(fake.publisherIDArgsForCall, struct {
	}{})
	stub := fake.PublisherIDStub
	fakeReturns := fake.publisherIDReturns
	fake.recordInvocation("PublisherID", []interface{}{})
	fake.publisherIDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubscribedTrack) PublisherIDCallCount() int {
	fake.publisherIDMutex.RLock()
	defer fake.publisherIDMutex.RUnlock()
	return len(fake.publisherIDArgsForCall)
}

func (fake *FakeSubscribedTrack) PublisherIDCalls(stub func() livekit.ParticipantID) {
	fake.publisherIDMutex.Lock()
	defer fake.publisherIDMutex.Unlock()
	fake.PublisherIDStub = stub
}

func (fake *FakeSubscribedTrack) PublisherIDReturns(result1 livekit.ParticipantID) {
	fake.publisherIDMutex.Lock()
	defer fake.publisherIDMutex.Unlock()
	fake.PublisherIDStub = nil
	fake.publisherIDReturns = struct {
		result1 livekit.ParticipantID
	}{result1}
}

func (fake *FakeSubscribedTrack) PublisherIDReturnsOnCall(i int, result1 livekit.ParticipantID) {
	fake.publisherIDMutex.Lock()
	defer fake.publisherIDMutex.Unlock()
	fake.PublisherIDStub = nil
	if fake.publisherIDReturnsOnCall == nil {
		fake.publisherIDReturnsOnCall = make(map[int]struct {
			result1 livekit.ParticipantID
		})
	}
	fake.publisherIDReturnsOnCall[i] = struct {
		result1 livekit.ParticipantID
	}{result1}
}

func (fake *FakeSubscribedTrack) PublisherIdentity() livekit.ParticipantIdentity {
	fake.publisherIdentityMutex.Lock()
	ret, specificReturn := fake.publisherIdentityReturnsOnCall[len(fake.publisherIdentityArgsForCall)]
	fake.publisherIdentityArgsForCall = append(fake.publisherIdentityArgsForCall, struct {
	}{})
	stub := fake.PublisherIdentityStub
	fakeReturns := fake.publisherIdentityReturns
	fake.recordInvocation("PublisherIdentity", []interface{}{})
	fake.publisherIdentityMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubscribedTrack) PublisherIdentityCallCount() int {
	fake.publisherIdentityMutex.RLock()
	defer fake.publisherIdentityMutex.RUnlock()
	return len(fake.publisherIdentityArgsForCall)
}

func (fake *FakeSubscribedTrack) PublisherIdentityCalls(stub func() livekit.ParticipantIdentity) {
	fake.publisherIdentityMutex.Lock()
	defer fake.publisherIdentityMutex.Unlock()
	fake.PublisherIdentityStub = stub
}

func (fake *FakeSubscribedTrack) PublisherIdentityReturns(result1 livekit.ParticipantIdentity) {
	fake.publisherIdentityMutex.Lock()
	defer fake.publisherIdentityMutex.Unlock()
	fake.PublisherIdentityStub = nil
	fake.publisherIdentityReturns = struct {
		result1 livekit.ParticipantIdentity
	}{result1}
}

func (fake *FakeSubscribedTrack) PublisherIdentityReturnsOnCall(i int, result1 livekit.ParticipantIdentity) {
	fake.publisherIdentityMutex.Lock()
	defer fake.publisherIdentityMutex.Unlock()
	fake.PublisherIdentityStub = nil
	if fake.publisherIdentityReturnsOnCall == nil {
		fake.publisherIdentityReturnsOnCall = make(map[int]struct {
			result1 livekit.ParticipantIdentity
		})
	}
	fake.publisherIdentityReturnsOnCall[i] = struct {
		result1 livekit.ParticipantIdentity
	}{result1}
}

func (fake *FakeSubscribedTrack) PublisherVersion() uint32 {
	fake.publisherVersionMutex.Lock()
	ret, specificReturn := fake.publisherVersionReturnsOnCall[len(fake.publisherVersionArgsForCall)]
	fake.publisherVersionArgsForCall = append(fake.publisherVersionArgsForCall, struct {
	}{})
	stub := fake.PublisherVersionStub
	fakeReturns := fake.publisherVersionReturns
	fake.recordInvocation("PublisherVersion", []interface{}{})
	fake.publisherVersionMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubscribedTrack) PublisherVersionCallCount() int {
	fake.publisherVersionMutex.RLock()
	defer fake.publisherVersionMutex.RUnlock()
	return len(fake.publisherVersionArgsForCall)
}

func (fake *FakeSubscribedTrack) PublisherVersionCalls(stub func() uint32) {
	fake.publisherVersionMutex.Lock()
	defer fake.publisherVersionMutex.Unlock()
	fake.PublisherVersionStub = stub
}

func (fake *FakeSubscribedTrack) PublisherVersionReturns(result1 uint32) {
	fake.publisherVersionMutex.Lock()
	defer fake.publisherVersionMutex.Unlock()
	fake.PublisherVersionStub = nil
	fake.publisherVersionReturns = struct {
		result1 uint32
	}{result1}
}

func (fake *FakeSubscribedTrack) PublisherVersionReturnsOnCall(i int, result1 uint32) {
	fake.publisherVersionMutex.Lock()
	defer fake.publisherVersionMutex.Unlock()
	fake.PublisherVersionStub = nil
	if fake.publisherVersionReturnsOnCall == nil {
		fake.publisherVersionReturnsOnCall = make(map[int]struct {
			result1 uint32
		})
	}
	fake.publisherVersionReturnsOnCall[i] = struct {
		result1 uint32
	}{result1}
}

func (fake *FakeSubscribedTrack) RTPSender() *webrtc.RTPSender {
	fake.rTPSenderMutex.Lock()
	ret, specificReturn := fake.rTPSenderReturnsOnCall[len(fake.rTPSenderArgsForCall)]
	fake.rTPSenderArgsForCall = append(fake.rTPSenderArgsForCall, struct {
	}{})
	stub := fake.RTPSenderStub
	fakeReturns := fake.rTPSenderReturns
	fake.recordInvocation("RTPSender", []interface{}{})
	fake.rTPSenderMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubscribedTrack) RTPSenderCallCount() int {
	fake.rTPSenderMutex.RLock()
	defer fake.rTPSenderMutex.RUnlock()
	return len(fake.rTPSenderArgsForCall)
}

func (fake *FakeSubscribedTrack) RTPSenderCalls(stub func() *webrtc.RTPSender) {
	fake.rTPSenderMutex.Lock()
	defer fake.rTPSenderMutex.Unlock()
	fake.RTPSenderStub = stub
}

func (fake *FakeSubscribedTrack) RTPSenderReturns(result1 *webrtc.RTPSender) {
	fake.rTPSenderMutex.Lock()
	defer fake.rTPSenderMutex.Unlock()
	fake.RTPSenderStub = nil
	fake.rTPSenderReturns = struct {
		result1 *webrtc.RTPSender
	}{result1}
}

func (fake *FakeSubscribedTrack) RTPSenderReturnsOnCall(i int, result1 *webrtc.RTPSender) {
	fake.rTPSenderMutex.Lock()
	defer fake.rTPSenderMutex.Unlock()
	fake.RTPSenderStub = nil
	if fake.rTPSenderReturnsOnCall == nil {
		fake.rTPSenderReturnsOnCall = make(map[int]struct {
			result1 *webrtc.RTPSender
		})
	}
	fake.rTPSenderReturnsOnCall[i] = struct {
		result1 *webrtc.RTPSender
	}{result1}
}

func (fake *FakeSubscribedTrack) SetPublisherMuted(arg1 bool) {
	fake.setPublisherMutedMutex.Lock()
	fake.setPublisherMutedArgsForCall = append(fake.setPublisherMutedArgsForCall, struct {
		arg1 bool
	}{arg1})
	stub := fake.SetPublisherMutedStub
	fake.recordInvocation("SetPublisherMuted", []interface{}{arg1})
	fake.setPublisherMutedMutex.Unlock()
	if stub != nil {
		fake.SetPublisherMutedStub(arg1)
	}
}

func (fake *FakeSubscribedTrack) SetPublisherMutedCallCount() int {
	fake.setPublisherMutedMutex.RLock()
	defer fake.setPublisherMutedMutex.RUnlock()
	return len(fake.setPublisherMutedArgsForCall)
}

func (fake *FakeSubscribedTrack) SetPublisherMutedCalls(stub func(bool)) {
	fake.setPublisherMutedMutex.Lock()
	defer fake.setPublisherMutedMutex.Unlock()
	fake.SetPublisherMutedStub = stub
}

func (fake *FakeSubscribedTrack) SetPublisherMutedArgsForCall(i int) bool {
	fake.setPublisherMutedMutex.RLock()
	defer fake.setPublisherMutedMutex.RUnlock()
	argsForCall := fake.setPublisherMutedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSubscribedTrack) Subscriber() types.LocalParticipant {
	fake.subscriberMutex.Lock()
	ret, specificReturn := fake.subscriberReturnsOnCall[len(fake.subscriberArgsForCall)]
	fake.subscriberArgsForCall = append(fake.subscriberArgsForCall, struct {
	}{})
	stub := fake.SubscriberStub
	fakeReturns := fake.subscriberReturns
	fake.recordInvocation("Subscriber", []interface{}{})
	fake.subscriberMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubscribedTrack) SubscriberCallCount() int {
	fake.subscriberMutex.RLock()
	defer fake.subscriberMutex.RUnlock()
	return len(fake.subscriberArgsForCall)
}

func (fake *FakeSubscribedTrack) SubscriberCalls(stub func() types.LocalParticipant) {
	fake.subscriberMutex.Lock()
	defer fake.subscriberMutex.Unlock()
	fake.SubscriberStub = stub
}

func (fake *FakeSubscribedTrack) SubscriberReturns(result1 types.LocalParticipant) {
	fake.subscriberMutex.Lock()
	defer fake.subscriberMutex.Unlock()
	fake.SubscriberStub = nil
	fake.subscriberReturns = struct {
		result1 types.LocalParticipant
	}{result1}
}

func (fake *FakeSubscribedTrack) SubscriberReturnsOnCall(i int, result1 types.LocalParticipant) {
	fake.subscriberMutex.Lock()
	defer fake.subscriberMutex.Unlock()
	fake.SubscriberStub = nil
	if fake.subscriberReturnsOnCall == nil {
		fake.subscriberReturnsOnCall = make(map[int]struct {
			result1 types.LocalParticipant
		})
	}
	fake.subscriberReturnsOnCall[i] = struct {
		result1 types.LocalParticipant
	}{result1}
}

func (fake *FakeSubscribedTrack) SubscriberID() livekit.ParticipantID {
	fake.subscriberIDMutex.Lock()
	ret, specificReturn := fake.subscriberIDReturnsOnCall[len(fake.subscriberIDArgsForCall)]
	fake.subscriberIDArgsForCall = append(fake.subscriberIDArgsForCall, struct {
	}{})
	stub := fake.SubscriberIDStub
	fakeReturns := fake.subscriberIDReturns
	fake.recordInvocation("SubscriberID", []interface{}{})
	fake.subscriberIDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubscribedTrack) SubscriberIDCallCount() int {
	fake.subscriberIDMutex.RLock()
	defer fake.subscriberIDMutex.RUnlock()
	return len(fake.subscriberIDArgsForCall)
}

func (fake *FakeSubscribedTrack) SubscriberIDCalls(stub func() livekit.ParticipantID) {
	fake.subscriberIDMutex.Lock()
	defer fake.subscriberIDMutex.Unlock()
	fake.SubscriberIDStub = stub
}

func (fake *FakeSubscribedTrack) SubscriberIDReturns(result1 livekit.ParticipantID) {
	fake.subscriberIDMutex.Lock()
	defer fake.subscriberIDMutex.Unlock()
	fake.SubscriberIDStub = nil
	fake.subscriberIDReturns = struct {
		result1 livekit.ParticipantID
	}{result1}
}

func (fake *FakeSubscribedTrack) SubscriberIDReturnsOnCall(i int, result1 livekit.ParticipantID) {
	fake.subscriberIDMutex.Lock()
	defer fake.subscriberIDMutex.Unlock()
	fake.SubscriberIDStub = nil
	if fake.subscriberIDReturnsOnCall == nil {
		fake.subscriberIDReturnsOnCall = make(map[int]struct {
			result1 livekit.ParticipantID
		})
	}
	fake.subscriberIDReturnsOnCall[i] = struct {
		result1 livekit.ParticipantID
	}{result1}
}

func (fake *FakeSubscribedTrack) SubscriberIdentity() livekit.ParticipantIdentity {
	fake.subscriberIdentityMutex.Lock()
	ret, specificReturn := fake.subscriberIdentityReturnsOnCall[len(fake.subscriberIdentityArgsForCall)]
	fake.subscriberIdentityArgsForCall = append(fake.subscriberIdentityArgsForCall, struct {
	}{})
	stub := fake.SubscriberIdentityStub
	fakeReturns := fake.subscriberIdentityReturns
	fake.recordInvocation("SubscriberIdentity", []interface{}{})
	fake.subscriberIdentityMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubscribedTrack) SubscriberIdentityCallCount() int {
	fake.subscriberIdentityMutex.RLock()
	defer fake.subscriberIdentityMutex.RUnlock()
	return len(fake.subscriberIdentityArgsForCall)
}

func (fake *FakeSubscribedTrack) SubscriberIdentityCalls(stub func() livekit.ParticipantIdentity) {
	fake.subscriberIdentityMutex.Lock()
	defer fake.subscriberIdentityMutex.Unlock()
	fake.SubscriberIdentityStub = stub
}

func (fake *FakeSubscribedTrack) SubscriberIdentityReturns(result1 livekit.ParticipantIdentity) {
	fake.subscriberIdentityMutex.Lock()
	defer fake.subscriberIdentityMutex.Unlock()
	fake.SubscriberIdentityStub = nil
	fake.subscriberIdentityReturns = struct {
		result1 livekit.ParticipantIdentity
	}{result1}
}

func (fake *FakeSubscribedTrack) SubscriberIdentityReturnsOnCall(i int, result1 livekit.ParticipantIdentity) {
	fake.subscriberIdentityMutex.Lock()
	defer fake.subscriberIdentityMutex.Unlock()
	fake.SubscriberIdentityStub = nil
	if fake.subscriberIdentityReturnsOnCall == nil {
		fake.subscriberIdentityReturnsOnCall = make(map[int]struct {
			result1 livekit.ParticipantIdentity
		})
	}
	fake.subscriberIdentityReturnsOnCall[i] = struct {
		result1 livekit.ParticipantIdentity
	}{result1}
}

func (fake *FakeSubscribedTrack) UpdateSubscriberSettings(arg1 *livekit.UpdateTrackSettings) {
	fake.updateSubscriberSettingsMutex.Lock()
	fake.updateSubscriberSettingsArgsForCall = append(fake.updateSubscriberSettingsArgsForCall, struct {
		arg1 *livekit.UpdateTrackSettings
	}{arg1})
	stub := fake.UpdateSubscriberSettingsStub
	fake.recordInvocation("UpdateSubscriberSettings", []interface{}{arg1})
	fake.updateSubscriberSettingsMutex.Unlock()
	if stub != nil {
		fake.UpdateSubscriberSettingsStub(arg1)
	}
}

func (fake *FakeSubscribedTrack) UpdateSubscriberSettingsCallCount() int {
	fake.updateSubscriberSettingsMutex.RLock()
	defer fake.updateSubscriberSettingsMutex.RUnlock()
	return len(fake.updateSubscriberSettingsArgsForCall)
}

func (fake *FakeSubscribedTrack) UpdateSubscriberSettingsCalls(stub func(*livekit.UpdateTrackSettings)) {
	fake.updateSubscriberSettingsMutex.Lock()
	defer fake.updateSubscriberSettingsMutex.Unlock()
	fake.UpdateSubscriberSettingsStub = stub
}

func (fake *FakeSubscribedTrack) UpdateSubscriberSettingsArgsForCall(i int) *livekit.UpdateTrackSettings {
	fake.updateSubscriberSettingsMutex.RLock()
	defer fake.updateSubscriberSettingsMutex.RUnlock()
	argsForCall := fake.updateSubscriberSettingsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSubscribedTrack) UpdateVideoLayer() {
	fake.updateVideoLayerMutex.Lock()
	fake.updateVideoLayerArgsForCall = append(fake.updateVideoLayerArgsForCall, struct {
	}{})
	stub := fake.UpdateVideoLayerStub
	fake.recordInvocation("UpdateVideoLayer", []interface{}{})
	fake.updateVideoLayerMutex.Unlock()
	if stub != nil {
		fake.UpdateVideoLayerStub()
	}
}

func (fake *FakeSubscribedTrack) UpdateVideoLayerCallCount() int {
	fake.updateVideoLayerMutex.RLock()
	defer fake.updateVideoLayerMutex.RUnlock()
	return len(fake.updateVideoLayerArgsForCall)
}

func (fake *FakeSubscribedTrack) UpdateVideoLayerCalls(stub func()) {
	fake.updateVideoLayerMutex.Lock()
	defer fake.updateVideoLayerMutex.Unlock()
	fake.UpdateVideoLayerStub = stub
}

func (fake *FakeSubscribedTrack) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addOnBindMutex.RLock()
	defer fake.addOnBindMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.downTrackMutex.RLock()
	defer fake.downTrackMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.isBoundMutex.RLock()
	defer fake.isBoundMutex.RUnlock()
	fake.isMutedMutex.RLock()
	defer fake.isMutedMutex.RUnlock()
	fake.isMuxedTrackMutex.RLock()
	defer fake.isMuxedTrackMutex.RUnlock()
	fake.mediaTrackMutex.RLock()
	defer fake.mediaTrackMutex.RUnlock()
	fake.needsNegotiationMutex.RLock()
	defer fake.needsNegotiationMutex.RUnlock()
	fake.onCloseMutex.RLock()
	defer fake.onCloseMutex.RUnlock()
	fake.publisherIDMutex.RLock()
	defer fake.publisherIDMutex.RUnlock()
	fake.publisherIdentityMutex.RLock()
	defer fake.publisherIdentityMutex.RUnlock()
	fake.publisherVersionMutex.RLock()
	defer fake.publisherVersionMutex.RUnlock()
	fake.rTPSenderMutex.RLock()
	defer fake.rTPSenderMutex.RUnlock()
	fake.setPublisherMutedMutex.RLock()
	defer fake.setPublisherMutedMutex.RUnlock()
	fake.subscriberMutex.RLock()
	defer fake.subscriberMutex.RUnlock()
	fake.subscriberIDMutex.RLock()
	defer fake.subscriberIDMutex.RUnlock()
	fake.subscriberIdentityMutex.RLock()
	defer fake.subscriberIdentityMutex.RUnlock()
	fake.updateSubscriberSettingsMutex.RLock()
	defer fake.updateSubscriberSettingsMutex.RUnlock()
	fake.updateVideoLayerMutex.RLock()
	defer fake.updateVideoLayerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSubscribedTrack) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ types.SubscribedTrack = new(FakeSubscribedTrack)
