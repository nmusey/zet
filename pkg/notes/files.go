package notes

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const NOTES_DIRECTORY_KEY = "ZET_NOTES_DIR"

func ReadNote(filename string) (string, error) {
    filepath, err := getFilepath(filename)
    if err != nil {
        return "", err
    }

    file, err := os.Open(filepath)
    if err != nil {
        return "", err
    }

    bytes, err := io.ReadAll(file)
    if err != nil {
        return "", err
    }

    return string(bytes), nil
}

func WriteNote(filename string, contents string) error {
    filepath, err := getFilepath(filename)
    if err != nil {
        return err
    }

    file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 666)
    if err != nil {
        return err
    }

    _, err = file.Write([]byte(contents))
    return err
}

func ListNotes() error {
    noteDirName, err := getNoteDir()
    dir, err := os.ReadDir(noteDirName)
    if err != nil {
        return err
    }

    // TODO make this handle folders recursively
    for _, file := range dir {
        fmt.Println(file.Name())
    }

    return nil
}

func DeleteNote(filename string) error {
    filepath, err := getFilepath(filename)
    if err != nil {
        return err
    }

    return os.Remove(filepath)
}

func getFilepath(filename string) (string, error) {
    dir, err := getNoteDir()
    if err != nil {
        return "", err
    }

    return dir + addMarkdownExtension(filename), nil
}

func addMarkdownExtension(filename string) string {
    dotIndex := strings.Index(filename, ".")
    if dotIndex == -1 || dotIndex == 0 {
        return filename + ".md"
    }

    return filename
}

func getNoteDir() (string, error) {
    dir := os.Getenv(NOTES_DIRECTORY_KEY)
    if len(dir) > 0 && dir[len(dir)-1] != '/' {
        dir += "/"
    }

    _, err := os.Stat(dir)
    return dir, err
}
