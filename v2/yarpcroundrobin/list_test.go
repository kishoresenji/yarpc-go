// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package yarpcroundrobin

// import (
// 	. "go.uber.org/yarpc/v2/yarpctest"
// )

// var (
// 	_noContextDeadlineError = yarpcerror.Newf(yarpcerror.CodeInvalidArgument, "can't wait for peer without a context deadline for a roundrobin peer list")
// )

// func newNotRunningError(err string) error {
// 	return yarpcerror.FailedPreconditionErrorf("roundrobin peer list is not running: %s", err)
// }

// func newUnavailableError(err error) error {
// 	return yarpcerror.UnavailableErrorf("roundrobin peer list timed out waiting for peer: %s", err.Error())
// }

// func TestRoundRobinList(t *testing.T) {
// 	type testStruct struct {
// 		msg string

// 		// PeerIDs that will be returned from the dialer's OnRetain with "Available" status
// 		retainedAvailablePeerIDs []string

// 		// PeerIDs that will be returned from the dialer's OnRetain with "Unavailable" status
// 		retainedUnavailablePeerIDs []string

// 		// PeerIDs that will be released from the dialer
// 		releasedPeerIDs []string

// 		// PeerIDs that will return "retainErr" from the dialer's OnRetain function
// 		errRetainedPeerIDs []string
// 		retainErr          error

// 		// PeerIDs that will return "releaseErr" from the dialer's OnRelease function
// 		errReleasedPeerIDs []string
// 		releaseErr         error

// 		// A list of actions that will be applied on the PeerList
// 		peerListActions []PeerListAction

// 		// PeerIDs expected to be in the PeerList's "Available" list after the actions have been applied
// 		expectedAvailablePeers []string

// 		// PeerIDs expected to be in the PeerList's "Unavailable" list after the actions have been applied
// 		expectedUnavailablePeers []string

// 		// PeerIDs expected to be in the PeerList's "Uninitialized" list after the actions have been applied
// 		expectedUninitializedPeers []string

// 		// Boolean indicating whether the PeerList is "running" after the actions have been applied
// 		expectedRunning bool

// 		// Boolean indicating whether peers should be shuffled
// 		shuffle bool
// 	}
// 	tests := []testStruct{
// 		{
// 			msg: "setup",
// 			retainedAvailablePeerIDs: []string{"1"},
// 			expectedAvailablePeers:   []string{"1"},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1"}},
// 			},
// 			expectedRunning: true,
// 		},
// 		{
// 			msg: "setup with disconnected",
// 			retainedAvailablePeerIDs:   []string{"1"},
// 			retainedUnavailablePeerIDs: []string{"2"},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1", "2"}},
// 			},
// 			expectedAvailablePeers:   []string{"1"},
// 			expectedUnavailablePeers: []string{"2"},
// 			expectedRunning:          true,
// 		},
// 		{
// 			msg: "start",
// 			retainedAvailablePeerIDs: []string{"1"},
// 			expectedAvailablePeers:   []string{"1"},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1"}},
// 				ChooseAction{
// 					ExpectedPeer: "1",
// 				},
// 			},
// 			expectedRunning: true,
// 		},
// 		{
// 			msg: "start stop",
// 			retainedAvailablePeerIDs:   []string{"1", "2", "3", "4", "5", "6"},
// 			retainedUnavailablePeerIDs: []string{"7", "8", "9"},
// 			releasedPeerIDs:            []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}},
// 				StopAction{},
// 				ChooseAction{
// 					ExpectedErr:         newUnavailableError(context.DeadlineExceeded),
// 					InputContextTimeout: 10 * time.Millisecond,
// 				},
// 			},
// 			expectedUninitializedPeers: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
// 			expectedRunning:            false,
// 		},
// 		{
// 			msg: "start many and choose",
// 			retainedAvailablePeerIDs: []string{"1", "2", "3", "4", "5", "6"},
// 			expectedAvailablePeers:   []string{"1", "2", "3", "4", "5", "6"},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1", "2", "3", "4", "5", "6"}},
// 				ChooseAction{ExpectedPeer: "1"},
// 				ChooseAction{ExpectedPeer: "2"},
// 				ChooseAction{ExpectedPeer: "3"},
// 				ChooseAction{ExpectedPeer: "4"},
// 				ChooseAction{ExpectedPeer: "5"},
// 				ChooseAction{ExpectedPeer: "6"},
// 				ChooseAction{ExpectedPeer: "1"},
// 			},
// 			expectedRunning: true,
// 		},
// 		{
// 			msg:                "update retain error",
// 			errRetainedPeerIDs: []string{"1"},
// 			retainErr:          yarpc.ErrInvalidPeerType{},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1"}, ExpectedErr: yarpc.ErrInvalidPeerType{}},
// 			},
// 			expectedRunning: true,
// 		},
// 		{
// 			msg: "update retain multiple errors",
// 			retainedAvailablePeerIDs: []string{"2"},
// 			errRetainedPeerIDs:       []string{"1", "3"},
// 			retainErr:                yarpc.ErrInvalidPeerType{},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{
// 					AddedPeerIDs: []string{"1", "2", "3"},
// 					ExpectedErr:  multierr.Combine(yarpc.ErrInvalidPeerType{}, yarpc.ErrInvalidPeerType{}),
// 				},
// 			},
// 			expectedAvailablePeers: []string{"2"},
// 			expectedRunning:        true,
// 		},
// 		{
// 			msg: "start stop release error",
// 			retainedAvailablePeerIDs: []string{"1"},
// 			errReleasedPeerIDs:       []string{"1"},
// 			releaseErr:               yarpc.ErrDialerHasNoReferenceToPeer{},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1"}},
// 				StopAction{
// 					ExpectedErr: yarpc.ErrDialerHasNoReferenceToPeer{},
// 				},
// 			},
// 			expectedUninitializedPeers: []string{"1"},
// 			expectedRunning:            false,
// 		},
// 		{
// 			msg: "start stop release multiple errors",
// 			retainedAvailablePeerIDs: []string{"1", "2", "3"},
// 			releasedPeerIDs:          []string{"2"},
// 			errReleasedPeerIDs:       []string{"1", "3"},
// 			releaseErr:               yarpc.ErrDialerHasNoReferenceToPeer{},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1", "2", "3"}},
// 				StopAction{
// 					ExpectedErr: multierr.Combine(
// 						yarpc.ErrDialerHasNoReferenceToPeer{},
// 						yarpc.ErrDialerHasNoReferenceToPeer{},
// 					),
// 				},
// 			},
// 			expectedUninitializedPeers: []string{"1", "2", "3"},
// 			expectedRunning:            false,
// 		},
// 		{
// 			msg: "choose before start",
// 			peerListActions: []PeerListAction{
// 				ChooseAction{
// 					ExpectedErr:         newNotRunningError("context finished while waiting for instance to start: context deadline exceeded"),
// 					InputContextTimeout: 10 * time.Millisecond,
// 				},
// 				ChooseAction{
// 					ExpectedErr:         newNotRunningError("context finished while waiting for instance to start: context deadline exceeded"),
// 					InputContextTimeout: 10 * time.Millisecond,
// 				},
// 			},
// 			expectedRunning: false,
// 		},
// 		{
// 			msg: "update before start",
// 			retainedAvailablePeerIDs: []string{"1"},
// 			expectedAvailablePeers:   []string{"1"},
// 			peerListActions: []PeerListAction{
// 				UpdateAction{AddedPeerIDs: []string{"1"}},
// 				StartAction{},
// 			},
// 			expectedRunning: true,
// 		},
// 		{
// 			msg: "start choose no peers",
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				ChooseAction{
// 					InputContextTimeout: 20 * time.Millisecond,
// 					ExpectedErr:         newUnavailableError(context.DeadlineExceeded),
// 				},
// 			},
// 			expectedRunning: true,
// 		},
// 		{
// 			msg: "start then add",
// 			retainedAvailablePeerIDs: []string{"1", "2"},
// 			expectedAvailablePeers:   []string{"1", "2"},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1"}},
// 				UpdateAction{AddedPeerIDs: []string{"2"}},
// 				ChooseAction{ExpectedPeer: "1"},
// 				ChooseAction{ExpectedPeer: "2"},
// 				ChooseAction{ExpectedPeer: "1"},
// 			},
// 			expectedRunning: true,
// 		},
// 		{
// 			msg: "start remove",
// 			retainedAvailablePeerIDs: []string{"1", "2"},
// 			expectedAvailablePeers:   []string{"2"},
// 			releasedPeerIDs:          []string{"1"},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1", "2"}},
// 				UpdateAction{RemovedPeerIDs: []string{"1"}},
// 				ChooseAction{ExpectedPeer: "2"},
// 			},
// 			expectedRunning: true,
// 		},
// 		{
// 			msg: "start add many and remove many",
// 			retainedAvailablePeerIDs: []string{"1", "2", "3-r", "4-r", "5-a-r", "6-a-r", "7-a", "8-a"},
// 			releasedPeerIDs:          []string{"3-r", "4-r", "5-a-r", "6-a-r"},
// 			expectedAvailablePeers:   []string{"1", "2", "7-a", "8-a"},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1", "2", "3-r", "4-r"}},
// 				UpdateAction{
// 					AddedPeerIDs: []string{"5-a-r", "6-a-r", "7-a", "8-a"},
// 				},
// 				UpdateAction{
// 					RemovedPeerIDs: []string{"5-a-r", "6-a-r", "3-r", "4-r"},
// 				},
// 				ChooseAction{ExpectedPeer: "1"},
// 				ChooseAction{ExpectedPeer: "2"},
// 				ChooseAction{ExpectedPeer: "7-a"},
// 				ChooseAction{ExpectedPeer: "8-a"},
// 				ChooseAction{ExpectedPeer: "1"},
// 			},
// 			expectedRunning: true,
// 		},
// 		{
// 			msg: "start add many and remove many with shuffle",
// 			retainedAvailablePeerIDs: []string{"1", "2", "3-r", "4-r", "5-a-r", "6-a-r", "7-a", "8-a"},
// 			releasedPeerIDs:          []string{"3-r", "4-r", "5-a-r", "6-a-r"},
// 			expectedAvailablePeers:   []string{"1", "2", "7-a", "8-a"},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1", "2", "3-r", "4-r"}},
// 				UpdateAction{
// 					AddedPeerIDs: []string{"5-a-r", "6-a-r", "7-a", "8-a"},
// 				},
// 				UpdateAction{
// 					RemovedPeerIDs: []string{"5-a-r", "6-a-r", "3-r", "4-r"},
// 				},
// 				ChooseAction{ExpectedPeer: "2"},
// 				ChooseAction{ExpectedPeer: "1"},
// 				ChooseAction{ExpectedPeer: "8-a"},
// 				ChooseAction{ExpectedPeer: "7-a"},
// 				ChooseAction{ExpectedPeer: "2"},
// 			},
// 			expectedRunning: true,
// 			shuffle:         true,
// 		},
// 		{
// 			msg: "add retain error",
// 			retainedAvailablePeerIDs: []string{"1", "2"},
// 			expectedAvailablePeers:   []string{"1", "2"},
// 			errRetainedPeerIDs:       []string{"3"},
// 			retainErr:                yarpc.ErrInvalidPeerType{},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1", "2"}},
// 				UpdateAction{
// 					AddedPeerIDs: []string{"3"},
// 					ExpectedErr:  yarpc.ErrInvalidPeerType{},
// 				},
// 				ChooseAction{ExpectedPeer: "1"},
// 				ChooseAction{ExpectedPeer: "2"},
// 				ChooseAction{ExpectedPeer: "1"},
// 			},
// 			expectedRunning: true,
// 		},
// 		{
// 			msg: "add duplicate peer",
// 			retainedAvailablePeerIDs: []string{"1", "2"},
// 			expectedAvailablePeers:   []string{"1", "2"},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1", "2"}},
// 				UpdateAction{
// 					AddedPeerIDs: []string{"2"},
// 					ExpectedErr:  yarpc.ErrPeerAddAlreadyInList("2"),
// 				},
// 				ChooseAction{ExpectedPeer: "1"},
// 				ChooseAction{ExpectedPeer: "2"},
// 				ChooseAction{ExpectedPeer: "1"},
// 			},
// 			expectedRunning: true,
// 		},
// 		{
// 			msg: "remove peer not in list",
// 			retainedAvailablePeerIDs: []string{"1", "2"},
// 			expectedAvailablePeers:   []string{"1", "2"},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1", "2"}},
// 				UpdateAction{
// 					RemovedPeerIDs: []string{"3"},
// 					ExpectedErr:    yarpc.ErrPeerRemoveNotInList("3"),
// 				},
// 				ChooseAction{ExpectedPeer: "1"},
// 				ChooseAction{ExpectedPeer: "2"},
// 				ChooseAction{ExpectedPeer: "1"},
// 			},
// 			expectedRunning: true,
// 		},
// 		{
// 			msg: "remove release error",
// 			retainedAvailablePeerIDs: []string{"1", "2"},
// 			errReleasedPeerIDs:       []string{"2"},
// 			releaseErr:               yarpc.ErrDialerHasNoReferenceToPeer{},
// 			expectedAvailablePeers:   []string{"1"},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1", "2"}},
// 				UpdateAction{
// 					RemovedPeerIDs: []string{"2"},
// 					ExpectedErr:    yarpc.ErrDialerHasNoReferenceToPeer{},
// 				},
// 				ChooseAction{ExpectedPeer: "1"},
// 				ChooseAction{ExpectedPeer: "1"},
// 			},
// 			expectedRunning: true,
// 		},
// 		// Flaky in CI
// 		// {
// 		// 	msg: "block until add",
// 		// 	retainedAvailablePeerIDs: []string{"1"},
// 		// 	expectedAvailablePeers:   []string{"1"},
// 		// 	peerListActions: []PeerListAction{
// 		// 		StartAction{},
// 		// 		ConcurrentAction{
// 		// 			Actions: []PeerListAction{
// 		// 				ChooseAction{
// 		// 					InputContextTimeout: 200 * time.Millisecond,
// 		// 					ExpectedPeer:        "1",
// 		// 				},
// 		// 				UpdateAction{AddedPeerIDs: []string{"1"}},
// 		// 			},
// 		// 			Wait: 20 * time.Millisecond,
// 		// 		},
// 		// 		ChooseAction{ExpectedPeer: "1"},
// 		// 	},
// 		// 	expectedRunning: true,
// 		// },
// 		// {
// 		// 	msg: "multiple blocking until add",
// 		// 	retainedAvailablePeerIDs: []string{"1"},
// 		// 	expectedAvailablePeers:   []string{"1"},
// 		// 	peerListActions: []PeerListAction{
// 		// 		StartAction{},
// 		// 		ConcurrentAction{
// 		// 			Actions: []PeerListAction{
// 		// 				ChooseAction{
// 		// 					InputContextTimeout: 200 * time.Millisecond,
// 		// 					ExpectedPeer:        "1",
// 		// 				},
// 		// 				ChooseAction{
// 		// 					InputContextTimeout: 200 * time.Millisecond,
// 		// 					ExpectedPeer:        "1",
// 		// 				},
// 		// 				ChooseAction{
// 		// 					InputContextTimeout: 200 * time.Millisecond,
// 		// 					ExpectedPeer:        "1",
// 		// 				},
// 		// 				UpdateAction{AddedPeerIDs: []string{"1"}},
// 		// 			},
// 		// 			Wait: 10 * time.Millisecond,
// 		// 		},
// 		// 		ChooseAction{ExpectedPeer: "1"},
// 		// 	},
// 		// 	expectedRunning: true,
// 		// },
// 		{
// 			msg: "block but added too late",
// 			retainedAvailablePeerIDs: []string{"1"},
// 			expectedAvailablePeers:   []string{"1"},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				ConcurrentAction{
// 					Actions: []PeerListAction{
// 						ChooseAction{
// 							InputContextTimeout: 10 * time.Millisecond,
// 							ExpectedErr:         newUnavailableError(context.DeadlineExceeded),
// 						},
// 						UpdateAction{AddedPeerIDs: []string{"1"}},
// 					},
// 					Wait: 20 * time.Millisecond,
// 				},
// 				ChooseAction{ExpectedPeer: "1"},
// 			},
// 			expectedRunning: true,
// 		},
// 		// Flaky in CI
// 		// {
// 		// 	msg: "block until new peer after removal of only peer",
// 		// 	retainedAvailablePeerIDs: []string{"1", "2"},
// 		// 	releasedPeerIDs:          []string{"1"},
// 		// 	expectedAvailablePeers:   []string{"2"},
// 		// 	peerListActions: []PeerListAction{
// 		// 		StartAction{},
// 		// 		UpdateAction{AddedPeerIDs: []string{"1"}},
// 		// 		UpdateAction{RemovedPeerIDs: []string{"1"}},
// 		// 		ConcurrentAction{
// 		// 			Actions: []PeerListAction{
// 		// 				ChooseAction{
// 		// 					InputContextTimeout: 200 * time.Millisecond,
// 		// 					ExpectedPeer:        "2",
// 		// 				},
// 		// 				UpdateAction{AddedPeerIDs: []string{"2"}},
// 		// 			},
// 		// 			Wait: 20 * time.Millisecond,
// 		// 		},
// 		// 		ChooseAction{ExpectedPeer: "2"},
// 		// 	},
// 		// 	expectedRunning: true,
// 		// },
// 		{
// 			msg: "no blocking with no context deadline",
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				ChooseAction{
// 					InputContext: context.Background(),
// 					ExpectedErr:  _noContextDeadlineError,
// 				},
// 			},
// 			expectedRunning: true,
// 		},
// 		{
// 			msg: "add unavailable peer",
// 			retainedAvailablePeerIDs:   []string{"1"},
// 			retainedUnavailablePeerIDs: []string{"2"},
// 			expectedAvailablePeers:     []string{"1"},
// 			expectedUnavailablePeers:   []string{"2"},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1"}},
// 				UpdateAction{AddedPeerIDs: []string{"2"}},
// 				ChooseAction{ExpectedPeer: "1"},
// 				ChooseAction{ExpectedPeer: "1"},
// 			},
// 			expectedRunning: true,
// 		},
// 		{
// 			msg: "remove unavailable peer",
// 			retainedUnavailablePeerIDs: []string{"1"},
// 			releasedPeerIDs:            []string{"1"},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1"}},
// 				UpdateAction{RemovedPeerIDs: []string{"1"}},
// 				ChooseAction{
// 					InputContextTimeout: 10 * time.Millisecond,
// 					ExpectedErr:         newUnavailableError(context.DeadlineExceeded),
// 				},
// 			},
// 			expectedRunning: true,
// 		},
// 		{
// 			msg: "notify peer is now available",
// 			retainedUnavailablePeerIDs: []string{"1"},
// 			expectedAvailablePeers:     []string{"1"},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1"}},
// 				ChooseAction{
// 					InputContextTimeout: 10 * time.Millisecond,
// 					ExpectedErr:         newUnavailableError(context.DeadlineExceeded),
// 				},
// 				NotifyStatusChangeAction{PeerID: "1", NewConnectionStatus: yarpc.Available},
// 				ChooseAction{ExpectedPeer: "1"},
// 			},
// 			expectedRunning: true,
// 		},
// 		{
// 			msg: "notify peer is still available",
// 			retainedAvailablePeerIDs: []string{"1"},
// 			expectedAvailablePeers:   []string{"1"},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1"}},
// 				ChooseAction{ExpectedPeer: "1"},
// 				NotifyStatusChangeAction{PeerID: "1", NewConnectionStatus: yarpc.Available},
// 				ChooseAction{ExpectedPeer: "1"},
// 			},
// 			expectedRunning: true,
// 		},
// 		{
// 			msg: "notify peer is now unavailable",
// 			retainedAvailablePeerIDs: []string{"1"},
// 			expectedUnavailablePeers: []string{"1"},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1"}},
// 				ChooseAction{ExpectedPeer: "1"},
// 				NotifyStatusChangeAction{PeerID: "1", NewConnectionStatus: yarpc.Unavailable},
// 				ChooseAction{
// 					InputContextTimeout: 10 * time.Millisecond,
// 					ExpectedErr:         newUnavailableError(context.DeadlineExceeded),
// 				},
// 			},
// 			expectedRunning: true,
// 		},
// 		{
// 			msg: "notify peer is still unavailable",
// 			retainedUnavailablePeerIDs: []string{"1"},
// 			expectedUnavailablePeers:   []string{"1"},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1"}},
// 				NotifyStatusChangeAction{PeerID: "1", NewConnectionStatus: yarpc.Unavailable},
// 				ChooseAction{
// 					InputContextTimeout: 10 * time.Millisecond,
// 					ExpectedErr:         newUnavailableError(context.DeadlineExceeded),
// 				},
// 			},
// 			expectedRunning: true,
// 		},
// 		{
// 			msg: "notify invalid peer",
// 			retainedAvailablePeerIDs: []string{"1"},
// 			releasedPeerIDs:          []string{"1"},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1"}},
// 				UpdateAction{RemovedPeerIDs: []string{"1"}},
// 				NotifyStatusChangeAction{PeerID: "1", NewConnectionStatus: yarpc.Available},
// 			},
// 			expectedRunning: true,
// 		},
// 		{
// 			// v: Available, u: Unavailable, a: Added, r: Removed
// 			msg: "notify peer stress test",
// 			retainedAvailablePeerIDs:   []string{"1v", "2va", "3vau", "4var", "5vaur"},
// 			retainedUnavailablePeerIDs: []string{"6u", "7ua", "8uav", "9uar", "10uavr"},
// 			releasedPeerIDs:            []string{"4var", "5vaur", "9uar", "10uavr"},
// 			expectedAvailablePeers:     []string{"1v", "2va", "8uav"},
// 			expectedUnavailablePeers:   []string{"3vau", "6u", "7ua"},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1v", "6u"}},

// 				// Added Peers
// 				UpdateAction{
// 					AddedPeerIDs: []string{"2va", "3vau", "4var", "5vaur", "7ua", "8uav", "9uar", "10uavr"},
// 				},

// 				ChooseMultiAction{ExpectedPeers: []string{"1v", "2va", "3vau", "4var", "5vaur"}},
// 				ChooseMultiAction{ExpectedPeers: []string{"1v", "2va", "3vau", "4var", "5vaur"}},

// 				// Change Status to Unavailable
// 				NotifyStatusChangeAction{PeerID: "3vau", NewConnectionStatus: yarpc.Unavailable},
// 				NotifyStatusChangeAction{PeerID: "5vaur", NewConnectionStatus: yarpc.Unavailable},

// 				ChooseMultiAction{ExpectedPeers: []string{"1v", "2va", "4var"}},
// 				ChooseMultiAction{ExpectedPeers: []string{"1v", "2va", "4var"}},

// 				// Change Status to Available
// 				NotifyStatusChangeAction{PeerID: "8uav", NewConnectionStatus: yarpc.Available},
// 				NotifyStatusChangeAction{PeerID: "10uavr", NewConnectionStatus: yarpc.Available},

// 				ChooseMultiAction{ExpectedPeers: []string{"1v", "2va", "4var", "8uav", "10uavr"}},
// 				ChooseMultiAction{ExpectedPeers: []string{"1v", "2va", "4var", "8uav", "10uavr"}},

// 				// Remove Peers
// 				UpdateAction{
// 					RemovedPeerIDs: []string{"4var", "5vaur", "9uar", "10uavr"},
// 				},

// 				ChooseMultiAction{ExpectedPeers: []string{"1v", "2va", "8uav"}},
// 				ChooseMultiAction{ExpectedPeers: []string{"1v", "2va", "8uav"}},
// 			},
// 			expectedRunning: true,
// 		},
// 		// Flaky in CI
// 		// {
// 		// 	msg: "block until notify available",
// 		// 	retainedUnavailablePeerIDs: []string{"1"},
// 		// 	expectedAvailablePeers:     []string{"1"},
// 		// 	peerListActions: []PeerListAction{
// 		// 		StartAction{},
// 		// 		UpdateAction{AddedPeerIDs: []string{"1"}},
// 		// 		ConcurrentAction{
// 		// 			Actions: []PeerListAction{
// 		// 				ChooseAction{
// 		// 					InputContextTimeout: 200 * time.Millisecond,
// 		// 					ExpectedPeer:        "1",
// 		// 				},
// 		// 				NotifyStatusChangeAction{PeerID: "1", NewConnectionStatus: yarpc.Available},
// 		// 			},
// 		// 			Wait: 20 * time.Millisecond,
// 		// 		},
// 		// 		ChooseAction{ExpectedPeer: "1"},
// 		// 	},
// 		// 	expectedRunning: true,
// 		// },
// 		{
// 			msg: "update no start",
// 			peerListActions: []PeerListAction{
// 				UpdateAction{AddedPeerIDs: []string{"1"}},
// 				UpdateAction{AddedPeerIDs: []string{"2"}},
// 				UpdateAction{RemovedPeerIDs: []string{"1"}},
// 			},
// 			expectedUninitializedPeers: []string{"2"},
// 			expectedRunning:            false,
// 		},
// 		{
// 			msg: "update after stop",
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				StopAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1"}},
// 				UpdateAction{AddedPeerIDs: []string{"2"}},
// 				UpdateAction{RemovedPeerIDs: []string{"1"}},
// 			},
// 			expectedUninitializedPeers: []string{"2"},
// 			expectedRunning:            false,
// 		},
// 		{
// 			msg: "remove peer not in list before start",
// 			peerListActions: []PeerListAction{
// 				UpdateAction{AddedPeerIDs: []string{"1", "2"}},
// 				UpdateAction{
// 					RemovedPeerIDs: []string{"3"},
// 					ExpectedErr:    yarpc.ErrPeerRemoveNotInList("3"),
// 				},
// 			},
// 			expectedUninitializedPeers: []string{"1", "2"},
// 			expectedRunning:            false,
// 		},
// 		{
// 			msg: "update before start",
// 			retainedAvailablePeerIDs: []string{"1", "2"},
// 			expectedAvailablePeers:   []string{"1", "2"},
// 			peerListActions: []PeerListAction{
// 				UpdateAction{AddedPeerIDs: []string{"1", "2"}},
// 				StartAction{},
// 			},
// 			expectedRunning: true,
// 		},
// 		{
// 			msg: "update before start, and stop",
// 			retainedAvailablePeerIDs: []string{"1", "2"},
// 			releasedPeerIDs:          []string{"1", "2"},
// 			peerListActions: []PeerListAction{
// 				UpdateAction{AddedPeerIDs: []string{"1", "2"}},
// 				StartAction{},
// 				StopAction{},
// 			},
// 			expectedUninitializedPeers: []string{"1", "2"},
// 			expectedRunning:            false,
// 		},
// 		{
// 			msg: "update before start, and update after stop",
// 			retainedAvailablePeerIDs: []string{"1", "2"},
// 			releasedPeerIDs:          []string{"1", "2"},
// 			peerListActions: []PeerListAction{
// 				UpdateAction{AddedPeerIDs: []string{"1", "2"}},
// 				StartAction{},
// 				StopAction{},
// 				UpdateAction{AddedPeerIDs: []string{"3"}, RemovedPeerIDs: []string{"1"}},
// 			},
// 			expectedUninitializedPeers: []string{"2", "3"},
// 			expectedRunning:            false,
// 		},
// 		{
// 			msg: "concurrent update and start",
// 			retainedAvailablePeerIDs: []string{"1", "2"},
// 			expectedAvailablePeers:   []string{"1", "2"},
// 			peerListActions: []PeerListAction{
// 				UpdateAction{AddedPeerIDs: []string{"1"}},
// 				ConcurrentAction{
// 					Actions: []PeerListAction{
// 						StartAction{},
// 						UpdateAction{AddedPeerIDs: []string{"2"}},
// 					},
// 				},
// 			},
// 			expectedRunning: true,
// 		},
// 		{
// 			msg: "concurrent update and stop",
// 			retainedAvailablePeerIDs: []string{"1", "2"},
// 			releasedPeerIDs:          []string{"1", "2"},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1", "2"}},
// 				ConcurrentAction{
// 					Actions: []PeerListAction{
// 						StopAction{},
// 						UpdateAction{RemovedPeerIDs: []string{"2"}},
// 						StopAction{},
// 					},
// 				},
// 			},
// 			expectedUninitializedPeers: []string{"1"},
// 			expectedRunning:            false,
// 		},
// 		{
// 			msg: "notify before start",
// 			peerListActions: []PeerListAction{
// 				UpdateAction{AddedPeerIDs: []string{"1", "2"}},
// 				NotifyStatusChangeAction{PeerID: "1", NewConnectionStatus: yarpc.Available, Unretained: true},
// 				NotifyStatusChangeAction{PeerID: "2", NewConnectionStatus: yarpc.Unavailable, Unretained: true},
// 			},
// 			expectedUninitializedPeers: []string{"1", "2"},
// 			expectedRunning:            false,
// 		},
// 		{
// 			msg: "notify after stop",
// 			retainedAvailablePeerIDs: []string{"1", "2"},
// 			releasedPeerIDs:          []string{"1", "2"},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1", "2"}},
// 				StopAction{},
// 				NotifyStatusChangeAction{PeerID: "1", NewConnectionStatus: yarpc.Available},
// 				NotifyStatusChangeAction{PeerID: "2", NewConnectionStatus: yarpc.Unavailable},
// 			},
// 			expectedUninitializedPeers: []string{"1", "2"},
// 			expectedRunning:            false,
// 		},
// 		{
// 			msg: "start with available and unavailable",
// 			retainedAvailablePeerIDs:   []string{"1", "2"},
// 			retainedUnavailablePeerIDs: []string{"3", "4"},
// 			expectedAvailablePeers:     []string{"1", "2"},
// 			expectedUnavailablePeers:   []string{"3", "4"},
// 			peerListActions: []PeerListAction{
// 				UpdateAction{AddedPeerIDs: []string{"1", "2", "3", "4"}},
// 				StartAction{},
// 			},
// 			expectedRunning: true,
// 		},
// 		{
// 			msg: "stop with available and unavailable",
// 			retainedAvailablePeerIDs:   []string{"1", "2"},
// 			retainedUnavailablePeerIDs: []string{"3", "4"},
// 			releasedPeerIDs:            []string{"1", "2", "3", "4"},
// 			peerListActions: []PeerListAction{
// 				StartAction{},
// 				UpdateAction{AddedPeerIDs: []string{"1", "2", "3", "4"}},
// 				StopAction{},
// 			},
// 			expectedUninitializedPeers: []string{"1", "2", "3", "4"},
// 			expectedRunning:            false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.msg, func(t *testing.T) {
// 			mockCtrl := gomock.NewController(t)
// 			defer mockCtrl.Finish()

// 			dialer := NewMockDialer(mockCtrl)

// 			// Healthy Dialer Retain/Release
// 			peerMap := ExpectPeerRetains(
// 				dialer,
// 				tt.retainedAvailablePeerIDs,
// 				tt.retainedUnavailablePeerIDs,
// 			)
// 			ExpectPeerReleases(dialer, tt.releasedPeerIDs, nil)

// 			// Unhealthy Dialer Retain/Release
// 			ExpectPeerRetainsWithError(dialer, tt.errRetainedPeerIDs, tt.retainErr)
// 			ExpectPeerReleases(dialer, tt.errReleasedPeerIDs, tt.releaseErr)

// 			opts := []ListOption{seed(0)}
// 			if !tt.shuffle {
// 				opts = append(opts, noShuffle)
// 			}
// 			pl := New(dialer, opts...)

// 			deps := ListActionDeps{
// 				Peers: peerMap,
// 			}
// 			ApplyPeerListActions(t, pl, tt.peerListActions, deps)

// 			assert.Equal(t, pl.NumAvailable(), len(tt.expectedAvailablePeers), "invalid available peerlist size")
// 			for _, expectedRingPeer := range tt.expectedAvailablePeers {
// 				ok := pl.Available(yarpc.Address(expectedRingPeer))
// 				assert.True(t, ok, fmt.Sprintf("expected peer: %s was not in available peerlist", expectedRingPeer))
// 			}

// 			assert.Equal(t, pl.NumUnavailable(), len(tt.expectedUnavailablePeers), "invalid unavailable peerlist size")
// 			for _, expectedUnavailablePeer := range tt.expectedUnavailablePeers {
// 				ok := !pl.Available(yarpc.Address(expectedUnavailablePeer))
// 				assert.True(t, ok, fmt.Sprintf("expected peer: %s was not in unavailable peerlist", expectedUnavailablePeer))
// 			}
// 		})
// 	}
// }

// type testDialer struct{}

// func (testDialer) RetainPeer(peerIdentifier yarpc.Identifier, _ yarpc.Subscriber) (yarpc.Peer, error) {
// 	return peerIdentifier.(*testPeer), nil
// }

// func (testDialer) ReleasePeer(yarpc.Identifier, yarpc.Subscriber) error {
// 	return nil
// }

// type testPeer struct {
// 	identifier          string
// 	pendingRequestCount int
// 	connectionStatus    yarpc.ConnectionStatus
// }

// func newTestPeer(
// 	identifier string,
// 	pendingRequestCount int,
// 	connectionStatus yarpc.ConnectionStatus,
// ) *testPeer {
// 	return &testPeer{
// 		identifier,
// 		pendingRequestCount,
// 		connectionStatus,
// 	}
// }

// func (p *testPeer) Identifier() string {
// 	return p.identifier
// }

// func (p *testPeer) Status() yarpc.Status {
// 	return yarpc.Status{
// 		PendingRequestCount: p.pendingRequestCount,
// 		ConnectionStatus:    p.connectionStatus,
// 	}
// }

// func (*testPeer) StartRequest() {}

// func (*testPeer) EndRequest() {}

// var noShuffle ListOption = func(c *listConfig) {
// 	c.shuffle = false
// }

// func seed(seed int64) ListOption {
// 	return func(c *listConfig) {
// 		c.seed = seed
// 	}
// }
