export FLOORS=4
libPath=$PWD
CGO_LDFLAGS="$libPath/driver/lib/simulation_elevator.a $libPath/driver/lib/libphobos2.a -lpthread -lm"
platform=`uname`
echo $platform
if [[ "$platform"  == "Linux" ]]; then
    CGO_LDFLAGS="$CGO_LDFLAGS -lcomedi"
fi
echo $CGO_LDFLAGS
export CGO_LDFLAGS
