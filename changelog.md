# Changelog
All notable changes to this project will be documented in this file.

The format of this file is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), 
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased] 
### Added
- Example pipeline with associated tasks to help people configuration this resource alongside glif.
### Changed
- Release script now also updates the tag of the resource in the example pipeline.
- Remove trailing comma when preparing 'ref' value in responses

## [1.3.9] - 2020-02-25
### Changed
- Modified 'out' models to match concourse Go struct VersionResult

## [1.3.8] - 2020-02-24
### Changed
- In script now takes into account the .version.ref path in the json request if the issue list is empty. 

## [1.3.7] - 2020-02-24

## [1.3.6] - 2020-02-24
### Added
- Output to &3 and to file the json result in the 'in' asset 
### Changed
- Refactored services method 'ExecuteAsLastStep' signature to take full parameters instead of just the context

## [1.3.5] - 2020-02-20
### Changed
- How the json is output via the 'jq' command in the out script

## [1.3.4] - 2020-02-20
### Changed
- Version reference in out script

## [1.3.3] - 2020-02-20
### Changed
- Version reference in out script

## [1.3.2] - 2020-02-20
### Changed
- Version reference in out script

## [1.3.1] - 2020-02-19
### Changed
- Updated test in 'configuration' package

## [1.3.0] - 2020-02-18
### Added
- Implemented the 'AddComment' context

## [1.2.0] - 2020-02-18
### Added
- 'forceOpen' functionality now implemented
- Added the 'keepGoing' flag which ignores error on 1 issues and exit current step list. This allows the pipeline to continue to the next issue, if any.
### Changed
- The 'out' asset now recognizes both an explicitly defined list and a path to a directory containing 'issues' files.
- Partial refactored the logic used for 'cross-steps' values
- What was used before to find out if an issue had a parent is now used to fetch various data prior to executing the pipeline.
- Error management now uses a custom type instead of built-in 'error' type. This allows for more information to be transmitted.

## [1.1.3] - 2020-02-12
### Added
- 'forceOpen' flag in the parameters object (not yet used but recognized)
- Internal mechanic to prevent the custom field ID extraction when reading an issue for its parent status

## [1.1.2] - 2020-02-11
### Changed
- Changed the whitespace removal for a simple field extraction instead

## [1.1.1] - 2020-02-11
### Added
- 'in' script for concourse integration
- Allows to read the status (open, in progress, etc...) of an issue
- Added structs that allows mapping of an 'in' response (version, metadata)
- New concept allows service to execute 'as last step' of a pipeline (useful for outputing results)
- Fixed bug of "Ticket doesn't exist" caused by trailing whitespaces in issues list
### Changed
- Renamed package containing auth info from 'status' to 'auth'
- Update to release script for readability

## [1.1.0] - 2020-01-21
### Added
- 'in' script for concourse integration
- Allows to read the status (open, in progress, etc...) of an issue
- Added structs that allows mapping of an 'in' response (version, metadata)
- New concept allows service to execute 'as last step' of a pipeleine (useful for outputing results)
### Changed
- Renamed package containing auth info from 'status' to 'auth'
- Update to release script for readability

## [1.0.5] - 2020-01-16
### Changed
- Compartmentalized the parameters according to their respective main context. This allows for clearer read and use.

## [1.0.4] - 2020-01-16
### Changed
- Change made in 1.0.3 regarding value from file should have been made for issue list instead
### Removed
- 1.0.3 change

## [1.0.3] - 2020-01-16
### Added
- When reading value from file we make sure the resulting string is cleaned (no brackets and such)

## [1.0.2] - 2020-01-15
### Changed
- Modified release.sh so that it now updates the version in Dockerfile's LABEL step
- Changed how HTTP 4xx and 5xx bodies are read and printed in the logger

## [1.0.1] - 2020-01-14
### Added
- Added the changelog.md file 
### Changed
- Modified version of release.sh to manage changelog.md and README.md

## [N/A] - 2020-01-14
### Changed
- Renaming from 'jira-api-resource' to 'jira-api-issue-resource'

## [1.0.0] - 2020-01-08
Initial version that supports
* Read issues
* Update issues
* 'force-on-parent' mechanic