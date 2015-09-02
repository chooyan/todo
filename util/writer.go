package util

import (
        "os"
        "bufio"
)

func Write(scheduleJson string) {
        var writer *bufio.Writer

        filePath := GetFilePath()
	os.Remove(filePath)

        writeFile, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0600)
        writer = bufio.NewWriter(writeFile)
        defer writeFile.Close()

        writer.Write([]byte(scheduleJson))
        writer.Flush()
}
