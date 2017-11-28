package job

import (
	"satis-go/satis/satisphp/api"
	"satis-go/satis/satisphp/db"
)

// Add or save a repo tp the repo collection
func NewSaveRepoJob(dbPath string, repo api.Repo) *SaveRepoJob {
	return &SaveRepoJob{
		dbPath:     dbPath,
		repository: repo,
		exitChan:   make(chan error, 1), //给channel增加缓冲区 1个 err 长度，然后在程序的最后让主线程休眠一秒
	}
	//或是:把ch<-1这一行代码放到子线程代码的后面
}

type SaveRepoJob struct {
	dbPath     string
	repository api.Repo
	exitChan   chan error
}

func (j SaveRepoJob) ExitChan() chan error {
	return j.exitChan
}
func (j SaveRepoJob) Run() error {
	dbMgr := db.SatisDbManager{Path: j.dbPath}

	if err := dbMgr.Load(); err != nil {
		return err
	}
	repos, err := j.doSave(j.repository, dbMgr.Db.Repositories)
	if err != nil {
		return err
	}
	dbMgr.Db.Repositories = repos

	if err := dbMgr.Write(); err != nil {
		return err
	}
	return nil
}
func (j SaveRepoJob) doSave(repo api.Repo, repos []db.SatisRepository) ([]db.SatisRepository, error) {
	repoEntity := db.SatisRepository{Type: repo.Type, Url: repo.Url}
	found := false
	for i, r := range repos {
		tmp := api.NewRepo(r.Type, r.Url)
		if tmp.Id == repo.Id {
			repos[i] = repoEntity
			found = true
		}
	}
	if !found {
		return append(repos, repoEntity), nil
	}

	return repos, nil
}
