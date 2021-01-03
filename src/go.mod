module github.com/jandedobbeleer/oh-my-posh3

go 1.15

require (
	github.com/Microsoft/go-winio v0.4.16
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/alecthomas/assert v0.0.0-20170929043011-405dbfeb8e38
	github.com/alecthomas/colour v0.1.0 // indirect
	github.com/alecthomas/repr v0.0.0-20201103221029-55c485bd663f // indirect
	github.com/distatus/battery v0.10.0
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/gookit/color v1.3.5
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/shirou/gopsutil v2.20.9+incompatible
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/stretchr/testify v1.6.1
	golang.org/x/sys v0.0.0-20201015000850-e3ed0017c211
	golang.org/x/text v0.3.3
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
	howett.net/plist v0.0.0-20200419221736-3b63eb3a43b5 // indirect
	muzzammil.xyz/jsonc v0.0.0-20201229145248-615b0916ca38
)

replace github.com/distatus/battery v0.10.0 => github.com/JanDeDobbeleer/battery v0.10.0-1

replace github.com/gookit/color v1.3.5 => github.com/JanDeDobbeleer/color v1.3.5-1
