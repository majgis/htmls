#!/usr/bin/env bash


node profiling/a.js &
node profiling/b.js &
htmls test.html 

kill %1
kill %2

