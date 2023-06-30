package logic

import (
	"bluebell_backend/dao/api"
	"bluebell_backend/models"
)

// GetGithubTrending 获取Github热榜项目
func GetGithubTrending(p *models.ParamGithubTrending) (data *models.GithubTrending, err error) {
	switch p.Language {
	case 0:
		data, err = api.GetGithubTrendingAll(p)
		//case 1:
		//	data, err = models.GetGithubTrendingGo(p.Since, p.Page, p.Size)
		//default:
		//	data, err = models.GetGithubTrendingAll(p.Since, p.Page, p.Size)
	}
	return
}
