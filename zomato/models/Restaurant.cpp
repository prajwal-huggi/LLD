#include<string>
#include<vector>
#include<iostream>

using namespace std;

class Restaurant{
    static int nextId;
    int restaurantId;
    string name;
    string loc;
    vector<MenuItem> menu;

    public:
    Restaurant(const string& name, const string& loc){
        this-> name= name;
        this-> loc= loc;
        this-> restaurantId=++nextId;

    }

    ~Restaurant() {
        // Optional: just for clarity or debug
        cout << "Destroying Restaurant: " << name << ", and clearing its menu." << endl;
        menu.clear();
    }

    string getName()const{
        return name;
    }

    void setName(const string& name){
        this-> name= name;
    }

    string getLoc()const{
        return loc;
    }

    void setLoc(const string& loc){
        this-> loc= loc;
    }

    const vector<MenuItem>& getMenu()const {
        return menu;
    }

    void setMenu(const MenuItem &item){
        menu.push_back(item);
    }

};

int Restaurant::nextId= 0;
