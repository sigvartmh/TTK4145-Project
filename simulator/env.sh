FLOORS=4
GOMAXPROCS=20
libPath=$PWD
if [[ "$platform"  == "Linux" ]]; then
    CGO_LDFLAGS="$CGO_LDFLAGS -lcomedi"
    CGO_LDFLAGS="$libPath/driver/lib/simulation_elevator.a $libPath/driver/lib/libphobos2.a -lpthread -lm"
elif [[ "$platform"  == "Darwin" ]]; then
    CGO_LDFLAGS="$libPath/driver/lib/osx/simulation_elevator.a $libPath/driver/lib/osx/libphobos2.a -lpthread -lm"
fi
platform=`uname`
echo $platform
if [[ "$platform"  == "Linux" ]]; then
    CGO_LDFLAGS="$CGO_LDFLAGS -lcomedi"
fi
echo "Linker flags: $CGO_LDFLAGS"
echo "Max numbers of floors: $FLOORS"
echo "Max amount of OS thread used: $GOMAXPROCS"
export GOMAXPROCS
export FLOORS
export CGO_LDFLAGS

