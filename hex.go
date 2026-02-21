package main

const (
	limit = 14
)

var (
	start1 = []byte{0x48, 0x8B, 0x12}
	start2 = []byte{0x48, 0x8D, 0x0D}
	start3 = []byte{0x48, 0x8B, 0xD0}
	start4 = []byte{0x48, 0x8D, 0x0D}

	startLength = len(start1)

	end       = []byte{0x85, 0xC0, 0x0F, 0x94, 0xC3, 0xE8}
	endEU5    = []byte{0x85, 0xC0, 0x0F, 0x94, 0xC1, 0x88}
	endLength = len(end)

	replacement       = []byte{0x31, 0xC0, 0x0F, 0x94, 0xC3, 0xE8}
	replacementEU5    = []byte{0x31, 0xC0, 0x0F, 0x94, 0xC1, 0x88}
	replacementLength = len(replacement)

	replacementMapPE = map[string][]byte{
		eu4exe:  replacement,
		eu5exe:  replacementEU5,
		hoi4exe: replacement,
	}

	replacementMapELF = map[string][]byte{
		eu4bin:  replacement,
		eu5bin:  replacementEU5,
		hoi4bin: replacement,
	}

	// Native HOI4 ELF sequence where bytes [2:] are patched from:
	// 85 C0 0F 94 C3 E8 -> 31 C0 0F 94 C3 E8
	elfHOI4Needle = []byte{0x31, 0xDB, 0x85, 0xC0, 0x0F, 0x94, 0xC3, 0xE8}
)
