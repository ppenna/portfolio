/*
 * MIT License
 *
 * Copyright(c) 2020 Pedro Henrique Penna <pedrohenriquepenna@gmail.com>
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

package drive

import (
	"fmt"
	"io"
	"os"
	"portfolio/internal/config"
)

type RemoteFile struct {
	id   string
	name string
	conn *RemoteConnection
}

// Downloads a remote file.
func (f *RemoteFile) Download() error {
	var out *os.File

	response, err := f.conn.srv.Files.Export(f.id, "text/csv").Download()
	if err != nil {
		return err
	}
	defer response.Body.Close()

	fmt.Printf("Downloading %s...\n", f.name)
	out, err = os.Create(config.DownloadsPath + f.name + ".csv")
	if err != nil {
		return err
	}
	defer out.Close()
	io.Copy(out, response.Body)

	return nil
}
