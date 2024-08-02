# go-repo-template

Repository [template](https://docs.github.com/en/repositories/creating-and-managing-repositories/creating-a-repository-from-a-template)
for new Go projects.

## Bootstrap

Click `Use this template` button:
![image](https://github.com/nobl9/go-repo-template/assets/48822818/a5edc131-00c8-46f5-8ae6-1b593cbb4714)

After you're done, you can run the following command to bootstrap the project
and make it ready to run:

```shell
export MODULE_NAME=<YOUR_NAME> # Example: replace <YOUR_NAME> with "sloctl".
grep -rl your-module-name | xargs sed -i "s/your-module-name/$MODULE_NAME/g"
rm -rf bootstrap
rm gitsync.json
```

In order for some automations to work, like
[Release Drafter](https://github.com/release-drafter/release-drafter),
we need a predefined set of labels.
If you create a new repository from this template, the labels will be
automatically transferred for you.
However, if you want to use these automations in an existing repository,
you'll need to create these labels.
We've provided a convenience script for that, located [here](./bootstrap/add-labels.sh).
Run the following:

```shell
./bootstrap/add-labels.sh <you-existing-repo-name>
```

If you wish to update existing labels, add `--force` to the `gh label create`
invocation in the script.

## Project structure

The existing folders such as `cmd`, `pkg` and `internal` serve as examples on
how to structure your project.
You should remove them and change their contents to your liking, but adhere to
the structure defined
in [golang standards](https://github.com/golang-standards/project-layout).

## Makefile

Makefile is used as the utility
You can quickly inspect the targets of Makefile by running:

```shell
make help
```

When writing new targets, make sure you document them with double `#` character
and place the comment directly above the target, like so:

```makefile
## Document me!
new-target:
  echo "Hello"
```

Want to include PlantUML diagrams in your project?
You can add these targets to Makefile:

```makefile
PLANTUML_JAR_URL := https://sourceforge.net/projects/plantuml/files/plantuml.jar/download
PLANTUML_JAR :=  $(BIN_DIR)/plantuml.jar
DIAGRAMS_PATH ?= .

## Generate PNG diagrams from PlantUML files.
generate/plantuml: $(PLANTUML_JAR)
	for path in $$(find $(DIAGRAMS_PATH) -name "*.puml" -type f); do \
  		echo "Generating PNG file(s) for $$path"; \
		java -jar $(PLANTUML_JAR) -tpng $$path; \
  	done

# If the plantuml.jar file isn't already present, download it.
$(PLANTUML_JAR):
	echo "Downloading PlantUML JAR..."
	curl -sSfL $(PLANTUML_JAR_URL) -o $(PLANTUML_JAR)
```

## Releasing

If you wish to ship binaries for your project, we recommend using [Goreleaser](https://goreleaser.com/).
Sloctl repository has some good examples on how to define Goreleaser
[config file](https://github.com/nobl9/sloctl/blob/main/.goreleaser.yml) and also,
how to use the tool in a [GitHub action](https://github.com/nobl9/sloctl/blob/main/.github/workflows/release.yml).

Sloctl also has examples of publishing Docker images to DockerHub.

## Gitsync

This repository is also used as a staple/root for other repositories to follow.
This means things like linter configs or CI/CD workflows in these repositories
are supposed to be kept in sync with this repository (with some variations).

This is achieved with a tool called [gitsync](https://github.com/nieomylnieja/gitsync).
Configuration file for the tool is [gitsync.json](./gitsync.json).

In order to see the diff between managed repositories run:

```shell
gitsync -c gitsync.json diff
```

In order to sync the changes for managed repositores run:

```shell
gitsync -c gitsync.json sync
```
