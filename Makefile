#!/usr/bin/make --no-print-directory --jobs=1 --environment-overrides -f

VERSION_TAGS += MAPS
MAPS_MK_SUMMARY := go-corelibs/maps
MAPS_MK_VERSION := v1.0.2

include CoreLibs.mk
