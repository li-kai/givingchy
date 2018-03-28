#!/bin/sh

psql -d testdb -U postgres -f seed.sql
psql -d testdb -U postgres -f categories.sql
psql -d testdb -U postgres -f comments.sql
psql -d testdb -U postgres -f payments.sql
psql -d testdb -U postgres -f projects.sql
psql -d testdb -U postgres -f users.sql
echo 'load done.'
