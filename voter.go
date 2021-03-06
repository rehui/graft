package graft

type Voter struct {
	CandidateServer
}

func (server *Server) ReceiveRequestVote(message RequestVoteMessage) (VoteResponseMessage, error) {
	defer server.Persist()
	if server.Term < message.Term && server.logUpToDate(message) {
		server.stepDown()
		server.Term = message.Term
		server.VotedFor = message.CandidateId
		server.ElectionTimer.Reset()

		return VoteResponseMessage{
			Term:        server.Term,
			VoteGranted: true,
		}, nil
	} else {
		return VoteResponseMessage{
			Term:        server.Term,
			VoteGranted: false,
		}, nil
	}
}
