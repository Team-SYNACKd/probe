/*
 * MIT License
 *
 * Copyright (c) 2019 Aniketh Girish
 * This file is part of libprobe.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package list

// ProbeList is Linked list routines
type ProbeList struct {
	next, prev *ProbeList
}

// ProbeListInit Initialises a new list (A Linked list implementation)
// \param[in] head A probe_list/linked list context
func ProbeListInit(head *ProbeList) {

}

// ProbeListAdd take the value for the new node and adds that to the linklist chain.
//
// \param[in] new A probe list context to refer new node to be added
// \param[in] head A probe list context
func ProbeListAdd(new *ProbeList, head *ProbeList) {

}

// ProbeListDelete delete's a probe list node according to the node context passed.
//
// \param[in] ctx A probe list context
func ProbeListDelete(*ProbeList) {

}

// ProbeListAddEnd Take the value for the new node and attaches that to the end of the linklist chain.
//
// \param[in] new A probe list context to refer new node that need to be added
// \param[in] head A probe list context
func ProbeListAddEnd(new *ProbeList, head *ProbeList) {

}

// ProbeListCheckEmpty checks whether the list is empty. Returns True if the list is empty.
//
// \param[in] head A probe list context
func ProbeListCheckEmpty(head *ProbeList) {

}

func _listHead() {

}

func _listEntry() {

}

func _listFirstEntry() {

}

func _listEach() {

}

func _listForEachSafe() {

}
