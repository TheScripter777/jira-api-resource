# jira-api-issue-resource

<sub>Some section needs to be updated</sub>

Version: 1.3.1

| Build status | Scrutinizer |
|--------------|-------------|
| [![Build Status](https://travis-ci.org/TurnsCoffeeIntoScripts/jira-api-issue-resource.svg?branch=master)](https://travis-ci.org/TurnsCoffeeIntoScripts/jira-api-issue-resource) | [![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/TurnsCoffeeIntoScripts/jira-api-issue-resource/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/TurnsCoffeeIntoScripts/jira-api-issue-resource/?branch=master) |

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/e6ea2afc744d4fbf8bffc65e794155f4)](https://www.codacy.com/app/TurnsCoffeeIntoScripts/jira-api-issue-resource?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=TurnsCoffeeIntoScripts/jira-api-issue-resource&amp;utm_campaign=Badge_Grade)   
<sub>Project certification on default branch (master)</sub>

This [Concourse](https://concourse-ci.org/) resource allows a pipeline to interface with a Jira REST API in order to manage (create/update/delete) issues.
It is intended to be as dynamic and generic as possible to allow a vast array of possible uses. In this regard it is important
to know that jira-api-issue-resource benefits greatly from being used with [glif](https://github.com/TurnsCoffeeIntoScripts/git-log-issue-finder).

# Table of content
1. [Resource Type Configuration](#Resource-Type-Configuration)
2. [Source Configuration](#Source-Configuration) 
    1. [Required Parameters Definition](#Required-Parameters-Definition)
    2. [Optionnal Parameters Definition](#Optionnal-Parameters-Definition)
    3. [Optionnal Flags Definition](#Optionnal-Flags-Definition) 
    4. [Context usage](#Context-Usage)
3. [Behavior](#Behavior)
    1. [Check](#Check)
    2. [In](#In)
    3. [Out](#Out)
4. [Contributing](#Contributing)

## Resource Type Configuration
``` yaml
resource_types:
    - name: jira-api-issue
      type: docker-image
      source:
          repository: turnscoffeeintoscripts/jira-api-issue-resource
          tag: latest
```

## Source Configuration
``` yaml
resources:
    - name: jira
      type: jira-api-issue
      source:
          url: https://...
          username: XXXX
          password: ((password-in-vault)
          context: <SEE_CONTEXT_USAGE>       
```

Firstly, here's a list of all required and optionnal parameters followed by a list of the optionnal flags:

### Required Parameters Definition

| Parameter      | Default Value | Description                                                        |
|----------------|---------------|--------------------------------------------------------------------|
| `url`          | nil           | The base URL of the Jira API                                       |
| `user`         | nil           | The username used to connect to the Jira API                       |
| `password`     | nil           | The password needed to connect to the Jira API                     |
| `context`      | nil           | The context of execution (see 'Context Usage bellow')              |

### Optionnal Parameters Definition
| Parameter             | Default Value | Description                                                       |
|-----------------------|---------------|-------------------------------------------------------------------|
| `loggingLevel`        | `INFO`        |                                                                   |
| `transitionName`      | `Reopened`    |                                                                   |
| `closedStatusName`    | `Closed`      |                                                                   |

### Optionnal Flags Definition
| Flag              | Description                                                                       |
|-------------------|-----------------------------------------------------------------------------------|
| `forceOnParent`   |                                                                                   |
| `forceOpen`       |                                                                                   |

### Context Usage
Here's the list of the available contexts that can be used. Each context will directly influence what operations will be
performed by the resource. 
1. [ReadIssue](#ReadIssue)
2. [ReadStatus](#ReadStatus)
3. [EditCustomField](#EditCustomField)
4. [AddComment](#AddComment)

#### ReadIssue
Documentation coming soon...

#### ReadStatus
Documentation coming soon...

#### EditCustomField
**This context allows the resource to be used in 'put' steps**. It also allows the edition of any free text element in a
Jira issue. These elements are referred to as 'custom fields'. 

Here's a simple example of the resource's configuration:
``` yaml
resources:
  - name: jira-build-number
    type: jira-api-issue
    source:
      url: https://jira....
      username: username1
      password: ((password-in-vault))
      context: EditCustomField
      custom_field_name: "Build Number"
      custom_field_type: "string"
```
As can be seen in this configuration neither the field value or the issue(s) are specified. Since this resource is meant
to be as dynamic as possible those values will be provided in the put step. The first example is done without the use of
glif. Therefore the issue(s) are directly specified in the parameters. The parameter `custom_field_value` is also used
meaning the value that will be put in the issue(s) is hardcoded in the put step configuration. 
``` yaml
jobs:
  - name: add-build-number
    serial: true
    public: false
    plan:
      ...
      - put: jira-build-number
        params:
          issues: "ABC-123 XYZ-1649 TEST-456"
          custom_field_value: "1.0.0"
```

Next is a more dynamic (and realistic) example. Glif is used meaning we have a dynamic list of issues. The custom value
is also specified via a file that a previous step will provide; perfect if you're using the [semver](https://github.com/concourse/semver-resource)
concourse resource. 
``` yaml
jobs:
  - name: add-build-number
    serial: true
    public: false
    plan:
      ...
      - get: version-rc # semver resource
      - task: glif
        file: path/to/glif/task/file.yml
      - put: jira-build-number
        params:
          issue_file_location: path/to/directory/
          custom_field_value_from_file: path/to/file/with/value.txt
 
```
<sub>You may want to read [glif](https://github.com/TurnsCoffeeIntoScripts/git-log-issue-finder) documentation to properly
setup the `glif` task in this example</sub>

Lets explain the two paramters `issue_file_location` and `custom_field_value_from_file`. 

The first parameter, `issue_file_location`, is the path to the directory in which there is one or more file (*.txt)
containing the list of issues.  
The second parameter, `custom_field_value_from_file`, is the path to the file containing the value to edit in said issue(s).

#### AddComment
**This context allows the resource to be used in 'put' steps**. It simply adds a comment to the specified issue(s).

Here's a simple example of the resource's configuration:
``` yaml
resources:
  - name: jira-comment
    type: jira-api-issue
    source:
      url: https://jira....
      username: username1
      password: ((password-in-vault))
      context: AddComment
```
The actual body of the comment (the text) is set in the 'params' section of the step:
``` yaml
jobs:
  - name: put-jira-comment
    serial: true
    public: false
    plan:
      ...
      - put: jira-comment
        params:
          issues: "ABC-123 XYZ-1649 TEST-456"
          comment_body: "Comment made from Concourse"
```
This step will post a the comment `Comment made from Concourse` to each of the following issues: ABC-123, XYZ-1649 and TEST-456. 

## Behavior
### Check
**NOOP**: does nothing.
### In
**NOOP**: does nothing. There are feature that will be coming soon.
### Out
Edit the issue(s) specified in the step parameters. Depending on the context defined in the resource various fields or
parameters will be updated. For more specific see the [context usage](#Context-Usage) section.

## Contributing
Anybody is welcome to contribute to this resource. You should checkout the develop `branch` and create your feature branch
from there. Only pull-requests made to the `develop` branch will be looked at and eventually accepted. Once `develop` is
stable and contains the desired feature a merge to `master` will be made and following that a release (tag, docker image).

For any questions or inquiries you can send an email at: [turns.coffee.into.scripts@gmail.com](mailto:turns.coffee.into.scripts@gmail.com) 