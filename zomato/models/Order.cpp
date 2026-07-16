#include<vector>
using namespace std;

class Order{
    int id;
    User* user;
    Restaurant* restaurant;
    vector<MenuItem> items;
    PaymentStrategy paymentStrategy;
    double total;
    string getType;

    public:

};
