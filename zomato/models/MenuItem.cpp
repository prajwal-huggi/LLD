class MenuItem{
    string code;
    string name;
    double price;

    public:
    MenuItem(const string& code, const string& name, const double& price){
        this-> code= code;
        this-> name= name;
        this-> price= price;
    }

    //getter and setter
    string getCode() const{
        return code;
    }

    void setCode(const string& c){
        this-> code= code;
    }

    string getName() const {
        return name;
    }

    void setName(const string &n) {
        name = n;
    }

    int getPrice() const {
        return price;
    }

    void setPrice(int p) {
        price = p;
    }
};
