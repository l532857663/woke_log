#!/bin/bash

# IP:Port
HOST="172.20.0.4:9000"

readLineFile(){
	if [ $2 == "" ]
	then
		savepath="."
	fi
	filename=$1
	echo "开始逐行读取文件：$filename"
	while read LINE
	do
		lineLen=${#LINE}
		idStr=${LINE:1:lineLen-2}
		`wkhtmltoimage --crop-w 870 http://$HOST/retailng/api/v1/online/wx/qrcode/create?id=$idStr $savepath/$idStr.png`
	done < $filename
}

# 解析的文件
file=$1
# 生成图片保存的位置 /home/w123/files/store_png
savePath=$2
# 生成图片的类型
saveType=$3
readLineFile $file $savePath $savetype
