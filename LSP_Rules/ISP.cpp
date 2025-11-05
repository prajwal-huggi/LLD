#include<iostream>
using namespace std;

class Shape{
    virtual double area()= 0;
    virtual double volume()= 0;
};

class Rectangle: public Shape{
    private:
    double side;

    public:
    Rectangle(double s): side(s){}
    double area() override {

    }
}

int main(){

    return 0;
}