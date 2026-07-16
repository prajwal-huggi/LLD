#include<iostream>
using namespace std;

class Cart{
    Restaurant* r;
    vector<MenuItem> items;

    public:
    Cart(){
        this.r= nullptr;
    }

    void addItem(MenuItem& item){
        if(!r){
            cout<<"Select the restaurant"<<endl;
            return;
        }

        items.push_back(item);
    }

    double getTotalCost() const{
        int ans= 0;
        for(auto &i: items) ans+= i.getPrice();

        return ans;
    }

    bool isEmpty() {
        return (!restaurant || items.empty());
    }

    void clear() {
        items.clear();
        restaurant = nullptr;
    }

    void setRestaurant(Restaurant* r){
        this.r= r;
    }

    Restaurant* getRestaurant() const{
        return r;
    }

    const vector<MenuItem>& getItems() const{
        return menu;
    }
}
