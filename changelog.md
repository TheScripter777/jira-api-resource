# Changelog
All notable changes to this project will be documented in this file.

The format of this file is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), 
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased] 
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
