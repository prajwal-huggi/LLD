#include<iostream>
#include<mutex>// Using this class inorder to keep the code thread free
using namespace std;

class SingleTon{
    private:
    static SingleTon* instance;
    static mutex mtx;

    SingleTon(){
        cout<<"Object Created"<<endl;
    }
    ~SingleTon()= delete;

    public:
    static SingleTon* getObject(){
        if(instance== nullptr){
            lock_guard<mutex> lock(mtx);
            if(instance== nullptr) instance= new SingleTon();
        }
        return instance;
    }
};

SingleTon* SingleTon::instance= nullptr;
mutex SingleTon:: mtx;

int main(){
    SingleTon* s1= SingleTon:: getObject();
    SingleTon* s2= SingleTon:: getObject();

    cout<<(s1== s2)<<endl;
    return 0;
}
