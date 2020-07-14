for pathname in assets/downloads/*;
do
	filename=`basename $pathname`

	cat $pathname                \
		| grep -v ",,,  ,,,,,,," \
		| tail -n +2             \
		| sed -e 's/R$ //g'      \
		| cut -d"," -f -10       \
	> assets/data/$filename

done
