package spi

type Display interface {
	Command(cmds ...byte)
	Data(data ...byte)

	Reset()

	Clear()
}
