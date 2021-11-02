# FAST-discord-token-checker
 ✨ Fastest Feature-packed Discord Token Checker written in GO ✨
 
 ## Overview 🔍
 This program is the fastest ever written Discord token checker which can validate login for discord authorization tokens from a list. This was written so people with multiple bot tokens can validate them quickly and efficiently! It simply sends requests to discord to determine if the token is valid, it also captures most attributes linked to the token [Email address, verification status, Phone number, Nitro Status, etc]. It can get upto 20,000 CPM (Checks per minute) Without using any Proxies! 
 
 ![Preview of the token checker](https://i.imgur.com/010UnEX.png)
 
 ## Disclaimer ⚠️
 The automation of User Discord accounts also known as self-bots is a violation of Discord Terms of Service & Community guidelines and will result in your account(s) being terminated. Discretion is adviced. I will not be responsible for your actions. 
 
 ## Features ✅
 - Fastest account checker made possibly by GO's Goroutines and efficient requests [20K+ CPM]
 - Proxyless 
 - Full Capture (UID/Username/Discriminator/Bio/Location & Language/MFA status/Email/Verification Status/ Phone number/ Nitro)
 - Free & Open source
 - Compatible with all major Operating systems and CPU architectures 

## Usage 💻
 - Build from Source or Download from [releases](https://github.com/V4NSH4J/FAST-discord-token-checker/releases)
 - Input your tokens in "tokens.txt"
 - Decide what to capture by setting your config file "config.json"
 - Run the binary
 - Working tokens will be outputted to "working.txt" and the tokens with capture would be outputted to "capture.txt" if enabled.
 
 ## Building from Source 🚧
 - [Install Golang](https://golang.org) and verify your installation
 - Open up a terminal window 
 - Navigate to the directory of the source code
 - Type "go build" into your console and a Binary should pop up
 
 ## Commonly Asked Question : Where is the Token Generator ❓ 
 There are a large number of programs and people claiming to be Token Generators and Checkers. You can try them out, rent servers and have millions of checks performed per minute but here's the simple answer, they will NEVER work. I would even go as far as saying that finding someone's token from "Generating and checking" is far far far unlikely than winning a national lottery multiple times consecutively. Read along the next section to understand what a Discord token is.
 
 ## Understanding Discord Tokens 🧠
 Discord account tokens (User accounts/ Bots) usually look something like this -> 

![Notepad_N0DhRyXvJl](https://user-images.githubusercontent.com/79518089/136600295-8968e59c-5dc1-487b-83fd-5f176e710bbe.png)

Notice that the token has 2 dots which divides the token into 3 segments. The first part of the token, before the first dot is a discord User ID or UID encoded in base 64. The discord UID is a "Snowflake" which is a uint64. The second part of the token also seems like some sort of a count in base64 and is often the same for tokens generated nearly at the same time. While the third part seems completely random for a unique token. These UIDs are unique numbers comprising of the time of action (time of creation in case of discord tokens) and some internal process IDs. You may not know it, but snowflakes are everywhere on discord! User IDs, Channel IDs, Server IDs, Message IDs, everything's discord snowflakes and hence you can find their time of creation. These snowflakes are also sent under the nonce field in requests while interacting with discord like sending messages or clicking buttons and are used for validating the action. 

![94722171abc49573d1a129e2264da4ad](https://user-images.githubusercontent.com/79518089/136601218-6f08cd18-4f15-4274-834f-77093f774382.png)

This image from discord API documentation very well explains how a snowflake can be converted to give the timestamp of creation. So in essence, a discord token is just an account's creation date, some internal process IDs and a bunch of random stuff thrown together. If someone does try to make a token generator for fun, the generated tokens should match this syntax or it can never give a valid token ever.
 
## Other interesting stuff by me
(Discord Invite Joiner)[https://github.com/V4NSH4J/discord-inviter-GO] - Joins given tokens to a server
(Discord Token Checker)[https://github.com/V4NSH4J/FAST-discord-token-checker] - Checks given tokens and records their information
(Discord Mass DM)[https://github.com/V4NSH4J/discord-mass-DM-GO] - DMs all users of a server or DM's a discord user from multiple accounts
(Dankgrinder)[https://github.com/V4NSH4J/dankgrinder] - An Advanced automation tool for Dankmemer
 
 ## Donations 🪙
I spend quite a lot of time in making High Quality & Open Source discord tools because hundreds of people get ripped-off everyday searching for this stuff. If this helped you out even in the slightest, Buy me a coffee and make my day! 
BTC: bc1qfmk95sqtw6sw2xc3kyaemcnltwcr5cs2phg2gh


 
 
 
 
