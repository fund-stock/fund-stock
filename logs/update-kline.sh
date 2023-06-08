#!/bin/bash

appdir=/home/yezhiming/binary/kline
appdirHistory=/home/yezhiming/binary/kline/history
appdirError=/home/yezhiming/binary/kline/error
appdirLogs=/home/yezhiming/binary/kline/logs
runname1=iqoption-go-binance
putfile1=binary-go-binance
runname2=iqoption-go-hsb
putfile2=binary-go-hsb
runname3=iqoption-go-huobi
putfile3=binary-go-huobi
#runname4=iqoption-go-olymptrade
#putfile4=binary-go-olymptrade
#runname5=iqoption-go-qxbroker
#putfile5=binary-go-qxbroker
runname6=iqoption-go-fxcm
putfile6=binary-go-fxcm
runname7=iqoption-go-iqoption
putfile7=binary-go-iqoption
runname8=iqoption-go-monitoring
putfile8=binary-go-monitoring
port=7107
app=test-iqoption-kline

if [ ! -d $appdir ]; then
  mkdir -p $appdir
fi

if [ ! -d $appdirHistory ]; then
  mkdir -p $appdirHistory
fi

if [ ! -d $appdirError ]; then
  mkdir -p $appdirError
fi

if [ ! -d $appdirLogs ]; then
  mkdir -p $appdirLogs
fi

date=$(date "+%Y-%m-%d-%H-%M-%S")

echo "请确认上传的目录为$appdir,上传包的名称为$putfile"
sleep 1

echo "开始更新项目"
ls -l $appdir | grep $runname1 | awk '{print $6,$7,$8,$9}'
ls -l $appdir | grep $runname2 | awk '{print $6,$7,$8,$9}'
ls -l $appdir | grep $runname3 | awk '{print $6,$7,$8,$9}'
#ls -l $appdir | grep $runname4 | awk '{print $6,$7,$8,$9}'
#ls -l $appdir | grep $runname5 | awk '{print $6,$7,$8,$9}'
ls -l $appdir | grep $runname6 | awk '{print $6,$7,$8,$9}'
ls -l $appdir | grep $runname7 | awk '{print $6,$7,$8,$9}'
ls -l $appdir | grep $runname8 | awk '{print $6,$7,$8,$9}'

if [ ! -f $appdir/$runname1 ]; then
  echo "$appdir/$runname1 文件不存在请检查后再操作"
  exit 0
fi
if [ ! -f $appdir/$runname2 ]; then
  echo "$appdir/$runname2 文件不存在请检查后再操作"
  exit 0
fi
if [ ! -f $appdir/$runname3 ]; then
  echo "$appdir/$runname3 文件不存在请检查后再操作"
  exit 0
fi
#if [ ! -f $appdir/$runname4 ]; then
#  echo "$appdir/$runname4 文件不存在请检查后再操作"
#  exit 0
#fi
#if [ ! -f $appdir/$runname5 ]; then
#  echo "$appdir/$runname5 文件不存在请检查后再操作"
#  exit 0
#fi
if [ ! -f $appdir/$runname6 ]; then
  echo "$appdir/$runname6 文件不存在请检查后再操作"
  exit 0
fi
if [ ! -f $appdir/$runname7 ]; then
  echo "$appdir/$runname7 文件不存在请检查后再操作"
  exit 0
fi
if [ ! -f $appdir/$runname8 ]; then
  echo "$appdir/$runname8 文件不存在请检查后再操作"
  exit 0
fi

if [ ! -f $appdir/$putfile1 ]; then
  echo "$appdir/$putfile1 文件不存在请检查后再操作"
  exit 0
fi
if [ ! -f $appdir/$putfile2 ]; then
  echo "$appdir/$putfile2 文件不存在请检查后再操作"
  exit 0
fi
if [ ! -f $appdir/$putfile3 ]; then
  echo "$appdir/$putfile3 文件不存在请检查后再操作"
  exit 0
fi
#if [ ! -f $appdir/$putfile4 ]; then
#  echo "$appdir/$putfile4 文件不存在请检查后再操作"
#  exit 0
#fi
#if [ ! -f $appdir/$putfile5 ]; then
#  echo "$appdir/$putfile5 文件不存在请检查后再操作"
#  exit 0
#fi
if [ ! -f $appdir/$putfile6 ]; then
  echo "$appdir/$putfile6 文件不存在请检查后再操作"
  exit 0
fi
if [ ! -f $appdir/$putfile7 ]; then
  echo "$appdir/$putfile7 文件不存在请检查后再操作"
  exit 0
fi
if [ ! -f $appdir/$putfile8 ]; then
  echo "$appdir/$putfile8 文件不存在请检查后再操作"
  exit 0
fi

echo "1、停止运行项目代码"
sleep 1
app_pid1=$(ps -ef | grep $runname1 | grep 01 | awk '{print $2}')
echo "old PID $app_pid1"
kill -9 $app_pid1

app_pid2=$(ps -ef | grep $runname2 | grep 02 | awk '{print $2}')
echo "old PID $app_pid2"
kill -9 $app_pid2

app_pid3=$(ps -ef | grep $runname3 | grep 03 | awk '{print $2}')
echo "old PID $app_pid3"
kill -9 $app_pid3
#
#app_pid4=$(ps -ef | grep $runname4 | grep 04 | awk '{print $2}')
#echo "old PID $app_pid4"
#kill -9 $app_pid4
#
#app_pid5=$(ps -ef | grep $runname5 | grep 05 | awk '{print $2}')
#echo "old PID $app_pid5"
#kill -9 $app_pid5

app_pid6=$(ps -ef | grep $runname6 | grep 06 | awk '{print $2}')
echo "old PID $app_pid6"
kill -9 $app_pid6

app_pid7=$(ps -ef | grep $runname7 | grep 07 | awk '{print $2}')
echo "old PID $app_pid7"
kill -9 $app_pid7

app_pid8=$(ps -ef | grep $runname8 | grep 08 | awk '{print $2}')
echo "old PID $app_pid8"
kill -9 $app_pid8

echo "2、移动旧代码到history目录"
sleep 1
mv $appdir/$runname1 $appdirHistory/$runname1"_"$date
mv $appdir/$runname2 $appdirHistory/$runname2"_"$date
mv $appdir/$runname3 $appdirHistory/$runname3"_"$date
#mv $appdir/$runname4 $appdirHistory/$runname4"_"$date
#mv $appdir/$runname5 $appdirHistory/$runname5"_"$date
mv $appdir/$runname6 $appdirHistory/$runname6"_"$date
mv $appdir/$runname7 $appdirHistory/$runname7"_"$date
mv $appdir/$runname8 $appdirHistory/$runname8"_"$date

echo "3、重命名新版本代码名称"
sleep 1
mv $appdir/$putfile1 $appdir/$runname1
mv $appdir/$putfile2 $appdir/$runname2
mv $appdir/$putfile3 $appdir/$runname3
#mv $appdir/$putfile4 $appdir/$runname4
#mv $appdir/$putfile5 $appdir/$runname5
mv $appdir/$putfile6 $appdir/$runname6
mv $appdir/$putfile7 $appdir/$runname7
mv $appdir/$putfile8 $appdir/$runname8

echo "4、启动项目代码"
cd $appdir
chmod 755 *
nohup $appdir/$runname1 --APP_PORT=01 >/dev/null 2>$appdirError/$runname1.log &
echo $! >$appdirLogs/$runname1-running.pid
nohup $appdir/$runname2 --APP_PORT=02 >/dev/null 2>$appdirError/$runname2.log &
echo $! >$appdirLogs/$runname2-running.pid
nohup $appdir/$runname3 --APP_PORT=03 >/dev/null 2>$appdirError/$runname3.log &
echo $! >$appdirLogs/$runname3-running.pid
#nohup $appdir/$runname4 --APP_PORT=04 >/dev/null 2>$appdirError/$runname4.log &
#echo $! >$appdirLogs/$runname4-running.pid
#nohup $appdir/$runname5 --APP_PORT=05 >/dev/null 2>$appdirError/$runname5.log &
#echo $! >$appdirLogs/$runname5-running.pid
nohup $appdir/$runname6 --APP_PORT=06 >/dev/null 2>$appdirError/$runname6.log &
echo $! >$appdirLogs/$runname6-running.pid
nohup $appdir/$runname7 --APP_PORT=07 >/dev/null 2>$appdirError/$runname7.log &
echo $! >$appdirLogs/$runname7-running.pid
nohup $appdir/$runname8 --APP_PORT=08 >/dev/null 2>$appdirError/$runname8.log &
echo $! >$appdirLogs/$runname8-running.pid

sleep 5
app_new_pid1=$(ps -ef | grep $runname1 | grep 01 | awk '{print $2}')
echo "$runname1 new PID $app_new_pid1"
app_new_pid2=$(ps -ef | grep $runname2 | grep 02 | awk '{print $2}')
echo "$runname2 new PID $app_new_pid2"
app_new_pid3=$(ps -ef | grep $runname3 | grep 03 | awk '{print $2}')
echo "$runname3 new PID $app_new_pid3"
#app_new_pid4=$(ps -ef | grep $runname4 | grep 04 | awk '{print $2}')
#echo "$runname4 new PID $app_new_pid4"
#app_new_pid5=$(ps -ef | grep $runname5 | grep 05 | awk '{print $2}')
#echo "$runname5 new PID $app_new_pid5"
app_new_pid6=$(ps -ef | grep $runname6 | grep 06 | awk '{print $2}')
echo "$runname6 new PID $app_new_pid6"
app_new_pid7=$(ps -ef | grep $runname7 | grep 07 | awk '{print $2}')
echo "$runname7 new PID $app_new_pid7"
app_new_pid8=$(ps -ef | grep $runname8 | grep 08 | awk '{print $2}')
echo "$runname8 new PID $app_new_pid8"

if [ $app_new_pid1 -gt 0 ]; then
  echo "$runname1运行成功"
  echo "$date 代码已更新，$runname1已运行" >>$appdirLogs/update.log
  echo -e "$date $app已更新 " | mail -v -s '更新代码通知' 'yezhiming@yuctime.com' >/dev/null 2>&1
else
  echo "$runname1运行失败，请检查原因"
  echo "$date $runname1运行失败，请检查原因" >>$appdirLogs/update.log
  echo -e "$date $app更新失败，请检查原因 " | mail -v -s '更新代码通知' 'yezhiming@yuctime.com' >/dev/null 2>&1
  exit 0
fi
