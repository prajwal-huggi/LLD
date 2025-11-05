#include<iostream>
using namespace std;

class Talkable{
    public:
    virtual void talk() = 0;
    virtual ~Talkable(){}
};

class NormalTalk: public Talkable{
    public:
    void talk() override{
        cout<<"Talking Normally"<<endl;
    }
};

class NoTalk: public Talkable{
    void talk() override{
        cout<<"Unable to talk"<<endl;
    }
};

class Flyable{
    public:
    virtual void fly() = 0;
    virtual ~Flyable(){}
};

class NormalFly: public Flyable{
    public:
    void fly() override{
        cout<<"Normal fly"<<endl;
    }
};

class NoFly: public Flyable{
    public:
    void fly() override{
        cout<<"Unable to Fly"<<endl;
    }
};

class Robot{
    Talkable* t;
    Flyable* f;
    
    public:
    Robot(Talkable* t, Flyable* f){
        this->t= t;
        this-> f= f;
    }

    void fly(){
        f-> fly();
    }

    void talk(){
        t-> talk();
    }

    virtual void projection()= 0;
};

class CompanionRobot: public Robot{
    public:
    CompanionRobot(Talkable* t, Flyable* f): Robot(t, f){}

    void projection(){
        cout<<"Projecting the companion robot"<<endl;
    }
};

int main(){

    Robot *robo1= new CompanionRobot(new NormalTalk(), new NoFly());
    robo1-> talk();
    robo1-> fly();

    return 0;
}
