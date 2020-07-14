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

// Directory
type RemoteDirectory struct {
	id   string
	name string
	conn *RemoteConnection
}

func GetRoot(conn *RemoteConnection) (*RemoteDirectory, error) {
	dir := &RemoteDirectory{
		name: "/",
		conn: conn,
	}

	return dir, nil
}

// Retrieve all files in a remote directory.
func (dir *RemoteDirectory) RetrieveFiles(remoteDir string) ([]*RemoteFile, error) {

	r, err := dir.conn.srv.Files.List().
		Fields("files(id, name)").Q(remoteDir + " in parents").Do()
	if err != nil {
		return nil, err
	}

	// Empty directory.
	if len(r.Files) == 0 {
		return nil, nil
	}

	remoteFiles := make([]*RemoteFile, 0)
	for _, i := range r.Files {
		f := &RemoteFile{
			id:   i.Id,
			name: i.Name,
			conn: dir.conn,
		}
		remoteFiles = append(remoteFiles, f)
	}

	return remoteFiles, nil
}
