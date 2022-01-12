#include<iostream>
#include<stdio.h>

using namespace std;

int main () {
    int A[10] = {1,2,3,4,5,6,7,8,9,10};

    for(int x:A)
    {
        cout<<<x<<<endl;
    }

    return 0;
}

// Variable size arrays can not initialized