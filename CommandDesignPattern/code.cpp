#include<iostream>
#include<vector>
using namespace std;

class ICommand{
    public:
    virtual void execute()= 0;
    virtual void undo()= 0;
    virtual ~ICommand(){}
};

class RemoteControl{
    vector<ICommand*> commands;
    vector<bool> isPressed;

    public:
    RemoteControl(int buttons){
        commands.resize(buttons, nullptr);
        isPressed.resize(buttons, false);
    }
    void setCommand(ICommand* command, int idx){
        commands[idx]= command;
    }

    void buttonPressed(int idx){
        if(!isPressed[idx]){
            isPressed[idx]= !isPressed[idx];
            commands[idx]-> execute();
        }
        else {
            isPressed[idx]= !isPressed[idx];
            commands[idx]-> undo();
        }
    }
};

class Light{
    public:
    void on(){
        cout<<"The light is switched on"<<endl;
    }
    void off(){
        cout<<"The light is switched off"<<endl;
    }
};

class LightCommand: public ICommand{
    Light* light;
    public:
    LightCommand(Light* light){
        this-> light= light;
    }
    void execute(){
        light-> on();
    }
    void undo(){
        light-> off();
    }
};

class Fan{
    public:
    void on(){

        cout<<"The fan is on"<<endl;
    }
    void off(){
        cout<<"The fan is off"<<endl;
    }
};

class FanCommand: public ICommand{
    Fan* fan;

    public:
    FanCommand(Fan* fan){
        this-> fan= fan;
    }
    void execute(){
        fan-> on();
    }
    void undo(){
        fan-> off();
    }
};

int main(){

    RemoteControl* remote= new RemoteControl(5);
    
    Light* bedroomLight= new Light();
    Fan* bedroomFan= new Fan();

    remote->setCommand(new FanCommand(bedroomFan), 1); //fan
    remote->setCommand(new LightCommand(bedroomLight), 2); //light

    cout<<"Testing light"<<endl;
    remote->buttonPressed(2);
    remote->buttonPressed(2);

    cout<<"Testing fan"<<endl;
    remote->buttonPressed(1);
    remote->buttonPressed(1);

    return 0;
}