# Vultr API in go
[![Build Status](https://travis-ci.org/askholme/vultr.svg?branch=master)](https://travis-ci.org/askholme/vultr)

A partly go implementation of the Vultr API.
This is being made for use with packer/terraform plugins, thus only servers and snapshots are supported with the
few methods needed. No support for startup scripts etc.

Also it's a pure go package (not CLI tool) - for that look elsewhere
