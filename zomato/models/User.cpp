#include<iostream>
using namespace std;

class User{
    int id;
    string name;
    string addr;
    Cart cart;

    public:
    User(int userId, const string& name, const string& address) {
        this->id = userId;
        this->name = name;
        this->addr = address;
        cart = new Cart();
    }

    ~User(){
        delete cart;
    }

    const string getName() const {
        return name;
    }

    void setName(const string name){
        this-> name= name;
    }

    const string getAddr() const{
        return this-> addr;
    }

    void setAddr(const string& addr){
        this-> addr= addr;
    }

    Cart* getCart() const{
        return cart;
    }
};
