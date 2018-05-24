#!/bin/bash
linux_build_name="Norcia_linux_amd64"
drawin_build_name="Norcia_drawin_amd64"
win_build_name="Norcia_win_amd64.exe"

echo "    _   _                _       ";
echo "   | \ | | ___  _ __ ___(_) __ _ ";
echo "   |  \| |/ _ \| '__/ __| |/ _  |";
echo "   | |\  | (_) | | | (__| | (_| |";
echo "   |_| \_|\___/|_|  \___|_|\__,_|";
echo "----------build script--------------";

echo"";
echo "start build for linux_amd64";
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${linux_build_name} Norcia.go;
echo "build for linux_amd64 successful";
echo "------------------------------------\n";

echo "start build for drawin_amd64";
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${drawin_build_name} Norcia.go;
echo "build for drawin_amd64 successful ";
echo "------------------------------------\n"

echo "start build for window_amd64";
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${win_build_name} Norcia.go;
echo "build for window_amd64 successful ";
echo "------------------------------------\n";

echo "finish build in ";
echo "\t ./${linux_build_name}";
echo "\t ./${drawin_build_name}";
echo "\t ./${win_build_name}"
echo "------------------------------------\n\n";