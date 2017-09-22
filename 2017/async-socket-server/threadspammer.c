#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

void* threadfunc(void* p) {
  printf("%ld: thread 0x%0x\n", (long)p, (unsigned)pthread_self());
  while (1) {
    usleep(50 * 1000);
  }
  return NULL;
}

int main(int argc, const char** argv) {
  int nthreads = 10;
  if (argc > 1) {
    nthreads = atoi(argv[1]);
  }
  printf("Running with nthreads = %d\n", nthreads);

  for (long i = 0; i < nthreads; ++i) {
    pthread_t t;

    pthread_create(&t, NULL, threadfunc, (void*)i);
    usleep(1000);
  }

  printf("... waiting ... \n");
  while (1) {
    usleep(200 * 1000);
  }

  return 0;
}
