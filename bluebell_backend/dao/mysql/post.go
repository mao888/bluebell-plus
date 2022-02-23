package mysql

import (
	"bluebell_backend/models"
	"database/sql"
	"strings"

	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"
)

/**
 * @Author huchao
 * @Description //TODO 创建帖子
 * @Date 19:53 2022/2/12
 **/
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
	return
}

/**
 * @Author huchao
 * @Description //TODO 根据Id查询帖子详情
 * @Date 21:53 2022/2/12
 **/
func GetPostByID(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	where post_id = ?`
	err = db.Get(post, sqlStr, pid)
	if err == sql.ErrNoRows {
		err = ErrorInvalidID
		return
	}
	if err != nil {
		zap.L().Error("query post failed", zap.String("sql", sqlStr), zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	return
}

/**
 * @Author huchao
 * @Description //TODO 根据给定的id列表查询帖子数据
 * @Date 22:55 2022/2/15
 **/
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

/**
 * @Author huchao
 * @Description //TODO 获取帖子列表
 * @Date 22:58 2022/2/12
 **/
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	ORDER BY create_time
	DESC 
	limit ?,?
	`
	posts = make([]*models.Post, 0, 2)	// 0：长度  2：容量
	err = db.Select(&posts, sqlStr,(page-1)*size,size)
	return

}
