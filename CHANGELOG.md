## v0.7.0-rc.0 / 2017-05-19

* [CHANGE] Rewrite UI from scratch improving UX
* [FEATURE] Add ability to update a silence on `api/v1/silences` POST endpoint (See #765)
* [FEATURE] Return alert status on `api/v1/alerts` GET endpoint
* [FEATURE] Serve silence state on `api/v1/silences` GET endpoint
* [ENHANCEMENT] Add ability to specify `filter` string and `silenced` flag on `api/v1/alerts` GET endpoint
* [ENHANCEMENT] Update `cache-control` to prevent caching for web assets in general.
* [BUGFIX] Enable builds for Solaris/Illumos

## v0.6.2 / 2017-05-09

* [BUGFIX] Correctly link to silences from alert again
* [BUGFIX] Correctly hide silenced/show active alerts in UI again
* [BUGFIX] Fix regression of alerts not being displayed until first processing
* [BUGFIX] Fix internal usage of wrong lock for silence markers
* [BUGFIX] Adapt amtool's API parsing to recent API changes
* [BUGFIX] Correctly marshal regexes in config JSON response
* [ENHANCEMENT] Anchor silence regex matchers to be consistent with Prometheus
* [ENHANCEMENT] Error if root route is using `continue` keyword

## v0.6.1 / 2017-04-28

* [BUGFIX] Fix incorrectly serialized hash for notification providers.
* [ENHANCEMENT] Add config hash metric.
* [ENHANCEMENT] Add processing status field to alerts.

## v0.6.0 / 2017-04-25

* [BUGFIX] Add `groupKey` to `alerts/groups` endpoint https://github.com/prometheus/alertmanager/pull/576
* [BUGFIX] Only notify on firing alerts https://github.com/prometheus/alertmanager/pull/595
* [BUGFIX] Correctly marshal regex's in config for routing tree https://github.com/prometheus/alertmanager/pull/602
* [BUGFIX] Prevent panic when failing to load config https://github.com/prometheus/alertmanager/pull/607
* [BUGFIX] Prevent panic when alertmanager is started with an empty `-mesh.peer` https://github.com/prometheus/alertmanager/pull/726
* [CHANGE] Add `DELETE` as accepted CORS method https://github.com/prometheus/alertmanager/pull/641
* [CHANGE] Rename VictorOps config variables https://github.com/prometheus/alertmanager/pull/667
* [CHANGE] Switch to using `gogoproto` for protobuf https://github.com/prometheus/alertmanager/pull/715
* [CHANGE] No longer generate releases for openbsd/arm https://github.com/prometheus/alertmanager/pull/732
* [ENHANCEMENT] Add `reReplaceAll` template function https://github.com/prometheus/alertmanager/pull/639
* [ENHANCEMENT] Expose mesh peers on status page https://github.com/prometheus/alertmanager/pull/644
* [ENHANCEMENT] Allow label-based filtering alerts/silences through API https://github.com/prometheus/alertmanager/pull/633
* [ENHANCEMENT] Include notifier type in logs and errors https://github.com/prometheus/alertmanager/pull/702
* [ENHANCEMENT] Add commandline tool for interacting with alertmanager https://github.com/prometheus/alertmanager/pull/636

## v0.5.1 / 2016-11-24

* [BUGFIX] Fix crash caused by race condition in silencing
* [ENHANCEMENT] Improve logging of API errors
* [ENHANCEMENT] Add metrics for the notification log

## v0.5.0 / 2016-11-01

This release requires a storage wipe. It contains fundamental internal
changes that came with implementing the high availability mode.

* [FEATURE] Alertmanager clustering for high availability
* [FEATURE] Garbage collection of old silences and notification logs
* [CHANGE] New storage format
* [CHANGE] Stricter silence semantics for consistent historical view

## v0.4.2 / 2016-09-02

* [BUGFIX] Fix broken regex checkbox in silence form
* [BUGFIX] Simplify inconsistent silence update behavior

## v0.4.1 / 2016-08-31

* [BUGFIX] Wait for silence query to finish instead of showing error
* [BUGFIX] Fix sorting of silences
* [BUGFIX] Provide visual feedback after creating a silence
* [BUGFIX] Fix styling of silences
* [ENHANCEMENT] Provide cleaner API silence interface

## v0.4.0 / 2016-08-23

* [FEATURE] Silences are now paginated in the web ui
* [CHANGE] Failure to start on unparsed flags

## 0.3.0 / 2016-07-07

* [CHANGE] Alerts are purely in memory and no longer persistent across restarts
* [FEATURE] Add SMTP LOGIN authentication mechanism

## 0.2.1 / 2016-06-23

* [ENHANCEMENT] Allow inheritance of route receiver
* [ENHANCEMENT] Add silence cache to silence provider
* [BUGFIX] Fix HipChat room number in integration URL

## 0.2.0 / 2016-06-17

This release uses a new storage backend based on BoltDB. You have to backup
and wipe your former storage path to run it.

* [CHANGE] Use BoltDB as data store.
* [CHANGE] Move SMTP authentification to configuration file
* [FEATURE] add /-/reload HTTP endpoint
* [FEATURE] Filter silenced alerts in web UI
* [ENHANCEMENT] reduce inhibition computation complexity
* [ENHANCEMENT] Add support for teams and tags in OpsGenie integration
* [BUGFIX] Handle OpsGenie responses correctly
* [BUGFIX] Fix Pushover queue length issue
* [BUGFIX] STARTTLS before querying auth mechanism in email integration

## 0.1.1 / 2016-03-15
* [BUGFIX] Fix global database lock issue
* [ENHANCEMENT] Improve SQLite alerts index
* [ENHANCEMENT] Enable debug endpoint

## 0.1.0 / 2016-02-23
This version is a full rewrite of the Alertmanager with a very different
feature set. Thus, there is no meaningful changelog.

Changes with respect to 0.1.0-beta2:
* [CHANGE] Expose same data structure to templates and webhook
* [ENHANCEMENT] Show generator URL in default templates and web UI
* [ENHANCEMENT] Support for Slack icon_emoji field
* [ENHANCEMENT] Expose incident key to templates and webhook data
* [ENHANCEMENT] Allow markdown in Slack 'text' field
* [BUGFIX] Fixed database locking issue

## 0.1.0-beta2 / 2016-02-03
* [BUGFIX] Properly set timeout for incoming alerts with fixed start time
* [ENHANCEMENT] Send source field in OpsGenie integration
* [ENHANCEMENT] Improved routing configuration validation
* [FEATURE] Basic instrumentation added

## 0.1.0-beta1 / 2016-01-08
* [BUGFIX] Send full alert group state on each update. Fixes erroneous resolved notifications.
* [FEATURE] HipChat integration
* [CHANGE] Slack integration no longer sends resolved notifications by default

## 0.1.0-beta0 / 2015-12-23
This version is a full rewrite of the Alertmanager with a very different
feature set. Thus, there is no meaningful changelog.

## 0.0.4 / 2015-09-09
* [BUGFIX] Fix version info string in startup message.
* [BUGFIX] Fix Pushover notifications by setting the right priority level, as
  well as required retry and expiry intervals.
* [FEATURE] Make it possible to link to individual alerts in the UI.
* [FEATURE] Rearrange alert columns in UI and allow expanding more alert details.
* [FEATURE] Add Amazon SNS notifications.
* [FEATURE] Add OpsGenie Webhook notifications.
* [FEATURE] Add `-web.external-url` flag to control the externally visible
  Alertmanager URL.
* [FEATURE] Add runbook and alertmanager URLs to PagerDuty and email notifications.
* [FEATURE] Add a GET API to /api/alerts which pulls JSON formatted
  AlertAggregates.
* [ENHANCEMENT] Sort alerts consistently in web UI.
* [ENHANCEMENT] Suggest to use email address as silence creator.
* [ENHANCEMENT] Make Slack timeout configurable.
* [ENHANCEMENT] Add channel name to error logging about Slack notifications.
* [ENHANCEMENT] Refactoring and tests for Flowdock notifications.
* [ENHANCEMENT] New Dockerfile using alpine-golang-make-onbuild base image.
* [CLEANUP] Add Docker instructions and other cleanups in README.md.
* [CLEANUP] Update Makefile.COMMON from prometheus/utils.

## 0.0.3 / 2015-06-10
* [BUGFIX] Fix email template body writer being called with parameters in wrong order.

## 0.0.2 / 2015-06-09

* [BUGFIX] Fixed silences.json permissions in Docker image.
* [CHANGE] Changed case of API JSON properties to initial lower letter.
* [CHANGE] Migrated logging to use http://github.com/prometheus/log.
* [FEATURE] Flowdock notification support.
* [FEATURE] Slack notification support.
* [FEATURE] Generic webhook notification support.
* [FEATURE] Support for "@"-mentions in HipChat notifications.
* [FEATURE] Path prefix option to support reverse proxies.
* [ENHANCEMENT] Improved web redirection and 404 behavior.
* [CLEANUP] Updated compiled web assets from source.
* [CLEANUP] Updated fsnotify package to its new source location.
* [CLEANUP] Updates to README.md and AUTHORS.md.
* [CLEANUP] Various smaller cleanups and improvements.
