# Mentor-Management-System

Mentors Management system is a people management app that enables proper.
coordination of mentors needed to execute projects, ranging from recruitment to off-boarding. Ensue to go through the app doc below to read more and follow all the instructions.

\*[Contributor's wiki](https://github.com/ALCOpenSource/Mentor-Management-System-Team-7/wiki)

## Requirements
To build and run this project, you will need:
-   Go v1.16 or later
-   [Docker](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-20-04)
-   [make(1)](https://man7.org/linux/man-pages/man1/make.1.html) utility
-   [Redis](https://redis.io/)
-   Mongo Database

## Deployment
The makefile included in this project provides several helpful commands to simplify the deployment and testing process. Set the [DB_SOURCE]("./app.env") environment variable to a valid Mongo connection string. 

#### Running the server

```bash
make redis
make server
```
will run a Redis container with the name "redis", mapping port 6379, and running in the background. Then compile and run the Go program contained in the file main.go.

#### Test

```bash
make test
```
will run all tests in the current directory and its subdirectories, display verbose output, and generate a coverage report.

## How to contributing to this project:

To get it up and running on your local machine, follow the steps below:

*   Fork this repo following this [guideline](https://docs.github.com/en/get-started/quickstart/fork-a-repo).
*   Clone the repo with the command `git clone`
*   Indicate your interest to work on any issue. "eg. I want to work on this issue or I am interested in this issue"
*   Open a feature branch from the 'develop' branch. eg feat/
*   Make sure the name is descriptive for your branch but not too long. Lead with what the the branch is doing eg new feature or bug but follow this pattern `type/branch-description` eg `feature/add-login-functionality`.
*   Ensure your branch is up to date with latest changes before pushing
*   Create a pull request against develop branch
*   Reference the issue you worked on in your PR
*   Open a pull request against the develop branch and request a review from your

##### Contributors

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->

<!-- prettier-ignore-start -->

<!-- markdownlint-disable -->

<!-- markdownlint-restore -->

<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->

[![All Contributors](https://img.shields.io/badge/all_contributors-8-orange.svg?style=flat-square)](#contributors)

<!-- ALL-CONTRIBUTORS-BADGE:END -->
