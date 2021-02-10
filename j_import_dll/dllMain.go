package main

// 10. golang中通过使用dll中的方法

import (
	"fmt"
	"golang.org/x/sys/windows"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

const MAX_PATH = 260

type HANDLE uintptr
type HWND HANDLE
type BROWSEINFO struct {
	HwndOwner      HWND
	PidlRoot       uintptr
	PszDisplayName *uint16
	LpszTitle      *uint16
	UlFlags        uint32
	Lpfn           uintptr
	LParam         uintptr
	IImage         int32
}

var (
	// Library
	libshell32 *windows.LazyDLL

	// Functions
	shBrowseForFolder   *windows.LazyProc
	shGetPathFromIDList *windows.LazyProc
)

func init() {
	// Library
	libshell32 = windows.NewLazySystemDLL("shell32.dll")
	// 打开系统选择文件的对话框
	shBrowseForFolder = libshell32.NewProc("SHBrowseForFolderW")
	shGetPathFromIDList = libshell32.NewProc("SHGetPathFromIDListW")
}

func SHBrowseForFolder(lpbi *BROWSEINFO) uintptr {
	ret, _, _ := syscall.Syscall(shBrowseForFolder.Addr(), 1,
		uintptr(unsafe.Pointer(lpbi)),
		0,
		0)

	return ret
}
func SHGetPathFromIDList(pidl uintptr, pszPath *uint16) bool {
	ret, _, _ := syscall.Syscall(shGetPathFromIDList.Addr(), 2,
		pidl,
		uintptr(unsafe.Pointer(pszPath)),
		0)

	return ret != 0
}

func MemSet(s unsafe.Pointer, c byte, n uintptr) {
	ptr := uintptr(s)
	var i uintptr
	for i = 0; i < n; i++ {
		pByte := (*byte)(unsafe.Pointer(ptr + i))
		*pByte = c
	}
}

func main() {
	var bi BROWSEINFO
	MemSet(unsafe.Pointer(&bi), 0, unsafe.Sizeof(bi))
	bi.UlFlags = 0x00004000 // 要能查看文件
	pidl := SHBrowseForFolder(&bi)
	var path [MAX_PATH]uint16
	fmt.Println("pidl", pidl)
	SHGetPathFromIDList(pidl, &path[0])
	var path1 []uint16
	path1 = path[:]
	decodeContent := utf16.Decode(path1)
	fmt.Println("decodeContent", string(decodeContent))
}
