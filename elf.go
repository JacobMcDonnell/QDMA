package main

import (
	"encoding/binary"
)

const (
	EI_NIDENT  = 16
	ET_NONE    = 0
	ET_REL     = 1
	ET_EXEC    = 2
	ET_DYN     = 3
	ET_CORE    = 4
	MIPS       = 0x08
	EHSIZE     = 52
	PHENSIZE   = 0x20
	SHENTSIZE  = 0x28
	PT_NULL    = 0
	PT_LOAD    = 1
	PT_DYNAMIC = 2
	PT_INTERP  = 3
	PT_NOTE    = 4
	PT_SHLIB   = 5
	PT_PHDR    = 6
	PT_TLS     = 7
	PF_READ    = 4
	PF_WRITE   = 2
	PF_EXEC    = 1
)

type ElfHeader struct {
	e_ident     [EI_NIDENT]byte
	e_type      uint16
	e_machine   uint16
	e_version   uint32
	e_entry     uint32
	e_phoff     uint32
	e_shoff     uint32
	e_flags     uint32
	e_ehsize    uint16
	e_phentsize uint16
	e_phnum     uint16
	e_shentsize uint16
	e_shnum     uint16
	e_shstrndx  uint16
}

type ProgramHeader struct {
	ptype  uint32
	offset uint32
	vaddr  uint32
	paddr  uint32
	filesz uint32
	memsz  uint32
	flags  uint32
	align  uint32
}

func EHInit(e_entry uint32, e_phnum uint16) ElfHeader {
	var e ElfHeader
	e.e_ident = [EI_NIDENT]byte{0x7f, 0x45, 0x4c, 0x46, // ELF magic number
		0x01,                                     // 32-bit format
		0x01,                                     // Little Endian
		0x01,                                     // ELF Version
		0x00,                                     // Target OS
		0x00,                                     // Target ABI
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // Padding
	e.e_type = ET_EXEC
	e.e_machine = MIPS
	e.e_version = 0x01
	e.e_entry = e_entry
	e.e_phoff = EHSIZE // Program headers will be after ELF Header
	e.e_shoff = 0
	e.e_ehsize = EHSIZE
	e.e_phentsize = PHENSIZE
	e.e_phnum = e_phnum
	e.e_shentsize = SHENTSIZE
	e.e_shnum = 0
	e.e_shstrndx = 0
	return e
}

/* NOTES: On the ToBytes methods
 * I understand that these two methods are not the prettiest, but I need
 * something to convert these structs to to byte arrays. From what I've read so
 * far there seems to be no way to do this. It appears that encoding/gob
 * ignores the size of data types, which is needed, and json serialization
 * would not produce the proper output.
 */

func (e *ElfHeader) ToBytes() []byte {
	bytes := make([]byte, EHSIZE)
	i := 0
	for _, b := range e.e_ident {
		bytes[i] = b
		i++
	}
	binary.NativeEndian.PutUint16(bytes[EI_NIDENT:], e.e_type)
	binary.NativeEndian.PutUint16(bytes[18:], e.e_machine)
	binary.NativeEndian.PutUint32(bytes[20:], e.e_version)
	binary.NativeEndian.PutUint32(bytes[24:], e.e_entry)
	binary.NativeEndian.PutUint32(bytes[28:], e.e_phoff)
	binary.NativeEndian.PutUint32(bytes[32:], e.e_shoff)
	binary.NativeEndian.PutUint32(bytes[36:], e.e_flags)
	binary.NativeEndian.PutUint16(bytes[40:], e.e_ehsize)
	binary.NativeEndian.PutUint16(bytes[42:], e.e_phentsize)
	binary.NativeEndian.PutUint16(bytes[44:], e.e_phnum)
	binary.NativeEndian.PutUint16(bytes[46:], e.e_shentsize)
	binary.NativeEndian.PutUint16(bytes[48:], e.e_shnum)
	binary.NativeEndian.PutUint16(bytes[50:], e.e_shstrndx)
	return bytes
}

func (p *ProgramHeader) ToBytes() []byte {
	bytes := make([]byte, PHENSIZE)
	binary.NativeEndian.PutUint32(bytes[0:], p.ptype)
	binary.NativeEndian.PutUint32(bytes[4:], p.offset)
	binary.NativeEndian.PutUint32(bytes[8:], p.vaddr)
	binary.NativeEndian.PutUint32(bytes[12:], p.paddr)
	binary.NativeEndian.PutUint32(bytes[16:], p.filesz)
	binary.NativeEndian.PutUint32(bytes[20:], p.memsz)
	binary.NativeEndian.PutUint32(bytes[24:], p.flags)
	binary.NativeEndian.PutUint32(bytes[28:], p.align)
	return bytes
}
