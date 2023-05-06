package mysql

import (
	"bluebell_backend/models"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"strings"

	"go.uber.org/zap"
)

// GetPostTotalCount 查询数据库帖子总数
func GetPostTotalCount() (count int64, err error) {
	sqlStr := `select count(post_id) from post`
	err = db.Get(&count, sqlStr)
	if err != nil {
		zap.L().Error("db.Get(&count, sqlStr) failed", zap.Error(err))
		return 0, err
	}
	return
}

// GetCommunityPostTotalCount 根据社区id查询数据库帖子总数
func GetCommunityPostTotalCount(communityID uint64) (count int64, err error) {
	sqlStr := `select count(post_id) from post where community_id = ?`
	err = db.Get(&count, sqlStr, communityID)
	if err != nil {
		zap.L().Error("db.Get(&count, sqlStr) failed", zap.Error(err))
		return 0, err
	}
	return
}

// CreatePost 创建帖子
func CreatePost(post *models.Post) (err error) {

	sqlStr := `insert into post(
	post_id, title, content, author_id, community_id)
	values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, post.PostID, post.Title,
		post.Content, post.AuthorId, post.CommunityID)
	if err != nil {
		zap.L().Error("insert post failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return nil
}

// GetPostByID 根据Id查询帖子详情
func GetPostByID(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id, title, content, author_id, community_id, status, create_time, update_time
	from post
	where post_id = ?`
	err = db.Get(post, sqlStr, pid)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(ErrorInvalidID)
		}
		zap.L().Error("query post failed", zap.String("sql", sqlStr), zap.Error(err))
		return nil, errors.New(ErrorQueryFailed)
	}
	return
}

// GetPostListByIDs 根据给定的id列表查询帖子数据
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	where post_id in (?)
	order by FIND_IN_SET(post_id, ?)`
	// 动态填充id
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return
	}
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}

// GetPostList 获取帖子列表
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	ORDER BY create_time
	DESC 
	limit ?,?
	`
	posts = make([]*models.Post, 0, 2) // 0：长度  2：容量
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}

// GetPostListByKeywords 根据关键词查询帖子列表
func GetPostListByKeywords(p *models.ParamPostList) (posts []*models.Post, err error) {
	// 根据帖子标题或者帖子内容模糊查询帖子列表
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	where title like ?
	or content like ?
	ORDER BY create_time
	DESC
	limit ?,?
	`
	// %keyword%
	p.Search = "%" + p.Search + "%"
	posts = make([]*models.Post, 0, 2) // 0：长度  2：容量
	err = db.Select(&posts, sqlStr, p.Search, p.Search, (p.Page-1)*p.Size, p.Size)
	return
}

// GetPostListTotalCount 根据关键词查询帖子列表总数
func GetPostListTotalCount(p *models.ParamPostList) (count int64, err error) {
	// 根据帖子标题或者帖子内容模糊查询帖子列表总数
	sqlStr := `select count(post_id)
	from post
	where title like ?
	or content like ?
	`
	// %keyword%
	p.Search = "%" + p.Search + "%"
	err = db.Get(&count, sqlStr, p.Search, p.Search)
	return
}
