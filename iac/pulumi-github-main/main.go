//revive:disable:package-comments,exported
package main

import (
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// GithubResources holds the created GitHub resources, making them accessible for testing and exporting.
type GithubResource struct {
	Repository *github.Repository
}

// defineInfrastructure defines the GitHub resources for the project.
// It is separated from main() to be independently testable.
func defineInfrastructure(ctx *pulumi.Context) (*GithubResource, error) {
	repositoryName := "strategy-optimizer"
	repositoryDescription := "A quantitative research lab for developing, backtesting, and optimizing trading strategies in Python for final deployment on MetaTrader 5 (MQL5)."
	repository, err := github.NewRepository(ctx, "newRepositoryStrategyOptimizer", &github.RepositoryArgs{
		DeleteBranchOnMerge: pulumi.Bool(true),
		Description:         pulumi.String(repositoryDescription),
		HasIssues:           pulumi.Bool(true),
		HasProjects:         pulumi.Bool(true),
		Name:                pulumi.String(repositoryName),
		Topics: pulumi.StringArray{
			pulumi.String("algorithmic-trading"),
			pulumi.String("backtesting"),
			pulumi.String("dagger"),
			pulumi.String("github"),
			pulumi.String("gitlab"),
			pulumi.String("golang"),
			pulumi.String("pulumi"),
			pulumi.String("python"),
			pulumi.String("trading-strategies"),
			pulumi.String("trading"),
			pulumi.String("vscode"),
		},
		Visibility: pulumi.String("public"),
		// VulnerabilityAlerts: pulumi.Bool(true),
	}, pulumi.Protect(false))
	if err != nil {
		return nil, err
	}

	_, err = github.NewBranchProtection(ctx, "branchProtection", &github.BranchProtectionArgs{
		RepositoryId:          repository.NodeId,
		Pattern:               pulumi.String("main"),
		RequiredLinearHistory: pulumi.Bool(true),
	}, pulumi.Protect(false))
	if err != nil {
		return nil, err
	}

	_, err = github.NewIssueLabel(ctx, "newIssueLabelGhActions", &github.IssueLabelArgs{
		Color:       pulumi.String("E66E01"),
		Description: pulumi.String("This issue is related to github-actions dependencies"),
		Name:        pulumi.String("github-actions dependencies"),
		Repository:  repository.Name,
	}, pulumi.Protect(false))
	if err != nil {
		return nil, err
	}

	_, err = github.NewIssueLabel(ctx, "newIssueLabelGoModules", &github.IssueLabelArgs{
		Color:       pulumi.String("9BE688"),
		Description: pulumi.String("This issue is related to go modules dependencies"),
		Name:        pulumi.String("go-modules dependencies"),
		Repository:  repository.Name,
	}, pulumi.Protect(false))
	if err != nil {
		return nil, err
	}

	_, err = github.NewActionsSecret(ctx, "newActionsSecretGLR", &github.ActionsSecretArgs{
		Repository: repository.Name,
		SecretName: pulumi.String("GITLAB_REPOSITORY"),
	}, pulumi.Parent(repository), pulumi.Protect(false))
	if err != nil {
		return nil, err
	}

	_, err = github.NewActionsSecret(ctx, "newActionsSecretGLT", &github.ActionsSecretArgs{
		Repository: repository.Name,
		SecretName: pulumi.String("GITLAB_TOKEN"),
	}, pulumi.Parent(repository), pulumi.Protect(false))
	if err != nil {
		return nil, err
	}

	_, err = github.NewActionsSecret(ctx, "newActionsSecretGLO", &github.ActionsSecretArgs{
		Repository: repository.Name,
		SecretName: pulumi.String("GITLAB_OWNER"),
	}, pulumi.Parent(repository), pulumi.Protect(false))
	if err != nil {
		return nil, err
	}

	return &GithubResource{
		Repository: repository,
	}, nil
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		resources, err := defineInfrastructure(ctx)
		if err != nil {
			return err
		}

		// Export outputs from the returned resources
		ctx.Export("repository", resources.Repository.Name)
		ctx.Export("repositoryUrl", resources.Repository.HtmlUrl)
		return nil
	})
}
