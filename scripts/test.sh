#!/usr/bin/env bash

go test -race ./daemon/...                # Run all daemon package tests ...
#go test -race ./lbrycrd/...               # Run all lbrycrd package tests ...
##go test -race ./db/...                    # Run all db package tests ...
go test -race ./datastore/...             # Run all datastore package tests ...