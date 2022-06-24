// Copyright 2016 The go-vgo Project Developers. See the COPYRIGHT
// file at the top-level directory of this distribution and at
// https://github.com/go-vgo/robotgo/blob/master/LICENSE
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// http://www.apache.org/licenses/LICENSE-2.0> or the MIT license
// <LICENSE-MIT or http://opensource.org/licenses/MIT>, at your
// option. This file may not be copied, modified, or distributed
// except according to those terms.

package robotgo

import ps "github.com/vcaesar/gops"

// Nps process struct
type Nps struct {
	Pid  int32
	Name string
}

// Pids get the all process id
func Pids() ([]int32, error) {
	return ps.Pids()
}

// PidExists determines whether a process with a given ID exists.
func PidExists(pid int32) (bool, error) {
	return ps.PidExists(pid)
}

// Process returns a struct describing all processes.
func Process() ([]Nps, error) {
	var npsArr []Nps
	nps, err := ps.Process()
	for i := 0; i < len(nps); i++ {
		np := Nps{
			nps[i].Pid,
			nps[i].Name,
		}

		npsArr = append(npsArr, np)
	}

	return npsArr, err
}

// FindName returns a process's name given its ID.
func FindName(pid int32) (string, error) {
	return ps.FindName(pid)
}

// FindNames returns a slice of all process names.
func FindNames() ([]string, error) {
	return ps.FindNames()
}

// FindIds finds the IDs of all processes named with a subset of "name" (case
// insensitive).
func FindIds(name string) ([]int32, error) {
	return ps.FindIds(name)
}

// FindPath finds a process's path given its ID.
func FindPath(pid int32) (string, error) {
	return ps.FindPath(pid)
}

// Run runs a cmd shell.
func Run(path string) ([]byte, error) {
	return ps.Run(path)
}

// Kill kills a process given its ID.
func Kill(pid int32) error {
	return ps.Kill(pid)
}
