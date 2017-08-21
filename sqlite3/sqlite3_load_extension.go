/*=================================
* Copyright(c)2015-2016 Yiqilai
* All rights reserved
* Inspired by mattn/go-sqlite3
*=================================*/

package sqlite3

/*
#include <sqlite3-binding.h>
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

func (c *SQLiteConn) loadExtensions(extensions []string) error {
	rv := C.sqlite3_enable_load_extension(c.db, 1)
	if rv != C.SQLITE_OK {
		return errors.New(C.GoString(C.sqlite3_errmsg(c.db)))
	}

	for _, extension := range extensions {
		cext := C.CString(extension)
		defer C.free(unsafe.Pointer(cext))
		rv = C.sqlite3_load_extension(c.db, cext, nil, nil)
		if rv != C.SQLITE_OK {
			return errors.New(C.GoString(C.sqlite3_errmsg(c.db)))
		}
	}

	rv = C.sqlite3_enable_load_extension(c.db, 0)
	if rv != C.SQLITE_OK {
		return errors.New(C.GoString(C.sqlite3_errmsg(c.db)))
	}
	return nil
}
