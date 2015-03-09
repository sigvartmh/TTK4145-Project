
#include "comedilib.h"

int comedi_data_read(comedi_t *it,unsigned int subd,unsigned int chan,
    unsigned int range,unsigned int aref,lsampl_t *SWIG_OUTPUT(data)){
    return 1;
};
int comedi_data_write(comedi_t *it,unsigned int subd,unsigned int chan,
    unsigned int range,unsigned int aref,lsampl_t data){
    return 1;
};
int comedi_dio_config(comedi_t *it,unsigned int subd,unsigned int chan,
    unsigned int dir){
    return 1;
};
int comedi_dio_read(comedi_t *it,unsigned int subd,unsigned int chan,
    unsigned int *SWIG_OUTPUT(bit)){
    return 1;
};
int comedi_dio_write(comedi_t *it,unsigned int subd,unsigned int chan,
    unsigned int bit){
    return 1;
};
comedi_t *comedi_open(const char *fn){
    return NULL;
};
