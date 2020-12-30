package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
	"syscall"
)

/*
	io.WriteCloser:
		WriteCloser is the interface that groups the basic Write and Close methods.
	・interface の中の　Write は Writer interface

	type WriteCloser interface {
    	Writer
    	Closer
	}

	type Writer interface {
    	Write(p []byte) (n int, err error)
	}
*/

func main() {
	//if err := ToUpper(os.Stdin, os.Stdout, os.Stderr); err != nil {
	//	log.Println("failed to complete command. \n", err)
	//}
	customFlag()
}

/*
	https://gist.github.com/RyotaBannai/ac3de36689cb5ca4ef93c68dabc33eda
	https://gist.github.com/mattes/d13e273314c3b3ade33f

	func isExist(err error) bool {
		err = underlyingError(err)
		return err == syscall.EEXIST || err == syscall.ENOTEMPTY || err == ErrExist
	}

	func isNotExist(err error) bool {
		err = underlyingError(err)
		return err == syscall.ENOENT || err == ErrNotExist
	}
*/
func IsFileExists(fileName string) (bool, error) {
	_, err := os.Stat(fileName)

	if err != nil {
		if os.IsNotExist(err) {
			// file does not exist, do something
			return false, err
		} else {
			// more serious errors
			return false, err
		}
	}
	return true, nil
}

func ToUpper(src io.Reader, dst io.Writer, errDst io.Writer) error {
	/*
		*flag.StringVar もある
		https://blog.y-yuki.net/entry/2017/05/01/000000
	*/
	filePath := flag.String("f", "", "file path to convert to characters.")
	flag.Parse()
	if _, err := IsFileExists(*filePath); err != nil {
		return err
	}
	cmd := exec.Command("sh", "-c",
		fmt.Sprintf("cat %s | tr [:lower:] [:upper:]", *filePath)) //  'tr a-z A-Z' also fine
	/*
		*os/exec.Command.StdinPipe() でコマンドの標準入力へ書き込むための io.WriteCloser を取得できるため、
		それに対して io.Copy() で入力をコピーする
	*/
	var (
		stdin_, _  = cmd.StdinPipe()
		stdout_, _ = cmd.StdoutPipe()
		stderr_, _ = cmd.StderrPipe()
		wg         sync.WaitGroup
	)

	if err := cmd.Start(); err != nil {
		return err
	}

	wg.Add(3)
	go func() {
		_, err := io.Copy(stdin_, src)
		/*
			外部コマンドがなんならかの原因により終了した場合（head とかで読み取りを終了した場合なども含む）、
			go から書き込みができないためのエラーが発生する（= os.PathError）

			https://qiita.com/Maki-Daisuke/items/fbb514acc7bb8cf4a36b
			・EPIPE は untyped constant として定義されているため、それ自体では 型付けされていない
				→ ただコンパイル時に比較などで使われている場合は自動で型推論をしてくれる（そのためエラーにならない）
		*/
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.EPIPE {
			// ignore EPIPE
		} else {
			log.Println("failed to write to STDIN", err)
		}
		stdin_.Close()
		wg.Done()
	}()
	go func() {
		if _, err := io.Copy(dst, stdout_); err != nil {
			log.Println("copy to dst error")
		}
		stdout_.Close()
		wg.Done()
	}()
	go func() {
		if _, err := io.Copy(errDst, stderr_); err != nil {
			log.Println("copy to errDst error")
		}
		stderr_.Close()
		wg.Done()
	}()
	wg.Wait()         // 標準入出力の goroutine の終了を待つ
	return cmd.Wait() // command の終了を待つ
}
