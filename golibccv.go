package golibccv

/*
#cgo CFLAGS: -I /Users/harm/prj/go/ccv/lib
#cgo LDFLAGS: -lm -ljpeg -lpng -lz -L/usr/X11/lib -lgsl -lfftw3f -lfftw3 -llinear -lblas -L /Users/harm/prj/go/ccv/lib -lccv
#include <ccv.h>
*/
import "C"

import (
	"unsafe"
)

type Image struct {
	image *C.ccv_dense_matrix_t
}

type cannyParams {
	size int
	lowThreshold int
	highThreshold int
}

type geometryFilteringParams struct {
	maxHeight int
	minHeight int
	minArea int
	letterOccludeThreshold int
	aspectRatio float64
	stdRatio float64
	thicknessRatio float64
}

type groupingParams struct {
	heightRatio float64
	intensityThreshold int
	distanceRatio float64
	intersectRatio float64
	elongateRatio float64
	letterThreshold int
}

type SwtParams struct {
	interval int
	minNeighbors int
	scaleInvariant int
	direction int
	sameWordThreshold float64
	/* canny parameters */
	cannyParams
	/* geometry filtering parameters */
	geometryFilteringParams
	/* grouping parameters */
	groupingParams
	/* break textline into words */
	breakdown int
	breakdownRatio float64
}

type ImageDiskType int
type ImageStreamType int
type ImageRawType int

const (
  ANY     = ImageDiskType(C.CCV_IO_ANY_FILE)
  BMP     = ImageDiskType(C.CCV_IO_BMP_FILE)
  JPEG    = ImageDiskType(C.CCV_IO_JPEG_FILE)
  PNG     = ImageDiskType(C.CCV_IO_PNG_FILE)
  BINAARY = ImageDiskType(C.CCV_IO_BINARY_FILE)
)

func ReadImage(src string, imageType ImageDiskType) (*Image, error) {
	image := new(Image)
	csrc := C.CString(src)
	defer C.free(unsafe.Pointer(csrc))
	C.ccv_read_impl(unsafe.Pointer(csrc), &image.image, C.CCV_IO_GRAY | C.int(imageType) , 0, 0, 0)
	return image, nil
}

func WriteImage(dst string, image Image, imageType ImageDiskType) error {
	cdst := C.CString(dst)
	defer C.free(unsafe.Pointer(cdst))
	lenght := C.int(0) // currently can't be set
	conf := C.int(0) // currently can't be set
	C.ccv_write(image.image, cdst, &lenght, C.int(imageType), unsafe.Pointer(&conf))
	return nil
}

func Swt(src Image, output Image, outputType int, params SwtParams) {
	
}

func SwtDetectWords(image Image, params SwtParams) {
}
