#!/bin/bash

if test $# -ne 1
then
    echo 'Usage: 1. ./load_function.sh [database name]'
    echo '       2. Press enter until end'
    echo 'eg. ./load_function.sh cs2102project'
else
    psql -d $1 -f ../1_log.sql
    psql -d $1 -f ../categories.sql
    psql -d $1 -f ../comments.sql
    psql -d $1 -f ../payments.sql
    psql -d $1 -f ../projects.sql
    psql -d $1 -f ../users.sql
    psql -d $1 -f ../tags.sql
    echo 'load done.'
fi