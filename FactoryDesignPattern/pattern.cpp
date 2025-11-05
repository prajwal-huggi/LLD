#include<iostream>
using namespace std;

class Burger{
    public:
    virtual void prepare()= 0;
    virtual ~Burger(){}
};

class BasicBurger: public Burger{
    void prepare() override{
        cout<<"Preparing the Basic Burger"<<endl;
    }
};

class StandardBurger: public Burger{
    void prepare() override{
        cout<<"Preparing the Standard Burger"<<endl;
    }
};

class PremiumBurger: public Burger{
    void prepare() override{
        cout<<"Preparing the Premium Burger"<<endl;
    }
};

class BurgerFactory{
    public:
    Burger* createBurger(string& type){
        if(type== "basic")
            return new BasicBurger();
        else if(type== "standard")
            return new StandardBurger();
        else if(type== "premium")
            return new PremiumBurger();
        else
            return nullptr;
    }
};

int main(){
    string type= "standard";

    BurgerFactory *bf= new BurgerFactory();

    Burger* burger= bf-> createBurger(type);

    burger-> prepare();

    return 0;
}