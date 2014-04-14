#!/bin/bash
go build
cp -rf ./web ./moremon ../moremon-openshift
goupx -s=true ../moremon-openshift/moremon
