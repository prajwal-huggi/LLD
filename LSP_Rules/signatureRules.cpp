// #include<bits/stdc++.h>
#include<iostream>
using namespace std;

class Parent{
    public:
    virtual void print(string s){
        cout<<"I am parent"<<endl;
    }
};

class Child: public Parent{
    public:
    void print(string s){
        cout<<"I am child class"<<endl;
    }

};

class Client{
    private:
    Parent *p;

    public:
    Client(Parent *p){
        this-> p= p;
    }

    void printMsg(){
        p->print("Hello");
    }
};

int main(){
    Parent *p= new Parent();
    Child *c= new Child();

    Client *client= new Client(p);

    client->printMsg();
    return 0;
}
