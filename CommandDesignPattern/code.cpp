#include <iostream>
#include <vector>
#include <memory>
using namespace std;

class ICommand {
public:
    virtual void execute() = 0;
    virtual void undo() = 0;
    virtual ~ICommand() = default;
};

class NoCommand : public ICommand {
public:
    void execute() override {}
    void undo() override {}
};

class RemoteControl {
    vector<unique_ptr<ICommand>> commands;
    vector<bool> isPressed;

public:
    RemoteControl(int buttons) {
        commands.resize(buttons);
        isPressed.resize(buttons, false);

        for (auto &cmd : commands) {
            cmd = make_unique<NoCommand>();
        }
    }

    void setCommand(unique_ptr<ICommand> command, int idx) {
        commands[idx] = std::move(command);
    }

    void buttonPressed(int idx) {
        if (!commands[idx]) return;

        if (!isPressed[idx]) {
            commands[idx]->execute();
            isPressed[idx] = true;
        } else {
            commands[idx]->undo();
            isPressed[idx] = false;
        }
    }
};

class Light {
public:
    void on()  { cout << "The light is switched on" << endl; }
    void off() { cout << "The light is switched off" << endl; }
};

class LightCommand : public ICommand {
    Light* light;
public:
    LightCommand(Light* light) : light(light) {}
    void execute() override { light->on(); }
    void undo() override { light->off(); }
};

class Fan {
public:
    void on()  { cout << "The fan is on" << endl; }
    void off() { cout << "The fan is off" << endl; }
};

class FanCommand : public ICommand {
    Fan* fan;
public:
    FanCommand(Fan* fan) : fan(fan) {}
    void execute() override { fan->on(); }
    void undo() override { fan->off(); }
};

int main() {
    RemoteControl remote(5);

    Light bedroomLight;
    Fan bedroomFan;

    remote.setCommand(make_unique<FanCommand>(&bedroomFan), 1);
    remote.setCommand(make_unique<LightCommand>(&bedroomLight), 2);

    cout << "Testing light" << endl;
    remote.buttonPressed(2);
    remote.buttonPressed(2);

    cout << "Testing fan" << endl;
    remote.buttonPressed(1);
    remote.buttonPressed(1);

    return 0;
}
