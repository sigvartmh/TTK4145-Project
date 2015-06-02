#Compiler type
CC=dmd
LDFLAGS= -lib -ofsimulation_elevator

all: simulator

simulator:
	$(CC) simulator_interface.d
	mv simulator_interface ../
	$(CC) simulation_elevator.d timer_event.d $(LDFLAGS)
	mv simulation_elevator.a ../

clean:
	rm *o
