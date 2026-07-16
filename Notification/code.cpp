#include<iostream>
#include<vector>
#include<mutex>
using namespace std;

class INotification{
    public:
    virtual string getContent() const= 0;
    virtual ~INotification(){}
};

class SimpleNotification: public INotification{
    private:
    string text;

    public:
    SimpleNotification(const string& msg){
        text= msg;
    }

    string getContent() const override{
        return text;
    }
};

class INotificationDecorator: public INotification{
    protected:
    INotification* notification;

    public:
    INotificationDecorator(INotification* n){
        notification= n;
    }

    virtual ~INotificationDecorator(){
        delete notification;
    }
};

class TimeStampDecorator: public INotificationDecorator{
    public:
    INotification* notification;
    string getContent() const override{
        return "timestamp: "+ notification-> getContent();
    }
};

class SignatureDecorator: public INotificationDecorator{
    private:
    string signature;

    public:
    SignatureDecorator(INotification* n, const string& sig): INotificationDecorator(n){
        signature= sig;
    }
    string getContent() const override{
        return notification-> getContent()+"\n"+ signature;
    }
};

/*============================
  Observer Pattern Components
=============================*/
class IObserver{
    public:
    virtual void update()= 0;
    virtual ~IObserver(){}
};

class IObservable{
    private:
    vector<IObserver> observers;
    
    public:
    virtual void add(IObserver* observer)= 0;
    virtual void remove(IObserver* observer)= 0;

    virtual void notify()= 0;
    virtual ~IObservable(){}
};

class NotificationObservable: public IObservable{
    private:
    vector<IObserver* > observers;
    INotification* currentNotification;

    public:
    void add(IObserver* observer){
        observers.push_back(observer);
    }
    void remove(IObserver* observer){
        observers.erase(remove(observers.begin(), observers.end(), observer), observers.end());
    }

};

// NotificationService which will be singleton
class NotificationService{
    private:
    static NotificationService* service;
    static mutex mtx;
    vector<INotification> notifications;

    NotificationService(){
        service= new NotificationService();
    }

    ~NotificationService()= delete;

    public:
    static NotificationService* getObject(){
        if(service== nullptr){
            lock_guard<mutex> lock(mtx);
            if(service== nullptr) service= new NotificationService();
        }

        return service;
    }
};

NotificationService* NotificationService:: service= nullptr;
mutex NotificationService:: mtx;

int main(){
    cout<<"Hello world";
    return 0;
}