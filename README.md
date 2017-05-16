[![CircleCI](https://img.shields.io/circleci/project/github/RedSparr0w/node-csgo-parser.svg)](https://circleci.com/gh/SignifAi/snap-plugin-processor-metric-repeat-filter)
[![Hex.pm](https://img.shields.io/hexpm/l/plug.svg)](https://github.com/SignifAi/snap-plugin-processor-metric-repeat-filter/blob/master/LICENSE)

# snap-plugin-processor-metric-repeat-filter
Collects state metrics from your Nagios installation
and presents them to the Snap system. The filter will pass to the publisher only new events and filter any duplications.

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Contributing](#contributing)
3. [License](#license-and-authors)
4. [Acknowledgements](#acknowledgements)

## Getting Started
### System Requirements 
* [golang 1.8+](https://golang.org/dl/) (needed only for building)

### Operating systems
All OSs currently supported by snap:
* Linux/amd64
* Darwin/amd64

### Installation

#### To build the plugin binary:
Fork https://github.com/SignifAi/snap-plugin-processor-metric-repeat-filter

Clone repo into `$GOPATH/src/github.com/SignifAi/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-processor-metric-repeat-filter.git
```


#### Building
The following provides instructions for building the plugin yourself if
you decided to download the source. We assume you already have a $GOPATH
setup for [golang development](https://golang.org/doc/code.html). The
repository utilizes [glide](https://github.com/Masterminds/glide) for
library management.

build:
  `make`

testing:
  `make test`

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)

#### Load the Plugin
Once the framework is up and running, you can load the plugin.

```
$ snaptel plugin load snap-plugin-collector-nagios
Plugin loaded
Name: metric-repeat-filter
Version: 1
Type: processor
Signed: false
Loaded Time: Sat, 18 Mar 2017 13:28:45 PDT
```

#### Task File
You need to create or update a task file to use the Nagios collector
plugin. We have provided an example, __examples/tasks/nagios-onlyone-task.yaml_ shown below. In
our example, we utilize the nagios collector so we have some data to
work with. There are no configuration settings to set.


```
---
  version: 1
  schedule:
    type: "simple"
    interval: "1s"
    count: 5
  max-failures: 10
  workflow:
    collect:
      config:
      metrics:
        /nagios/*/acknowledged: {} 
        /nagios/*/state: {}
        /nagios/*/services/*/acknowledged: {}
        /nagios/*/services/*/state: {}
      process:
        - plugin_name: "metric-repeat-filter"
          publish:
            - plugin_name: file
              config:
                path: /tmp/mrf-test.log
```

Once the task file has been created, you can create and watch the task.

```
$ snaptel task create -t examples/tasks/nagios-status-task.yaml
Using task manifest to create task
Task created
ID: 72869b36-def6-47c4-9db2-822f93bb9d1f
Name: Task-72869b36-def6-47c4-9db2-822f93bb9d1f
State: Running

$ snaptel task list
ID                                       NAME
STATE     ...
72869b36-def6-47c4-9db2-822f93bb9d1f
Task-72869b36-def6-47c4-9db2-822f93bb9d1f    Running
```

## Contributing  - We love contributions!

The most immediately helpful way you can benefit this plug-in is by cloning the repository, adding some further examples and submitting a pull request.

## License
Released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
* Author: [@SignifAi](https://github.com/SignifAi/)
* Info: www.signifai.io
