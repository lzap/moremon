#!/bin/bash
go build
cp -rf ./web ./moremon ../moremon-openshift
