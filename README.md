# 6cord

## Screenshots

#### With the channel browser

![](http://u.cubeupload.com/diamondburned/i804yI.png)

#### Hidden channel browser

![](http://u.cubeupload.com/diamondburned/1jGJNU.png)

#### Fuzzy commands

![](http://u.cubeupload.com/diamondburned/H86qdB.png)

#### Fuzzzy emojis

![](http://u.cubeupload.com/diamondburned/dX90YD.png)

## Installation

### 1. [From CI (only when the tick-mark is green)](https://gitlab.com/diamondburned/6cord/builds/artifacts/master/file/6cord?job=compile)

### 2. `go get -u gitlab.com/diamondburned/6cord`

## Behaviors

- From input, hit arrow up to go to autocompletion. Arrow up again to go to the message box.
- In the message box
  - Arrow up/down and Page up/down will be used for scrolling
  - Any other key focuses back to input
- Tab to hide channels, focusing on input
- Tab again to show channels, focusing on the channel list
- To clear the keyring, feed `6cord` a new token with `-t`

## Todo

- [ ] [Fix paste not working](https://github.com/rivo/tview/issues/133) (workaround: Ctrl + V)
    - [x] Better paste library with image support (Linux only)
- [x] Syntax highlighting, better markdown parsing
- [x] Message Delete and Edit
- [x] Full reaction support
- [ ] Typing events
	- [x] The client sends the typing event
	- [ ] The client receives and indicates typing events
- [x] Commands
    - [x] `/goto`
    - [x] `/edit`
    - [x] `s//` with regexp
    - [x] `/exit`, `/shrug`
    - [x] Autocompletion for those commands
	- (refer to the screenshot)
- [x] Use TextView instead of List for Messages
	- [x] Consider tv.Write() bytes
	- [x] Proper inline message edit renders
	- ~~Split messages into Primitives or find a way to edit them individually (cordless does this, too much effort)~~
- [x] Fetch nicknames and colors (16-bit hex to 256 cols somehow...)
	- [x] Async should be for later, when Split messages is done
	- [x] Add a user store
- [ ] Implement embed SIXEL images
    - [ ] Port library to [termui](https://github.com/gizak/termui)
    - [ ] Work on [issue #213](https://github.com/gizak/termui/issues/213)
- [x] Implement inline emojis
- [x] Implement auto-completion popups
	- Behavior: all keys except Enter and Esc belongs to the Input Field
	- Esc closes the popup, Enter puts the popup content into the box
	- When 0 results, hide dialog
	- Show dialog when: `@`, `#` and potentially `:` (`:` is pointless as I don't plan on adding emoji inputs any time soon)
	- Auto-completed items:
    	- Mentions `@`
    	- Stock emojis `:`
		- Commands `/`
		- Channels `#`
		- Messages `~`
- [x] An actual channel browser
- [x] Message acknowledgements (read/unread)
	- Isn't fully working yet, channel overrides are still janky
- [ ] Message mentions
	- Partially working (only counts future mentions)
	- Todo: past mentions using the endpoint
- [x] Scrolling up gets more messages
- [ ] Port current user stores into only Discord state caches
- [ ] Voice support (partially atm)
	- [x] Show who's in, muted, deafened and ignored
	- [ ] [Actual microphone handling](https://github.com/gordonklaus/portaudio/blob/master/examples/record.go)
	- [ ] [Auto volume](https://dsp.stackexchange.com/questions/46147/how-to-get-the-volume-level-from-pcm-audio-data)
		- Basically, I need to time so that an array of PCM int16s will contain data for 400ms
		- Then, I'll need to either root-mean-square it or calculate decibels 
		- Finally, I will compare the calculated value to the one in `config.go`
		- If it's louder, send it over to the buffer
	- ~~Keyboard event handling~~
- [x] Fix `discordgo` spasming out when a goroutine panics
	- A solution could be `./6cord 2> /dev/null`
- [ ] ~~Confirm Windows compatibility~~
	- `/upload` fuzzy match doesn't work, wontfix

## Credits

- XTerm from 
	- https://invisible-island.net/xterm/
	- https://gist.github.com/saitoha/7822989
- Fishy ([RumbleFrog](https://github.com/rumblefrog)) for his
	- [discordgo fork](https://github.com/rumblefrog/discordgo)
	- [Channel sort lib ~~that he stole from my shittercord~~](https://gist.github.com/rumblefrog/c9ebd9fb84a8955495d4fb7983345530)
- Some people on unixporn and nix nest (ym555, tdeo, ...)
- [cordless](https://github.com/Bios-Marcel/cordless) [author](https://github.com/Bios-Marcel) for some of the functions

