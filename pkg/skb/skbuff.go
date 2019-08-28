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

 package skb
 
 import (
	linkedList "github.com/Aniketh01/probe/pkg/list"
 )

 // ProbeSkbuff sk_buff routines
 type ProbeSkbuff struct {
	list linkedList.ProbeList
	head, end, data, tail, payload *uint8
	protocol uint16
	len, dlen, seq, endseq uint32
	refcount int
 }

 type probeSkbuffHead struct {
	 list linkedList.ProbeList
	 qlen uint32
 }

// ProbeSkbuffInit Initialises a sk_buff.
func ProbeSkbuffInit(size uint) *ProbeSkbuff {
}

func ProbeSkbuffPeek(head *probeSkbuffHead) *ProbeSkbuff {

}

func ProbeSkBuffDeinit(skb *ProbeSkbuff) {

}

func ProbeSkbuffPush(skb *ProbeSkbuff, len uint) uint8 {

}

func ProbeSkbuffReserve(skb *ProbeSkbuff, len uint) {

}

func ProbeSkbuffResetHead(skb *ProbeSkbuff) {

}

func ProbeSkbuffGetHead(skb *ProbeSkbuff) uint8 {

}

/*
 * Probe sk buff datagram/head queue function design
*/

func ProbeSkbuffQueueInit(head *probeSkbuffHead) {

}

func ProbSkbuffQueueDeinit(head *probeSkbuffHead) {
}

func ProbeSkbuffDequeue(head *probeSkbuffHead) *ProbeSkbuff {

}

func ProbeSkbuffQueueAdd(head *probeSkbuffHead, new *ProbeSkbuff, next *ProbeSkbuff) {

}

func ProbeSkbuffQueueAddEnd(head *probeSkbuffHead, new *ProbeSkbuff) {

}

func ProbeSkbuffQueueCheckEmpty(head *probeSkbuffHead) int {

}

func ProbeSkbuffQueueGetLen(head *probeSkbuffHead) uint32 {

}