package main

import (
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		tempData, err := os.CreateTemp("", "data*")
		Check(err)
		dName := tempData.Name()
		defer os.Remove(dName)

		tempRoData, err := os.CreateTemp("", "rodata*")
		Check(err)
		rName := tempRoData.Name()
		defer os.Remove(rName)

		tempText, err := os.CreateTemp("", "text*")
		Check(err)
		tName := tempText.Name()
		defer os.Remove(tName)

		// First pass to find the address of all labels and encode the data
		LabelFind(arg, tempData, tempRoData)

		var names [3]string
		var PHeaders [4]ProgramHeader
		i := 0
		n := 0

		if SectionPos[TEXT] > 0 {
			PHeaders[i] = ProgramHeader{PT_LOAD, 0, 0, 0,
				uint32(SectionPos[TEXT]), uint32(SectionPos[TEXT]), PF_READ +
					PF_EXEC, 0}
			names[i] = tName
			i++
			n++
			// Second pass to assemble the instructions.
			Assemble(arg, tempText)
		}

		if SectionPos[DATA] > 0 {
			PHeaders[i] = ProgramHeader{PT_LOAD, 0, 0, 0,
				uint32(SectionPos[DATA]), uint32(SectionPos[DATA]), PF_READ +
					PF_WRITE, 0}
			names[i] = dName
			i++
			n++
		}

		if SectionPos[RODATA] > 0 {
			PHeaders[i] = ProgramHeader{PT_LOAD, 0, 0, 0,
				uint32(SectionPos[RODATA]), uint32(SectionPos[RODATA]), PF_READ, 0}
			names[i] = rName
			i++
			n++
		}

		if SectionPos[BSS] > 0 {
			PHeaders[i] = ProgramHeader{PT_LOAD, 0, 0, 0, 0,
				uint32(SectionPos[BSS]), PF_READ + PF_WRITE, 0}
			i++
		}

		var fOffset, mOffset uint32 = 0, 0
		for j, p := range PHeaders[:i] {
			PHeaders[j].offset = uint32(i*PHENSIZE) + fOffset + EHSIZE
			fOffset += p.filesz
			PHeaders[j].vaddr = mOffset
			mOffset += p.memsz
		}

		EHeader := EHInit(0, uint16(i))
		WriteBinary(arg, EHeader, PHeaders[:i], names[:n])
	}
}

func WriteBinary(path string, e ElfHeader, ps []ProgramHeader, ns []string) {
	out, err := os.Create(strings.ReplaceAll(path, ".asm", ".bin"))
	Check(err)
	defer out.Close()
	_, err = out.Write(e.ToBytes())
	Check(err)
	for _, p := range ps {
		_, err = out.Write(p.ToBytes())
		Check(err)
	}
	for _, name := range ns {
		b, err := os.ReadFile(name)
		Check(err)
		_, err = out.Write(b)
		Check(err)
	}
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
