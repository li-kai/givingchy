#!/bin/bash

if test $# -ne 1
then
    echo 'Usage: 1. ./load_function.sh [database name]'
    echo '       2. Press enter until end'
    echo 'eg. ./load_function.sh cs2102project'
else
    psql -s $1 -f categories.sql
    psql -s $1 -f comments.sql
    psql -s $1 -f payments.sql
    psql -s $1 -f projects.sql
    psql -s $1 -f users.sql
    echo 'load done.'
fi