#!/bin/bash
cd $(dirname $0)
pwd
sleep 1
(sleep 1; open "http://localhost:8080") &
exec ./localhttp
