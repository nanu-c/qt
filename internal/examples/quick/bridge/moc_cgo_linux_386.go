package main

/*
#cgo CPPFLAGS: -pipe -O2 -g -pipe -Wall -Wp,-D_FORTIFY_SOURCE=2 -fexceptions -fstack-protector --param=ssp-buffer-size=4 -Wformat -Wformat-security -m32 -msse -msse2 -march=i686 -mfpmath=sse -mtune=generic -fno-omit-frame-pointer -fasynchronous-unwind-tables -fPIC -fvisibility=hidden -fvisibility-inlines-hidden -Wall -W -D_REENTRANT -fPIE
#cgo CPPFLAGS: -DQT_NO_DEBUG -DQT_CORE_LIB
#cgo CPPFLAGS: -I/srv/mer/targets/SailfishOS-i486/usr/share/qt5/mkspecs/linux-g++ -I/srv/mer/targets/SailfishOS-i486/usr/include -I/srv/mer/targets/SailfishOS-i486/usr/include/sailfishapp -I/srv/mer/targets/SailfishOS-i486/usr/include/mdeclarativecache5 -I/srv/mer/targets/SailfishOS-i486/usr/include/qt5
#cgo CPPFLAGS: -I/srv/mer/targets/SailfishOS-i486/usr/include/qt5/QtCore

#cgo LDFLAGS: -Wl,-O1 -Wl,-rpath-link,/srv/mer/targets/SailfishOS-i486/usr/lib -Wl,-rpath-link,/srv/mer/targets/SailfishOS-i486/lib
#cgo LDFLAGS: -rdynamic -L/srv/mer/targets/SailfishOS-i486/usr/lib -L/srv/mer/targets/SailfishOS-i486/lib -lsailfishapp -lmdeclarativecache5 -lQt5Core -lGLESv2 -lpthread
*/
import "C"
