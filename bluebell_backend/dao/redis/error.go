package redis

import "errors"

var (
	ErrorVoteTimeExpire = errors.New("已过投票时间")
	ErrorVoted          = errors.New("已经投过票了")
	ErrVoteRepested     = errors.New("投票时间已过")
)
