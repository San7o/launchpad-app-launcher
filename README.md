# launchpad-app-launcher

I created an application launcher utilizing my Novation Launchpad S MIDI controller. This launcher is a custom program designed to enhance my workflow by allowing me to quickly access and execute various applications, tools, or commands through the physical interface of the Launchpad S.

![image](https://github.com/San7o/launchpad-app-launcher/assets/81651399/a48c857e-4c1a-448b-96a4-3ef03d189e07)

# Current Buttons
With `0,0` being bottom left, and `8, 8` top right:


### Poweroff
- `0, 8`: shutdown driver


## Terminal Appliactions
- `0, 7`: ranger

- `1, 7`: kitty

- `2, 7`: nvidia-smi -l

- `3, 7`: btop

### Desktop Applications

- `0, 6`: vivaldi

## Run the Application
You will need go to run this, then just run
```bash
./build
```
And violà

All the commands are configured for my personal linux configuration, many commands might not work for you but the code is easily customizable and readable


