/*
 Playlist structures tests.

 Copyleft 2013-2015 Alexander I.Grafov aka Axel <grafov@gmail.com>
 Copyleft 2013-2015 library authors (see AUTHORS file).

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.

 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.

 ॐ तारे तुत्तारे तुरे स्व
*/
package m3u8

import (
	"testing"
)

func CheckType(t *testing.T, p Playlist) {
	t.Logf("%T implements Playlist interface OK\n", p)
}

// Create new media playlist.
func TestNewMediaPlaylist(t *testing.T) {
	_, e := NewMediaPlaylist(1, 2)
	if e != nil {
		t.Fatalf("Create media playlist failed: %s", e)
	}
}
