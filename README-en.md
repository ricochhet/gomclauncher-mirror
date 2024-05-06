# gomclauncher
![Go](https://github.com/xmdhs/gomclauncher/workflows/Go/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/xmdhs/gomclauncher)](https://goreportcard.com/report/github.com/xmdhs/gomclauncher)
A simple command line launcher for minecraft. Supports automatic download and validation of minecraft game files and genuine logins, supports launching fabric, quilt and forge after installing via the installer, supports linux windows and mac (mac not tested yet).
## Usage
Use `-h` to get instructions on how to use the relevant parameters.

Example `. /gml-linux -h`

Start the game `. /gml-linux -run 1.16.1 -username xmdhs`

Start the game with the specified java `. /gml-linux -run 1.16.1 -username xmdhs -javapath “. /java”`

Start the game and turn off detection of launcher update detection, game file validation, version isolation `. /gml-linux -run 1.16.1 -username xmdhs -test=f -independent=f -update=f`

First genuine login `. /gml-linux -run 1.16.1 -email example@example.com -password example`

Second time `. /gml-linux -run 1.16.1 -email example@example.com` The initiator does not save your password, it saves the accessToken for the next passwordless login.

Logging in with a Microsoft account `. /gml-linux -run 1.16.1 -email example@example.com -ms`

First external login `. /gml-linux -run 1.16.1 -email example@example.com -password example -yggdrasil example.com` Full api address is not required, the launcher will automatically complete it according to the protocol.

The second `. /gml-linux -run 1.16.1 -email example@example.com -yggdrasil example.com` 

View all saved licensed/external logins `. /gml-linux -list`

Delete saved external logins `. /gml-linux -email example@example.com -yggdrasil example.com -remove` 

Remove saved genuine logins `. /gml-linux -email example@example.com -remove` 

Remove saved genuine Microsoft logins `. /gml-linux -email example@example.com -ms -remove` 

Customize startup jvm parameters `. /gml-linux -run 1.16.1 -username xmdhs -flag “-XX:+AggressiveOpts -XX:+UseCompressedOops”`

Download the game and specify the mirror download source and set the number of concurrent processes used to 32 `. /gml-linux -downver 1.16.1 -type=bmclapi -int 32`.

Download the game and mix the two download sources `. /gml-linux -downver 1.16.1 -type “bmclapi|vanilla”`.

See all official versions available for download `. /gml-linux -verlist release`, `release` is the version type, which can be obtained with the following command.

View other optional release types `. /gml-linux -verlist ? `.

Remove unused files from the assets/objects folder `. /gml-linux -tidy `.

View the launcher version `. /gml-linux -v`

Translated with DeepL.com (free version)
## Screenshot
![image.png](https://i.loli.net/2020/07/02/E7ZcBCGfo1v46kI.png)

## Resources
BMCLAPI https://bmclapidoc.bangbang93.com/  
authlib-injector https://github.com/yushijinhun/authlib-injector  