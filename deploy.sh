#! /bin/bash

WORK_DIR=/home/streaking/streaking
EXE_NAME=streaking

cd $WORK_DIR
git pull origin master
kill -9 `cat pid` > /dev/null 2>&1
rm pid > /dev/null 2>&1
rm $WORK_DIR/$EXE_NAME
go build -o $WORK_DIR/$EXE_NAME
PORT=8080 BASE_URL=http://streakingapp.com nohup $WORK_DIR/$EXE_NAME > $WORK_DIR/log 2>&1 &
echo $! > pid
