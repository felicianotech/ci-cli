How I image the `ci` command being used.

ci browse
	opens the CI page in a browser

	ci browse --ci
		opens the CI page in a browser

	ci browse --circleci
		opens the CircleCI page in a browser

	ci browse --travis
		opens the Travis CI page in a browser

	ci browse --vcs
		opens the GitHub/GitLab page in a browser

	ci browse --github
		opens the GitHub page in a browser

	ci browse --gitlab
		opens the GitLab page in a browser

ci orb
	a CircleCI specific command that deals with orbs

	ci orb diff
		prints a diff between two versions of an orb
		https://github.com/CircleCI-Public/circleci-cli/pull/320
	
	ci orb init
		creates a scaffold of files for a new orb
		https://github.com/CircleCI-Public/circleci-cli/issues/284

ci config
	a grouping of config related commands

	ci config validate
		validate (yaml and schema) a CI config

	ci config create
		create a new template for a CI config. Can optionally use a template file
		https://github.com/CircleCI-Public/circleci-cli/issues/189
