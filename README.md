Go bindings for [libccv](libccv.org).

Very much a work in progress.

## Compiling a shared library from libccv

Cgo won't work with static libraries, which is the default with libccv. You'll need to build your own.
To do this I edited the `makefile` in the `lib` directory of libccv. I've added the flag `-fPIC` to the CFLAGS variable:

    CFLAGS = -fPIC -O3 -ffast-math -Wall `cat .DEF`# -fprofile-arcs -ftest-coverage

And I more or less copy pasted the `libccv.a` make target and changed the command:

    libccv.so: ccv_cache.o ccv_memory.o 3rdparty/sha1/sha1.o 3rdparty/kissfft/kiss_fft.o 3rdparty/kissfft/kiss_fftnd.o 3rdparty/kissfft/kiss_fftr.o 3rdparty/kissfft/kiss_fftndr.o 3rdparty/kissfft/kissf_fft.o 3rdparty/kissfft/kissf_fftnd.o 3rdparty/kissfft/kissf_fftr.o 3rdparty/kissfft/kissf_fftndr.o 3rdparty/dsfmt/dSFMT.o 3rdparty/sfmt/SFMT.o ccv_io.o ccv_numeric.o ccv_algebra.o ccv_util.o ccv_basic.o ccv_resample.o ccv_transform.o ccv_classic.o ccv_daisy.o ccv_sift.o ccv_bbf.o ccv_mser.o ccv_swt.o ccv_dpm.o ccv_tld.o ccv_ferns.o
      $(CC) `cat .LN` -shared -o $@ $^

This yields a `libccv.so` file. Awesome.
