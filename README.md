# XDOT Helper


This is a library that contains easy to use functions to directly start communicating the [Multitech XDOT Devkit](https://www.multitech.com/brands/micro-xdot-devkit).

> The software is tested on [MTMDK-XDOT-EU1-A00](https://www.multitech.com/models/94558024LF) but is compatible with other models.

## AT Commands

The AT commands supported by the XDOT are available [here](https://www.multitech.com/documents/publications/manuals/s000643.pdf).

## Usage

This library exposes a [Command](./api/command.proto) api which can be used to build AT commands. 
These commands can be chained in a sequence and executed sequentially. The following is an example of a [registration](./pkg/templates/registration/registration.go) sequence.

```
	s := sequence.New(device, 1, true)
	s.AddCommand(pbapi.Command{Request: XDOTLoraATFactoryReset, WaitPeriod: 1, LinesInResponse: 0})
	s.AddCommand(pbapi.Command{Request: XDOTLoraATWriteAppEUI + appEUI, WaitPeriod: 1, LinesInResponse: 1})
	s.AddCommand(pbapi.Command{Request: XDOTLoraATWriteNwkKey + nwkKey, WaitPeriod: 1, LinesInResponse: 1})
	s.AddCommand(pbapi.Command{Request: XDOTLoraATSaveConfig, WaitPeriod: 1, LinesInResponse: 0})
```

This sequence can simple be executed using `Sequence.Execute()`.

## License
The contents of this repository are released `as-is` under the terms of the [APACHE 2.0 License](LICENSE). 