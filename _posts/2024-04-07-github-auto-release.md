---
title: "Create Auto Github Release"
date: 2024-04-07
categories: [Github]
tags: [Github, Release, Auto Release, Tag, Github Tag]
---

# Github Auto Release

## What is Release?

When developers work on software projects, they often reach points where they want to make a specific version of their software available to users or other developers and at that point they create a release of that specific software. GitHub Releases are a way to package and distribute software releases on the GitHub platform. 

### Creating a release involves several steps in sequence_
1. Create github tag. 
2. Create release of that specific tag. 
3. Fill the release information. 
4. Publish the release. 

`The most common way of publishing release is in X.Y.Z format.`

Here_
- X for Major Release. 
- Y for Minor Release. 
- Z for Patch Release. 

There is also release types like Pre-Release/Beta, Release Candidate(RC) or Alpha Release being used frequently. The choice of release type depends on factors such as the development process, project requirements, and the level of stability and maturity desired for the software. 

For now, we will just focus on `Major` and `Minor` Release type. 

Using Github Action we can automate the release process. Lets jump into the process. 

### Our Target
We are going to use two different pipelines for creating the release process. 
1. `Minor Update` - This will be automated. Once some feature branch is merged to the main/master, a auto minor-release update will held. 
2. `Major Update`- As Major update depends on several process, we don't want to update the major version frequently. Rather we will use manual button trigger option, where once when the developer team want, they can update the major version of the release by clicking the button from th pipeline. 

#### Pre-requisite
To automate the release-creation process, we need 2 values to be set manually. As we don't want to expose any sensitive information in the source-code, we will use github secrets. 

1. Go to the github repository. 
2. Go to Settings -> Secrets and Variables -> Actions.
3. Create two secrets using name `USER_EMAIL` and `USER_NAME`. 
4. Set the secret value accordingly. 

#### Pipeline 
- Create `.github` folder in the root directory of the project. 
- Create `workflows` folder in the `.github` folder.

#### Minor Release Creation Pipeline.
- Create a `auto-release.yaml` file inside `workflows` folder. 

In `auto-release.yaml` file paste the following code_

```
name: Release Workflow

on:
  push:
    branches:
      - main

jobs:
  release:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout Code
        uses: actions/checkout@v3

      - name: Set Git Config
        run: |
          git config --global user.email "${{ secrets.USER_EMAIL }}"
          git config --global user.name "${{ secrets.USER_NAME }}"
        shell: bash

      - name: Check for Existing Tags and Create
        id: check_tags
        run: |
          tags=$(git ls-remote --tags origin)
          if [ -z "$tags" ]; then
            echo "No tags found."
            git tag -a "v1.0" -m "Initial release"
            git push origin "v1.0"
            echo "::set-output name=new_tag::v1.0"
          else
            echo "Tags found:"
            remote_tags=$(git ls-remote --tags origin | awk '{print $2}' | awk -F '/' '{print $3}' | sort -V | tail -n 1)
            remote_tag="${remote_tags::-3}"
            echo $remote_tag
            echo "Existing tag is $remote_tag"

            # Internal Field Separator
            IFS='.' read -r major minor <<< "${remote_tag:1}"

            minor=$((minor + 1))
            new_tag_version="v$major.$minor"
            echo $new_tag_version

            git tag -a "${new_tag_version}" -m "Release -${new_tag_version}"
            git push origin "${new_tag_version}"

            echo "::set-output name=new_tag::${new_tag_version}"
          fi
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}

      - name: Create Release
        uses: actions/create-release@v1
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        with:
          tag_name: ${{ steps.check_tags.outputs.new_tag }}
          release_name: Release ${{ steps.check_tags.outputs.new_tag }}
          body: |
            Get the Release Details__
```

##### How it works?
When a branch is merged to the main branch, a new minor release update will take place. 




#### Major Release Creation Pipeline.

- Create another `update-major.yaml` file inside `workflows` folder. 

In `update-major.yaml` file paste the following code_

```
name: Update Major Release

on:
  workflow_dispatch:

jobs:
  release:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout Code
        uses: actions/checkout@v3
      
      - name: Exit if the branch is not main
        run: |
          if [[ "${{ github.ref }}" != "refs/heads/main" ]]; then
            echo "Branch is not main, exiting."
            exit 1
          fi

      - name: Set Git Config
        run: |
          git config --global user.email "${{ secrets.USER_EMAIL }}"
          git config --global user.name "${{ secrets.USER_NAME }}"
        shell: bash

      - name: Check for Existing Tags and Create
        id: check_tags
        run: |
          tags=$(git ls-remote --tags origin)
          if [ -z "$tags" ]; then
            echo "No tags found."
            git tag -a "v1.0" -m "Initial release"
            git push origin "v1.0"
            echo "::set-output name=new_tag::v1.0"
          else
            echo "Tags found:"
            remote_tags=$(git ls-remote --tags origin | awk '{print $2}' | awk -F '/' '{print $3}' | sort -V | tail -n 1)
            remote_tag="${remote_tags::-3}"
            echo $remote_tag
            echo "Existing tag is $remote_tag"

            # Internal Field Separator
            IFS='.' read -r major minor <<< "${remote_tag:1}"

            major=$((major + 1))
            minor=0
            new_tag_version="v$major.$minor"
            echo $new_tag_version

            git tag -a "${new_tag_version}" -m "Release -${new_tag_version}"
            git push origin "${new_tag_version}"

            echo "::set-output name=new_tag::${new_tag_version}"
          fi
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}

      - name: Create Release
        uses: actions/create-release@v1
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        with:
          tag_name: ${{ steps.check_tags.outputs.new_tag }}
          release_name: Release ${{ steps.check_tags.outputs.new_tag }}
          body: |
            Get the Release Details__
```


##### How it works?
If a major update is needed_
- Navigate to the `Actions` button from the project repository. 
- From the left panel, Select `Update Major Release` button. 
- From the workflow right side, select `Run Workflow` button, use `main` branch and hit `Run Workdlow` button again. 
- A major update will occur. 





