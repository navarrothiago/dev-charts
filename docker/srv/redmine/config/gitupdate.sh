#!/bin/bash
cd /usr/src/redmine/git/repos

targetarray=($( find . -maxdepth 2 -type d -name '*.git' -not -name '*.wiki.git' ))

cd /usr/src/redmine/git/mirrors

sourcearray=($( find ../repos -maxdepth 2 -type d -name '*.git' -not -name '*.wiki.git' ))

rm -rf /usr/src/redmine/git/mirrors/*

for index in ${!targetarray[*]} ; do
	git clone --mirror ${sourcearray[$index]} ${targetarray[$index]}
done
