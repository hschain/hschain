#!/bin/bash

Sleepfun(){

	ti1=`date +%s`    #获取时间戳
	ti2=`date +%s`
	i=$(($ti2 - $ti1 ))
	while [[ "$i" -ne "$1" ]]
	do
		ti2=`date +%s`
		i=$(($ti2 - $ti1 ))
	done
}



Init(){

	num=0
	while [ $num -ne $1 ]
	do
		 ./hscli/hscli keys add node$num << EOF 
12345678
EOF
    echo $(./hscli/hscli keys show node$num -a)
		Sleepfun 1	
	
		num=$(($num + 1))
	done
}



txsend(){
	num=0
	while [ $num -ne $1 ]
	do
    ./hscli/hscli tx send $(./hscli/hscli keys show node -a) $(./hscli/hscli keys show node$num -a) 1000000uhst --chain-id=test -y  << EOF 
12345678
EOF
        echo $(./hscli/hscli keys show node -a) "<====================>" $(./hscli/hscli keys show node$num -a)
		Sleepfun 5
		num=$(($num + 1))
	done
}


#Init 5

txsend 5


./hscli/hscli rest-server --chain-id=test
