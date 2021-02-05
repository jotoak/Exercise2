#include <pthread.h>
#include <stdio.h>
//I think that the mutex look is best when it it limetes the acess to the varible. And no wating thread is posted.
int i = 0;
pthread_mutex_t lock;
// Note the return type: void*
void* incrementingThreadFunction(){
    pthread_mutex_lock(&lock);
    for (int j = 0; j < 1000000; j++) {
	// TODO: sync access to i
	i++;
    }
    pthread_mutex_unlock(&lock);
    return NULL;
}

void* decrementingThreadFunction(){
    pthread_mutex_lock(&lock);
    for (int j = 0; j < 1000000; j++) {
	// TODO: sync access to i
	i--;
    }
    pthread_mutex_unlock(&lock);
    return NULL;
}


int main(){
    if (pthread_mutex_init(&lock, NULL) != 0)
    {
        printf("\n mutex init failed\n");
        return 1;
    }
    pthread_t incrementingThread, decrementingThread;

    pthread_create(&incrementingThread, NULL, incrementingThreadFunction, NULL);
    pthread_create(&decrementingThread, NULL, decrementingThreadFunction, NULL);
    
    pthread_join(incrementingThread, NULL);
    pthread_join(decrementingThread, NULL);
    pthread_mutex_destroy(&lock);
    printf("The magic number is: %d\n", i);
    return 0;
}
