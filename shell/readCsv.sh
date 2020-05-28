#!/bin/bash

# IP:Port
HOST="172.20.0.4:9000"

readLineFile(){
	filename=$1
	echo "开始逐行读取文件：$filename"
	while read LINE
	do
		echo $LINE
		dataArr=(${LINE//,/ })
		`wkhtmltoimage --crop-w 600 http://$HOST/retailng/api/v1/online/wx/qrcode/create?scene=${dataArr[0]} $2/${dataArr[1]}.png`
	done < $filename
}

# 解析的文件
file=$1
# 生成图片保存的位置 /home/w123/files/store_png
savePath=$2
readLineFile $file $savePath
