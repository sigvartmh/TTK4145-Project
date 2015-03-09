/*
    comedilib_version.h
    header file for comedilib's version number

    COMEDI - Linux Control and Measurement Device Interface
    Copyright (C) 1997-2000 David A. Schleef <ds@schleef.org>

    This program is free software; you can redistribute it and/or modify
    it under the terms of the GNU Lesser General Public License as
    published by the Free Software Foundation; either version 2 of the
    License, or (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU Lesser General Public
    License along with this program; if not, write to the Free Software
    Foundation, Inc., 59 Temple Place, Suite 330, Cambridge, MA 02111, USA.

*/

#ifndef _COMEDILIB_VERSION_H
#define _COMEDILIB_VERSION_H

/* Note that this header file first appeared in comedilib 0.10.0, so
 * the header file and macros won't exist in earlier versions unless
 * retro-fitted by a third party packager. */

#define COMEDILIB_VERSION_MAJOR	0
#define COMEDILIB_VERSION_MINOR	10
#define COMEDILIB_VERSION_MICRO	0

/**
 * COMEDILIB_CHECK_VERSION:
 * @major: major version
 * @minor: minor version
 * @micro: micro version
 *
 * Evaluates to %TRUE when the comedilib version (as indicated by
 * COMEDILIB_VERSION_MAJOR, COMEDILIB_VERSION_MINOR,
 * COMEDILIB_VERSION_MICRO) is at least as great as the given version.
 */
#define COMEDILIB_CHECK_VERSION(major, minor, micro)	\
	(COMEDILIB_VERSION_MAJOR > (major) || \
	 (COMEDILIB_VERSION_MAJOR == (major) && \
	  (COMEDLIB_VERSION_MINOR > (minor) || \
	   (COMEDILIB_VERSION_MINOR == (minor) && \
	    COMEDILIB_VERSION_MICRO >= (micro)))))

#endif	/* _COMEDILIB_VERSION_H */
