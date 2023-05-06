package logic

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/dao/redis"
	"bluebell_backend/models"
	"bluebell_backend/pkg/snowflake"
	"fmt"
	"strconv"

	"go.uber.org/zap"
)

// CreatePost 创建帖子
func CreatePost(post *models.Post) (err error) {
	// 1、 生成post_id(生成帖子ID)
	postID, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("snowflake.GetID() failed", zap.Error(err))
		return
	}
	post.PostID = postID
	// 2、创建帖子 保存到数据库
	if err := mysql.CreatePost(post); err != nil {
		zap.L().Error("mysql.CreatePost(&post) failed", zap.Error(err))
		return err
	}
	community, err := mysql.GetCommunityNameByID(fmt.Sprint(post.CommunityID))
	if err != nil {
		zap.L().Error("mysql.GetCommunityNameByID failed", zap.Error(err))
		return err
	}
	// redis存储帖子信息
	if err := redis.CreatePost(
		post.PostID,
		post.AuthorId,
		post.Title,
		TruncateByWords(post.Content, 120),
		community.CommunityID); err != nil {
		zap.L().Error("redis.CreatePost failed", zap.Error(err))
		return err
	}
	return

}

// GetPostById 根据Id查询帖子详情
func GetPostById(postID int64) (data *models.ApiPostDetail, err error) {
	// 查询并组合我们接口想用的数据
	// 查询帖子信息
	post, err := mysql.GetPostByID(postID)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(postID) failed",
			zap.Int64("postID", postID),
			zap.Error(err))
		return nil, err
	}
	// 根据作者id查询作者信息
	user, err := mysql.GetUserByID(post.AuthorId)
	if err != nil {
		zap.L().Error("mysql.GetUserByID() failed",
			zap.Uint64("postID", post.AuthorId),
			zap.Error(err))
		return
	}
	// 根据社区id查询社区详细信息
	community, err := mysql.GetCommunityByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityByID() failed",
			zap.Uint64("community_id", post.CommunityID),
			zap.Error(err))
		return
	}
	// 根据帖子id查询帖子的投票数
	voteNum, err := redis.GetPostVoteNum(postID)

	// 接口数据拼接
	data = &models.ApiPostDetail{
		Post:               post,
		CommunityDetailRes: community,
		AuthorName:         user.UserName,
		VoteNum:            voteNum,
	}
	return data, nil
}

// GetPostList 获取帖子列表
func GetPostList(page, size int64) ([]*models.ApiPostDetail, error) {
	postList, err := mysql.GetPostList(page, size)
	if err != nil {
		zap.L().Error("mysql.GetPostList() failed")
		return nil, err
	}
	data := make([]*models.ApiPostDetail, 0, len(postList)) // data 初始化
	for _, post := range postList {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserByID(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserByID() failed",
				zap.Uint64("postID", post.AuthorId),
				zap.Error(err))
			continue
		}
		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityByID() failed",
				zap.Uint64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}
		// 接口数据拼接
		postDetail := &models.ApiPostDetail{
			Post:               post,
			CommunityDetailRes: community,
			AuthorName:         user.UserName,
		}
		data = append(data, postDetail)
	}
	return data, nil
}

// GetPostList2 升级版帖子列表接口：按 创建时间 或者 分数排序
func GetPostList2(p *models.ParamPostList) (*models.ApiPostDetailRes, error) {
	var res models.ApiPostDetailRes
	// 从mysql获取帖子列表总数
	total, err := mysql.GetPostTotalCount()
	if err != nil {
		return nil, err
	}
	res.Page.Total = total
	// 1、根据参数中的排序规则去redis查询id列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data")
		return &res, nil
	}
	zap.L().Debug("GetPostList2", zap.Any("ids: ", ids))
	// 2、提前查询好每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return nil, err
	}

	// 3、根据id去数据库查询帖子详细信息
	// 返回的数据还要按照我给定的id的顺序返回  order by FIND_IN_SET(post_id, ?)
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return nil, err
	}
	res.Page.Page = p.Page
	res.Page.Size = p.Size
	res.List = make([]*models.ApiPostDetail, 0, len(posts))
	// 4、组合数据
	// 将帖子的作者及分区信息查询出来填充到帖子中
	for idx, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserByID(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserByID() failed",
				zap.Uint64("postID", post.AuthorId),
				zap.Error(err))
			user = nil
		}
		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityByID() failed",
				zap.Uint64("community_id", post.CommunityID),
				zap.Error(err))
			community = nil
		}
		// 接口数据拼接
		postDetail := &models.ApiPostDetail{
			VoteNum:            voteData[idx],
			Post:               post,
			CommunityDetailRes: community,
			AuthorName:         user.UserName,
		}
		res.List = append(res.List, postDetail)
	}
	return &res, nil
}

// GetCommunityPostList 根据社区id去查询帖子列表
func GetCommunityPostList(p *models.ParamPostList) (*models.ApiPostDetailRes, error) {
	var res models.ApiPostDetailRes
	// 从mysql获取该社区下帖子列表总数
	total, err := mysql.GetCommunityPostTotalCount(p.CommunityID)
	if err != nil {
		return nil, err
	}
	res.Page.Total = total
	// 1、根据参数中的排序规则去redis查询id列表
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetCommunityPostList(p) return 0 data")
		return &res, nil
	}
	zap.L().Debug("GetPostList2", zap.Any("ids", ids))
	// 2、提前查询好每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return nil, err
	}
	// 3、根据id去数据库查询帖子详细信息
	// 返回的数据还要按照我给定的id的顺序返回  order by FIND_IN_SET(post_id, ?)
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return nil, err
	}
	res.Page.Page = p.Page
	res.Page.Size = p.Size
	res.List = make([]*models.ApiPostDetail, 0, len(posts))
	// 4、根据社区id查询社区详细信息
	// 为了减少数据库的查询次数，这里将社区信息提前查询出来
	community, err := mysql.GetCommunityByID(p.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityByID() failed",
			zap.Uint64("community_id", p.CommunityID),
			zap.Error(err))
		community = nil
	}
	for idx, post := range posts {
		// 过滤掉不属于该社区的帖子
		if post.CommunityID != p.CommunityID {
			continue
		}
		// 根据作者id查询作者信息
		user, err := mysql.GetUserByID(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserByID() failed",
				zap.Uint64("postID", post.AuthorId),
				zap.Error(err))
			user = nil
		}
		// 接口数据拼接
		postDetail := &models.ApiPostDetail{
			VoteNum:            voteData[idx],
			Post:               post,
			CommunityDetailRes: community,
			AuthorName:         user.UserName,
		}
		res.List = append(res.List, postDetail)
	}
	return &res, nil
}

// GetPostListNew 将两个查询帖子列表逻辑合二为一的函数
func GetPostListNew(p *models.ParamPostList) (data *models.ApiPostDetailRes, err error) {
	// 根据请求参数的不同,执行不同的业务逻辑
	if p.CommunityID == 0 {
		// 查所有
		data, err = GetPostList2(p)
	} else {
		// 根据社区id查询
		data, err = GetCommunityPostList(p)
	}
	if err != nil {
		zap.L().Error("GetPostListNew failed", zap.Error(err))
		return nil, err
	}
	return data, nil
}

// PostSearch 搜索业务-搜索帖子
func PostSearch(p *models.ParamPostList) (*models.ApiPostDetailRes, error) {
	var res models.ApiPostDetailRes
	// 根据搜索条件去mysql查询符合条件的帖子列表总数
	total, err := mysql.GetPostListTotalCount(p)
	if err != nil {
		return nil, err
	}
	res.Page.Total = total
	// 1、根据搜索条件去mysql分页查询符合条件的帖子列表
	posts, err := mysql.GetPostListByKeywords(p)
	if err != nil {
		return nil, err
	}
	// 查询出来的帖子总数可能为0
	if len(posts) == 0 {
		return &models.ApiPostDetailRes{}, nil
	}
	// 2、查询出来的帖子id列表传入到redis接口获取帖子的投票数
	ids := make([]string, 0, len(posts))
	for _, post := range posts {
		ids = append(ids, strconv.Itoa(int(post.PostID)))
	}
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return nil, err
	}
	res.Page.Size = p.Size
	res.Page.Page = p.Page
	// 3、拼接数据
	res.List = make([]*models.ApiPostDetail, 0, len(posts))
	for idx, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserByID(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserByID() failed",
				zap.Uint64("postID", post.AuthorId),
				zap.Error(err))
			user = nil
		}
		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityByID() failed",
				zap.Uint64("community_id", post.CommunityID),
				zap.Error(err))
			community = nil
		}
		// 接口数据拼接
		postDetail := &models.ApiPostDetail{
			VoteNum:            voteData[idx],
			Post:               post,
			CommunityDetailRes: community,
			AuthorName:         user.UserName,
		}
		res.List = append(res.List, postDetail)
	}
	return &res, nil
}
