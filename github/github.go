package github

import "github.com/crosbymichael/octokat"

type GitHub struct {
	AuthToken string
	User      string
}

func (g GitHub) Client() *octokat.Client {
	gh := octokat.NewClient()
	gh = gh.WithToken(g.AuthToken)
	return gh
}

func nameWithOwner(repo *octokat.Repository) octokat.Repo {
	return octokat.Repo{
		Name:     repo.Name,
		UserName: repo.Owner.Login,
	}
}

func bot(user octokat.User) bool {
	return user.Login == "GordonTheTurtle"
}
