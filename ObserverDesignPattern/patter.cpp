#include<iostream>
#include<vector>
#include<algorithm>
using namespace std;

class ISubscriber{
    public:
    virtual void update()= 0;
    virtual ~ISubscriber(){}
};

class IChannel{
    public:
    virtual void subscribe(ISubscriber* s)= 0;
    virtual void unsubscribe(ISubscriber* s)= 0;
    virtual void notify()= 0;
    virtual ~IChannel(){}
};

class Channel: public IChannel{
    vector<ISubscriber*> subscribers;
    string name;
    string latestVideo;

    public:
    Channel(const string& name){
        this-> name= name;
    }
    void subscribe(ISubscriber* s) override{
        if(find(subscribers.begin(), subscribers.end(), s)== subscribers.end())
        subscribers.push_back(s);
    }

    void unsubscribe(ISubscriber* s)override{
        auto it= find(subscribers.begin(), subscribers.end(), s);
        if (it != subscribers.end())
            subscribers.erase(it);
        // if(find(subscribers.begin(), subscribers.end(), s)!= subscirbers.end())
        // subscribers.erase(s);
    }

    void notify()override{
        for(auto &i: subscribers){
            i->update();
        }
    }

    void addVideo(const string& title){
        this->latestVideo= title;
        cout<<"Name: "<<name<<" latestVideo: "<<this->latestVideo<<endl;
        notify();
    }

    string getVideoData() const{
        // cout<<"Latest Videos is: "<<latestVideo<<endl;
        return latestVideo;
    }
};

class Subscriber: public ISubscriber{
    Channel* channel;
    string name;
    public:
    Subscriber(const string& name, Channel* channel){
        this-> name= name;
        this-> channel= channel;
    }

    void update() override{
        cout<<"Hey! "<<name<<" Checkout the new video "<<this->channel->getVideoData()<<endl;
    }
};

int main(){
    Channel* beast= new Channel("beast");

    Subscriber* subscribe1= new Subscriber("subscriber1", beast);
    Subscriber* subscribe2= new Subscriber("subscriber2", beast);

    beast->subscribe(subscribe1);
    beast->subscribe(subscribe2);

    beast->addVideo("Video1");

    beast-> unsubscribe(subscribe1);
    beast-> addVideo("Video2");
    
    return 0;
}
