#include<iostream>
using namespace std;

// abstract Burger
class Burger{
    public:
    virtual void prepare()= 0;
    virtual ~Burger(){}
};

//Created 6 types of burger
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

class BasicWheatBurger: public Burger{
    void prepare() override{
        cout<<"Preparing the Basic Wheat Burger"<<endl;
    }
};

class StandardWheatBurger: public Burger{
    void prepare() override{
        cout<<"Preparing the Standard Wheat Burger"<<endl;
    }
};

class PremiumWheatBurger: public Burger{
    void prepare() override{
        cout<<"Preparing the Premium Wheat Burger"<<endl;
    }
};

//2-> Second product
class GarlicBread{
    public:
    virtual void prepare()= 0;
    virtual ~GarlicBread(){}
};

class NormalGarlicBread: public GarlicBread{
    void prepare() override{
        cout<<"Preparing normal garlic bread"<<endl;
    }
};
class PremiumGarlicBread: public GarlicBread{
    void prepare() override{
        cout<<"Preparing premium garlic bread"<<endl;
    }
};

// Creating the main factory
class MainFactory{
    public:
    virtual Burger* parepareBurger(string &type)= 0;
    virtual GarlicBread* prepareGarlicBread(string &type)= 0; 
};

class SinghBurger: public MainFactory{
    Burger* parepareBurger(string& type) override {
        if (type == "basic") {
            return new BasicBurger();
        } else if (type == "standard") {
            return new StandardBurger();
        } else if (type == "premium") {
            return new PremiumBurger();
        } else {
            cout << "Invalid burger type! " << endl;
            return nullptr;
        }
    }

    GarlicBread* prepareGarlicBread(string& type) override {
        if (type == "basic") {
            return new NormalGarlicBread();
        } 
        else {
            cout << "Invalid Garlic bread type! " << endl;
            return nullptr;
        }
    }
};

class KingBurger: public MainFactory{
    Burger* parepareBurger(string& type) override {
        if (type == "basic") {
            return new BasicWheatBurger();
        } else if (type == "standard") {
            return new StandardWheatBurger();
        } else if (type == "premium") {
            return new PremiumWheatBurger();
        } else {
            cout << "Invalid burger type! " << endl;
            return nullptr;
        }
    }

    GarlicBread* prepareGarlicBread(string& type) override {
        if (type == "premium") {
            return new PremiumGarlicBread();
        } 
        else {
            cout << "Invalid Garlic bread type! " << endl;
            return nullptr;
        }
    }

};

int main(){
    string burgerType = "basic";
    string garlicBreadType = "premium";

    MainFactory* mealFactory = new KingBurger();

    Burger* burger = mealFactory->parepareBurger(burgerType);
    GarlicBread* garlicBread = mealFactory->prepareGarlicBread(garlicBreadType);

    burger->prepare();
    garlicBread->prepare();

    return 0;
}