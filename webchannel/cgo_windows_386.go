package webchannel

/*
#cgo CPPFLAGS: -pipe -fno-keep-inline-dllexport -O2 -Wall -Wextra
#cgo CPPFLAGS: -DUNICODE -DQT_NO_DEBUG -DQT_CORE_LIB -DQT_NETWORK_LIB -DQT_QML_LIB -DQT_WEBCHANNEL_LIB
#cgo CPPFLAGS: -IC:/Qt/Qt5.5.1/5.5/mingw492_32/include -IC:/Qt/Qt5.5.1/5.5/mingw492_32/mkspecs/win32-g++
#cgo CPPFLAGS: -IC:/Qt/Qt5.5.1/5.5/mingw492_32/include/QtCore -IC:/Qt/Qt5.5.1/5.5/mingw492_32/include/QtNetwork -IC:/Qt/Qt5.5.1/5.5/mingw492_32/include/QtQml -IC:/Qt/Qt5.5.1/5.5/mingw492_32/include/QtWebChannel

#cgo LDFLAGS: -Wl,-s -Wl,-subsystem,windows -mthreads
#cgo LDFLAGS: -LC:/Qt/Qt5.5.1/5.5/mingw492_32/lib -lQt5Core -lQt5Network -lQt5Qml -lQt5WebChannel -lmingw32 -lqtmain -lshell32
*/
import "C"