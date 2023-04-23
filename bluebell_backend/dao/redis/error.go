package redis

import "errors"

var (
	ErrorVoteTimeExpire = errors.New("已过投票时间")
	ErrorVoted          = errors.New("已经投过票了")
	ErrVoteRepeated     = errors.New("不允许重复投票")
)
