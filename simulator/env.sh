export FLOORS=3
libPath=$PWD
CGO_LDFLAGS="$libPath/driver/lib/simulation_elevator.a $libPath/driver/lib/libphobos2.a -lpthread -lcomedi -lm"
echo $CGO_LDFLAGS
export CGO_LDFLAGS
