<h1 align="center">
  gitxpose
  <br>
</h1>
<h4 align="center">Find the All exposed API, token before hacker's do</h4>
<p align="center">
  <a href="https://pkg.go.dev/github.com/mrmahile/gitxpose"><img src="https://pkg.go.dev/badge/github.com/mrmahile/gitxpose.svg"></a>
  <a href="https://github.com/R0X4R/Pinaak/issues"><img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat"></a>
  <a href="https://goreportcard.com/report/github.com/mrmahile/gitxpose"><img src="https://goreportcard.com/badge/github.com/mrmahile/gitxpose"></a>
  <a href="https://github.com/mrmahile/gitxpose/blob/master/LICENSE"><img src="https://img.shields.io/badge/License-MIT-yellow.svg"></a>
  <a href="https://twitter.com/mrmahile"><img src="https://img.shields.io/badge/twitter-%40mrmahile-blue.svg"></a>
</p>

<p align="center">
  <a href="#installation">Installation</a> •
  <a href="#usage">Usage</a> •
  <a href="#workflow">Workflow</a>
</p>

## Prerequisites
Before installing gitxpose, ensure you have:

```
TruffleHog (for vulnerability scanning):
git clone https://github.com/trufflesecurity/trufflehog.git
cd trufflehog; go install

Notify (for Discord notifications):
go install -v github.com/projectdiscovery/notify/cmd/notify@latest
```

## Installation
```
go install github.com/mrmahile/gitxpose@latest
```

## Download prebuilt binaries
```
wget https://github.com/mrmahile/gitxpose/releases/download/v0.0.1/gitxpose-linux-amd64-0.0.1.tgz
tar -xvzf gitxpose-linux-amd64-0.0.1.tgz
rm -rf gitxpose-linux-amd64-0.0.1.tgz
mv gitxpose ~/go/bin/gitxpose
```
Or download [binary release](https://github.com/mrmahile/gitxpose/releases) for your platform.

## Compile from source
```
git clone --depth 1 github.com/mrmahile/gitxpose.git
cd gitxpose; go install
```

## Usage
```yaml
           _  __
   ____ _ (_)/ /_ _  __ ____   ____   _____ ___
  / __  // // __/| |/_// __ \ / __ \ / ___// _ \
 / /_/ // // /_ _>  < / /_/ // /_/ /(__  )/  __/
 \__, //_/ \__//_/|_|/ .___/ \____//____/ \___/
/____/              /_/
                        Current GitXpose version v0.0.1

A longer description of your application.

Usage:
  gitxpose [flags]
  gitxpose [command]

Available Commands:
  code         Fetch code from multiple commits
  commit       Show commit logs
  completion   Generate the autocompletion script for the specified shell
  download     Clone Git repositories with a custom directory name and parallel option
  help         Help about any command
  leaksmoniter Monitor GitHub repositories for leaks and vulnerabilities
  member       Fetch GitHub member name of a single ORG or multiple ORGS using a list of orgnames
  org          Fetch GitHub repositories of a single ORG or multiple ORGS using a list of orgnames
  user         Fetch GitHub repositories of a single USER or multiple USERS using a list of usernames
  vuln         Scan repositories for vulnerabilities using TruffleHog

Flags:
  -h, --help      help for gitxpose
  -u, --update    update gitxpose to latest version
  -v, --version   Print the version of the tool and exit.

Use "gitxpose [command] --help" for more information about a command.
```

## Usage Examples

<details>
  <summary>
    <b>gitxpose org -h</b>
  </summary>

```yaml
           _  __
   ____ _ (_)/ /_ _  __ ____   ____   _____ ___
  / __  // // __/| |/_// __ \ / __ \ / ___// _ \
 / /_/ // // /_ _>  < / /_/ // /_/ /(__  )/  __/
 \__, //_/ \__//_/|_|/ .___/ \____//____/ \___/
/____/              /_/
                        Current GitXpose version v0.0.1

Examples:
$ echo "IBM" | gitxpose org -c -o output.json
$ cat orgnames.txt | gitxpose org -c -o output.json
$ echo "IBM" | gitxpose org --delay 1ns
$ echo "IBM" | gitxpose org --config custompath/config.yaml -t custompath/github-token.txt

Usage:
  gitxpose org [flags]

Flags:
      --config string   path to the config.yaml file (default "$HOME/.config/gitxpose/config.yaml")
  -c, --custom-field    Custom Fields JSON output
      --delay string    Delay duration between requests (e.g., 1ns, 1us, 1ms, 1s, 1m) (default "-1ns")
  -h, --help            help for org
  -o, --output string   File to save the output.
  -t, --token string    Path to the file containing GitHub tokens, 1 token per line (default "$HOME/.config/gitxpose/github-token.txt")
```
</details>

<details>
  <summary>
    <b>gitxpose member -h</b>
  </summary>

```yaml
           _  __
   ____ _ (_)/ /_ _____ ___   ____   ____   ___   ____   __  __ ____ ___
  / __  // // __// ___// _ \ / __ \ / __ \ / _ \ / __ \ / / / // __  __ \
 / /_/ // // /_ / /   /  __// /_/ // /_/ //  __// / / // /_/ // / / / / /
 \__, //_/ \__//_/    \___// .___/ \____/ \___//_/ /_/ \__,_//_/ /_/ /_/
/____/                    /_/
                                         Current gitxpose version v0.0.1

Examples:
$ echo "IBM" | gitxpose member -c -o output.json
$ cat orgnames.txt | gitxpose member -c -o output.json
$ echo "IBM" | gitxpose member --delay 1ns
$ echo "IBM" | gitxpose member --config custompath/config.yaml -t custompath/github-token.txt

Usage:
  gitxpose member [flags]

Flags:
      --config string   path to the config.yaml file (default "$HOME/.config/gitxpose/config.yaml")
  -c, --custom-field    Custom Fields JSON output
      --delay string    Delay duration between requests (e.g., 1ns, 1us, 1ms, 1s, 1m) (default "-1ns")
  -h, --help            help for member
  -o, --output string   File to save the output.
  -t, --token string    Path to the file containing GitHub tokens, 1 token per line (default "$HOME/.config/gitxpose/github-token.txt")
```
</details>

<details>
  <summary>
    <b>gitxpose user -h</b>
  </summary>

```yaml
           _  __
   ____ _ (_)/ /_ _____ ___   ____   ____   ___   ____   __  __ ____ ___
  / __  // // __// ___// _ \ / __ \ / __ \ / _ \ / __ \ / / / // __  __ \
 / /_/ // // /_ / /   /  __// /_/ // /_/ //  __// / / // /_/ // / / / / /
 \__, //_/ \__//_/    \___// .___/ \____/ \___//_/ /_/ \__,_//_/ /_/ /_/
/____/                    /_/
                                         Current gitxpose version v0.0.1

Fetch GitHub repositories of a single USER or multiple USERS using a list of usernames

Examples:
$ echo "mrmahile" | gitxpose user -c -o output.json
$ cat usernames.txt | gitxpose user -c -o output.json
$ echo "mrmahile" | gitxpose user --delay 1ns
$ echo "mrmahile" | gitxpose user --config custompath/config.yaml -t custompath/github-token.txt

Usage:
  gitxpose user [flags]

Flags:
      --config string   path to the config.yaml file (default "$HOME/.config/gitxpose/config.yaml")
  -c, --custom-field    Custom Fields JSON output
      --delay string    Delay duration between requests (e.g., 1ns, 1us, 1ms, 1s, 1m) (default "-1ns")
  -h, --help            help for user
  -o, --output string   File to save the output.
  -t, --token string    Path to the file containing GitHub tokens, 1 token per line (default "$HOME/.config/gitxpose/github-token.txt")
```
</details>

<details>
  <summary>
    <b>gitxpose download -h</b>
  </summary>

```yaml
           _  __
   ____ _ (_)/ /_ _____ ___   ____   ____   ___   ____   __  __ ____ ___
  / __  // // __// ___// _ \ / __ \ / __ \ / _ \ / __ \ / / / // __  __ \
 / /_/ // // /_ / /   /  __// /_/ // /_/ //  __// / / // /_/ // / / / / /
 \__, //_/ \__//_/    \___// .___/ \____/ \___//_/ /_/ \__,_//_/ /_/ /_/
/____/                    /_/
                                         Current gitxpose version v0.0.1

Clone Git repositories and customize the directory name to username-repositoryname with an option to clone in parallel.

Examples:
$ echo "https://github.com/mrmahile/gitxpose.git" | gitxpose download
$ cat reponames.txt | gitxpose download
$ cat reponames.txt | gitxpose download -o ~/allgithubrepo/download
$ cat reponames.txt | gitxpose download -p 100
$ cat reponames.txt | gitxpose download -d 1

Usage:
  gitxpose download [flags]

Flags:
  -d, --depth int                 Create a shallow clone with a history truncated to the specified number of commits, use -d 0 if you want all commits (default 1)
  -h, --help                      help for download
  -o, --output-directory string   Directory to clone the repositories into (default "/root/allgithubrepo/download")
  -p, --parallel int              Number of repositories to clone in parallel (default 10)
```
</details>

<details>
  <summary>
    <b>gitxpose commit -h</b>
  </summary>

```yaml
           _  __
   ____ _ (_)/ /_ _____ ___   ____   ____   ___   ____   __  __ ____ ___
  / __  // // __// ___// _ \ / __ \ / __ \ / _ \ / __ \ / / / // __  __ \
 / /_/ // // /_ / /   /  __// /_/ // /_/ //  __// / / // /_/ // / / / / /
 \__, //_/ \__//_/    \___// .___/ \____/ \___//_/ /_/ \__,_//_/ /_/ /_/
/____/                    /_/
                                         Current gitxpose version v0.0.1

This command retrieves git commit logs based on date and time parameters.

Examples:
$ gitxpose commit -i ~/allgithubrepo/download -d 50s -t before -o ~/allgithubrepo/commits
$ gitxpose commit -i ~/allgithubrepo/download -d 5h -t before -o ~/allgithubrepo/commits
$ gitxpose commit -i ~/allgithubrepo/download -d 1d -t after -o ~/allgithubrepo/commits
$ gitxpose commit -i ~/allgithubrepo/download -d all -o ~/allgithubrepo/commits

Date Options:
50s     # 50 seconds
40m     # 40 minutes
5h      # 5 hours
1d      # 1 day
2w      # 2 weeks
3M      # 3 months
1y      # 1 year
all     # All commits

Usage:
  gitxpose commit [input] [flags]

Flags:
  -d, --date string     Specify the date range for the commits (e.g., 50s, 40m, 5h, 1d, 2w, 3M, 1y, all) (default "all")
  -h, --help            help for commit
  -i, --input string    Specify the input directory containing Git repositories (default "/root/allgithubrepo/download")
  -o, --output string   Specify the output directory for commit logs (default "/root/allgithubrepo/commits")
  -t, --time string     Specify 'before' or 'after' the given date
```
</details>

<details>
  <summary>
    <b>gitxpose code -h</b>
  </summary>

```yaml
           _  __
   ____ _ (_)/ /_ _____ ___   ____   ____   ___   ____   __  __ ____ ___
  / __  // // __// ___// _ \ / __ \ / __ \ / _ \ / __ \ / / / // __  __ \
 / /_/ // // /_ / /   /  __// /_/ // /_/ //  __// / / // /_/ // / / / / /
 \__, //_/ \__//_/    \___// .___/ \____/ \___//_/ /_/ \__,_//_/ /_/ /_/
/____/                    /_/
                                         Current gitxpose version v0.0.1

This command fetches code from multiple commits based on a list in commits.txt for each repository.

Examples:
$ gitxpose code -i ~/allgithubrepo/download -o ~/allgithubrepo/commits

Usage:
  gitxpose code [flags]

Flags:
  -h, --help            help for code
  -i, --input string    Specify the input directory containing Git repositories (default "/root/allgithubrepo/download")
  -o, --output string   Specify the output directory for storing commit code (default "/root/allgithubrepo/commits")
```
</details>

<details>
  <summary>
    <b>gitxpose vuln -h</b>
  </summary>

```yaml
           _  __
   ____ _ (_)/ /_ _____ ___   ____   ____   ___   ____   __  __ ____ ___
  / __  // // __// ___// _ \ / __ \ / __ \ / _ \ / __ \ / / / // __  __ \
 / /_/ // // /_ / /   /  __// /_/ // /_/ //  __// / / // /_/ // / / / / /
 \__, //_/ \__//_/    \___// .___/ \____/ \___//_/ /_/ \__,_//_/ /_/ /_/
/____/                    /_/
                                         Current gitxpose version v0.0.1

This command scans multiple repositories for vulnerabilities using TruffleHog
and saves the results in the specified output directory.

Examples:
$ gitxpose vuln
$ gitxpose vuln -i ~/allgithubrepo/commits -o ~/allgithubrepo/commits

Usage:
  gitxpose vuln [flags]

Flags:
  -h, --help            help for vuln
  -i, --input string    Input directory containing repositories code (default "/root/allgithubrepo/commits")
  -o, --output string   Output directory for vulnerability reports (default "/root/allgithubrepo/commits")
```
</details>

<details>
  <summary>
    <b>gitxpose leaksmoniter -h</b>
  </summary>

```yaml
           _  __
   ____ _ (_)/ /_ _____ ___   ____   ____   ___   ____   __  __ ____ ___
  / __  // // __// ___// _ \ / __ \ / __ \ / _ \ / __ \ / / / // __  __ \
 / /_/ // // /_ / /   /  __// /_/ // /_/ //  __// / / // /_/ // / / / / /
 \__, //_/ \__//_/    \___// .___/ \____/ \___//_/ /_/ \__,_//_/ /_/ /_/
/____/                    /_/
                                         Current gitxpose version v0.0.1

A comprehensive tool to monitor GitHub organizations, users, and members
for potential leaks and vulnerabilities using trufflehog scanning.

Features:
- Fetch repositories from organizations, users, and their members
- Clone repositories with configurable depth and parallelism
- Extract commits and code changes
- Scan for vulnerabilities using trufflehog
- Send notifications to Discord

Examples:
  # Complete automated workflow including vulnerability scanning
  echo "Shopify" | gitxpose leaksmoniter --scan-repo org --date 24h

  # Scan individual user repositories
  echo "mrmahile" | gitxpose leaksmoniter --scan-repo user

  # Scan both org and member repositories
  cat orgnames.txt | gitxpose leaksmoniter --scan-repo org,member

  # With Discord notifications for vulnerabilities
  cat orgnames.txt | gitxpose leaksmoniter --scan-repo org,member --notifyid allvuln

  # With custom base directory
  cat orgnames.txt | gitxpose leaksmoniter --scan-repo org --download-dir ~/myrepos

  # High parallelism for faster cloning
  cat orgnames.txt | gitxpose leaksmoniter --parallel 20 --depth 10

  # Scan recent repositories only (last 7 days)
  echo "google" | gitxpose leaksmoniter --scan-repo org --date 7d

  # Comprehensive scan with all options
  echo "microsoft" | gitxpose leaksmoniter --scan-repo org,member,user --date 30d --parallel 15 --notifyid my-webhook

Usage:
  gitxpose leaksmoniter [flags]

Flags:
  -D, --date string           Specify the date range for repositories (e.g., 50s, 40m, 5h, 1d, 2w, 3M, 1y, all) (default "24h")
  -d, --delay string          Delay between requests (e.g., 1ns, 1us, 1ms, 1s, 1m) (default "-1ns")
  -z, --depth int             Git clone depth (default 5)
  -o, --download-dir string   Base directory for downloads, commits, code, and vulnerabilities
  -h, --help                  help for leaksmoniter
  -n, --notifyid string       Send verified vulnerabilities to Discord (default "allvuln")
  -p, --parallel int          Repositories to clone in parallel (default 10)
  -s, --scan-repo string      Scan type: org, member, user (default "org,member")
  -t, --token string          GitHub tokens file, 1 token per line (default "$HOME/.config/gitxpose/github-token.txt")
```
</details>

## Demo

## Workflow
<p align="center"> 
<a href="Workflow/gitxpose.png" target="_blank"> 
<img src="Workflow/gitxpose.png"/>
</a>  
</p>
