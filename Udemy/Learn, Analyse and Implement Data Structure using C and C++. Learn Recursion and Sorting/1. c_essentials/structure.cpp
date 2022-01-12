#include<iostream>
#include<stdio.h>

using namespace std;

struct Rectangle {
    int length;
    int right;
}

main() {


    struct Rectagle r;


    struct Rectangle r{10, 20};


    printf("Area of rectangle is %d", r.length * r.right);
}