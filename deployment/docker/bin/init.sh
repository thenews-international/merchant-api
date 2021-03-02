#!/usr/bin/env ash

echo 'Deleting mysql-client...'
apk del mysql-client

echo 'Start application...'
/merchant/bin/app