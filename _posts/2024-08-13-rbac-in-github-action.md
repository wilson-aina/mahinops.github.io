---
title: "RBAC in Github Action"
date: 2024-08-13
categories: [Github]
tags: [Github, Composite Action, RBAC, Github Permission, Workflow Permission, Github Action, CI/CD Pipeline, Pipeline, DevOps, SRE, CI/CD, Continuous Development, Continuous Integration]
---

# RBAC in Github Action

## Why? 
 RBAC (Role-Based Access Control) in GitHub Actions, when implemented using a permissions.yml file, allows you to define specific GitHub usernames that are authorized to trigger particular pipelines. By listing allowed users, this approach ensures that only those with explicit permission can execute the workflows. Applied as a composite action, this method provides granular control over who can initiate specific pipelines, enhancing security and maintaining strict access controls within your CI/CD processes.

## Why Composite Action? 
 This approach promotes reusability, maintainability, and consistency across different workflows. By encapsulating a series of steps into a composite action, you can reduce redundancy and ensure that common tasks are executed in a standardized manner. 

## How Can We Do That?
 To implement RBAC in Github Action, I have created a certain folder and files. The structure is given below_

```
.github/
├── actions
│   └── check-permissions
│       └── action.yml
└── workflows
|   └── test-workflow.yml
└── permissions.yml
```

In `action.yml`, I’ve defined a composite action that checks permissions to determine if a user is authorized to trigger the GitHub Action job. 

In `test-workflow.yml`, this composite action is used to enforce those permission checks before proceeding with the workflow.

The `permissions.yml` file contains the allower github usernames.

The file contains the code below to ensure the RBAC_

`action.yml` contains_
```yml
name: 'Check Permissions'
description: 'Check if the user triggering the workflow is allowed to proceed.'
inputs:
  permissions-file:
    description: 'Path to the permissions.yml file'
    required: true
  actor:
    description: 'GitHub actor triggering the workflow'
    required: true
runs:
  using: 'composite'
  steps:
  - name: Set up Python
    uses: actions/setup-python@v4

  - name: Install PyYAML
    run: pip install pyyaml
    shell: bash

  - name: Check Permissions
    run: |
      python -c "
      import yaml
      with open('${{ inputs.permissions-file }}', 'r') as f:
          permissions = yaml.safe_load(f)
      if '${{ inputs.actor }}' not in permissions['allowed_users']:
          print('-----------------ERROR--------------------')
          print('User ${{ inputs.actor }} is not allowed to trigger this workflow.')
          print('------------------------------------------')
          exit(1)
      "
    shell: bash
```

`test-workflow.yml` file contains_

```yml
name: Example Workflow
on:
  workflow_dispatch:
jobs:
  example-job:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Check Permissions
        uses: ./.github/actions/check-permissions
        with:
          permissions-file: 'permissions.yml'
          actor: '${{ github.actor }}'

      - name: Echo Something
        run: echo "Hello World!"
```

And the allowed github users list in the `permissions.yml` file_
```yml
allowed_users:
- test-user
- mokhlesurr031
```

## Github Repository
You can visit this repo for the complete structure => [Click Here](https://github.com/mokhlesurr031/github-action-by-permission)




