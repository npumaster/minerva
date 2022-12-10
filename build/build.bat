cd %~dp0

cd ../
go build -o ./build/minerva.exe
cd %~dp0

cd ../server
go build -o ../build/minerva-server.exe
cd %~dp0
