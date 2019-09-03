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

ci projects list
	list all projects on the CI provider that you follow
	
ci project
	a group of commands related to an individual project
	
	ci project build
		kick off a new build for the current branch
		--retry will rebuild the last build on the current branch
	
	ci project add-ssh
		add a local SSH key to the project's SSH settings
	
	ci project add
		follow/add this project on your CI provider
	
	ci project remove
		unfollow/remove this project on your CI provider
		For some providers, this jus removes you from it. The project may still build for others
