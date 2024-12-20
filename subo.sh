git add . 
git commit -m "ültim commit"
git push

go build main.go

rm main.zip
tar.gz -a -cf main.zip main